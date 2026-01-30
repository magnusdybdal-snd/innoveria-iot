# Prerequisite
- docker and docker compose
https://docs.docker.com/compose/install/

# Development enviroment

Development enviroment with hot code reloading

```bash
docker compose up -d # runs all services

docker compose down # Stops all docker instances

docker ps # see all docker instances

docker compose logs -f # see all logging from every service (detach with "d")

docker compose logs <service-name> -f # see specific service
```

# Production deployment
