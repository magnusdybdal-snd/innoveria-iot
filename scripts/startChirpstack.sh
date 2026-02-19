#!/bin/sh
set -e

echo "Importing LoRaWAN devices..."

chirpstack -c /etc/chirpstack import-device-profiles -d /opt/lorawan-devices || true

echo "Starting ChirpStack..."

# Start in background
chirpstack -c /etc/chirpstack &
CHIRP_PID=$!

echo "Waiting for ChirpStack to initialize..."
sleep 15

echo "Creating API key..."

# Ensure directory exists
mkdir -p /secrets

# Create API key (ignore if already exists)
KEY_FILE="/secrets/chirpstack-api-key"

# Checks to see if key file does not exist, and generates an api key
if [ ! -f "$KEY_FILE" ]; then
  echo "Generating API key..."
  chirpstack -c /etc/chirpstack create-api-key \
    --name innoveria-iot > "$KEY_FILE"
  echo "API key stored in $KEY_FILE"
else
  echo "API key already exists"
fi

# Keep container running
wait $CHIRP_PID
