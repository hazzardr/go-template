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
| Database | SQLite via `modernc.org/sqlite` (goose migrations; sqlc uses `database/sql`)
| Queries | `sqlc` — `db/schema.sql` + `db/query.sql` → `generated/domain` |
| Migrations | `goose` — SQL files in `migrations/` (`-- +goose Up/Down`) |
| Lint | `golangci-lint` (`.golangci.yml`) |
| Discord | `disgoorg/disgo` — bot in `internal/discord`, wired via `bot` Cobra command |

## Layout

- `main.go` — entrypoint; sets up slog, calls `cli.Execute()`
- `cmd/cli/` — Cobra commands (`root.go`, `serve.go`); add new commands here
- `internal/discord/` — disgo bot (`bot.go`: client + command wiring, `ping.go`: slash command). Add new slash commands here + to the `commands` slice, then run `mise exec -- go run . bot sync`
- `db/` — sqlc schema and queries (source of truth for `sqlc generate`); SQLite
- `migrations/` — goose migration files (sqlite3 dialect)
- `generated/` — codegen output; never edit by hand
- `mise.toml` — pinned tool versions and task definitions
- `.env` — local config (copy from `.env.example`; holds `DB_URL`, `DISCORD_TOKEN`)

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
mise run db:migration:create -- NAME sql -dir migrations
mise run db:migration:status
```

## Gotchas

- New sqlc queries need a `-- name: GetX :many` annotation or sqlc fails.
- The generated `domain` code uses `database/sql` with no driver baked in —
  when wiring the DB, open with `sql.Open("sqlite", DB_URL)` after
  `go get modernc.org/sqlite` (pure-Go, no cgo). The sqlite file lives in
  `data/` (gitignored); set its path via `DB_URL`.
- If generated code fails to compile with missing module errors, the codegen
  pulled in a new dependency: run `go get <module> && go mod tidy`.
- mise config must be trusted after edits: `mise trust`.
- `goose` and `cobra-cli` are installed through mise's `go:` backend, not the
  mise registry — don't "fix" them to registry names that don't exist.
- disgo releases move fast and change APIs between versions — check `go doc
  github.com/disgoorg/disgo/discord` for the current API instead of copying
  from older projects or training data.
