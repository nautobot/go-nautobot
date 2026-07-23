# Protected-file edits for `v3-nautobot-3.0` (Support Nautobot 3.0)

**Why:** the v3 pipeline (adopted from PR #37) reads the Nautobot version from `development/local_dev.env`, and its `release.yml` no longer wires the workflow `tag` input into generation. As shipped, CI/release would build against **3.0.1** (not the chosen **3.0.11**) and the release `tag` input would be decorative. These two edits make CI/release build against 3.0.11 and restore input-driven releases (matching how the 2.4 release worked).

Both files are protected (env file + CI workflow), so I can't edit them — please apply manually, then ping me.

---

## Change 1 of 2 — `development/local_dev.env`

**What:** bump the pinned Nautobot version from `3.0.1` to `3.0.11` (the latest 3.0.x, and what the committed bindings were generated from). `ci.yml` has no input, so it generates/tests against whatever this file pins.

**Find this block:**
```
NAUTOBOT_VER="3.0.1"
PYTHON_VER="3.12"
```

**Replace with:**
```
NAUTOBOT_VER="3.0.11"
PYTHON_VER="3.12"
```

(Only the version number on line 1 changes. `PYTHON_VER` stays `3.12`.)

---

## Change 2 of 2 — `.github/workflows/release.yml`

**What:** add a job-level `env:` block so the workflow `tag` input drives the Nautobot version used for generation. A runner env var overrides `--env-file local_dev.env` during docker-compose interpolation, so dispatching with `tag=3.0.11` will generate 3.0.11 regardless of the file pin. Without this, the `tag` input only appears in the Slack message and has no effect on what gets built.

**Find this block** (job header, lines ~13–17):
```
jobs:
  release-go-nautobot:
    runs-on: ubuntu-latest
    steps:
      - name: "Get go-nautobot generator code"
```

**Replace with** (inserts a 2-line `env:` block; `env:` is indented 4 spaces, the key 6 spaces):
```
jobs:
  release-go-nautobot:
    runs-on: ubuntu-latest
    env:
      NAUTOBOT_VER: "${{ github.event.inputs.tag }}"
    steps:
      - name: "Get go-nautobot generator code"
```

---

## Optional (nice-to-have, not required for correctness)

These are pre-existing carry-overs from #37; skip if you want to keep this PR tight. Both are in `.github/workflows/release.yml`:

- **Slack announces the wrong tag.** Line ~44 links `v${{ github.event.inputs.tag }}` (e.g. `v3.0.11` — the *Nautobot* version), but the real git tag is the computed go-nautobot version (`v3.0.0-beta`). A proper fix exposes the computed tag as a job `output` and references it in the slack-notify job; leaving it is purely cosmetic.
- **Deprecated action.** `actions/checkout@v3` (both jobs) rides on the deprecated Node 20 runtime — bump to `actions/checkout@v4` when convenient.

---

## Verification after applying

```
git status --short
```
You should see exactly:
```
 M development/local_dev.env
 M .github/workflows/release.yml
```

```
git diff development/local_dev.env .github/workflows/release.yml
```
Expected hunks:
- `local_dev.env`: one line, `3.0.1` → `3.0.11`.
- `release.yml`: a 2-line insertion (`env:` + `NAUTOBOT_VER: "${{ github.event.inputs.tag }}"`) between `runs-on: ubuntu-latest` and `steps:`.

**If anything else changed, undo and reapply** — these two files should have no other diffs.

## Why this works

- `ci.yml` runs `cd development && make` with no input, so it always builds the version in `local_dev.env`. Bumping it to `3.0.11` makes CI regenerate/test the same version the committed bindings came from.
- `release.yml` also runs `make`, but the new job-level `NAUTOBOT_VER` env var (sourced from the `tag` input) is exported into the runner shell, and docker-compose interpolation gives the shell env precedence over `--env-file`. So `gh workflow run release.yml -f tag=3.0.11` generates 3.0.11 and tags the computed `v3.0.0-beta`, exactly like the 2.4 release used its input.

## Next steps (once you confirm applied)

I'll `git add` both files onto `v3-nautobot-3.0`, commit them as a follow-up (`Wire release/CI to Nautobot 3.0.11`), and push to the PR branch. The PR itself will already be open (bindings commit) — this just makes its CI/release plumbing coherent.
