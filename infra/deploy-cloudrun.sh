#!/usr/bin/env bash
# Deploy backend + auth-service to Cloud Run. Local dev (docker-compose) is untouched.
#
# One-time setup before first run:
#   gcloud artifacts repositories create uzazi --repository-format=docker --location=$REGION
#   gcloud auth configure-docker ${REGION}-docker.pkg.dev
#
#   Create secrets (values from Neon / Google OAuth console — never commit these):
#     echo -n "$NEON_CONNECTION_STRING"      | gcloud secrets create uzazi-database-url        --data-file=-
#     echo -n "$(openssl rand -base64 32)"   | gcloud secrets create uzazi-better-auth-secret   --data-file=-
#     echo -n "$GOOGLE_CLIENT_ID"            | gcloud secrets create uzazi-google-client-id     --data-file=-
#     echo -n "$GOOGLE_CLIENT_SECRET"        | gcloud secrets create uzazi-google-client-secret --data-file=-
#
#   Run migrations against Neon once:
#     migrate -path db/migrations -database "$NEON_CONNECTION_STRING" up
#
# Usage: PROJECT_ID=your-gcp-project ./infra/deploy-cloudrun.sh

set -euo pipefail

PROJECT_ID="${PROJECT_ID:?set PROJECT_ID}"
REGION="${REGION:-us-central1}"
REPO="${REGION}-docker.pkg.dev/${PROJECT_ID}/uzazi"

gcloud config set project "$PROJECT_ID" >/dev/null

# --- auth-service (needs repo-root build context) ---
AUTH_IMAGE="${REPO}/auth-service:$(git rev-parse --short HEAD)"
docker build -f services/auth-service/Dockerfile -t "$AUTH_IMAGE" .
docker push "$AUTH_IMAGE"

gcloud run deploy uzazi-auth \
  --image "$AUTH_IMAGE" \
  --region "$REGION" \
  --allow-unauthenticated \
  --set-secrets DATABASE_URL=uzazi-database-url:latest,BETTER_AUTH_SECRET=uzazi-better-auth-secret:latest,GOOGLE_CLIENT_ID=uzazi-google-client-id:latest,GOOGLE_CLIENT_SECRET=uzazi-google-client-secret:latest

AUTH_URL=$(gcloud run services describe uzazi-auth --region "$REGION" --format='value(status.url)')

# auth-service needs to know its own URL — set it now that we have it
gcloud run services update uzazi-auth \
  --region "$REGION" \
  --set-env-vars BETTER_AUTH_URL="$AUTH_URL",NEXT_PUBLIC_AUTH_URL="$AUTH_URL"

# --- backend (self-contained context) ---
gcloud run deploy uzazi-backend \
  --source backend \
  --region "$REGION" \
  --allow-unauthenticated \
  --set-env-vars AUTH_JWKS_URL="${AUTH_URL}/api/auth/jwks",AI_PROVIDER="${AI_PROVIDER:-none}",CORS_ALLOWED_ORIGINS="${CORS_ALLOWED_ORIGINS:?set CORS_ALLOWED_ORIGINS to your frontend origin(s)}" \
  --set-secrets DATABASE_URL=uzazi-database-url:latest

BACKEND_URL=$(gcloud run services describe uzazi-backend --region "$REGION" --format='value(status.url)')

echo "auth-service: $AUTH_URL"
echo "backend:      $BACKEND_URL"
echo "Point your frontends' NEXT_PUBLIC_AUTH_URL / NEXT_PUBLIC_BACKEND_URL / PUBLIC_* env vars at these."
