# Frontend Architecture — Feature-Sliced Design (FSD)

Reference: [feature-sliced.design](https://feature-sliced.design/docs/get-started/overview)

---

## Core concept

The frontend is structured around **FSD** — a methodology that organises code by business domain and responsibility rather than by technical type. Code is split into **layers**, each layer into **slices** (domain concepts), and each slice into **segments** (technical role).

**The golden rule:** a layer may only import from layers _below_ it.

```
app → pages → widgets → features → entities → shared
```

A `page` can import from `widgets`, `features`, `entities`, and `shared`.
A `feature` can import from `entities` and `shared` — but never from `pages` or `widgets`.

---

## Layer overview

| Layer | Import alias | Purpose |
|-------|-------------|---------|
| `app` | `@app` | Entry point, providers, routing, global styles |
| `pages` | `@pages` | Full-page components composed from lower layers |
| `widgets` | `@widgets` | Large self-contained UI blocks reused across pages |
| `features` | `@features` | User-facing interactions (forms, actions) |
| `entities` | `@entities` | Domain models and their data, logic, and display |
| `shared` | `@shared` | Framework-agnostic utilities, UI primitives, config |

Static files live in `src/assets/` (alias `@assets`) — accessible from any layer.

---

## Segment conventions

Every slice (e.g. `entities/sensor`) is divided into segments:

| Segment | Contains |
|---------|---------|
| `api/` | HTTP calls — functions that talk to the backend |
| `model/` | Stateful logic — hooks (`use*.ts`), TypeScript interfaces/schemas |
| `lib/` | Pure utility functions — no state, no side effects |
| `ui/` | React components that belong to this slice |

Each slice exposes a **public API** via `index.ts`. Consumers import from the slice root, not from internal paths:

```ts
// correct
import { useSensors, sortSensors } from "@entities/sensor";

// avoid — reaches into internals
import { useSensors } from "@entities/sensor/model/useSensors";
```

---

## Current file structure

```
src/
├── app/                          # Entry point and global setup
│   ├── main.tsx                  # Root component: providers, theme, router
│   ├── routes/index.tsx          # Route definitions
│   ├── providers/styles/         # Global CSS
│   ├── model/                    # (reserved) Global state — auth, user session
│   └── process/                  # (reserved) App-level async processes
│
├── assets/                       # Static files (SVG, PNG). Alias: @assets
│
├── entities/                     # Domain models
│   ├── (entity_name)/
│   │   ├── api/                  # HTTP calls — functions that talk to the backend
│   │   ├── model/                # Stateful logic — hooks (`use*.ts`), TypeScript interfaces/schemas
│   │   ├── lib/                  # Pure utility functions — no state, no side effects
│   │   ├── ui/                   # SensorMainInfo, SensorAllInfoPopUp, SensorsGenInfo
│   │   └── index.ts              # React components that belong to this slice
│
├── features/                     # User interactions
│   └── (feature_name)/
│       ├── ui/                   # Form modal for adding/other features
│       └── index.ts
│
├── widgets/                      # Composite UI blocks
│   ├── menu/                     # Menu.tsx — sidebar navigation
│   └── dashboard/                # AddBox.tsx — dashboard placeholder widget
│
├── pages/                        # Full-page components
│   ├── Home.tsx
│   ├── (Device_page).tsx
│   ├── (Admin_page).tsx
│   ├── StatusPage.tsx            # 404 page
│
└── shared/                       # No domain logic — reusable across everything
    ├── api/                      # Axios instance + generic apiRequest<T>() helper
    ├── config/
    │   ├── theme/                # MUI dark/light themes + ThemeContext
    │   └── navigation/           # mainPageList, subPageList config
    ├── lib/
    │   ├── formatter/            # Formats strings
    ├── mocks/                    # Dev mock data (sensors, gateways)
    └── ui/                       # Generic UI primitives
        └── (each component has ComponentName.tsx + index.ts)
```

---

## Where to put new code

| What you are building | Where it goes |
|-----------------------|--------------|
| New domain concept (e.g. `alarm`) | `entities/alarm/` with `api/`, `model/`, `lib/`, `ui/` |
| New user action (e.g. `deleteDevice`) | `features/deleteDevice/ui/` |
| Large composite section (e.g. `SensorDashboardPanel`) | `widgets/sensorDashboardPanel/` |
| New page | `pages/PageName.tsx` + register in `app/routes/index.tsx` |
| Reusable UI with no domain logic | `shared/ui/ComponentName/` |
| Generic helper (e.g. date formatter) | `shared/lib/` |
| App-wide state (e.g. auth, current user) | `app/model/` |

---

## Testing convention

Tests live **alongside the code they test**, inside the same segment folder:

```
entities/sensor/lib/sortSensors.ts
entities/sensor/lib/sortSensors.test.ts   ← unit test for pure logic
entities/sensor/model/useSensors.ts
entities/sensor/model/useSensors.test.ts  ← hook test (requires @testing-library/react)
entities/sensor/ui/SensorMainInfo.tsx
entities/sensor/ui/SensorMainInfo.test.tsx
```

Run all tests: `bun test`
