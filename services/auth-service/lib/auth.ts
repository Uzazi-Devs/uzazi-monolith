import { betterAuth } from "better-auth";
import { bearer, jwt } from "better-auth/plugins";
import { Pool } from "pg";

// BetterAuth is the ONLY credential issuer in uzazi. It reads the SAME tables
// defined in db/migrations (the single source of truth), NOT its own
// auto-generated schema.
//
// Workflow when auth tables change:
//   1. edit db/migrations/
//   2. run `bun run auth:generate` to see BetterAuth's expected schema
//   3. hand-fold any diff INTO db/migrations/ (never keep a second schema)
//   4. keep the field names below in sync
export const auth = betterAuth({
  database: new Pool({ connectionString: process.env.DATABASE_URL }),
  baseURL: process.env.BETTER_AUTH_URL ?? "http://localhost:3000",
  secret: process.env.BETTER_AUTH_SECRET,

  emailAndPassword: { enabled: true },

  socialProviders: process.env.GOOGLE_CLIENT_ID
    ? {
        google: {
          clientId: process.env.GOOGLE_CLIENT_ID,
          clientSecret: process.env.GOOGLE_CLIENT_SECRET as string,
        },
      }
    : undefined,

  // jwt() publishes a JWKS at /api/auth/jwks — the Go backend verifies tokens
  // against it. bearer() lets non-cookie clients authenticate with
  // `Authorization: Bearer <token>`.
  plugins: [jwt(), bearer()],
});
