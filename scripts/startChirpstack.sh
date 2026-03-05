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

  echo "Seeding ChirpStack with known dev tenant and application..."

  PGPASSWORD=chirpstack psql -h chirpstack-postgres -U chirpstack -d chirpstack <<SQL
INSERT INTO tenant (id, created_at, updated_at, name, description,
    can_have_gateways, max_gateway_count, max_device_count,
    private_gateways_up, private_gateways_down, tags)
VALUES ('d0000000-0000-0000-0000-000000000001', NOW(), NOW(),
    'Innoveria Dev', '', true, 0, 0, false, false, '{}')
ON CONFLICT (id) DO NOTHING;

INSERT INTO application (id, tenant_id, created_at, updated_at, name, description, mqtt_tls_cert, tags)
VALUES ('e0000000-0000-0000-0000-000000000001',
    'd0000000-0000-0000-0000-000000000001',
    NOW(), NOW(), 'Innoveria Dev App', '', '\x', '{}')
ON CONFLICT (id) DO NOTHING;
SQL
  echo "ChirpStack seed done."

  # Keep container running
  wait $CHIRP_PID