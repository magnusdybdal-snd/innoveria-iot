FROM chirpstack/chirpstack:4

# TODO: Handle stable commit version

USER root

RUN apk add --no-cache git \
    && git clone https://github.com/vinjdev/chirpstack-device-profiles.git /opt/lorawan-devices

USER 1000
