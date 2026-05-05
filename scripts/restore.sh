#!/usr/bin/env bash
set -euo pipefail

# Restore a single database from a backup file.
# Usage: ./scripts/restore.sh <db-service> <backup-file>
#
# Examples:
#   ./scripts/restore.sh collection-db backups/collection-db_2026-05-04.dump.gz
#   ./scripts/restore.sh device-db     backups/device-db_2026-05-04.dump.gz
#
# For TimescaleDB (collection-db): automatically tears down and recreates the
# container and volume so the restore goes into a clean empty database.
# For all other databases: restores in-place using --clean --if-exists.
#
# The dependent application service is stopped before restore and restarted after.

DB_SERVICE="${1:?Usage: $0 <db-service> <backup-file>}"
BACKUP_FILE="${2:?}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
ENV_FILE="$PROJECT_DIR/.env"

if [[ ! -f "$BACKUP_FILE" ]]; then
  echo "ERROR: backup file not found: $BACKUP_FILE" >&2
  exit 1
fi

# ---------------------------------------------------------------------------
# Load credentials — .env in production, dev defaults as fallback
# ---------------------------------------------------------------------------
if [[ -f "$ENV_FILE" ]]; then
  set -o allexport
  # shellcheck source=/dev/null
  source <(grep -E '^[A-Z_]+=.+' "$ENV_FILE")
  set +o allexport
  COMPOSE_FILE="${COMPOSE_FILE:-docker-compose.prod.yml}"
else
  COMPOSE_FILE="${COMPOSE_FILE:-docker-compose.yml}"
fi

export COMPOSE_FILE

# ---------------------------------------------------------------------------
# Database definitions per service
# ---------------------------------------------------------------------------
declare -A DB_NAME DB_USER DB_PASS IS_TIMESCALE DEPENDENT_SERVICE

DB_NAME[collection-db]="${COLLECTION_DB_NAME:-collection}"
DB_USER[collection-db]="${COLLECTION_DB_USER:-collection}"
DB_PASS[collection-db]="${COLLECTION_DB_PASSWORD:-collection}"
IS_TIMESCALE[collection-db]="yes"
DEPENDENT_SERVICE[collection-db]="collection-service"

DB_NAME[device-db]="${DEVICE_DB_NAME:-device}"
DB_USER[device-db]="${DEVICE_DB_USER:-device}"
DB_PASS[device-db]="${DEVICE_DB_PASSWORD:-device}"
IS_TIMESCALE[device-db]="no"
DEPENDENT_SERVICE[device-db]="device-service"

DB_NAME[auth-db]="${AUTH_DB_NAME:-auth}"
DB_USER[auth-db]="${AUTH_DB_USER:-auth}"
DB_PASS[auth-db]="${AUTH_DB_PASSWORD:-auth}"
IS_TIMESCALE[auth-db]="no"
DEPENDENT_SERVICE[auth-db]="auth-service"

DB_NAME[context-db]="${CONTEXT_DB_NAME:-context}"
DB_USER[context-db]="${CONTEXT_DB_USER:-context}"
DB_PASS[context-db]="${CONTEXT_DB_PASSWORD:-context}"
IS_TIMESCALE[context-db]="no"
DEPENDENT_SERVICE[context-db]="context-service"

DB_NAME[erp-db]="${ERP_DB_NAME:-erp}"
DB_USER[erp-db]="${ERP_DB_USER:-erp}"
DB_PASS[erp-db]="${ERP_DB_PASSWORD:-erp}"
IS_TIMESCALE[erp-db]="no"
DEPENDENT_SERVICE[erp-db]="erp-service"

DB_NAME[chirpstack-db]="chirpstack-db"
DB_USER[chirpstack-db]="${CHIRPSTACK_DB_USER:-chirpstack}"
DB_PASS[chirpstack-db]="${CHIRPSTACK_DB_PASSWORD:-chirpstack}"
IS_TIMESCALE[chirpstack-db]="no"
DEPENDENT_SERVICE[chirpstack-db]="chirpstack"

if [[ -z "${DB_NAME[$DB_SERVICE]+x}" ]]; then
  echo "ERROR: unknown service '$DB_SERVICE'" >&2
  echo "Valid services: collection-db device-db auth-db context-db erp-db chirpstack-db" >&2
  exit 1
fi

DB="${DB_NAME[$DB_SERVICE]}"
USER="${DB_USER[$DB_SERVICE]}"
PASSWORD="${DB_PASS[$DB_SERVICE]}"
TIMESCALE="${IS_TIMESCALE[$DB_SERVICE]}"
DEPENDENT="${DEPENDENT_SERVICE[$DB_SERVICE]}"

cd "$PROJECT_DIR"

echo "=== Restoring $DB_SERVICE from $(basename "$BACKUP_FILE") ==="
echo "  Database   : $DB"
echo "  TimescaleDB: $TIMESCALE"
echo "  Stops first: $DEPENDENT"
echo ""

# ---------------------------------------------------------------------------
# Stop the dependent application service
# ---------------------------------------------------------------------------
echo "  Stopping $DEPENDENT..."
docker compose stop "$DEPENDENT" 2>/dev/null || true

# ---------------------------------------------------------------------------
# TimescaleDB: tear down and recreate for a clean restore
# ---------------------------------------------------------------------------
if [[ "$TIMESCALE" == "yes" ]]; then
  echo "  Tearing down $DB_SERVICE (TimescaleDB requires a fresh database)..."

  # Capture the volume name before removing the container
  CONTAINER_ID=$(docker compose ps -q "$DB_SERVICE" 2>/dev/null | head -1 || true)
  if [[ -n "$CONTAINER_ID" ]]; then
    VOLUME=$(docker inspect "$CONTAINER_ID" \
      --format '{{range .Mounts}}{{if eq .Type "volume"}}{{.Name}}{{end}}{{end}}' 2>/dev/null \
      | awk '{print $1}')
  else
    VOLUME=""
  fi

  docker compose stop "$DB_SERVICE"
  docker compose rm -f "$DB_SERVICE"

  if [[ -n "$VOLUME" ]]; then
    echo "  Removing volume $VOLUME..."
    docker volume rm "$VOLUME"
  fi

  echo "  Recreating $DB_SERVICE..."
  docker compose up -d "$DB_SERVICE"

  # Wait until the database accepts connections
  echo "  Waiting for $DB_SERVICE to be ready..."
  for i in $(seq 1 30); do
    if docker compose exec -T -e PGPASSWORD="$PASSWORD" "$DB_SERVICE" \
        psql -U "$USER" -d "$DB" -c "SELECT 1;" &>/dev/null; then
      echo "  Ready."
      break
    fi
    if [[ $i -eq 30 ]]; then
      echo "ERROR: $DB_SERVICE did not become ready in time" >&2
      exit 1
    fi
    sleep 2
  done

  echo "  Running timescaledb_pre_restore()..."
  docker compose exec -T -e PGPASSWORD="$PASSWORD" "$DB_SERVICE" \
    psql -U "$USER" -d "$DB" -c "SELECT timescaledb_pre_restore();"

  echo "  Restoring data..."
  gunzip -c "$BACKUP_FILE" \
    | docker compose exec -T -e PGPASSWORD="$PASSWORD" "$DB_SERVICE" \
      pg_restore -U "$USER" -d "$DB" --no-tablespaces

  echo "  Running timescaledb_post_restore()..."
  docker compose exec -T -e PGPASSWORD="$PASSWORD" "$DB_SERVICE" \
    psql -U "$USER" -d "$DB" -c "SELECT timescaledb_post_restore();"

# ---------------------------------------------------------------------------
# Standard PostgreSQL: restore in-place
# ---------------------------------------------------------------------------
else
  echo "  Restoring data..."
  gunzip -c "$BACKUP_FILE" \
    | docker compose exec -T -e PGPASSWORD="$PASSWORD" "$DB_SERVICE" \
      pg_restore -U "$USER" -d "$DB" --clean --if-exists
fi

# ---------------------------------------------------------------------------
# Restart the dependent service
# ---------------------------------------------------------------------------
echo "  Starting $DEPENDENT..."
docker compose start "$DEPENDENT"

echo ""
echo "=== Restore complete ==="
