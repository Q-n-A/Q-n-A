FROM golang:1.17.5-alpine AS builder

RUN go install github.com/jwilder/dockerize@latest
RUN go install github.com/cosmtrek/air@latest

FROM golang:1.17.5-alpine as hotReload
WORKDIR /Q-n-A

COPY --from=builder /go/bin/air /usr/local/bin/
COPY --from=builder /go/bin/dockerize /usr/local/bin/

EXPOSE 9000

ENTRYPOINT dockerize -timeout 60s -wait tcp://mariadb:3306 air
