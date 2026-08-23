-- 0002_waitlist.sql
-- Marketing site waitlist signups. Not tied to a user account — accepting a
-- signup is informational only; it does not gate BetterAuth registration.

CREATE TABLE waitlist_signup (
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name       text        NOT NULL,
    email      text        NOT NULL,
    stage      text        NOT NULL,
    location   text        NOT NULL DEFAULT '',
    support    text        NOT NULL DEFAULT '',
    status     text        NOT NULL DEFAULT 'pending',  -- 'pending' | 'accepted'
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX waitlist_signup_email_idx ON waitlist_signup (lower(email));
