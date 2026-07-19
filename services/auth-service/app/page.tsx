"use client";

import { useState } from "react";
import { authClient } from "@/lib/auth-client";

export default function Page() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [msg, setMsg] = useState("");

  async function signUp() {
    const res = await authClient.signUp.email({ email, password, name: email });
    setMsg(res.error ? res.error.message ?? "error" : "signed up");
  }
  async function signIn() {
    const res = await authClient.signIn.email({ email, password });
    setMsg(res.error ? res.error.message ?? "error" : "signed in");
  }

  return (
    <main style={{ fontFamily: "sans-serif", maxWidth: 420, margin: "4rem auto" }}>
      <h1>uzazi · auth-service</h1>
      <p>BetterAuth — the single source of truth for sign-in.</p>
      <input placeholder="email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <br />
      <input
        placeholder="password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <br />
      <button onClick={signUp}>Sign up</button>
      <button onClick={signIn}>Sign in</button>
      <p>{msg}</p>
    </main>
  );
}
