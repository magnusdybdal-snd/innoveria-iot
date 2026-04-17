# Payload Schema System

Sensors send payloads with arbitrary field names. `AccAmp`, `bus1`, `temperature` — each
sensor model or installation may differ. This document describes how the system makes sense
of that raw data so context-service can aggregate it meaningfully.

---

## Core Concepts

### Three tables, one goal

```
measurement_type       — the canonical vocabulary (what things are called)
payload_schema         — profile-level mapping  (fixed sensors: profile → key → type)
sensor_metric          — per-sensor mapping     (configurable sensors: sensor → key → type)
sensor_profile_config  — marks which profiles require per-installation configuration
```

Context-service speaks only in `measurement_type` slugs. It asks device-service
"what payload key gives me `accumulated_current` for device X?" and extracts that field
from the raw JSONB in collection-service.

---

### MeasurementType — the vocabulary

A controlled list of canonical slugs. Every label in the system must reference a slug
from this table. Admin-managed, system-level (no company scope).

```
slug                  display_name            default_unit   deprecated
accumulated_current   Accumulated Current     A              false
temperature           Temperature             °C             false
nitrogen_ppm          Nitrogen                ppm            false
```

- Slugs are immutable — they are the cross-service contract
- Types are never deleted, only deprecated (hides from dropdowns, existing references stay valid)

---

### PayloadSchema — profile-level mapping (fixed sensors)

Maps raw payload keys to measurement types for a Chirpstack device profile.
All sensors on that profile inherit these mappings automatically.

```
chirpstack_profile_id   payload_key   measurement_type        unit
abc-123                 AccAmp        accumulated_current     A
abc-123                 Amp           instantaneous_current   A
abc-123                 Temp          temperature             °C
```

- Only fully labeled rows are used — no draft concept
- Admin sets this up once per hardware model

---

### SensorMetric — per-sensor mapping (configurable sensors)

Maps raw payload keys to measurement types for one specific sensor.
Used for devices like the UC300 where field meaning depends on per-installation wiring.

```
sensor_id   payload_key   measurement_type   unit
uuid-x      bus1          nitrogen_ppm       ppm
uuid-x      bus2          pressure_bar       bar
```

---

### SensorProfileConfig — fixed vs configurable signal

A small overlay table that marks which Chirpstack profiles require per-installation
configuration. No row = fixed schema (the default for all 158 community profiles).

```sql
CREATE TABLE device.sensor_profile_config (
    chirpstack_profile_id text PRIMARY KEY,
    configurable_schema   boolean NOT NULL DEFAULT false
);
```

Only profiles like the UC300 ever get a row here. Everything else inherits the default silently.
This table is completely decoupled from the Chirpstack repo sync.

---

## Two Sensor Types

### Fixed-schema sensors (EM340, distance, humidity/temp, current)

The hardware model always outputs the same fields. Same codec, same keys, every sensor.

- Config lives at: `payload_schema` (profile level)
- Set up by: system admin, once per hardware model
- Requires: at least one sensor on that profile to have sent a reading (keys come from live data)
- Sensors inherit mappings automatically via `GetEffectiveMetrics` fallback

### Configurable-schema sensors (Milesight UC300)

Fields depend on per-installation wiring. Two UC300s at different factories will have
`bus1` meaning completely different things.

- Config lives at: `sensor_metric` (per sensor)
- Set up by: tenant operator per installation — or admin for demo/early setup
- No profile-level template, by design
- Profile is marked in `sensor_profile_config` with `configurable_schema = true`

---

## GetEffectiveMetrics — the resolution logic

`SensorMetricService.GetEffectiveMetrics(deviceEUI)` is the single call any service
makes to get the payload key mappings for a sensor:

1. Look up per-sensor `sensor_metric` rows
2. If any exist → return them (configurable sensor path)
3. If none → look up labeled `payload_schema` rows for the sensor's `chirpstack_profile_id`
4. If still nothing → return empty slice (sensor not yet configured)

Context-service calls this endpoint before extracting fields from raw measurements.
An empty result means the sensor has no usable mappings yet.

---

## Service Boundaries — Decoupling

| Direction | Allowed | Notes |
|---|---|---|
| collection-service → device-service | No | Collection stores raw data only |
| device-service → collection-service | No | UI bridges them client-side |
| context-service → device-service | Yes | Calls `GET /sensors/{eui}/metrics` |
| context-service → collection-service | Yes | Calls measurement endpoints |
| Admin → tenant sensor data | No | Admin uses sample-eui endpoint — gets one EUI, nothing else |

The UI is the bridge for the labeling flow: it fetches payload keys from collection-service
and saves labels to device-service. No service-to-service coupling.

---

## Frontend Overview

Four UI pieces need to be built. Two are admin-only, two are tenant-facing (lower priority).
This section describes what each screen looks like and how it behaves. API calls are in the detailed sections below.

---

### A — Measurement Type Manager (Admin)

A settings page where the admin manages the vocabulary that all labeling depends on.

Looks like a simple table with a slug, display name, unit, and status column.
Each row has a "Deprecate" button. 

A "New type" button opens a small form at the top.

- Slug field auto-formats as you type: lowercase, spaces become underscores
- Slug is locked after save — it cannot be edited, only deprecated. This is because changing it will conflict with already labeled measurement / context calculations
- Deprecated rows are greyed out or hidden. Also hidden from all dropdowns elsewhere in the system (own endpoint takes care of this)
- No delete — types stay forever once created

---

### B — Payload Schema Labeling Page (Admin)

The main admin tool for making sensors meaningful. One page that handles both sensor types.

At the top: a profile dropdown listing all Chirpstack profiles.

Once a profile is selected, a form appears below containing:

**1. Toggle: "This profile requires per-installation configuration"**
Pre-filled from the database if already set for this profile, otherwise off by default.
This distinguishes a fixed sensor (EM340) from a configurable bus converter (UC300).

**2a. If toggle is OFF (fixed sensor — e.g. EM340):**
The labeling table appears immediately, pre-filled with live payload keys fetched from any sensor on this profile:
The Unit is already configured on the measurement_type page and not selectable

```
Payload key       Measurement type            Unit
AccAmp            [ Accumulated Current ▼ ]   [ A  ]
Amp               [ Instantaneous Current ▼ ] [ A  ]
Temp              [ Temperature ▼ ]            [ °C ]
```

Measurement type dropdown shows display names from the vocabulary. Save button at the bottom.
If no sensor on this profile has sent data yet, show a notice instead of the table:
"No readings yet — deploy a sensor first."

**2b. If toggle is ON (configurable sensor — e.g. UC300):**
A sensor picker appears — a dropdown of sensors using this profile, shown by name.
Admin picks the specific sensor they want to configure.
The same labeling table appears below, populated with that sensor's payload keys.
Save writes to that sensor only, not the profile.

---

### C — Sensor List badge (Tenant, lower priority)

On the existing sensor list, each sensor gets a small status indicator showing whether it is configured or not. No separate page — just a badge/icon on the existing list rows. Sensors where `metrics_configured = false` show a warning badge.

---

### D — Sensor Detail — Metric Configuration (Tenant, lower priority)

An additional section on the existing sensor detail page, visible only for configurable sensors that are not yet configured.

Shows the same labeling table as screen B, pre-filled with the sensor's live payload keys. Tenant picks measurement types from the vocabulary and saves. Once saved, the section switches to a read-only summary with an edit button.

If no reading has arrived yet, show: "No readings yet — check back after the first transmission."

---

## Admin Flow — Fixed Sensor (e.g. EM340)

```
1. Admin installs the sensor, it sends at least one reading
2. Opens payload schema page, selects the Chirpstack profile
3. configurable_schema box is NOT checked
4. UI calls GET /sensors/sample-eui?chirpstack_profile_id={id}
   → gets one EUI from any sensor on that profile
5. UI calls GET /collection/payload-tags?device_eui={eui}
   → gets the actual payload keys: ["AccAmp", "Amp", "Temp"]
6. Admin labels each key by selecting a measurement type slug from the vocabulary
7. Saves → PUT /payload-schema/{chirpstack_profile_id}
8. Done — all sensors on this profile now return these mappings via GetEffectiveMetrics
```

If no sensor on the profile has sent data yet, step 5 returns empty and the UI shows:
"No readings available — deploy a sensor and wait for its first transmission."

---

## Admin Flow — Configurable Sensor (e.g. UC300, demo/early setup)

```
1. Admin installs the sensor, it sends at least one reading
2. Opens payload schema page, selects the Chirpstack profile
3. Admin checks configurable_schema box
   → PUT /sensor-profile-config/{chirpstack_profile_id}  (sets configurable_schema = true)
   → devEUI picker appears
4. Admin selects the specific sensor EUI
5. UI calls GET /collection/payload-tags?device_eui={eui}
   → gets keys: ["bus1", "bus2", "bus3"]
6. Admin labels each key with a measurement type slug
7. Saves → PUT /sensors/{eui}/metrics  (per-sensor, not profile-level)
8. Done — that specific sensor is configured. Other UC300s still need per-sensor setup.
```

---

## Tenant Flow — Configurable Sensor (production, if we have time)

For UC300 deployments where the tenant labels their own sensors.
The backend already fully supports this — it is a frontend addition only.

```
1. Tenant sees sensor flagged as unconfigured in sensor list (metrics_configured = false)
2. Opens sensor detail page
3. UI calls GET /collection/payload-tags?device_eui={eui}
   → if empty: "No readings yet — check back after first transmission"
   → if has keys: shows labeling form
4. Tenant labels each key with a measurement type slug
5. Saves → PUT /sensors/{eui}/metrics
6. Sensor is now configured
```

---

## metrics_configured — sensor status

No boolean column is stored on the sensor. Status is derived server-side in the sensor
list query using a subquery:

```
metrics_configured = true  if sensor has sensor_metric rows
                        OR  sensor's profile has labeled payload_schema rows
metrics_configured = false otherwise
```

This avoids N+1 calls in the UI while keeping the sensor table clean.
Requires a backend change to the sensor list query and response DTO.

---

## What Is Not Yet Implemented

### 1. metrics_configured on sensor list
Add derived `metrics_configured` field to the sensor list endpoint response.
Computed via subquery — no schema change needed.

### 2. Sensor list filter by chirpstack_profile_id
Add `?chirpstack_profile_id=` as an optional query param to `GET /sensors`.
Used by the configurable sensor devEUI picker — currently the UI fetches all sensors and filters client-side.
Consistent with the existing `production_resource_id` filter pattern.

### 3. Context-service wiring
Context-service does not yet call `GET /sensors/{eui}/metrics` before aggregating.
Will be handled in a separate branch.

---

## API Reference

### Measurement Types (device-service)
| Method | Path | Description | Access |
|---|---|---|---|
| GET | `/api/v1/device/measurement-types` | List active types | All |
| GET | `/api/v1/device/measurement-types/all` | List including deprecated | Admin |
| POST | `/api/v1/device/measurement-types` | Create new type | Admin |
| PATCH | `/api/v1/device/measurement-types/{slug}/deprecate` | Deprecate a type | Admin |

### Payload Schema (device-service)
| Method | Path | Description | Access |
|---|---|---|---|
| GET | `/api/v1/device/payload-schema/{profile_id}` | All labeled rows for a profile | Admin |
| PUT | `/api/v1/device/payload-schema/{profile_id}` | Save labels for a profile | Admin |

### Sensor Profile Config (device-service)
| Method | Path | Description | Access |
|---|---|---|---|
| GET | `/api/v1/device/sensor-profile-config/{profile_id}` | Get config for a profile | Admin |
| PUT | `/api/v1/device/sensor-profile-config/{profile_id}` | Set configurable_schema flag | Admin |

### Sensor Metrics (device-service)
| Method | Path | Description | Access |
|---|---|---|---|
| GET | `/api/v1/device/sensors/{eui}/metrics` | Effective metrics (per-sensor → profile fallback) | All |
| PUT | `/api/v1/device/sensors/{eui}/metrics` | Save per-sensor labels | Admin + Tenant |
| GET | `/api/v1/device/sensors/sample-eui` | One EUI from any sensor on a profile (`?chirpstack_profile_id`) | Admin |

### Collection-service (used by UI)
| Method | Path | Description |
|---|---|---|
| GET | `/api/v1/collection/payload-tags?device_eui={eui}` | Distinct payload keys from recent readings |

---

## UI Elements Required

### 1. Measurement Type Manager (Admin)

Manage the canonical vocabulary.

**On load:** `GET /measurement-types/all`

**Create:** `POST /measurement-types` with `{ slug, display_name, description?, default_unit? }`
- Normalize slug on input: lowercase, trim, spaces → underscores

**Deprecate:** `PATCH /measurement-types/{slug}/deprecate`
- Confirm dialog — existing schema rows still reference the slug

**Notes:**
- Deprecated types shown greyed out, excluded from all labeling dropdowns
- Slug is read-only after creation

---

### 2. Payload Schema Labeling Page (Admin)

Label payload keys for a profile. Entry point for both sensor types.

**On load:**
- `GET /sensor-profiles` — populate profile dropdown
- `GET /measurement-types` — populate measurement type dropdowns

**On profile selected:**
- `GET /sensor-profile-config/{profile_id}` — check if configurable_schema is set
- If NOT configurable:
  - `GET /sensors/sample-eui?chirpstack_profile_id={profile_id}` → get one EUI
  - `GET /collection/payload-tags?device_eui={eui}` → get keys
  - Show labeling form with keys pre-filled, measurement type dropdown per key
  - **Save:** `PUT /payload-schema/{profile_id}` with all labeled rows
- If configurable:
  - Show devEUI picker (dropdown of sensors using this profile — call `GET /sensors`, filter client-side by `chirpstack_profile_id`, display by name)
  - `GET /collection/payload-tags?device_eui={selected_eui}` → get keys
  - Show labeling form
  - **Save:** `PUT /sensors/{eui}/metrics`

**Toggle configurable_schema:**
- `PUT /sensor-profile-config/{profile_id}` with `{ configurable_schema: true/false }`

**Empty state (no readings yet):**
"No readings available for this profile — deploy a sensor and wait for its first transmission."

---

## 3 and 4 is lower prio, only if we have time. This is to enabel the tenant to label their own per installation sensors, instead of having the admin do it.

### 3. Sensor List — Configuration Status (Tenant)

Show `metrics_configured` badge per sensor. Requires backend change to sensor list response.

No additional API calls — status comes from the list endpoint directly.

---

### 4. Sensor Detail — Metric Configuration (Tenant)

For configurable sensors where the tenant does the labeling themselves.

**On load:**
- `GET /sensors/{eui}/metrics` — if non-empty, show read-only view with edit button
- If empty: `GET /collection/payload-tags?device_eui={eui}`
  - If empty: "No readings yet"
  - If has keys: show labeling form

**Save:** `PUT /sensors/{eui}/metrics`

**Notes:**
- Measurement type dropdown shows `display_name`, submits `slug`
- This UI is a frontend-only addition — all backend endpoints already exist
