# Architecture and Infrastructure Notes

## 1. Microservice Architecture Overview

FabrikkPuls uses a microservice architecture with the following components:

- **Frontend** — React web app, the user interface and dashboard
- **API Gateway** — Stateless orchestrator. Handles routing, JWT authentication, and saga coordination for multi-service flows. (No dedicated database?? Or run Auth in the same one).
- (**Auth / Tenant Service** — Manages users, companies, factory areas, and roles)
- **Device Service** — Abstraction layer between the API Gateway and ChirpStack (or another LoRaWAN network server). Owns the sensor and gateway registry. Handles all communication with ChirpStack, so the rest of the system has no direct dependency on ChirpStack.
- **ChirpStack** — Open-source LoRaWAN Network Server. Manages IoT device registration, tenants, and decoding. Device Service uses its API for CRUD operations on devices. ([chirpstack.io/docs](https://www.chirpstack.io/docs/))
- **Collection Service** — Subscribes to MQTT topics and stores sensor measurements in TimescaleDB. Its sole responsibility is to receive and store time-series data.
- **ERP Service** — Adapter/cache for Monitor ERP (can be extended for other ERP systems). Polls and caches orders and production resources.
- **Context Service** — Combines sensor data + ERP data to calculate business insights (e.g., nitrogen consumption per order, per hour)
- **MQTT Broker** — Message bus (e.g., Mosquitto) between ChirpStack and Collection Service

### Data Flow

```
Frontend → API Gateway (REST)
API Gateway → Each service (internal REST)
API Gateway → Device Service → ChirpStack (for sensor/gateway CRUD)
LoRaWAN Gateway → ChirpStack → MQTT Broker → Collection Service
Monitor ERP ← ERP Service (polling)
Collection Service + ERP Service → Context Service (reads from both)
```

### Cross-Service Queries

When the frontend needs data spanning multiple services (e.g., "Sensor X on Machine Y with Order Z context"), the API Gateway calls each service and aggregates the responses. This is the API Composition pattern. ([microservices.io — Database per Service](https://microservices.io/patterns/data/database-per-service.html))

### Why Device Service as a Separate Layer in Front of ChirpStack

The API Gateway never calls ChirpStack directly. All ChirpStack communication goes through Device Service. Benefits:

- **Replaceability** — If we switch from ChirpStack to another LoRaWAN server, we only change Device Service. The rest of the system is unaffected.
- **Natural home for sensor/gateway metadata** — Device Service owns the business context around sensors (which machine a sensor monitors, measurement type, name) that ChirpStack doesn't know about.
- **Cleaner separation of concerns** — Collection Service doesn't need to handle sensor registration and can focus solely on time-series data.

---

## 2. Database Strategy: Database-per-Service

Each microservice owns its own database. No shared databases. This ensures loose coupling — changes to one service's database don't affect others, and each service can use the database technology best suited to its needs. ([AWS — Database-per-service pattern](https://docs.aws.amazon.com/prescriptive-guidance/latest/modernization-data-persistence/database-per-service.html))

### Our Approach

**Separate database instances** (one container per database) + **TimescaleDB as a separate instance** (requires the timescale extension). Six database containers in total:

| Container | Technology | Database | What it stores |
|-----------|-----------|----------|----------------|
| auth-db | postgres:16 | auth_db | Users, companies, factory areas, roles |
| device-db | postgres:16 | device_db | Sensor registry, gateway registry, chirpstackId mapping |
| erp-db | postgres:16 | erp_db | Cached orders, production resources, sync status |
| context-db | postgres:16 | context_db | Computed context data |
| collection-db | timescale/timescaledb:latest-pg16 | collection_db | Sensor measurements (hypertable) |
| chirpstack-db | postgres:16 | chirpstack | LoRaWAN devices, tenants (managed by ChirpStack) |

### What Each Database Stores

**Auth DB (PostgreSQL)**
- User (users, roles, passwordHash)
- Company (companies, chirpstackTenantId, ERP API key)
- Factory Area (factory areas)
- Role enum (FactoryWorker, FactorySuperUser, Admin)

**Device DB (PostgreSQL)**
- Sensor — business metadata: sensorId, deviceEUI, productionResourceId (link to machine), measurementType, name, companyId
- Gateway — gatewayId, gatewayEUI, name, companyId
- ChirpStack ID mappings (internal sensorId ↔ ChirpStack device ID)

**Collection DB (TimescaleDB)**
- SensorMeasurement — time-series data from sensors (hypertable, partitioned by timestamp)
- No sensor or gateway tables — this is now owned by Device Service

**ERP DB (PostgreSQL)**
- ProductionResource — cached machines from ERP (ERPId, name, status: ACTIVE/INACTIVE)
- Order — cached orders (orderNumber, productName, quantity, startTime, endTime, operations)
- SyncMetadata — fetchedDate, syncStatus

**Context DB (PostgreSQL)**
- ContextData (contextType, orderId, value, unit, calculatedDate)
- (Dashboard configuration (per user/company))

**ChirpStack DB (managed by ChirpStack)**
- LoRaWAN devices (EUI, keys, profiles)
- Tenants (corresponds to our Company)

### Multitenancy

Every table has `companyId`. All queries filter on this. ChirpStack tenants map directly to our Company concept — each company gets a ChirpStack tenant, and we store `chirpstackTenantId` on the Company entity. ([ChirpStack — Tenants documentation](https://www.chirpstack.io/docs/chirpstack/use/tenants.html))

`companyId` (our own) is used everywhere internally — in auth-db, device-db, erp-db, context-db, collection-db. All tables filter on this. It is the only ID our services know about.

`chirpstackTenantId` is only used when Device Service talks to the ChirpStack API. E.g., "create this sensor under tenant X" — it looks up `chirpstackTenantId` from the Company table and sends it to ChirpStack.

If ChirpStack is replaced, it is essential that we have our own internal ID and don't rely on ChirpStack attributes for our own business logic.

### Collection Service and Tenant Mapping

Collection Service doesn't need to know anything about sensors or gateways. It only needs to map `chirpstackTenantId → companyId` to tag incoming sensor data with the correct company. ChirpStack includes `tenantId` in MQTT messages.

This mapping is fetched from Auth Service at startup and cached in memory. The number of companies is small (5–50) and changes infrequently.

### Sensor Metadata: ChirpStack vs. Device Service

- **ChirpStack** is the source of truth for LoRaWAN configuration (device profiles, network keys, device status)
- **Device Service** owns the business context (which machine a sensor monitors, what it measures, business-level configuration)
- `deviceEUI` is the shared key between them

### Soft Deletion for ERP Data

If a machine is deleted from the ERP system, we mark it as `status: INACTIVE` — never hard-delete. Historical context data still references it. Sensors linked to it show a warning in the UI.

### Saga Pattern for Cross-Service Operations

Multi-service flows (like "add sensor") are coordinated using the saga pattern. If a step fails, compensating transactions undo the previous steps. ([microservices.io — Saga pattern](https://microservices.io/patterns/data/saga.html), [Microsoft Azure — Saga design pattern](https://learn.microsoft.com/en-us/azure/architecture/patterns/saga))

With Device Service, the saga for sensor creation is now internal to Device Service (not in the API Gateway), which is cleaner — the API Gateway doesn't need to know anything about ChirpStack.

---

## 3. Key Flows

### Adding a Gateway

```
Frontend: POST /api/gateways { gatewayEUI, name }
  → API Gateway
    → Device Service
      1. Call ChirpStack API: register gateway under the company's tenant
      2. Save to device-db: gatewayEUI, name, companyId
      3. Return success
      (If step 2 fails → compensate by deleting from ChirpStack)
```

### Adding a Sensor

```
Step 1: Frontend requests machine list
  → API Gateway: GET /erp/production-resources?companyId=X
  → Frontend displays dropdown with machines

Step 2: User fills in sensor details + selects machine
  → Frontend: POST /api/sensors { deviceEUI, name, measurementType, productionResourceId }
  → API Gateway
    → Device Service
      1. Call ChirpStack API: register device under the company's tenant
      2. Save to device-db: deviceEUI, name, measurementType,
         productionResourceId, companyId
      3. Return success
      (If step 2 fails → compensate by deleting from ChirpStack)
```

The machine list comes from ERP Service, but the sensor ↔ machine link is stored in Device Service. ERP Service doesn't know about sensors.

### Sensor Data Pipeline (MQTT)

```
Physical sensor → LoRaWAN Gateway → ChirpStack (decodes data)
  → MQTT Broker (ChirpStack publishes decoded data, incl. tenantId)
  → Collection Service (subscribes, maps tenantId → companyId,
     writes to TimescaleDB)
```

---

## 4. TimescaleDB

TimescaleDB is a PostgreSQL extension for time-series data. It is not a separate database — it is regular Postgres with additional functionality. ([TimescaleDB documentation](https://docs.timescale.com/use-timescale/latest/compression/))

### Setup

```sql
-- Step 1: Create a regular table with all columns
CREATE TABLE sensor_measurement (
    measurement_id UUID DEFAULT gen_random_uuid(),
    device_eui VARCHAR NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    unit VARCHAR NOT NULL,
    company_id UUID NOT NULL
);

-- Step 2: Convert to hypertable (partitioned by time)
SELECT create_hypertable('sensor_measurement', 'timestamp');
```

The `timestamp` parameter tells TimescaleDB "use this column as the partition key." All columns are still there, fully queryable. Under the hood, data is divided into time-based chunks. When you query `WHERE timestamp > NOW() - INTERVAL '7 days'`, only the relevant chunk is scanned.

### Compression

TimescaleDB has built-in columnar compression that can reduce storage by 90%+. ([TimescaleDB — About compression](https://docs.timescale.com/use-timescale/latest/compression/about-compression/))

```sql
-- Enable compression
ALTER TABLE sensor_measurement SET (
    timescaledb.compress,
    timescaledb.compress_segmentby = 'device_eui, company_id',
    timescaledb.compress_orderby = 'timestamp DESC'
);

-- Auto-compress data older than 7 days
SELECT add_compression_policy('sensor_measurement', INTERVAL '7 days');
```

### Retention and Continuous Aggregates

For long-term storage: keep raw data for a few months, pre-compute hourly/daily aggregates, then drop the raw data.

```sql
-- Continuous aggregate: hourly averages
CREATE MATERIALIZED VIEW sensor_measurement_hourly
WITH (timescaledb.continuous) AS
SELECT
    device_eui,
    company_id,
    time_bucket('1 hour', timestamp) AS bucket,
    AVG(value) AS avg_value,
    MIN(value) AS min_value,
    MAX(value) AS max_value,
    COUNT(*) AS count
FROM sensor_measurement
GROUP BY device_eui, company_id, bucket;

-- Optional: drop raw data older than 1 year
SELECT add_retention_policy('sensor_measurement', INTERVAL '1 year');
```

### Storage Estimates

20 sensors × 2 measurements/min × 24h = ~57,600 rows/day → ~2–3 GB/year per factory (uncompressed). With compression (90%+) this becomes ~200–300 MB/year.

---

## 5. Hosting and Infrastructure

### Docker Compose on a Single Server

Everything runs in Docker containers on a single server, orchestrated by a single `docker-compose.yml`:

```yaml
services:
  # --- Frontend ---
  frontend:
    build: ./frontend
    ports: ["3000:3000"]

  # --- API Gateway ---
  api-gateway:
    build: ./api-gateway
    ports: ["8080:8080"]

  # --- Application Services ---
  device-service:
    build: ./device-service
  collection-service:
    build: ./collection-service
  erp-service:
    build: ./erp-service
  context-service:
    build: ./context-service

  # --- ChirpStack ---
  chirpstack:
    image: chirpstack/chirpstack:4
    ports: ["8090:8080"]
    depends_on:
      - chirpstack-db
      - chirpstack-redis
      - mosquitto
  chirpstack-db:
    image: postgres:16
    volumes:
      - chirpstack-db-data:/var/lib/postgresql/data
  chirpstack-redis:
    image: redis:7-alpine
    volumes:
      - chirpstack-redis-data:/data

  # --- MQTT ---
  mosquitto:
    image: eclipse-mosquitto:2

  # --- Databases ---
  auth-db:
    image: postgres:16
    volumes:
      - auth-db-data:/var/lib/postgresql/data
  device-db:
    image: postgres:16
    volumes:
      - device-db-data:/var/lib/postgresql/data
  erp-db:
    image: postgres:16
    volumes:
      - erp-db-data:/var/lib/postgresql/data
  context-db:
    image: postgres:16
    volumes:
      - context-db-data:/var/lib/postgresql/data
  collection-db:
    image: timescale/timescaledb:latest-pg16
    volumes:
      - collection-db-data:/var/lib/postgresql/data

volumes:
  chirpstack-db-data:
  chirpstack-redis-data:
  auth-db-data:
  device-db-data:
  erp-db-data:
  context-db-data:
  collection-db-data:
```

14 containers in total: 6 application services (incl. frontend and API Gateway), ChirpStack + its Postgres + Redis, Mosquitto, and 5 databases (auth, device, erp, context, collection).

Named volumes ensure data survives container restarts.

### Server Requirements

Minimum 8 GB RAM (ideally 16 GB). A VPS like Hetzner (~€7–10/month) is recommended, since we need a public IP for the LoRaWAN gateway to reach ChirpStack and for the web frontend to be accessible.

---

## 6. Do We Need Kubernetes for This Project?

### The Short Version

We have ~3 months, no Kubernetes experience, and Docker Compose gives us everything we need at our scale. Kubernetes would cost time on learning and debugging that takes away from development.

### Kubernetes and Databases

Kubernetes is designed for stateless services. Running databases in Kubernetes introduces specific challenges:

- **Risk of data loss** — If persistent volumes aren't configured correctly and a pod is rescheduled, it can start with an empty disk. ([Google Cloud — To run or not to run a database on Kubernetes](https://cloud.google.com/blog/products/databases/to-run-or-not-to-run-a-database-on-kubernetes-what-to-consider))
- **Corruption on restart** — Postgres needs a graceful shutdown; Kubernetes' time-based termination policies can force-kill a database pod mid-write. ([CockroachDB — Kubernetes: The state of stateful apps](https://www.cockroachlabs.com/blog/kubernetes-state-of-stateful-apps/))
- **Storage provisioning complexity** — PersistentVolumes, PersistentVolumeClaims, StorageClasses — concepts to learn that have nothing to do with our product
- **Networking** — Pod IPs change on restart; databases need stable connections via StatefulSets and headless services
- **Backups** — Kubernetes doesn't back up persistent volumes; you have to set up pg_dump cron jobs or volume snapshots yourself

With Docker Compose, data lives in a named volume on disk. Simple, predictable, hard to mess up.

### What We Don't Lose

Our architecture is already microservice-based with proper separation. Kubernetes is just an operational deployment choice — it doesn't make the architecture better. The same containers can be deployed to Kubernetes in the future if needed.
