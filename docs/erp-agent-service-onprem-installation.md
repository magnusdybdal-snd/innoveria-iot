# ERP Agent Service On-Prem Installation

This guide installs `erp-agent-service` as a standalone on-prem service, separate from the rest of the platform.

## What this service needs

- Network access to Monitor ERP (`https://<MONITOR_ERP_HOST>:<MONITOR_ERP_PORT>`)
- Network access to `erp-service` ingest API (`http(s)://<ERP_SERVICE_HOST>:<PORT>/api/v1/erp/ingest/...`)
- A valid platform admin JWT access token used as `JWT_TOKEN`

## 1) Build the Docker image

Run from repository root (the Dockerfile copies both `pkg/` and `services/erp-agent-service/`):

```bash
docker build -f infra/production/erp-agent-service.Dockerfile -t innoveria/erp-agent-service:onprem .
```

## 2) Create environment file

Create `erp-agent-service.onprem.env`:

```env
GO_ENV=production

# JWT from platform admin session (see step 3)
JWT_TOKEN=<paste_access_jwt_here>

# Monitor ERP connection
MONITOR_ERP_HOST=<monitor-host>
MONITOR_ERP_PORT=<monitor-port>
MONITOR_ERP_COMPANY_NUMBER=<company-number>
MONITOR_ERP_USERNAME=<monitor-username>
MONITOR_ERP_PASSWORD=<monitor-password>
MONITOR_ERP_FORCE_RELOGIN=false

# ERP service endpoint (where agent pushes data)
ERP_SERVICE=http://<erp-service-host>:8080

# Optional tuning
POLLING_INTERVAL=10m
CYCLE_TIMEOUT=2m
MAX_BACKOFF_TIME=5m
PORT=8080
```

Notes:
- `MOCK_MONITOR` must not be enabled in production.
- `PORT` is currently not used for inbound HTTP in this service, but leaving default is safe.

## 3) Get `Company ERP Token` from the website (platform admin)

Use a platform admin account in the web app and extract a valid **access JWT**:

1. Log in to the website as platform admin.
2. Open browser DevTools.
3. Find an authenticated API request (Network tab) and copy the `Authorization` header value.
4. Remove the `Bearer ` prefix and paste the token into `JWT_TOKEN`.

Security recommendations:
- Treat this token as a secret.
- Do not commit it to git or store it in plain-text docs.
- Rotate the token when access changes.

## 4) Run container

```bash
docker run -d \
  --name erp-agent-service \
  --restart unless-stopped \
  --env-file erp-agent-service.onprem.env \
  innoveria/erp-agent-service:onprem
```

## 5) Verify operation

```bash
docker logs -f erp-agent-service
```

Expected behavior:
- Service starts worker loop.
- It logs successful sync cycles on interval.
- Data is posted to `erp-service` ingest endpoints.

## Troubleshooting

- `unauthorized ERP token` in logs:
  - `JWT_TOKEN` is invalid/expired, or signed by another environment.
  - Get a fresh platform admin access token and restart container.

- Monitor login/query failures:
  - Verify `MONITOR_ERP_*` values.
  - Verify network/TLS reachability from the on-prem host.

- No data appears in ERP:
  - Verify `ERP_SERVICE` points to reachable `erp-service`.
  - Verify `erp-service` runs with matching `ERP_AGENT_JWT_SECRET` for the token issuer environment.
