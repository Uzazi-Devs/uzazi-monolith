"use client";

import { useEffect, useState } from "react";
import { authClient } from "@/lib/auth-client";

const BACKEND = process.env.NEXT_PUBLIC_BACKEND_URL ?? "http://localhost:8080";

export default function Page() {
  const [health, setHealth] = useState("checking…");
  const { data: session } = authClient.useSession();

  useEffect(() => {
    fetch(`${BACKEND}/healthz`)
      .then((r) => r.json())
      .then((d) => setHealth(d.status))
      .catch(() => setHealth("unreachable"));
  }, []);

  return (
    <main style={{ fontFamily: "sans-serif", maxWidth: 480, margin: "4rem auto" }}>
      <h1>uzazi · app</h1>
      <p>Mother-facing app shell.</p>
      <p>
        Backend <code>/healthz</code>: <b>{health}</b>
      </p>
      <p>Session: {session ? session.user.email : "signed out"}</p>
      {session ? (
        <button onClick={() => authClient.signOut()}>Sign out</button>
      ) : (
        <button onClick={() => authClient.signIn.social({ provider: "github" })}>
          Sign in with GitHub
        </button>
      )}
    </main>
  );
}
