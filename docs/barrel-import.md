# Barrel export in React

## What is barreling

A barrel file is a module that exists solely to re-export things from other modules. The goal is cleaner import paths — instead of navigating deep relative paths to find a component, you import from a single well-known location.

## How it works

This project uses a **folder-per-component** pattern. Each component lives in its own folder with a thin `index.ts` that re-exports it. This gives the ergonomics of barreling without the performance cost of a single top-level barrel that re-exports everything.

### Structure

```
src/components/
  menuBox/
    menuBox.tsx      # component implementation: export function MenuBox() { ... }
    index.ts         # public API: export { MenuBox } from "./menuBox.tsx";
```

### Importing components

Use the `@/` path alias (points to `src/`) and import from the folder — no file extension needed:

```ts
// ✅ Correct — named import, resolves via index.ts
import { MenuBox } from "@/components/menuBox";

// ❌ Wrong — bypasses the public API
import { MenuBox } from "@/components/menuBox/menuBox.tsx";

// ❌ Wrong — no top-level barrel exists
import { MenuBox } from "@/components";
```

### Why named exports?

Components use named exports (`export function Foo`) rather than default exports (`export default function Foo`). This enforces consistent naming across the codebase — the name is locked at the export site, not chosen freely at each import site. Benefits:

- Consistent names everywhere — can't accidentally import `MenuBox` as `Banana`
- IDE autocomplete suggests the correct name
- Easier to search for where a component is used
- Refactoring tools reliably update all import references

### Adding a new component

1. Create a folder under `src/components/<ComponentName>/`
2. Add the component file: `<ComponentName>.tsx` using a **named export**:
   ```ts
   export function ComponentName() { ... }
   ```
3. Add `index.ts` re-exporting by name:
   ```ts
   export { ComponentName } from "./ComponentName.tsx";
   ```

### Path alias

The `@/` alias is configured in both `tsconfig.app.json` and `vite.config.ts`, pointing to `src/`. Use it everywhere to avoid `../` chains:

```ts
// ✅ Preferred
import Menu from "@/Menu.tsx";
import { Color } from "@/Theme/color.tsx";

// ❌ Avoid
import Menu from "../../Menu.tsx";
```

## Restrictions
> NB! @mui components should **NOT** be barreled due to performance issues during development. [mui docs](https://mui.com/material-ui/guides/minimizing-bundle-size/#avoid-barrel-imports)

### Example
```js
// ✅ Preferred
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';

// ❌ Slower in dev
import { Button, TextField } from '@mui/material';
```
Enforcing the avoidance of barrel-importing mui components is enforced with ESLint: (*eslint.config.js*)
```json
"rules": {
    "no-restricted-imports": [
      "error",
      {
        "patterns": [{ "regex": "^@mui/[^/]+$" }]
      }
    ]
  }
```