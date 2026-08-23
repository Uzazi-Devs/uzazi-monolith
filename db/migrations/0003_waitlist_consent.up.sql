-- 0003_waitlist_consent.up.sql
-- The waitlist form requires a consent checkbox before it will submit, but
-- nothing recorded that consent was actually given. Record when.

ALTER TABLE waitlist_signup ADD COLUMN consented_at timestamptz NOT NULL DEFAULT now();
