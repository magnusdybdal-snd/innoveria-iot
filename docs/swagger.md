# Swagger API Documentation

We use [swaggo/swag](https://github.com/swaggo/swag) to generate Swagger 2.0 docs from annotations in handler code.

## Accessing the docs

| Endpoint | URL |
|---|---|
| Merged UI (all services) | `http://localhost:8081/swagger/` |
| Device service only | `http://localhost:8083/swagger/` |
| Collection service only | `http://localhost:8082/swagger/` |

The merged UI at the API gateway is the primary place to view all endpoints. It requires all services to be running.

---

## How to annotate your handlers

Every handler function needs a comment block directly above the `func` keyword. Example:

```go
// GetSensors returns all sensors.
//
// @Summary      List all sensors
// @Tags         sensors
// @Produce      json
// @Success      200  {object}  dto.SensorListResponse
// @Failure      500
// @Router       /sensors [get]
func GetSensors(svc domain.SensorService) http.HandlerFunc {
```

### Key annotations

| Annotation | Description |
|---|---|
| `@Summary` | Short title shown in the UI |
| `@Tags` | Groups endpoints — use the resource name (e.g. `sensors`, `gateways`) |
| `@Accept` | Request body format — use `json` |
| `@Produce` | Response format — use `json` |
| `@Param` | Describes an input parameter (see below) |
| `@Success` | Success response code and optional body type |
| `@Failure` | Error response codes |
| `@Router` | Path **relative to the service basePath** and HTTP method |

### @Param format

```
@Param  <name>  <in>  <type>  <required>  "<description>"
```

- `in` can be `path`, `query`, or `body`
- For body params, `type` is a DTO struct (e.g. `dto.CreateSensorRequest`)

Examples:
```go
// @Param  id          path   string                   true  "Sensor ID"
// @Param  device_eui  query  string                   true  "Device EUI"
// @Param  body        body   dto.CreateSensorRequest  true  "Request payload"
```

### Response body types

- Single object: `{object} dto.SensorResponse`
- List/array: `{array} dto.SensorResponse`

---

## Adding swagger to a new service

**1. Add the service to the Makefile** (`/Makefile`):
```makefile
SERVICES         := ... services/your-service
ANNOTATED_SERVICES := ... services/your-service
```

**2. Install dependencies** (from repo root):
```bash
make swagger-deps
```

**3. Add the global API info block** in `cmd/main.go` above `func main()`:
```go
// @title       Your Service API
// @version     1.0
// @description Short description of what this service does.

// @host        localhost:<port>
// @BasePath    /api/v1/your-service
```

**4. Annotate all handler functions** — see examples above.

**5. Register the Swagger UI route** in `internal/server/router.go`:
```go
import (
    _ "innoveria-iot/your-service/docs"
    httpSwagger "github.com/swaggo/http-swagger"
)

mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
```

> **Note:** The `docs` import will show a compile error (red underline) until you run `make swagger-gen` in step 6 — the `docs/` package does not exist yet. This is expected.

After adding the imports, run:
```bash
go mod tidy
```
from the service directory to resolve any missing dependencies before building.

**6. Generate the docs** (from repo root):
```bash
make swagger-gen
```

**7. Wire the new service into the gateway merger.**

The gateway merges specs from each service at runtime. You need to update three places:

**a) `services/api-gateway/internal/config/config.go`** — add the service URL to the config struct and load it from an environment variable:
```go
YourSvcURL string

// in Load():
YourSvcURL: env.Get("YOUR_SERVICE", "http://your-service:8080"),
```

**b) `services/api-gateway/internal/handlers/swagger_handler.go`** — add the new service as a parameter and fetch+merge its spec:
```go
// Add the parameter:
func MergedSwaggerSpec(deviceSvcURL, collSvcURL, yourSvcURL string) http.HandlerFunc {

    // Fetch the new service spec (same pattern as existing services):
    yourSpec, err := fetchSpec(yourSvcURL)
    if err != nil {
        http.Error(w, "failed to fetch your-service spec", http.StatusInternalServerError)
        return
    }

    // Merge its paths and definitions (add alongside the existing loops):
    for path, val := range yourSpec.Paths {
        merged.Paths[yourSpec.BasePath+path] = val
    }
    maps.Copy(merged.Definitions, yourSpec.Definitions)
}
```

**c) `services/api-gateway/internal/server/router.go`** — pass the new URL when calling `MergedSwaggerSpec`:
```go
mux.HandleFunc("GET /swagger/doc.json", handlers.MergedSwaggerSpec(cfg.DeviceSvcURL, cfg.CollSvcURL, cfg.YourSvcURL))
```

---

## Regenerating docs after changes

Any time you add or modify handler annotations, regenerate the docs:

```bash
make swagger-gen
```

Commit the updated `docs/` folder along with your annotation changes.

---

## Marking required vs optional fields in DTOs

To make required and optional fields visible in the Swagger UI Model view, tag struct fields with `binding:"required"`:

```go
type CreateSensorRequest struct {
    Name      string  `json:"name"       binding:"required"` // shows as required (*) in UI
    CompanyID string  `json:"company_id" binding:"required"` // shows as required (*) in UI
    Notes     *string `json:"notes"`                         // optional — nil if omitted from request
}
```

- Use `binding:"required"` (not `validate:"required"` — swaggo does not recognise that tag)
- Use `*string` for optional fields — a missing field in the JSON payload sets the pointer to `nil`
- Add an inline comment on optional fields explaining why they can be omitted

---

## Notes

- `@Router` paths are relative to `@BasePath` — do not repeat the base path
- The `docs/` folder is generated code — do not edit it manually
- The `swag` CLI is registered as a Go tool in each service's `go.mod` — no separate install needed
