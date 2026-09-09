# FabrikkPuls

## Bachelor project 2026 — Group 203

### Overview

FabrikkPuls is a multi-tenant IoT platform that combines factory sensor data with ERP data to surface production insights (e.g. resource consumption per order). LoRaWAN sensors report through ChirpStack and MQTT into a Go microservice backend — API gateway, auth, device, collection, context, and ERP-integration services — backed by TimescaleDB, with a React/TypeScript frontend.

### Clone the project

```bash
# before cloning
git clone --recursive git@github.com:magnusdybdal-snd/innoveria-iot.git

# after cloning
git submodule update --init --recursive
```

### Running the project

Prerequisites: [Docker Compose](https://docs.docker.com/compose/install/) and [Tilt](https://docs.tilt.dev/index.html).

```bash
docker compose up -d   # start all services
tilt up                # dashboard with hot code reloading
```

See `docs/Installation.md` for production deployment and additional commands.
