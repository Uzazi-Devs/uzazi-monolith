FROM oven/bun:1-alpine AS build
WORKDIR /app
COPY package.json bun.lock ./
COPY services/auth-service/package.json services/auth-service/
COPY frontends/app-nextjs/package.json frontends/app-nextjs/
COPY frontends/marketing-astro/package.json frontends/marketing-astro/
RUN bun install --frozen-lockfile
COPY services/auth-service services/auth-service
WORKDIR /app/services/auth-service
RUN bun run build

FROM oven/bun:1-alpine
WORKDIR /app/services/auth-service
COPY --from=build --chown=bun:bun /app/node_modules /app/node_modules
COPY --from=build --chown=bun:bun /app/services/auth-service ./
USER bun
EXPOSE 3000
CMD ["bun", "run", "start"]
