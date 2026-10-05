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
