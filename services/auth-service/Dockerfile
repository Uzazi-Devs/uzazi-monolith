FROM oven/bun:1-alpine AS build
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile
COPY . .
RUN bun run build

FROM oven/bun:1-alpine
WORKDIR /app
COPY --from=build --chown=bun:bun /app ./
USER bun
EXPOSE 3000
CMD ["bun", "run", "start"]
