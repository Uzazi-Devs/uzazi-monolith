FROM oven/bun:1-alpine AS build
WORKDIR /app

COPY package.json bun.lock ./
COPY services/auth-service/package.json services/auth-service/
COPY frontends/app-nextjs/package.json frontends/app-nextjs/
COPY frontends/marketing-astro/package.json frontends/marketing-astro/
RUN bun install --frozen-lockfile

COPY frontends/marketing-astro frontends/marketing-astro
WORKDIR /app/frontends/marketing-astro
RUN bun run build

FROM caddy:2-alpine AS runner
COPY --from=build /app/frontends/marketing-astro/dist /usr/share/caddy
EXPOSE 4321
CMD ["caddy", "file-server", "--root", "/usr/share/caddy", "--listen", ":4321"]
