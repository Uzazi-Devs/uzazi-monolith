FROM oven/bun:1-alpine AS build
WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

COPY package.json bun.lock ./
COPY services/auth-service/package.json services/auth-service/
COPY frontends/app-nextjs/package.json frontends/app-nextjs/
COPY frontends/marketing-astro/package.json frontends/marketing-astro/
RUN bun install --frozen-lockfile

COPY frontends/app-nextjs frontends/app-nextjs
WORKDIR /app/frontends/app-nextjs
RUN bun run build

FROM oven/bun:1-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3001
ENV HOSTNAME="0.0.0.0"

COPY --from=build --chown=bun:bun /app/frontends/app-nextjs/.next/standalone ./
COPY --from=build --chown=bun:bun /app/frontends/app-nextjs/.next/static ./frontends/app-nextjs/.next/static

USER bun
EXPOSE 3001
CMD ["bun", "run", "frontends/app-nextjs/server.js"]
