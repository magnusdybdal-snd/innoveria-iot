FROM chirpstack/chirpstack:4

# TODO: clone the forked version here
# TODO: Handle stable commit version

USER root

RUN apk add --no-cache git \
    && git clone https://github.com/chirpstack/chirpstack-device-profiles.git /opt/lorawan-devices

USER 1000
