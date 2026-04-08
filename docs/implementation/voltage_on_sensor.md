# Voltage on Sensor

Electricity sensors (current clamps etc.) need a voltage value stored per sensor so
context-service can calculate watt hours using `P = V × I`.

Voltage is a property of the electrical installation, not the hardware model — two identical
current sensors may be installed on 230V and 400V systems respectively. It therefore lives
on the sensor, not the profile.

---

## Data Model

Add a nullable `voltage` integer column to the sensor table:

```sql
ALTER TABLE device.sensor ADD COLUMN voltage integer;
```

```
null   → not an electricity sensor (default)
230    → single phase
400    → three phase
```

---

## Backend Changes

- Migration: add `voltage integer` (nullable) to `device.sensor`
- Domain: add `Voltage *int` to `Sensor` struct
- Repository: include `voltage` in select and insert/update queries
- DTO: add `Voltage *int` to sensor request and response DTOs
- Handler: no logic change — just pass the value through

---

## Frontend — Add/Edit Sensor Form

Add the following to the sensor registration and edit forms:

**Checkbox:** "Electricity sensor"
- Unchecked by default
- Pre-checked on edit if `voltage` is non-null

**When checked:** a dropdown appears:
```
Voltage
[ 230V ]
[ 400V ]
```

**When unchecked:** dropdown is hidden, voltage is submitted as null.

The checkbox is a UI-only concept — the backend only sees `voltage: 230 | 400 | null`.

---

## Usage in Context-service

When calculating energy consumption for a sensor with a non-null voltage:

```
P (watts)     = voltage × current
Wh            = P × hours in period
```

Note: this assumes power factor = 1 (cos φ). Accurate enough for a demo. Revisit
when precision matters — real power factor typically ranges 0.8–1 for industrial loads.
