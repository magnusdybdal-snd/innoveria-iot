# Wiring the Repository Layer into Services
> Branch: `107-device-service-wire-repository-into-services`

This document walks through connecting the already-implemented repository layer into the service layer. The repositories exist and work — this is purely about dependency injection and filling in the TODO stubs.

---

## The mental model

Before touching any code, understand the flow:

```
HTTP request
    → handler (parses HTTP, calls service)
        → service (business logic: talks to DB + ChirpStack, merges results)
            → repository (SQL only, no business logic)
            → chirpstack client (REST calls only)
        ← service returns domain struct
    ← handler encodes response
```

The services currently skip the repository step entirely. Your job is to add it back in.

---

## Step 1 — Add `Delete` to `GatewayRepository`

**Why first?** It's a self-contained two-file change with no dependencies on anything else. Good warmup.

**What:** The `GatewayRepository` interface in `domain/gateway.go` is missing a `Delete` method. The service needs to delete from our DB after deleting from ChirpStack.

**Where:**
- `domain/gateway.go` — add `Delete` to the `GatewayRepository` interface
- `repository/gateway_repository.go` — implement it

**How to implement:** Look at `UpdateState` in `gateway_repository.go`. `Delete` follows the exact same pattern:
- `Exec` with a `DELETE WHERE gateway_id = $1`
- Check `RowsAffected() == 0` and return an error if nothing was deleted
- Wrap the error with `fmt.Errorf`

---

## Step 2 — Update the service constructors

**Why:** Services need to receive their repositories as dependencies. This is dependency injection — the service doesn't create its own DB connection, it receives one. This makes testing possible and keeps concerns separated.

**Where:**
- `service/gateway_service.go`
- `service/sensor_service.go`

**What to do in each file:**
1. Add the repository fields to the struct (look at how `cc *chirpstackrest.Client` is already there as a field — same idea)
2. Update the `New...` constructor function to accept the repos as parameters and store them on the struct

For `GatewayServiceImpl` you need: `GatewayRepository` and `CompanyConfigRepository`
For `SensorServiceImpl` you need: `SensorRepository` and `CompanyConfigRepository`

Use the **domain interfaces** as the types (e.g., `domain.GatewayRepository`), not the concrete `*repository.GatewayRepository`. This is the interface segregation principle — the service only knows about the contract, not the implementation.

**The code won't compile yet** after this step — `server.go` still calls the old constructors. That's fine, fix it in step 5.

---

## Step 3 — Fix `gateway_service.go` methods

Work through each method and replace the TODO stubs.

### `GetAll` — the merge pattern

This is the most important one to understand. Read this before coding:

ChirpStack knows about connectivity status. Your DB knows about your internal metadata (company mapping, internal ID). Neither source alone gives you the full picture. The service merges them.

The flow:
1. Call `chirpstack.GetAllGateways(ctx, 100)` — 100 is a reasonable limit until auth gives you a real count. Leave a `// TODO: get real count from DB once auth exists` comment.
2. For each ChirpStack gateway in the response, call `gatewayRepo.FindByEUI(ctx, csGateway.GatewayEUI)` to get the matching DB record.
3. Build the `domain.Gateway` by taking DB fields (Id, CompanyId, Name, State) and runtime fields from ChirpStack (Status, LastSeenAt). You'll need to update the mapper or do the merge inline — your choice.

Note on the mapper TODOs: `MapChirpstackGateway` in `service/mappers/gateway_mapper.go` has two TODOs (`// TODO: Change this to internal database id` and `// TODO: look up company mapping in DB`). Now that you have the DB data available in the service, you can resolve these. Consider updating the mapper to also accept a `domain.Gateway` (the DB record) so the merge logic lives in one place.

### `Create`
1. Look up `companyConfigRepo.FindByCompanyID(ctx, payload.CompanyId)` to get the real `ChirpstackTenantID`
2. Pass that to `MapCreateChirpstackGateway` instead of `payload.CompanyId` (the current code is wrong — it passes the internal company UUID as a ChirpStack tenant ID)
3. After ChirpStack succeeds, call `gatewayRepo.Create(ctx, payload)` to persist to your DB

### `Update`
1. `gatewayRepo.FindByID(ctx, gatewayId)` — verify the gateway exists in your DB first. If not found, return an appropriate error.
2. `companyConfigRepo.FindByCompanyID(ctx, dbGateway.CompanyId)` — get the ChirpStack tenant ID
3. Call ChirpStack with the correct tenant ID

### `Delete`
Currently broken — it passes the internal UUID directly to ChirpStack as a gateway EUI, which is wrong.
1. `gatewayRepo.FindByID(ctx, gatewayID)` — get the GatewayEUI
2. Call `chirpstack.DeleteGateway(ctx, dbGateway.GatewayEUI)` — now using the correct identifier
3. Call `gatewayRepo.Delete(ctx, gatewayID)` — remove from your DB

Order matters: delete from ChirpStack first. If that fails, you still have your DB record and can retry. If you delete from DB first and ChirpStack fails, you've lost the mapping.

---

## Step 4 — Fix `sensor_service.go` methods

### `GetAll` — same merge pattern as gateways

The broken line is `applicationId := ""`. ChirpStack requires an applicationId to list devices.

Flow:
1. `companyConfigRepo.FindByCompanyID(ctx, companyId)` — get `ChirpstackApplicationID`
2. `chirpstack.GetAllSensors(ctx, 100, cfg.ChirpstackApplicationID)`
3. For each ChirpStack sensor, `sensorRepo.FindByEUI(ctx, csSensor.DeviceEUI)` to get DB record
4. Merge and return

**The companyId placeholder problem:** Without auth there's no company ID in the request context. Define a package-level constant for now:
```
const demoCompanyID = "a0000000-0000-0000-0000-000000000001"
```
With a `// TODO: replace with company ID from auth context` comment. Be explicit that this is temporary — don't hide it.

### `Create`
Leave it for now. It's empty and not needed for the current scope.

---

## Step 5 — Wire everything in `server.go`

Now that the constructors have new signatures, `server.go` won't compile. Fix it here.

After the DB is initialised (after `db.RunSeeds`), instantiate the three repositories:
```
companyConfigRepo = repository.NewCompanyConfigRepository(database)
gatewayRepo       = repository.NewGatewayRepository(database)
sensorRepo        = repository.NewSensorRepository(database)
```

Then pass them into the service constructors.

You'll need to add the `repository` package to the imports. The package path follows the same pattern as the other imports already in the file.

---

## Checklist

- [ ] `domain/gateway.go` — `Delete` added to `GatewayRepository` interface
- [ ] `repository/gateway_repository.go` — `Delete` implemented
- [ ] `service/gateway_service.go` — struct updated, constructor updated, all 4 methods fixed
- [ ] `service/sensor_service.go` — struct updated, constructor updated, `GetAll` fixed
- [ ] `service/mappers/gateway_mapper.go` — TODO comments resolved
- [ ] `service/mappers/sensor_mapper.go` — TODO comments resolved (Id, CompanyID from DB)
- [ ] `server.go` — repos instantiated, new constructor signatures used
- [ ] Code compiles (`go build ./...` from the service root)

---

## How to verify it works

Once the code compiles and the service is running:

- `GET /gateways` — should return gateways with correct internal IDs and company mapping, plus live status from ChirpStack
- `GET /sensors` — should no longer fail on empty applicationId; returns sensors with DB metadata merged with ChirpStack status
- Check logs for any `find gateway by eui: no rows` errors — means a device exists in ChirpStack but not in your DB (expected for unregistered devices, handle gracefully)
