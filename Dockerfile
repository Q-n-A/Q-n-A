FROM golang:1.17.5-alpine AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /Q-n-A -ldflags '-s -w'

FROM alpine:3.15.0 AS runner
EXPOSE 9000

COPY --from=builder /Q-n-A .

HEALTHCHECK CMD ./Q-n-A healthcheck || exit 1
ENTRYPOINT ./Q-n-A
