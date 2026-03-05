FROM oven/bun:latest AS builder

WORKDIR /app

COPY ./web/package.json ./web/bun.lock ./
RUN bun install

COPY ./web ./

RUN bun run build

FROM caddy:2-alpine

COPY --from=builder /app/dist /usr/share/caddy
COPY infra/production/Caddyfile /etc/caddy/Caddyfile
