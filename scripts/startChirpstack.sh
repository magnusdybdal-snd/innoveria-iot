#!/bin/sh
  set -e

  echo "Importing LoRaWAN devices..."

  chirpstack -c /etc/chirpstack import-device-profiles -d /opt/lorawan-devices || true

  echo "Starting ChirpStack..."

  # Start in background
  chirpstack -c /etc/chirpstack &
  CHIRP_PID=$!

  # Ensure directory and key file path are set before the retry loop
  mkdir -p /secrets
  KEY_FILE="/secrets/chirpstack-api-key"

  echo "Waiting for ChirpStack to be ready..."
  until chirpstack -c /etc/chirpstack create-api-key \
    --name innoveria-iot 2>/dev/null | \
    grep '^token:' | \
    awk '{print $2}' > "$KEY_FILE" && [ -s "$KEY_FILE" ]; do
      echo "Not ready yet, retrying..."
      rm -f "$KEY_FILE"
      sleep 2
  done
  echo "API key stored in $KEY_FILE"

  if [ "$CHIRPSTACK_SEED" = "true" ]; then
    echo "Seeding ChirpStack with known dev tenant and application..."
    PGPASSWORD=chirpstack psql -h chirpstack-postgres -U chirpstack -d chirpstack < /scripts/seed-chirpstack.sql
    echo "ChirpStack seed done."
  fi

  # Keep container running
  wait $CHIRP_PID
