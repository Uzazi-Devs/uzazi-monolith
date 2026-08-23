-- 0001_init.sql
-- SINGLE SOURCE OF TRUTH for every table in uzazi.
--
-- The "user"/"session"/"account"/"verification" tables mirror the output of
-- `npx @better-auth/cli generate` (run once in services/auth-service, then
-- hand-folded here). Column names are quoted camelCase to match BetterAuth's
-- default field mapping — do NOT rename them without also updating
-- services/auth-service/lib/auth.ts. The Go backend reads these same tables
-- through sqlc (see db/sqlc.yaml).

-- ---------- BetterAuth tables (owned by services/auth-service) ----------

CREATE TABLE "user" (
    "id"            text PRIMARY KEY,
    "name"          text        NOT NULL,
    "email"         text        NOT NULL UNIQUE,
    "emailVerified" boolean     NOT NULL DEFAULT false,
    "image"         text,
    "createdAt"     timestamptz NOT NULL DEFAULT now(),
    "updatedAt"     timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "session" (
    "id"        text PRIMARY KEY,
    "userId"    text        NOT NULL REFERENCES "user"("id") ON DELETE CASCADE,
    "token"     text        NOT NULL UNIQUE,
    "expiresAt" timestamptz NOT NULL,
    "ipAddress" text,
    "userAgent" text,
    "createdAt" timestamptz NOT NULL DEFAULT now(),
    "updatedAt" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "account" (
    "id"                    text PRIMARY KEY,
    "userId"                text        NOT NULL REFERENCES "user"("id") ON DELETE CASCADE,
    "accountId"             text        NOT NULL,
    "providerId"            text        NOT NULL,
    "accessToken"           text,
    "refreshToken"          text,
    "idToken"               text,
    "accessTokenExpiresAt"  timestamptz,
    "refreshTokenExpiresAt" timestamptz,
    "scope"                 text,
    "password"              text,
    "createdAt"             timestamptz NOT NULL DEFAULT now(),
    "updatedAt"             timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "verification" (
    "id"         text PRIMARY KEY,
    "identifier" text        NOT NULL,
    "value"      text        NOT NULL,
    "expiresAt"  timestamptz NOT NULL,
    "createdAt"  timestamptz NOT NULL DEFAULT now(),
    "updatedAt"  timestamptz NOT NULL DEFAULT now()
);

-- ---------- App tables (owned by the Go backend, snake_case) ----------

CREATE TABLE health_record (
    id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      text        NOT NULL REFERENCES "user"("id") ON DELETE CASCADE,
    kind         text        NOT NULL,              -- 'triage' | 'clinical'
    triage_level text,                              -- 'green' | 'yellow' | 'red'
    notes        text        NOT NULL DEFAULT '',
    created_at   timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE community_message (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    text        NOT NULL REFERENCES "user"("id") ON DELETE CASCADE,
    room       text        NOT NULL,
    body       text        NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE media_object (
    id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      text        NOT NULL REFERENCES "user"("id") ON DELETE CASCADE,
    bucket       text        NOT NULL,
    object_key   text        NOT NULL,
    content_type text        NOT NULL,
    created_at   timestamptz NOT NULL DEFAULT now()
);
