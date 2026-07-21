# uzazi

A healthcare SaaS platform for maternal care, built as a **modular monolith**
backend plus a small set of web frontends and one dedicated auth service.
Web-only stack: **Go, Astro, Next.js**, with **BetterAuth** for
authentication. No mobile, no Firebase.

## Repo layout

```
uzazi/
├── db/
│   ├── migrations/            # SINGLE SOURCE OF TRUTH — plain SQL, one runner
│   └── sqlc.yaml              # generates typed Go code into backend/internal/db
├── backend/                   # Go modular monolith (health, community, ai, media, auth-verify)
│   ├── cmd/server/            # main.go — wires modules together via interfaces
│   └── internal/
│       ├── auth/              # JWT VERIFICATION only (never issues credentials)
│       ├── health/            # clinical / triage records
│       ├── community/         # chat / forums
│       ├── ai/                # InferenceProvider (Gemma stub, swappable)
│       ├── media/             # Cloud Storage object refs
│       ├── db/                # sqlc-generated structs + queries
│       └── shared/            # config, logger, db pool, middleware
├── services/
│   └── auth-service/          # Next.js app running BetterAuth (the ONLY issuer)
├── frontends/
│   ├── marketing-astro/       # public landing page
│   └── app-nextjs/            # mother-facing app
└── infra/
    ├── docker-compose.yml     # everything, one command
    └── caddy/Caddyfile        # reverse proxy by subdomain
```

## Architecture notes

### Shared schema — the single source of truth

`db/migrations/` is the **only** place table definitions live. Two consumers
read from it, but neither owns it:

- the **Go backend** via [sqlc](https://sqlc.dev) → typed structs/queries in
  `backend/internal/db`;
- the **auth-service**, where BetterAuth is configured against these same
  tables (`services/auth-service/lib/auth.ts`), **not** its own auto-generated
  schema.

One migration runner (`migrate/migrate` in compose) applies the SQL. No ORM
auto-migrates on its own — that's what keeps the two consumers from diverging.

**To add or change a table:**

1. Edit `db/migrations/` (add a new numbered file, e.g. `0002_....sql`).
2. Run `sqlc generate` from `db/` — regenerates `backend/internal/db`.
   **Never edit generated code by hand.**
3. If the change touches auth tables (`user`/`session`/`account`/`verification`),
   run `bun run auth:generate` in `services/auth-service`, fold any diff back
   into `db/migrations/`, and update the field mapping in `lib/auth.ts`.

CI (`backend-ci`) runs `sqlc generate` and fails if the committed code is stale.

### Why auth-service is separate from the backend

BetterAuth is a TypeScript library — it can't run inside the Go process. It
lives in its own Next.js service and is the **single credential issuer** (email
+ OAuth, session + JWT). The Go backend's `internal/auth` only **verifies**
tokens: it checks signature + expiry against the auth-service JWKS
(`/api/auth/jwks`) and looks up the user via sqlc. It never signs or issues
anything.

### Backend module wiring

Modules expose Go interfaces (`health.Service`, `community.Service`, …) and are
wired together in `cmd/server/main.go`. They call each other in-process, not
over HTTP. `GET /healthz` is always available.

## Running locally

```bash
cd infra
docker compose up --build
```

This brings up postgres, redis, the migration runner, the backend, the
auth-service, both frontends, and Caddy. Migrations run automatically
before the backend and auth-service start.

| Service          | Direct URL              | Via Caddy               |
| ---------------- | ----------------------- | ----------------------- |
| Marketing        | http://localhost:4321   | http://uzazi.localhost  |
| Mother app       | http://localhost:3001   | http://app.localhost    |
| Auth (BetterAuth)| http://localhost:3000   | http://auth.localhost   |
| Backend API      | http://localhost:8080   | http://api.localhost    |

Each frontend shows the backend `/healthz` status and offers
a BetterAuth sign-in round trip.

## Branch & PR workflow

- **No direct pushes to `main`.** Branch protection requires a PR with **1
  approval** and passing status checks (admins included).
- The active GitHub `protect-main` ruleset applies to the default branch and
  blocks deletion, force-pushes, and direct updates. Changes must go through a
  PR with 1 approval, matching CODEOWNERS where configured, and required checks.
- Two backup guards remain in the repo:
  - **Local pre-push hook** — after cloning, run once:
    ```bash
    git config core.hooksPath .githooks
    ```
    This blocks `git push` to `main` from your machine.
  - **`protect-main` workflow** — CI fails on any commit that lands on `main`
    without a merged PR, so accidents are flagged immediately.
- Branch naming: `feature/<name>`, `fix/<name>`, `chore/<name>`.
- Every PR uses the template checklist. PRs touching `db/migrations/` require
  **@deluxe** approval (see `.github/CODEOWNERS`) because a schema change there
  can break the backend and auth-service simultaneously.

## Team ownership

| Area                            | Owner(s)                          |
| ------------------------------- | --------------------------------- |
| `db/migrations/`                | Deluxe                            |
| `backend/`                      | Deluxe, Vincent                   |
| `services/auth-service/`        | Deluxe, Vincent                   |
| `frontends/marketing-astro/`    | Deluxe                            |
| `frontends/app-nextjs/`         | Roro, Isaac                       |
| `.github/`                      | Deluxe                            |
