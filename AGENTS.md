# AGENTS.md

## Pull request and git push rules

These rules apply to all work in this repository.

1. Never push directly to `main`. Create a branch and merge through a pull request.
2. Name branches `feature/<name>`, `fix/<name>`, or `chore/<name>`.
3. Do not force-push or delete `main`.
4. Before pushing, run the tests and checks relevant to every changed area. Record the commands and results in the PR under **How tested**.
5. Push only the working branch. Do not push unrelated commits or another contributor's changes.
6. Complete every applicable section and checkbox in `.github/pull_request_template.md`, including screenshots for UI changes.
7. A PR needs passing required status checks and at least one approving review before merge.
8. Code-owner review is required for owned paths. Changes under `db/migrations/` specifically require approval from `@deluxesande`.
9. For `db/migrations/` changes, read the shared-schema guidance in `README.md`, run `sqlc generate`, and update `services/auth-service/lib/auth.ts` when auth tables change.
10. Never bypass branch protection, required checks, approvals, or CODEOWNERS. Do not merge a PR while any required gate is pending or failing.
11. Agents must not push, open a PR, merge, rewrite remote history, or delete a remote branch unless the user explicitly requests that remote action.

## Recommended workflow

```bash
git switch main
git pull --ff-only
git switch -c chore/<name>   # or feature/<name>, fix/<name>
# make and validate focused changes
git add <changed-files>
git commit -m "Concise description"
git push -u origin HEAD
# open a PR into main and complete the repository PR template
```

Keep each PR focused. If unrelated work is discovered, put it in a separate branch and PR.
