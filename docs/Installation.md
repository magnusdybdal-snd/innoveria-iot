# Prerequisite
- docker and docker compose: https://docs.docker.com/compose/install/
- Tilt: https://docs.tilt.dev/index.html

Related guides:
- `docs/erp-agent-service-onprem-installation.md` for standalone on-prem ERP agent deployment

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
Most important command (start production):

NOTE: Its a prerequisite to set up enviroment variables in a file called .env.production
    - See .env.example on how to set it up


As both dev and prod is in root. Its recommended to run prod with -p to create a new docker project
```bash
docker compose -p app-prod --env-file .env.production -f docker-compose.prod.yml up -d --build
```

Useful commands:

```bash
docker compose -f docker-compose.prod.yml ps
docker compose -f docker-compose.prod.yml logs -f
docker compose -f docker-compose.prod.yml down
docker compose -f docker-compose.prod.yml down -v # deletes volumes (database data)
```
