FROM oven/bun:latest

WORKDIR /app

COPY ./web/package.json ./web/bun.lock ./
RUN bun install --production

COPY ./web ./

EXPOSE 3000

CMD ["bun", "run", "start"]
