# Deploying to Cloud Run

Local dev (`infra/docker-compose.yml`) is unaffected — this covers the free-tier
production path: Cloud Run for `backend` + `auth-service`, Neon for Postgres,
and Vercel / Cloudflare Pages for the frontends. Run `infra/deploy-cloudrun.sh`
after completing steps 1-3 below.

## 1. Neon

- Create a Neon project + database, grab the pooled connection string
  (`postgres://...neon.tech/uzazi?sslmode=require`).
- Run migrations against it once:
  ```bash
  migrate -path db/migrations -database "$NEON_URL" up
  ```

## 2. GCP Secret Manager

One-time setup, values never go in the repo:

| secret name | value |
|---|---|
| `uzazi-database-url` | the Neon connection string |
| `uzazi-better-auth-secret` | `openssl rand -base64 32` |
| `uzazi-google-client-id` | from Google Cloud Console OAuth client |
| `uzazi-google-client-secret` | same |
| `uzazi-admin-user` | your choice, for `/admin` basic auth |
| `uzazi-admin-pass` | `openssl rand -base64 24` |

## 3. Google OAuth console

Add the Cloud Run auth-service URL as an authorized redirect URI:
`https://uzazi-auth-xxxx.run.app/api/auth/callback/google`. You won't know the
exact URL until the first `uzazi-auth` deploy, so this is a post-deploy step.

## 4. Run the deploy

```bash
PROJECT_ID=your-gcp-project \
CORS_ALLOWED_ORIGINS=https://your-app-domain,https://your-marketing-domain \
./infra/deploy-cloudrun.sh
```

`CORS_ALLOWED_ORIGINS` needs your real frontend domains up front. The
Cloud Run URLs for the backend/auth-service are resolved automatically inside
the script and printed at the end.

## 5. Frontend hosts

Set these as dashboard env vars (not `.env` files — those stay as local-dev
defaults only):

- **app-nextjs (Vercel)**: Project Settings → Environment Variables
  - `NEXT_PUBLIC_AUTH_URL` = auth-service URL
  - `NEXT_PUBLIC_BACKEND_URL` = backend URL
- **marketing-astro (Cloudflare Pages)**: Settings → Environment variables
  - `PUBLIC_WAITLIST_FORM_ENDPOINT` = `<backend URL>/waitlist`

`PUBLIC_AUTH_URL` / `PUBLIC_BACKEND_URL` are declared in
`infra/docker-compose.yml` for marketing-astro but unused in its source
(only `PUBLIC_WAITLIST_FORM_ENDPOINT` is read) — skip setting them.
