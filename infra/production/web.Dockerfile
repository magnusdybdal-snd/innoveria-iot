FROM oven/bun:latest AS builder

WORKDIR /app

COPY ./web/package.json ./web/bun.lock ./
RUN bun install

COPY ./web ./

RUN bun run build

FROM node:20-alpine

WORKDIR /app

RUN npm install -g serve

COPY --from=builder /app/dist ./dist

EXPOSE 3000

CMD ["serve", "-s", "dist", "-l", "3000"]

