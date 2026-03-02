# Database Migrations with Goose

We use [goose](https://github.com/pressly/goose) for database migrations. Migrations run automatically on service startup — no manual steps needed.

Install the tool to make new migrations files
```bash
go install github.com/pressly/goose/v3/cmd/goose@v3.27.0
```

## How it works

- Each service manages its own migrations in `internal/db/migrations/`
- On startup, goose checks which migrations have not been applied and runs them
- Applied migrations are tracked in the `goose_db_version` table in each database
- Seeds in `internal/db/seeds/` are applied based on `GO_ENV`:
  - `development` — runs `seeds/development.sql`
  - `production` — skips seeding entirely

## Creating a new migration

If you need to update the db schema (Add table, alter table (add row, remove row, rename etc))

Run this from the migrations directory in the service you are wokring on:

```bash
goose create <name> sql
```

Example:

```bash
goose create add_sensor_index sql
```

This generates a timestamped file like `20260224130000_add_sensor_index.sql` with the goose annotations pre-filled:

```sql
-- +goose Up


-- +goose Down

```

Fill in the `Up` section with your changes and the `Down` section with how to reverse them.

## Notes

- Never edit or delete an existing migration file — write a new one instead
- The `Down` section is for rollbacks and is only ever run manually in development
- If a service has no seed file for the current environment, it logs a warning and continues normally
