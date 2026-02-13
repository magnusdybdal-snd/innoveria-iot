#!/bin/sh
set -e

echo "Importing LoRaWAN devices..."

chirpstack -c /etc/chirpstack import-device-profiles -d /opt/lorawan-devices || true

echo "Starting ChirpStack..."

exec chirpstack -c /etc/chirpstack
