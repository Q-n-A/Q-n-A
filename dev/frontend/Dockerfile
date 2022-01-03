FROM logica0419/protoc-node:16.13.1 AS builder
WORKDIR /build

COPY ./Q-n-A_UI /build
RUN npm ci --unsafe-perm
RUN npm run gen
RUN npm run build

FROM caddy:2.4.6-alpine AS runner
EXPOSE 80

COPY --from=builder /build/dist /usr/share/caddy
