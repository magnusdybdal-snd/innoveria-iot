FROM chirpstack/chirpstack:4

USER root

RUN apk add --no-cache postgresql-client

COPY infra/config/chirpstack-device-profiles /opt/lorawan-devices

USER 1000
