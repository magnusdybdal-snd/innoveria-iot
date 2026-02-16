# Linting and Formatting

## Installation

After cloning the repo, run the following from the `web/` directory:

```bash
cd web
bun install
```

This installs all dependencies and automatically runs the `prepare` script, which sets up the Husky git hooks.

> **Note:** The `.git` directory is at the repo root (one level above `web/`). Husky is configured to handle this via `"prepare": "cd .. && husky web/.husky"`.

---

## Frontend (TypeScript / React)

### Overview

We use three tools together to ensure consistent code quality:

| Tool | Purpose |
|------|---------|
| **ESLint** | Catches code errors and enforces rules (e.g. unused variables, React hooks rules) |
| **Prettier** | Automatically formats code (indentation, line breaks, quotes, etc.) |
| **lint-staged** | Runs ESLint and Prettier only on staged files at commit time |

### Commands

Run from the `web/` directory:

```bash
bun run lint         # Lint the entire project
bun run format       # Format the entire project
bun run formatcheck  # Check formatting without modifying files
```

### Configuration

- **ESLint:** `web/eslint.config.js` — flat config with TypeScript, React Hooks, and React Refresh plugins
- **Prettier:** `web/.prettierrc` — uses default settings (empty config)
- **lint-staged:** Defined in `web/package.json` under `"lint-staged"`

---

## Backend (Go)

### Overview

We use golangci-lint (v2) for both linting and format checking of Go code:

| Tool | Purpose |
|------|---------|
| **gofmt** | Standard Go formatter — enforces idiomatic formatting (like Prettier for Go) |
| **golangci-lint** | Meta-linter that runs multiple linters in one pass (govet, errcheck, staticcheck, etc.) |

golangci-lint is declared as a `tool` in each service's `go.mod` file (Go 1.24+ feature). This means **no manual installation is needed** — running `go tool golangci-lint` auto-downloads and caches the exact pinned version. All team members use the same version.

### Commands

Run from the repo root:

```bash
# Lint a specific service
cd services/api-gateway && go tool golangci-lint run ./...
cd services/collection-service && go tool golangci-lint run ./...

# Format Go files manually
gofmt -w services/
```

### Configuration

- **golangci-lint:** `/.golangci.yml` (v2 format) — auto-discovered from any subdirectory
- **Enabled linters:** govet, staticcheck, errcheck, ineffassign, unused, gocritic
- **Enabled formatters:** gofmt

### Adding a new Go service

1. Add the `tool` directive to the new service's `go.mod`:
   ```
   tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint
   ```
2. Run `go mod tidy` in the service directory
3. Add the service path to the `GO_SERVICES` array in `web/.husky/pre-commit`
4. Add the service path to the `GO_SERVICES` variable in `services/.gitlab-ci.yml`

---

## Pre-commit Hook

The pre-commit hook (`web/.husky/pre-commit`) runs automatically on every `git commit`:

### What happens on commit?

1. You run `git commit`
2. Husky triggers the pre-commit hook
3. **Frontend:** lint-staged runs on staged files:
   - `*.{ts,tsx}` → ESLint with `--fix`, then Prettier
   - `*.{json,css,md,html}` → Prettier
4. **Go** (only when `.go` files are staged):
   - `gofmt -w` auto-formats the staged files and re-stages them
   - `golangci-lint run ./...` runs on each service
5. If any linter finds errors that cannot be auto-fixed, the **commit fails** and you must fix the errors manually

### Bypassing the pre-commit hook

Sometimes you need to commit without linting, e.g. for a WIP commit or unfinished work. Use the `--no-verify` flag:

```bash
git commit -m "wip: unfinished changes" --no-verify
```

This skips all git hooks. **Use sparingly** — code should be linted and formatted before it is merged into `dev` or `main`.

---

## CI/CD Pipeline

The GitLab CI pipeline runs on every **merge request** as a safety net. It checks all files in the project (not just staged ones like the pre-commit hook).

### Pipeline stages

```
check → build → test
```

### Jobs

| Job | Stage | Image | What it does |
|-----|-------|-------|--------------|
| `web-check` | check | `oven/bun:latest` | ESLint + Prettier format check |
| `web-build` | build | `oven/bun:latest` | TypeScript + Vite production build |
| `web-test` | test | `oven/bun:latest` | `bun test` |
| `go-check` | check | `golang:1.25` | golangci-lint on all Go services |
| `go-build` | build | `golang:1.25` | `go build ./...` per service |
| `go-test` | test | `golang:1.25` | `go test ./...` per service |

### Configuration files

| File | Purpose |
|------|---------|
| `.gitlab-ci.yml` | Root pipeline — defines stages and includes the others |
| `web/.gitlab-ci.yml` | Frontend jobs (web-check, web-build, web-test) |
| `services/.gitlab-ci.yml` | Backend jobs (go-check, go-build, go-test) |

### Key difference: local vs CI

| | Local (pre-commit) | CI (pipeline) |
|---|---|---|
| **Scope** | Only staged files | All files in the project |
| **Formatting** | Auto-fixes and re-stages | Only checks — never modifies files |
| **When** | Every commit | Every merge request |

### Why golangci-lint is installed differently in CI

Locally, `go tool golangci-lint` compiles the linter from source (managed by `go.mod`). In CI, we use the [official install script](https://golangci-lint.run/docs/install/#install-from-binary) to download a pre-built binary, which is significantly faster (~5s vs ~60s).

The version in `services/.gitlab-ci.yml` (`GOLANGCI_LINT_VERSION`) must match the version in each service's `go.mod` to ensure consistency.
