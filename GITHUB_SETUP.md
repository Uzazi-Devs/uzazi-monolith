# GitHub setup

Exact `gh` commands to create the repo, protect `main`, and add the team.
Replace `ISAAC_LOGIN` / `RORO_LOGIN` / `VINCENT_LOGIN` once confirmed, and
update `.github/CODEOWNERS` (the placeholder handles are marked with a TODO).

## 1. Create the repo + push initial commit

```bash
cd /home/deluxe/Documents/Projects/uzazi
git init -b main
git add -A
git commit -m "Initial commit: uzazi monorepo scaffold"
gh repo create deluxesande/uzazi --private --source=. --remote=origin --push
```

## 2. Branch protection on `main` (PR + 1 approval, admins included, no direct pushes)

```bash
gh api -X PUT repos/deluxesande/uzazi/branches/main/protection \
  -H "Accept: application/vnd.github+json" \
  --input - <<'JSON'
{
  "required_status_checks": null,
  "enforce_admins": true,
  "required_pull_request_reviews": {
    "required_approving_review_count": 1,
    "require_code_owner_reviews": true
  },
  "restrictions": null,
  "allow_force_pushes": false,
  "allow_deletions": false
}
JSON
```

`require_code_owner_reviews: true` enforces the CODEOWNERS rules (incl.
`db/migrations/` → @deluxesande).

## 3. Require status checks — run AFTER the first CI run

The five workflows are **path-filtered**: a PR touching only the backend never
triggers the Astro/Svelte/Next checks. If you hard-require all of them, those
PRs can never merge. So enable required checks only once you know the check
names, and only for checks you want on every PR.

```bash
gh api -X PATCH repos/deluxesande/uzazi/branches/main/protection/required_status_checks \
  -H "Accept: application/vnd.github+json" \
  -f strict=true \
  -f 'contexts[]=build'   # add the specific check names you want as gates
```

## 4. Add collaborators (push access)

```bash
gh api -X PUT repos/deluxesande/uzazi/collaborators/ISAAC_LOGIN   -f permission=push
gh api -X PUT repos/deluxesande/uzazi/collaborators/RORO_LOGIN    -f permission=push
gh api -X PUT repos/deluxesande/uzazi/collaborators/VINCENT_LOGIN -f permission=push
```
