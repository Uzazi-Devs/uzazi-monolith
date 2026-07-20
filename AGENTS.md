# AGENTS.md

## Tooling

Use Bun, not npm, for all JavaScript/TypeScript work in this repository:

- `bun install` instead of `npm install`
- `bun run <script>` instead of `npm run <script>`
- `bunx <pkg>` instead of `npx <pkg>`

## Secrets and environment variables

Secrets are injected with [Doppler](https://www.doppler.com/), not `.env` files:

- Run services through Doppler, e.g. `doppler run -- bun run dev` or `doppler run -- go run ./cmd/...`.
- Never commit secrets, `.env` files, or hardcoded credentials. Read config from environment variables only.
- Never print, log, or echo secret values, including in debug output or CI logs.
- When adding a new required env var, document its name (never its value) in the relevant README and call it out in the PR description so it can be added to the Doppler project.
- Local fallback defaults in code are fine only for non-secret dev values (ports, localhost URLs).

## Styling

Use Tailwind CSS (v4) for all frontend styling. Every frontend under
`frontends/` has Tailwind set up:

- `app-nextjs`: via `@tailwindcss/postcss` (`postcss.config.mjs`), imported in `app/globals.css`
- `admin-svelte`: via `@tailwindcss/vite` (`vite.config.ts`), imported in `src/app.css`
- `marketing-astro`: via `@tailwindcss/vite` (`astro.config.mjs`), imported in `src/styles/global.css`

Prefer Tailwind utility classes over inline `style` attributes or ad hoc CSS.
Tailwind v4 needs no `tailwind.config.js`; add theme tokens with `@theme` in the CSS entry file.

## Code organization

One function per file, named after the operation it performs. This applies to
HTTP handlers, service methods, and analysis/processing functions alike:

- Name the file after the operation: `get.go`, `create.go` / `post.go`,
  `update.go` / `put.go`, `delete.go`, `list.go`, `analyze.go`, etc.
- Group the files by module/resource, e.g.
  `backend/internal/health/get.go`, `backend/internal/health/create.go`,
  `backend/internal/community/post.go`.
- Shared types, interfaces, and constructors for a module stay in a small
  `service.go` (or equivalent) in that module's directory.
- Same rule for TypeScript: one route handler or server function per file,
  e.g. Next.js route handlers split by method where the framework allows.
- Keep helpers used by only one function in that function's file. Promote a
  helper to a shared file only when a second caller appears.

## Pull request and git push rules

These rules apply to all work in this repository.

1. Never push directly to `main`. Create a branch and merge through a pull request.
2. Name branches `feature/<name>`, `fix/<name>`, or `chore/<name>`.
3. Do not force-push or delete `main`.
4. Before pushing, run the tests and checks relevant to every changed area. Record the commands and results in the PR under **How tested**.
5. Push only the working branch. Do not push unrelated commits or another contributor's changes.
6. Complete every applicable section and checkbox in `.github/pull_request_template.md`, including screenshots for UI changes.
7. A PR needs passing required status checks and at least one approving review before merge.
8. Code-owner review is required for owned paths. Changes under `db/migrations/` specifically require approval from `@deluxesande`.
9. For `db/migrations/` changes, read the shared-schema guidance in `README.md`, run `sqlc generate`, and update `services/auth-service/lib/auth.ts` when auth tables change.
10. Never bypass branch protection, required checks, approvals, or CODEOWNERS. Do not merge a PR while any required gate is pending or failing.
11. Agents must not push, open a PR, merge, rewrite remote history, or delete a remote branch unless the user explicitly requests that remote action.

## Recommended workflow

```bash
git switch main
git pull --ff-only
git switch -c chore/<name>   # or feature/<name>, fix/<name>
# make and validate focused changes
git add <changed-files>
git commit -m "Concise description"

# STOP HERE unless the user has explicitly authorized the remote actions below.
# Pushing and opening a PR are remote operations covered by rule 11.
git push -u origin HEAD
# open a PR into main and complete the repository PR template
```

Keep each PR focused. If unrelated work is discovered, put it in a separate branch and PR.
