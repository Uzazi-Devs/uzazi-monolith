# Deploying to Production

Local dev (`infra/docker-compose.yml`) is unaffected. Production path:
**Render** for `backend` (Go), **Vercel** for `services/auth-service`,
`frontends/app-nextjs`, and `frontends/marketing-astro`, **Neon** for
Postgres, **Doppler** (`uzazi-monolith` project) for secrets — synced to
Render/Vercel via their dashboard integrations, never pasted into either
platform's UI or the repo by hand.

(Cloud Run was the original plan — `infra/deploy-cloudrun.sh` is kept for
reference but isn't the active path: none of the available GCP billing
accounts were usable.)

## 1. Neon

- Create a Neon project + database, grab the pooled connection string
  (`postgres://...neon.tech/uzazi?sslmode=require`).
- Run migrations against it once:
  ```bash
  migrate -path db/migrations -database "$NEON_URL" up
  ```

## 2. Doppler secrets

Project `uzazi-monolith`, config `prd` (dashboard.doppler.com). Set:

| secret | value |
|---|---|
| `DATABASE_URL` | the Neon pooled connection string |
| `BETTER_AUTH_SECRET` | `openssl rand -base64 32` |
| `GOOGLE_CLIENT_ID` | from Google Cloud Console OAuth client |
| `GOOGLE_CLIENT_SECRET` | same |

Then connect the Render and Vercel integrations from Doppler's Integrations
tab so both platforms pull these automatically.

`/admin/*` routes are gated by a BetterAuth JWT with `role: "admin"`, not a
separate credential — sign up normally, then promote the account once:
```sql
UPDATE "user" SET role = 'admin' WHERE email = 'you@example.com';
```

## 3. Google OAuth console

Add the Vercel auth-service URL as an authorized redirect URI:
`https://<auth-service>.vercel.app/api/auth/callback/google`. You won't know
the exact URL until the first `auth-service` deploy, so this is a
post-deploy step.

## 4. Render — backend

- Authorize the Render GitHub App for `Uzazi-Devs/uzazi-monolith` (one-time,
  Render dashboard).
- Create a Web Service: repo `uzazi-monolith`, branch `main`, root directory
  `backend`, Docker runtime (`backend/Dockerfile`), Free plan.
- Env vars sync from Doppler. Also set `CORS_ALLOWED_ORIGINS` (your frontend
  domains) and `AUTH_JWKS_URL` (`<auth-service URL>/api/auth/jwks`) — the
  latter is a post-deploy update once the auth-service URL is known.

## 5. Vercel — auth-service + frontends

`vercel link` inside each of `services/auth-service`, `frontends/app-nextjs`,
and `frontends/marketing-astro`, connecting each to its own Vercel project
pointed at the same GitHub repo (different root directory per project),
branch `main`.

Env vars sync from Doppler, plus per-project:

- **auth-service**: `BETTER_AUTH_URL`, `NEXT_PUBLIC_AUTH_URL` = its own
  Vercel URL (known after first deploy)
- **app-nextjs**: `NEXT_PUBLIC_AUTH_URL`, `NEXT_PUBLIC_BACKEND_URL`
- **marketing-astro**: `PUBLIC_WAITLIST_FORM_ENDPOINT` =
  `<backend URL>/waitlist`

`PUBLIC_AUTH_URL` / `PUBLIC_BACKEND_URL` are declared in
`infra/docker-compose.yml` for marketing-astro but unused in its source
(only `PUBLIC_WAITLIST_FORM_ENDPOINT` is read) — skip setting them.

## Status

- [x] Render CLI + Vercel CLI installed and authenticated
- [ ] Doppler `prd` secrets populated
- [ ] Render + Vercel GitHub App authorized for `uzazi-monolith`
- [ ] `feat/backend-jwt-admin-auth` merged to `main`
- [ ] Render backend service created
- [ ] Vercel projects created (auth-service, app-nextjs, marketing-astro)
