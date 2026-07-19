<script lang="ts">
  import { onMount } from "svelte";
  import { env } from "$env/dynamic/public";
  import { authClient } from "$lib/auth-client";

  const BACKEND = env.PUBLIC_BACKEND_URL ?? "http://localhost:8080";
  let health = $state("checking…");
  const session = authClient.useSession();

  onMount(async () => {
    try {
      const r = await fetch(`${BACKEND}/healthz`);
      health = (await r.json()).status;
    } catch {
      health = "unreachable";
    }
  });
</script>

<main style="font-family: sans-serif; max-width: 520px; margin: 4rem auto;">
  <h1>uzazi · admin</h1>
  <p>Health-worker / admin dashboard shell.</p>
  <p>Backend <code>/healthz</code>: <b>{health}</b></p>
  <p>Session: {$session.data ? $session.data.user.email : "signed out"}</p>
  <button onclick={() => authClient.signIn.social({ provider: "github" })}>Sign in</button>
</main>
