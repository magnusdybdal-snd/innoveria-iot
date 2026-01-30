# Prerequisite
- docker and docker compose: https://docs.docker.com/compose/install/
- Tilt: https://docs.tilt.dev/index.html

# Development enviroment

Development enviroment with hot code reloading

```bash
docker compose up -d # runs all services

docker compose down # Stops all docker instances

docker ps # see all docker instances

docker compose logs -f # see all logging from every service (detach with "d")

docker compose logs <service-name> -f # see specific service

docker compose down --rmi all # Remove all data
```

Tilt will showcase all services in a dashboard
after running docker compose -d do:
```bash
tilt up   # start
tilt down # stop
```

# Production deployment
