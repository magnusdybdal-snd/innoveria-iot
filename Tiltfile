docker_compose("./docker-compose.yml")

dc_resource(
    'web',
    labels=['frontend']
)

dc_resource(
    'api-gateway',
    labels=['api']
)

dc_resource(
    'chirpstack',
    labels=['chirpstack']
)

dc_resource(
    'chirpstack-gateway-bridge',
    labels=['chirpstack']
)

dc_resource(
    'chirpstack-rest-api',
    labels=['chirpstack']
)

dc_resource(
    'postgres',
    labels=['chirpstack']
)

dc_resource(
    'redis',
    labels=['chirpstack']
)

dc_resource(
    'mosquitto',
    labels=['chirpstack']
)
