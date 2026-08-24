import { betterAuth } from "better-auth";
import { bearer, jwt, haveIBeenPwned, admin } from "better-auth/plugins";
import { Pool } from "pg";

const googleClientId = process.env.GOOGLE_CLIENT_ID;
const googleClientSecret = process.env.GOOGLE_CLIENT_SECRET;

//Google auth is Optional ,but both values must be provided together
if (
  (googleClientId && !googleClientSecret) ||
  (!googleClientId && googleClientSecret)
) {
  throw new Error(
    "Both GOOGLE_CLIENT_ID and GOOGLE_CLIENT_SECRET must be set together.",
  );
}

// BetterAuth is the ONLY credential issuer in uzazi. It reads the SAME tables
// defined in db/migrations (the single source of truth), NOT its own
// auto-generated schema.
//
// Workflow when auth tables change:
//   1. edit db/migrations/
//   2. run `bun run auth:generate` to see BetterAuth's expected schema
//   3. hand-fold any diff INTO db/migrations/ (never keep a second schema)
//   4. keep the field names below in sync
const socialProviders =
  googleClientId && googleClientSecret
    ? {
        google: {
          clientId: googleClientId,
          clientSecret: googleClientSecret,
        },
      }
    : {};

export const auth = betterAuth({
  database: new Pool({ connectionString: process.env.DATABASE_URL }),
  baseURL: process.env.BETTER_AUTH_URL ?? "http://localhost:3000",
  secret: process.env.BETTER_AUTH_SECRET,

  emailAndPassword: { enabled: true },

  socialProviders,

  rateLimit: {
    enabled: true,
    window: 60,
    max: 20,
  },

  // jwt() publishes a JWKS at /api/auth/jwks — the Go backend verifies tokens
  // against it. bearer() lets non-cookie clients authenticate with
  // `Authorization: Bearer <token>`.
  plugins: [jwt(), bearer(), admin(), haveIBeenPwned()],
});
