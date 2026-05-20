# Stress Test Results — YYYY-MM-DD

## Environment

| Property | Value |
|---|---|
| Date | YYYY-MM-DD |
| Branch | `git rev-parse --short HEAD` |
| Docker Compose services | `docker compose ps` output |
| Host machine | e.g. MacBook Pro M2, 16 GB RAM |

---

## Test 1 — Sensor Throughput (Req 2: 600 measurements/hour)

### Scenario A — Baseline (600/hour)
```
Devices=10  interval=60
```
| Metric | Value |
|---|---|
| Messages published | |
| Messages received by collection-service | |
| Messages dropped | |
| collection-service CPU (avg) | |
| collection-service RAM (avg) | |
| collection-db CPU (avg) | |
| Errors / warnings in logs | |

### Scenario B — 5× load (3 000/hour)
```
Devices=50  interval=60
```
| Metric | Value |
|---|---|
| Messages published | |
| Messages received | |
| Messages dropped | |
| collection-service CPU (avg) | |
| collection-service RAM (avg) | |
| Errors / warnings in logs | |

### Scenario C — Stress ceiling (36 000/hour)
```
Devices=60  interval=6
```
| Metric | Value |
|---|---|
| Messages published | |
| Messages received | |
| Messages dropped | |
| collection-service CPU (avg) | |
| collection-service RAM (avg) | |
| Errors / warnings in logs | |

**Passes Req 2?** YES / NO  
**Breaking point:** (first scenario where messages were dropped)

---

## Test 2 — HTTP Response Time, Low Load (Req 3: p95 < 1s)

```
k6 run --scenario baseline ...
```

| Metric | Value |
|---|---|
| VUs | 10 |
| Duration | 3 min |
| Total requests | |
| p50 response time | |
| p95 response time | |
| p99 response time | |
| Error rate | |
| Threshold breached | YES / NO |

Paste k6 summary output here:
```
<k6 output>
```

**Passes Req 3?** YES / NO

---

## Test 3 — 300 Concurrent Users (Req 1 + Req 3)

```
k6 run --scenario concurrent_300 ...
```

| Metric | Value |
|---|---|
| Peak VUs | 300 |
| Duration | 8 min (2 ramp-up + 5 hold + 1 ramp-down) |
| Total requests | |
| p50 response time | |
| p95 response time | |
| p99 response time | |
| Max response time | |
| Error rate | |
| http_req_failed rate | |
| Threshold p95 < 1000ms | PASS / FAIL |
| api-gateway CPU (peak) | |
| api-gateway RAM (peak) | |

Paste k6 summary output here:
```
<k6 output>
```

**Passes Req 1?** YES / NO  
**Passes Req 3 under load?** YES / NO

---

## Test 4 — Combined Load (Sensors + Users)

Run sensor simulator at 600/hour simultaneously with concurrent_300.

| Metric | Value |
|---|---|
| Sensor messages/hour | 600 |
| Peak HTTP VUs | 300 |
| HTTP p95 response time | |
| Sensor messages dropped | |
| Error rate | |

**Notes / observations:**

---

## Summary

| Requirement | Result | Notes |
|---|---|---|
| Req 1: 300 concurrent users | PASS / FAIL | |
| Req 2: 600 measurements/hour | PASS / FAIL | |
| Req 3: Response < 1s (p95) | PASS / FAIL | |

## Issues found

- 

## Recommendations

- 
