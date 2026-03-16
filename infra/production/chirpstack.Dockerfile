FROM chirpstack/chirpstack:4

USER root

COPY infra/config/chirpstack-device-profiles /opt/lorawan-devices

USER 1000
