# go-template

Starter template for Go services. See [AGENTS.md](AGENTS.md) for the full
tech stack, repo layout, and agent operating instructions.

```bash
mise install       # install pinned tools (go, sqlc, goose, golangci-lint, cobra-cli, lefthook)
mise exec -- lefthook install  # install git hooks
cp .env.example .env
mise run build
mise run run
```

## Releasing

Pushing a `v*` tag runs `.forgejo/workflows/release.yml`, which tests the code
and publishes a Forgejo release with GoReleaser. It authenticates with the
job's built-in token, so the repo needs no Actions secret for it.

One-time setup per repo created from this template:

1. Replace `go-template` everywhere, including `project_name` and
   `release.gitea.name` in `.goreleaser.yml`.
2. Check the config: `mise run release:check`, and `mise run release:snapshot`
   to build the archives into `dist/` without publishing.

Then `git tag v0.1.0 && git push origin v0.1.0`.

## Deploying

Fill this in per project: it's the contract the deploy (`brian/hosted`, the
`release_binary` role) relies on. Keep it current when any of it changes.

- **Releases:** each tag publishes `go-template_Linux_x86_64.tar.gz` (and
  `_arm64`), with the `go-template` binary at the archive root, and
  `go-template_<version>_checksums.txt` (the tag without its `v`) holding their
  SHA-256 sums.
- **Run command:** for example `go-template serve` or `go-template bot serve`.
- **Configuration** is environment variables only (`.env.example` explains
  each one):

  | Variable | In production |
  |---|---|
  | `EXAMPLE_VAR` | What to set, and whether it's a secret |

- **Serving:** the listen address and port, or "none" for a worker or bot
  that only makes outbound connections.
- **Health:** an HTTP endpoint (for example `GET /healthz`), or how failure
  shows instead (for example, exits non-zero on bad config, so systemd's
  restart loop is the signal).
- **Data:** paths it writes and how to back them up safely, or "none".
- **Migrations:** whether they run on startup, and how to roll back.
- **Logs:** go to stderr; note anything security-relevant they include.
- **Stopping:** what it does on SIGTERM.
