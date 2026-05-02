# Stress Testing

Tests the three non-functional requirements:

| # | Requirement | Tool |
|---|---|---|
| Req 1 | 300 concurrent users | k6 |
| Req 2 | 600 sensor measurements/hour | collection-simulator |
| Req 3 | Backend response time < 1s (p95) | k6 |

---

## Prerequisites

```bash
# k6 must be installed
k6 version

# Docker Compose stack must be running
docker compose up -d
```

You need a valid user account in the running system.

---

## Test 1 — Sensor Throughput (Req 2)

Uses the `collection-simulator` with varying device counts and intervals.

**Formula:** `messages/hour = Devices × (3600 / interval)`

| Scenario | Command | msg/hour | Target |
|---|---|---|---|
| Baseline | `Devices=10 interval=60` | 600 | Req 2 ✓ |
| 5× load | `Devices=50 interval=60` | 3 000 | Headroom |
| Stress | `Devices=60 interval=6` | 36 000 | Breaking point |

**Run:**
```bash
# Start the collection stack first
docker compose up -d mosquitto collection-db collection-service

# Baseline (600/hour)
docker compose run --rm -e Devices=10 -e interval=60 collection-simulator

# 5× load (3 000/hour)
docker compose run --rm -e Devices=50 -e interval=60 collection-simulator

# Stress ceiling (36 000/hour)
docker compose run --rm -e Devices=60 -e interval=6 collection-simulator
```

**Monitor while running:**
```bash
# Watch for dropped messages or errors
docker compose logs -f collection-service | grep -iE "warn|drop|full|error"

# Resource usage
docker stats collection-service collection-db mosquitto --no-stream
```

---

## Test 2 — HTTP Response Time, Low Load (Req 3)

Runs 10 virtual users for 3 minutes against the key frontend endpoints.  
Validates that response time stays under 1s before adding concurrency stress.

**Run:**
```bash
k6 run \
  --scenario baseline \
  --env K6_USERNAME=your@email.com \
  --env K6_PASSWORD=yourpassword \
  --env BASE_URL=http://localhost:8081 \
  docs/tests/k6/http-load-test.js
```

**Pass criteria:** `p(95) < 1000ms`

---

## Test 3 — 300 Concurrent Users (Req 1 + Req 3)

Ramps from 0 → 300 virtual users over 2 minutes, holds for 5 minutes, then ramps down.  
Each VU simulates a logged-in user hitting the sensor dashboard, production view, and user profile.

**Run:**
```bash
k6 run \
  --scenario concurrent_300 \
  --env K6_USERNAME=your@email.com \
  --env K6_PASSWORD=yourpassword \
  --env BASE_URL=http://localhost:8081 \
  docs/tests/k6/http-load-test.js
```

**Pass criteria:**
- `p(95) < 1000ms`
- `http_req_failed < 1%`

---

## Test 4 — Combined Load (Req 1 + 2 + 3 together)

The realistic production scenario: sensor data flowing in while users are active.

**Run in two terminals:**

Terminal 1 — sensor load:
```bash
docker compose run --rm -e Devices=10 -e interval=60 collection-simulator
```

Terminal 2 — HTTP load:
```bash
k6 run \
  --scenario concurrent_300 \
  --env K6_USERNAME=your@email.com \
  --env K6_PASSWORD=yourpassword \
  --env BASE_URL=http://localhost:8081 \
  docs/tests/k6/http-load-test.js
```

---

## Recording Results

Copy `docs/tests/results/TEMPLATE.md` to a new file named by date:

```bash
cp docs/tests/results/TEMPLATE.md docs/tests/results/2026-05-02-run-01.md
```

Fill in the values from k6 output and `docker stats` for each scenario.

---

## Endpoints Under Test

| Group | Endpoint | Why |
|---|---|---|
| user_info | `GET /api/v1/auth/me` | Every page load |
| sensor_dashboard | `GET /api/v1/device/sensors` | Main dashboard |
| sensor_dashboard | `GET /api/v1/collection/latest?device_eui=X` | Live sensor readings |
| production_view | `GET /api/v1/erp/production-resources` | Production page |
| production_view | `GET /api/v1/erp/orders` | Orders overview |
