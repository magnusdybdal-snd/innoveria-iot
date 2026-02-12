# Linting and Formatting

## Installation

After cloning the repo, run the following from the `web/` directory:

```bash
cd web
bun install
```

This installs all dependencies and automatically runs the `prepare` script, which sets up the Husky git hooks.

> **Note:** The `.git` directory is at the repo root (one level above `web/`). Husky is configured to handle this via `"prepare": "cd .. && husky web/.husky"`.

## Overview

We use three tools together to ensure consistent code quality:

| Tool | Purpose |
|------|---------|
| **ESLint** | Catches code errors and enforces rules (e.g. unused variables, React hooks rules) |
| **Prettier** | Automatically formats code (indentation, line breaks, quotes, etc.) |
| **lint-staged** | Runs ESLint and Prettier only on staged files at commit time |

### What happens on commit?

1. You run `git commit`
2. Husky triggers the pre-commit hook
3. lint-staged runs on your staged files:
   - `*.{ts,tsx}` → ESLint with `--fix`, then Prettier
   - `*.{json,css,md,html}` → Prettier
4. If ESLint finds errors that cannot be auto-fixed, the **commit fails** and you must fix the errors manually

## Commands

Run from the `web/` directory:

```bash
# Lint the entire project
bun run lint

# Format the entire project
bun run format

# Check formatting without modifying files
bun run formatcheck
```

## Configuration

- **ESLint:** `web/eslint.config.js` — flat config with TypeScript, React Hooks, and React Refresh plugins
- **Prettier:** `web/.prettierrc` — uses default settings (empty config)
- **lint-staged:** Defined in `web/package.json` under `"lint-staged"`

## Bypassing the pre-commit hook

Sometimes you need to commit without linting, e.g. for a WIP commit or unfinished work. Use the `--no-verify` flag:

```bash
git commit -m "wip: unfinished changes" --no-verify
```

This skips all git hooks. **Use sparingly** — code should be linted and formatted before it is merged into `dev` or `main`.
