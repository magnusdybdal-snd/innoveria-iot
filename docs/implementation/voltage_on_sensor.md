# Voltage on Sensor

Electricity sensors (current clamps etc.) need a voltage value stored per sensor so
context-service can calculate watt hours using `P = V × I`.

Voltage is a property of the electrical installation, not the hardware model — two identical
current sensors may be installed on 230V and 400V systems respectively. It therefore lives
on the sensor, not the profile.

---

## Data Model

Two columns are added to the sensor table:

```sql
ALTER TABLE device.sensor ADD COLUMN electricity_sensor boolean NOT NULL DEFAULT false;
ALTER TABLE device.sensor ADD COLUMN voltage integer;
```

`electricity_sensor` is the source of truth for whether a sensor is electrical.
`voltage` is only meaningful when `electricity_sensor` is true.

```
electricity_sensor = false  → not an electricity sensor (default)
electricity_sensor = true, voltage = 230  → single phase
electricity_sensor = true, voltage = 400  → three phase
```

The combination `electricity_sensor = false` with a non-null `voltage` is an invalid state
and cannot be reached through the API.

---

## Backend Changes

- Migration: add `electricity_sensor boolean NOT NULL DEFAULT false` and `voltage integer` (nullable) to `device.sensor`
- Domain: add `ElectricitySensor *bool` and `Voltage *int` to `Sensor` struct (`*bool` preserves nil = "not provided" in partial update payloads)
- Repository: include both columns in all select, insert, and update queries
- DTO:
  - `SensorResponse`: `electricity_sensor bool`, `voltage *int`
  - `CreateSensorRequest`: `electricity_sensor bool`, `voltage *int`
  - `UpdateSensorRequest`: `electricity_sensor *bool`, `voltage *int`
- Handler: validation guards on both create and update
- Service: merge logic enforces the following rules

---

## Update Rules

| Payload | Effect |
|---|---|
| `{"electricity_sensor": true, "voltage": 230\|400}` | Mark as electrical, set voltage |
| `{"electricity_sensor": false}` | Unmark as electrical, voltage is cleared automatically |
| `{"voltage": 230\|400}` | Change voltage — only valid if sensor is already electrical |
| `{"electricity_sensor": true}` without voltage | **400 Bad Request** |
| `{"voltage": ...}` on a non-electricity sensor | **Error** |

---

## Frontend — Add/Edit Sensor Form

Add the following to the sensor registration and edit forms:

**Checkbox:** "Electricity sensor"
- Unchecked by default
- Pre-checked on edit if `electricity_sensor` is true

**When checked:** a dropdown appears:
```
Voltage
[ 230V ]
[ 400V ]
```

**When unchecked:** dropdown is hidden, `electricity_sensor: false` is submitted (voltage is cleared server-side).

Note: the backend now has an explicit `electricity_sensor` boolean — the frontend should send it directly rather than inferring from `voltage`.

---

## Usage in Context-service

When calculating energy consumption for a sensor where `electricity_sensor` is true:

```
P (watts)     = voltage × current
Wh            = P × hours in period
```

Note: this assumes power factor = 1 (cos φ). Accurate enough for a demo. Revisit
when precision matters — real power factor typically ranges 0.8–1 for industrial loads.
