-- 0004_admin_role.up.sql
-- services/auth-service now runs BetterAuth's admin() plugin, which reads
-- and writes these columns on every sign-up/session. Without them here,
-- BetterAuth's queries fail and no user gets saved. See db/migrations/
-- 0001_init.up.sql header for the hand-fold workflow.

ALTER TABLE "user"
    ADD COLUMN "role"       text,
    ADD COLUMN "banned"     boolean NOT NULL DEFAULT false,
    ADD COLUMN "banReason"  text,
    ADD COLUMN "banExpires" timestamptz;

ALTER TABLE "session"
    ADD COLUMN "impersonatedBy" text;
