FROM chirpstack/chirpstack:4

USER root

RUN apk add --no-cache git postgresql-client \
    && git clone https://github.com/vinjdev/chirpstack-device-profiles.git /opt/lorawan-devices

USER 1000
