FROM golang:1.18.3-alpine AS builder

RUN go install github.com/jwilder/dockerize@latest && \
  go install github.com/cosmtrek/air@latest

FROM golang:1.18.3-alpine as hotReload
WORKDIR /Q-n-A

RUN useradd -m -u 1000 Q-n-A
USER Q-n-A

COPY go.* .
RUN go mod download

COPY --from=builder /go/bin/air /usr/local/bin/
COPY --from=builder /go/bin/dockerize /usr/local/bin/

EXPOSE 9090
EXPOSE 9001
EXPOSE 6060

ENTRYPOINT ["dockerize", "-timeout", "60s", "-wait", "tcp://mariadb:3306", "air"]
