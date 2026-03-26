# syntax=docker/dockerfile:1.7
FROM node:20-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV NODE_OPTIONS="--max-old-space-size=4096"
WORKDIR /app
RUN corepack enable

FROM base AS deps
COPY package.json pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm-store,target=/pnpm/store \
    pnpm install --frozen-lockfile --store-dir=/pnpm/store

FROM deps AS build
COPY . .
RUN pnpm run build


FROM nginx:alpine
LABEL MAINTAINER="bypanghu@163.com"
LABEL POWER="gin-vue-admin"
ENV BACKEND=localhost \
    BACKEND_API_PORT=8888 \
    BACKEND_WS_PORT=8888 \
    GZIP=on \
    GZIP_STATIC=on
RUN apk add --no-cache gettext
RUN mkdir -p /etc/nginx/conf.d
COPY --from=build /app/nginx/default.conf.template /etc/nginx/default.conf.template
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["/bin/sh", "-c", "envsubst '${BACKEND} ${BACKEND_API_PORT} ${BACKEND_WS_PORT} ${GZIP} ${GZIP_STATIC}' < /etc/nginx/default.conf.template > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"]
