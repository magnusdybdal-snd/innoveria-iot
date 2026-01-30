FROM oven/bun:1.1-alpine

WORKDIR /app

COPY ./web/package.json ./web/bun.lock ./

RUN bun install

COPY ./web ./

EXPOSE 3000

CMD ["bun", "run", "dev"]
