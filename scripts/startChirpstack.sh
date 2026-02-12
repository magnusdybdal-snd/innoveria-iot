#!/bin/sh
set -e

echo "Importing device profiles..."
chirpstack -c /etc/chirpstack import-device-profiles -d /profiles || true

echo "Starting ChirpStack..."
exec chirpstack -c /etc/chirpstack


