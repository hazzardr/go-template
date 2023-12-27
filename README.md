# go-template

## Tools

### APIs
* HTTP server via labstack/echo
* Server codegen via deepmap/oapi-codegen

### Bootstrapping / Config
* CLI via Cobra
* Config via Viper

### Domain and Storage
* Postgres DB via pgx
* DB Migrations via go-migrate
* sqlc for generating go code from sql

```bash
make doctor
make deps # if required
make build
make run
```
# Needed locally:

