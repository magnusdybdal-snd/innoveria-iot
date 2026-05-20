#!/usr/bin/env bash
set -euo pipefail

# Daily backup script for all innoveria-iot PostgreSQL databases.
# Run via cron, e.g.:  0 2 * * * /path/to/scripts/backup.sh
#
# Requires: docker compose up (all db services must be running).
# Everything is auto-detected from the running containers — no flags needed.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"

BACKUP_DIR="${BACKUP_DIR:-$PROJECT_DIR/backups}"
RETENTION_DAYS=7
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")

# ---------------------------------------------------------------------------
# Auto-detect compose project name and compose file from running containers.
# Both are stored as Docker labels on every container compose manages.
# ---------------------------------------------------------------------------
CONTAINER_ID=$(docker ps -q --filter "name=collection-db" | head -1)
if [[ -n "$CONTAINER_ID" ]]; then
  COMPOSE_PROJECT_NAME="${COMPOSE_PROJECT_NAME:-$(docker inspect "$CONTAINER_ID" \
    --format '{{index .Config.Labels "com.docker.compose.project"}}' 2>/dev/null)}"
  DETECTED_COMPOSE_FILE=$(docker inspect "$CONTAINER_ID" \
    --format '{{index .Config.Labels "com.docker.compose.project.config_files"}}' 2>/dev/null || true)
else
  COMPOSE_PROJECT_NAME="${COMPOSE_PROJECT_NAME:-}"
  DETECTED_COMPOSE_FILE=""
fi

# Derive which compose file and env file to use from what's actually running.
if [[ "$DETECTED_COMPOSE_FILE" == *"prod"* ]]; then
  COMPOSE_FILE="${COMPOSE_FILE:-docker-compose.prod.yml}"
  ENV_FILE="${PROJECT_DIR}/.env.production"
else
  COMPOSE_FILE="${COMPOSE_FILE:-docker-compose.yml}"
  ENV_FILE="${PROJECT_DIR}/.env"
fi

export COMPOSE_FILE COMPOSE_PROJECT_NAME

# ---------------------------------------------------------------------------
# Load credentials — env file if present, dev defaults as fallback.
# ---------------------------------------------------------------------------
if [[ -f "$ENV_FILE" ]]; then
  set -o allexport
  # shellcheck source=/dev/null
  source <(grep -E '^[A-Z_]+=.+' "$ENV_FILE")
  set +o allexport
fi

COLLECTION_DB_USER="${COLLECTION_DB_USER:-collection}"
COLLECTION_DB_PASSWORD="${COLLECTION_DB_PASSWORD:-collection}"
COLLECTION_DB_NAME="${COLLECTION_DB_NAME:-collection}"

DEVICE_DB_USER="${DEVICE_DB_USER:-device}"
DEVICE_DB_PASSWORD="${DEVICE_DB_PASSWORD:-device}"
DEVICE_DB_NAME="${DEVICE_DB_NAME:-device}"

AUTH_DB_USER="${AUTH_DB_USER:-auth}"
AUTH_DB_PASSWORD="${AUTH_DB_PASSWORD:-auth}"
AUTH_DB_NAME="${AUTH_DB_NAME:-auth}"

CONTEXT_DB_USER="${CONTEXT_DB_USER:-context}"
CONTEXT_DB_PASSWORD="${CONTEXT_DB_PASSWORD:-context}"
CONTEXT_DB_NAME="${CONTEXT_DB_NAME:-context}"

ERP_DB_USER="${ERP_DB_USER:-erp}"
ERP_DB_PASSWORD="${ERP_DB_PASSWORD:-erp}"
ERP_DB_NAME="${ERP_DB_NAME:-erp}"

CHIRPSTACK_DB_USER="${CHIRPSTACK_DB_USER:-chirpstack}"
CHIRPSTACK_DB_PASSWORD="${CHIRPSTACK_DB_PASSWORD:-chirpstack}"

# ---------------------------------------------------------------------------
# Database definitions: "compose_service|db_name|db_user|password|timescaledb"
# Set the last field to "timescaledb" for TimescaleDB databases.
# ---------------------------------------------------------------------------
DATABASES=(
  "collection-db|${COLLECTION_DB_NAME}|${COLLECTION_DB_USER}|${COLLECTION_DB_PASSWORD}|timescaledb"
  "device-db|${DEVICE_DB_NAME}|${DEVICE_DB_USER}|${DEVICE_DB_PASSWORD}|"
  "auth-db|${AUTH_DB_NAME}|${AUTH_DB_USER}|${AUTH_DB_PASSWORD}|"
  "context-db|${CONTEXT_DB_NAME}|${CONTEXT_DB_USER}|${CONTEXT_DB_PASSWORD}|"
  "erp-db|${ERP_DB_NAME}|${ERP_DB_USER}|${ERP_DB_PASSWORD}|"
  "chirpstack-db|chirpstack-db|${CHIRPSTACK_DB_USER}|${CHIRPSTACK_DB_PASSWORD}|"
)

# ---------------------------------------------------------------------------
# Run backups
# ---------------------------------------------------------------------------
mkdir -p "$BACKUP_DIR"
cd "$PROJECT_DIR"

echo "=== Backup started at $TIMESTAMP ==="

for entry in "${DATABASES[@]}"; do
  IFS='|' read -r service db user password is_timescale <<< "$entry"
  outfile="$BACKUP_DIR/${service}_${TIMESTAMP}.dump.gz"

  echo "  Backing up $db ($service) -> $(basename "$outfile")"

  # TimescaleDB requires --no-tablespaces to avoid internal schema conflicts on restore.
  if [[ "$is_timescale" == "timescaledb" ]]; then
    extra_flags="--no-tablespaces"
  else
    extra_flags=""
  fi

  if docker compose exec -T -e PGPASSWORD="$password" "$service" \
      pg_dump -U "$user" -d "$db" -Fc $extra_flags 2>/dev/null \
    | gzip > "$outfile"; then
    echo "    OK ($(du -sh "$outfile" | cut -f1))"
  else
    echo "    FAILED — removing partial file" >&2
    rm -f "$outfile"
  fi
done

# ---------------------------------------------------------------------------
# Delete backups older than RETENTION_DAYS
# ---------------------------------------------------------------------------
echo "  Pruning backups older than $RETENTION_DAYS days..."
find "$BACKUP_DIR" -name "*.dump.gz" -mtime +"$RETENTION_DAYS" -delete

echo "=== Backup finished ==="
