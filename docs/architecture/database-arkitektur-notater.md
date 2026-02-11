# Arkitektur og infrastrukturnotater

## 1. Oversikt over mikrotjenestearkitektur

Fabrikkpuls bruker en mikrotjenestearkitektur med følgende komponenter:

- **Frontend** — React-webapp, brukergrensesnittet og dashbordet
- **API Gateway** — Tilstandsløs orkestrator. Håndterer ruting, JWT-autentisering og saga-koordinering for flertjenesteflyter (f.eks. "legg til sensor"). (Ingen egen database?? Eller kjøre Auth i samme).
- (**Auth / Tenant Service** — Administrerer brukere, selskaper, fabrikkområder og roller)
- **ChirpStack** — Open-source LoRaWAN Network Server. Administrerer IoT-enhetsregistrering, tenanter og dekoding. Vi bruker API-et for CRUD-operasjoner på enheter. ([chirpstack.io/docs](https://www.chirpstack.io/docs/))
- **Collection Service** — Abonnerer på MQTT-topics, lagrer sensormålinger i TimescaleDB, og vedlikeholder et lettvekts sensor-/gateway-register med forretningsmetadata
- **ERP Service** — Adapter/cache for Monitor ERP (kan utvides for andre ERP-systemer). Poller og cacher ordrer og produksjonsressurser
- **Context Service** — Kombinerer sensordata + ERP-data for å beregne forretningsinnsikt (f.eks. nitrogenforbruk per ordre, per time)
- **MQTT Broker** — Meldingsbuss (f.eks. Mosquitto) mellom ChirpStack og Collection Service

### Dataflyt

```
Frontend → API Gateway (REST)
API Gateway → Hver tjeneste (intern REST)
LoRaWAN Gateway → ChirpStack → MQTT Broker → Collection Service
Monitor ERP → ERP Service (polling)
Collection Service + ERP Service → Context Service (leser fra begge)
```

### Spørringer på tvers av tjenester

Når frontenden trenger data som spenner over flere tjenester (f.eks. "Sensor X på Maskin Y med Ordre Z-kontekst"), kaller API Gateway hver tjeneste og aggregerer svarene. Dette er API Composition-mønsteret. ([microservices.io — Database per Service](https://microservices.io/patterns/data/database-per-service.html))

---

## 2. Databasestrategi: Database-per-tjeneste

Hver mikrotjeneste eier sin egen database. Ingen delte databaser. Dette sikrer løs kobling — endringer i én tjenestes database påvirker ikke andre, og hver tjeneste kan bruke databaseteknologien som passer best til sine behov. ([AWS — Database-per-service pattern](https://docs.aws.amazon.com/prescriptive-guidance/latest/modernization-data-persistence/database-per-service.html))

### Vår tilnærming

**Alternativ 1** (separate databaseinstanser) + **TimescaleDB som separat instans**
(trenger timescale-utvidelsen). Totalt fem databasecontainere for våre tjenester:

| Container | Teknologi | Database | Hva som lagres |
|-----------|-----------|----------|----------------|
| auth-db | postgres:16 | auth_db | Brukere, selskaper, fabrikkområder, roller |
| erp-db | postgres:16 | erp_db | Cachede ordrer, produksjonsressurser, synkstatus |
| context-db | postgres:16 | context_db | Beregnet kontekstdata, dashbordkonfig |
| collection-db | timescale/timescaledb:latest-pg16 | collection_db | Sensormåling (hypertable), sensorregister, gatewayregister |
| chirpstack-db | postgres:16 | chirpstack | LoRaWAN-enheter, tenanter (administrert av ChirpStack) |

### Hva hver database lagrer

**Auth DB (PostgreSQL)**
- Bruker (brukere, roller, passordHash)
- Selskap (selskap, chirpstackTenantId, ERP API-nøkkel)
- Fabrikkområde (fabrikkområder)
- Rolle-enum (Fabrikkarbeider, Fabrikk-Superbruker, Admin)

**Collection DB (TimescaleDB)**
- Sensormåling — tidsseriedata fra sensorer (hypertable, partisjonert etter tidspunkt)
- Sensor — forretningsmetadata: deviceEUI, produksjonsressursId (kobling til maskin), måletype, navn, selskapId
- Gateway — gatewayEUI, navn, selskapId

**ERP DB (PostgreSQL)**
- Produksjonsressurs — cachede maskiner fra ERP (ERPId, navn, status: AKTIV/INAKTIV)
- Ordre — cachede ordrer (ordrenummer, produktnavn, antall, startTid, sluttTid, operasjoner)
- SyncMetadata — hentetDato, synkroniseringsstatus

**Context DB (PostgreSQL)**
- Kontekstdata (kontekstType, ordreId, verdi, enhet, beregnetDato)
- (Dashbordkonfigurasjon (per bruker/selskap))

**ChirpStack DB (administrert av ChirpStack)**
- LoRaWAN-enheter (EUI, nøkler, profiler)
- Tenanter (tilsvarer vårt Selskap)

### Multitenancy

Hver tabell har `selskapId`. Alle spørringer filtrerer på dette. ChirpStack-tenanter tilsvarer direkte vårt Selskap-konsept — hvert selskap får en ChirpStack-tenant, og vi lagrer `chirpstackTenantId` på Selskap-entiteten. ([ChirpStack — Tenants-dokumentasjon](https://www.chirpstack.io/docs/chirpstack/use/tenants.html))

selskapId (vår egen) brukes overalt internt — i auth-db, erp-db, context-db, collection-db. Alle tabeller filtrerer på denne. Det er den eneste ID-en tjenestene våre kjenner til.

chirpstackTenantId brukes kun når API Gateway snakker med ChirpStack API. F.eks. "opprett denne sensoren under tenant X" — da slår den opp chirpstackTenantId fra Selskap-tabellen og sender den til ChirpStack.

Hvis Chirpstack byttes ut er det essensielt at vi har egen intern ID og ikke belager oss på attributter fra Chirpstack for egen buisness logikk

### Sensormetadata: ChirpStack vs. Collection Service

- **ChirpStack** er kilden til sannhet for LoRaWAN-konfigurasjon (enhetsprofiler, nettverksnøkler, enhetsstatus)
- **Collection Service** eier forretningskonteksten (hvilken maskin en sensor overvåker, hva den måler, forretningsnivå-konfigurasjon)
- `deviceEUI` er den delte nøkkelen mellom dem
- Vi beholder en lettvekts lokal kopi i Collection Service slik at datapipelinen vår ikke har en hard kjøretidsavhengighet til ChirpStack

### Myk sletting for ERP-data

Hvis en maskin slettes fra ERP-systemet, markerer vi den som `status: INAKTIV` — aldri hard-delete. Historisk kontekstdata refererer fortsatt til den. Sensorer koblet til den viser en advarsel i brukergrensesnittet.

### Saga-mønsteret for operasjoner på tvers av tjenester

Flertjenesteflyter (som "legg til sensor") koordineres av API Gateway ved hjelp av saga-mønsteret. Hvis et steg feiler, angrer kompenserende transaksjoner de foregående stegene. ([microservices.io — Saga pattern](https://microservices.io/patterns/data/saga.html), [Microsoft Azure — Saga design pattern](https://learn.microsoft.com/en-us/azure/architecture/patterns/saga)).
Dette må implementeres med en form for try/catch hele veien slik at man kan gå tilbake på operasjoner om noe i rekken feiler.

---

## 3. Nøkkelflyter

### Legge til gateway

```
Frontend: POST /api/gateways { gatewayEUI, navn }
  → API Gateway
    1. ChirpStack API: POST gateway (registrer under selskapets tenant)
    2. Collection Service: POST /gateways (lagre gatewayEUI, navn, selskapId)
    3. Returner suksess
    (Hvis steg 2 feiler → kompenser ved å slette fra ChirpStack)
```

### Legge til sensor

```
Steg 1: Frontend ber om maskinliste
  → API Gateway: GET /erp/produksjonsressurser?selskapId=X
  → Frontend viser nedtrekksmeny med maskiner

Steg 2: Bruker fyller inn sensordetaljer + velger maskin
  → Frontend: POST /api/sensors { deviceEUI, navn, måletype, produksjonsressursId }
  → API Gateway
    1. ChirpStack API: POST device (registrer under selskapets tenant)
    2. Collection Service: POST /sensors
       { deviceEUI, navn, måletype, produksjonsressursId, selskapId }
    3. Returner suksess
    (Hvis steg 2 feiler → kompenser ved å slette fra ChirpStack)
```

Maskinlisten kommer fra ERP Service, men koblingen sensor ↔ maskin lagres i Collection Service. ERP Service vet ikke om sensorer.

### Sensordata-pipeline (MQTT)

```
Fysisk sensor → LoRaWAN Gateway → ChirpStack (dekoder data)
  → MQTT Broker (ChirpStack publiserer dekodet data)
  → Collection Service (abonnerer, skriver til TimescaleDB)
```

---

## 4. TimescaleDB

TimescaleDB er en PostgreSQL-utvidelse for tidsseriedata. Det er ikke en separat database — det er vanlig Postgres med ekstra funksjonalitet. ([TimescaleDB-dokumentasjon](https://docs.timescale.com/use-timescale/latest/compression/))

### Oppsett

```sql
-- Steg 1: Opprett en vanlig tabell med alle kolonner
CREATE TABLE sensormaaling (
    maaling_id UUID DEFAULT gen_random_uuid(),
    device_eui VARCHAR NOT NULL,
    tidspunkt TIMESTAMPTZ NOT NULL,
    verdi DOUBLE PRECISION NOT NULL,
    enhet VARCHAR NOT NULL,
    selskaps_id UUID NOT NULL
);

-- Steg 2: Konverter til hypertable (partisjonert etter tid)
SELECT create_hypertable('sensormaaling', 'tidspunkt');
```

`tidspunkt`-parameteren forteller TimescaleDB "bruk denne kolonnen som partisjonsnøkkel." Alle kolonner er fortsatt der, fullt spørrbare. Under panseret deles data inn i tidsbaserte biter (chunks). Når du spør `WHERE tidspunkt > NOW() - INTERVAL '7 days'`, skannes kun den relevante biten.

Sensor- og Gateway-tabellene er vanlige PostgreSQL-tabeller — ingen spesiell behandling nødvendig.

### Komprimering

TimescaleDB har innebygd kolonnebasert komprimering som kan redusere lagring med 90%+. ([TimescaleDB — Om komprimering](https://docs.timescale.com/use-timescale/latest/compression/about-compression/))

```sql
-- Aktiver komprimering
ALTER TABLE sensormaaling SET (
    timescaledb.compress,
    timescaledb.compress_segmentby = 'device_eui, selskaps_id',
    timescaledb.compress_orderby = 'tidspunkt DESC'
);

-- Auto-komprimer data eldre enn 7 dager
SELECT add_compression_policy('sensormaaling', INTERVAL '7 days');
```

### Oppbevaring og kontinuerlige aggregater

For langtidslagring: behold rådata i noen måneder, forhåndsberegn time-/dagaggregater, slett deretter rådataen.

```sql
-- Kontinuerlig aggregat: timegjennomsnitt
CREATE MATERIALIZED VIEW sensormaaling_hourly
WITH (timescaledb.continuous) AS
SELECT
    device_eui,
    selskaps_id,
    time_bucket('1 hour', tidspunkt) AS bucket,
    AVG(verdi) AS avg_verdi,
    MIN(verdi) AS min_verdi,
    MAX(verdi) AS max_verdi,
    COUNT(*) AS antall
FROM sensormaaling
GROUP BY device_eui, selskaps_id, bucket;

-- Valgfritt: slett rådata eldre enn 1 år
SELECT add_retention_policy('sensormaaling', INTERVAL '1 year');
```

### Lagringsestimater

20 sensorer × 2 målinger/min × 24t = ~57 600 rader/dag → ~2–3 GB/år per fabrikk (ukomprimert). Med komprimering (90%+) blir dette ~200–300 MB/år.

---

## 5. Hosting og infrastruktur

### Docker Compose på én server

Alt kjører i Docker-containere på én server, orkestrert av en enkelt `docker-compose.yml`:

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
  erp-db-data:
  context-db-data:
  collection-db-data:
```

Navngitte volumer sikrer at data overlever omstart av containere.

### Serverkrav

Minimum 8 GB RAM (ideelt 16 GB). En VPS som Hetzner (~€7–10/mnd) anbefales, siden vi trenger en offentlig IP for at LoRaWAN-gatewayen skal nå ChirpStack og for at web-frontenden skal være tilgjengelig.

---

## 6. Trenger vi Kubernetes for dette prosjektet?

### Kortversjonen

Vi har ~3 måneder, ingen Kubernetes-erfaring, og Docker Compose gir oss alt vi trenger i vår skala. Kubernetes bil koste tid med læring og feilsøking som tar vekk fra utvikling.

### Kubernetes og databaser

Kubernetes er designet for tilstandsløse tjenester. Å kjøre databaser i Kubernetes introduserer spesifikke utfordringer:

- **Risiko for datatap** — Hvis persistente volumer ikke er konfigurert riktig og en pod omplanlegges, kan den starte med tom disk. ([Google Cloud — To run or not to run a database on Kubernetes](https://cloud.google.com/blog/products/databases/to-run-or-not-to-run-a-database-on-kubernetes-what-to-consider))
- **Korrupsjon ved omstart** — Postgres trenger kontrollert avslutning; Kubernetes' tidsbaserte avslutningspolicyer kan tvangsdrepe en databasepod midt i en skriveoperasjon. ([CockroachDB — Kubernetes: The state of stateful apps](https://www.cockroachlabs.com/blog/kubernetes-state-of-stateful-apps/))
- **Kompleksitet rundt lagringsprovisionering** — PersistentVolumes, PersistentVolumeClaims, StorageClasses — konsepter å lære som ikke har noe med produktet vårt å gjøre
- **Nettverk** — Pod-IP-er endres ved omstart; databaser trenger stabile tilkoblinger via StatefulSets og headless services
- **Sikkerhetskopiering** — Kubernetes sikkerhetskopierer ikke persistente volumer; du må sette opp pg_dump cron-jobber eller volumsnapshots selv

Med Docker Compose ligger data i et navngitt volum på disk. Enkelt, forutsigbart, vanskelig å rote til.

### Hva vi ikke mister

Arkitekturen vår er allerede mikrotjenestebasert med riktig separasjon. Kubernetes er bare et operativt distribueringsvalg — det gjør ikke arkitekturen bedre. De samme containerne kan distribueres til Kubernetes i fremtiden ved behov.
