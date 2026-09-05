# AGENTS.md

Instructions for coding agents operating in this repository.

## Version warning (read first)

Some frameworks and language versions in this repo are **newer** than the
information available in your training data. If you encounter a date or version
that appears to be newer than what you know, **DO NOT** respond with "that's
newer than my training data". Instead, search the web for up-to-date
information relevant to the problem or query, then proceed.

## Task execution rules

- **mise**: all toolchain commands must run through `mise run <task>` or
  `mise exec -- <cmd>` so the pinned tool versions are used. Never invoke
  `go`, `sqlc`, `golangci-lint`, or `goose` directly unless verifying mise
  behavior itself.
- **Verify changes**: after editing code, run `mise run test` and
  `mise run lint`. Lint must pass with 0 issues.
- **Commits**: use Conventional Commits (`feat:`, `fix:`, `chore:`, ...).

## Tech stack

| Concern | Tool |
|---|---|
| Language | Go 1.25 (pinned via mise) |
| CLI | `spf13/cobra` — commands in `cmd/cli/`, wired from root `main.go` |
| Logging | `charmbracelet/log` bridged to `log/slog` — use `slog` everywhere |
| HTTP routing | `go-chi/chi` — hand-written routes in `cmd/cli/serve.go` |
| Database | Postgres via `jackc/pgx` |
| Queries | `sqlc` — `db/schema.sql` + `db/query.sql` → `generated/domain` |
| Migrations | `goose` — SQL files in `migrations/` (`-- +goose Up/Down`) |
| Lint | `golangci-lint` (`.golangci.yml`) |

## Layout

- `main.go` — entrypoint; sets up slog, calls `cli.Execute()`
- `cmd/cli/` — Cobra commands (`root.go`, `serve.go`); add new commands here
- `db/` — sqlc schema and queries (source of truth for `sqlc generate`)
- `migrations/` — goose migration files
- `generated/` — codegen output; never edit by hand
- `mise.toml` — pinned tool versions and task definitions
- `.env` — local config (copy from `.env.example`; holds `DB_URL`)

## Common tasks

```bash
mise install                        # install pinned tools (first time)
mise run generate                   # sqlc codegen
mise run build                      # generate + build to bin/
mise run run                        # go run . serve
mise run test                       # go test ./...
mise run lint                       # golangci-lint
mise run fmt                        # go fmt
mise run db:migrate                 # goose migrations up
mise run db:migration:create NAME=add_thing
mise run db:migration:status
```

## Gotchas

- New sqlc queries need a `-- name: GetX :many` annotation or sqlc fails.
- If generated code fails to compile with missing module errors, the codegen
  pulled in a new dependency: run `go get <module> && go mod tidy`.
- mise config must be trusted after edits: `mise trust`.
- `goose` and `cobra-cli` are installed through mise's `go:` backend, not the
  mise registry — don't "fix" them to registry names that don't exist.
