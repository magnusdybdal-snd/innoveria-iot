# Sensor Schema Type Design

## The Problem

When a sensor is added to the system, the admin UI needs to know how its payload keys should be
configured — at the **profile level** (shared across all sensors of that hardware type) or at the
**per-sensor level** (unique to each physical installation). Without this distinction the UI has no
way to route the admin to the correct labeling flow.

### Two sensor types

| Type | Example | Who configures | Where stored |
|------|---------|----------------|--------------|
| Fixed-schema | AC current sensor | System admin | `device.payload_schema` (per Chirpstack profile ID) |
| Configurable | Bus converter (UC300) | Tenant admin | `device.sensor_metric` (per sensor EUI) |

Fixed-schema sensors always send the same payload keys regardless of installation. Configurable
sensors are wired differently per site — the payload keys depend on what is physically connected.

---

## The Labeling Flow

### "Is this sensor configured?" check

Call `GetEffectiveMetrics` (device service). It checks:
1. Per-sensor `SensorMetric` rows first
2. Falls back to profile-level `payload_schema` rows

If it returns empty, the sensor needs configuration.

### Discovering payload keys

The collection service exposes an endpoint (`GET /collection/payload-tags?device_eui=<eui>`) that
returns the distinct keys seen in actual measurement data for that sensor. This is the source of
truth for what keys a sensor actually sends — not a spec or codec guess.

The sensor must have sent at least one reading before discovery works. For a brand new sensor type
this means deploying one sensor and waiting for a transmission before labeling.

### Labeling flows

**Fixed-schema (profile level):**
```
GET /collection/payload-tags?device_eui=<eui>           → discover keys from data
POST /device/payload-schema/{chirpstack_profile_id}/discover  → store as drafts for whole profile
PUT  /device/payload-schema/{chirpstack_profile_id}          → label all keys
```
Once labeled, every future sensor on that profile is automatically configured. This only needs to
happen once per device type.

**Configurable (per-sensor):**
```
GET /collection/payload-tags?device_eui=<eui>   → discover keys from data
PUT /device/sensors/{eui}/metrics               → label keys for this sensor only
```
Each sensor installation is configured independently.

---

## Edge Cases

**No data yet** — collection service returns empty keys. UI should tell the admin to wait for the
first reading. Caught by the payload-tags endpoint returning an empty list.

**Resetting per-sensor config to fall back to profile** — a `DELETE /device/sensors/{eui}/metrics`
endpoint is needed. Without it, once a sensor has SensorMetrics it can never fall back to the
profile schema. Not yet implemented.

**Partial labeling** — intentionally allowed. The admin can choose to skip keys they don't care
about. Skipped keys are silently ignored by `GetEffectiveMetrics` — context receives only the
labeled keys and has no knowledge that others exist. The raw payload stored in the collection
service still contains all keys (e.g. `{"bus1": 10, "bus2": 20, "bus3": 30}`), but context only
extracts the labeled ones (`bus1 → nitrogen`, `bus3 → temp`). `bus2` sits in the JSONB unused.

Skipped keys can be labeled at any time via the same PUT endpoint. The change takes effect
immediately on the next `GetEffectiveMetrics` call. Historical readings already contain the data —
it becomes visible retroactively with no backfill needed.

The UI should make unlabeled keys visually obvious so the admin consciously skips rather than
accidentally forgets. No backend count validation is needed — the API accepts whatever is submitted.

**Tenant admin configures profile sensor before admin does** — if a profile has no
`payload_schema` rows yet, a tenant admin could accidentally label it per-sensor first. Per-sensor
takes priority in `GetEffectiveMetrics`, so the profile schema would be ignored for that sensor.
Solved by ensuring profile type is registered before sensors are deployed (see solution below).

---

## The Solution: `sensor_profile_type` Table

Rather than building a full profile management UI (codecs, profile creation, etc.), a small
mapping table is sufficient:

```sql
CREATE TABLE device.sensor_profile_type (
    chirpstack_profile_id text PRIMARY KEY,
    configurable          boolean NOT NULL
);
-- configurable = true  → per-sensor flow (tenant admin labels)
-- configurable = false → profile-level flow (system admin labels)
```

### Workflow for onboarding a new device type

1. Add the profile and codec to the GitLab profile repo (existing step — syncs to Chirpstack)
2. Admin opens profile config UI → selects profile from Chirpstack dropdown → marks
   configurable/not → saves to `sensor_profile_type`
3. Sensors of that type are deployed — system knows which flow to use before any sensor arrives

If a profile has no entry in `sensor_profile_type`, the UI warns "this profile type has not been
configured" rather than guessing.

### Role separation (future)

Currently all users are super users, so both flows are available to everyone. When role-based auth
is added:

- `configurable = false` → only system admin can label (profile flow)
- `configurable = true`  → tenant admin labels (per-sensor flow); UI shows "awaiting labeling"
  for their sensors that have no SensorMetrics yet

The backend does not change when roles are introduced — only which UI elements are visible per role.

---

## UI Elements Required

**1. Measurement types page** — list all measurement type slugs, add new ones, deprecate existing.
Simple vocabulary management.

**2. Sensor profile configuration page** — lists profiles already registered in `sensor_profile_type`
with their type (profile-level / per-sensor). Add new: dropdown of all Chirpstack profile IDs
filtered to exclude already-configured ones, plus a toggle for configurable/not. This must be done
before sensors of that type are deployed.

**3. Sensor list — "needs labeling" indicator** — the existing `GET /device/sensors` query is
extended with two LEFT JOINs to check configuration status in bulk (one query, no per-sensor
follow-up calls):

```sql
SELECT
    s.sensor_id,
    s.device_eui,
    s.name,
    CASE
        WHEN sm.sensor_id IS NOT NULL THEN true
        WHEN ps.chirpstack_profile_id IS NOT NULL THEN true
        ELSE false
    END AS is_configured
FROM device.sensor s
LEFT JOIN (
    SELECT DISTINCT sensor_id FROM device.sensor_metric
) sm ON sm.sensor_id = s.sensor_id
LEFT JOIN (
    SELECT DISTINCT chirpstack_profile_id
    FROM device.payload_schema
    WHERE measurement_type IS NOT NULL
) ps ON ps.chirpstack_profile_id = s.chirpstack_profile_id
WHERE s.company_id = $1
```

`is_configured: bool` is added to the sensor list response. When false, the UI flags the sensor
for labeling.

Clicking "label" on a flagged sensor routes to the correct flow based on `sensor_profile_type`:
- `configurable = false` → profile labeling flow (clears all sensors on that profile at once)
- `configurable = true` → per-sensor labeling flow (clears only this sensor)

---

## What Is Not Needed

- "Label before first reading" — only saves time on the very first sensor of a new type. Requires
  codec knowledge from the admin. Not worth the complexity; waiting for one reading is acceptable.
- A full profile management UI with codec editing — Chirpstack + the git submodule already handles
  this. The `sensor_profile_type` table is the only metadata we need to own.
