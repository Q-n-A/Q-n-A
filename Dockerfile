FROM logica0419/protoc-go:1.1.0 AS builder
WORKDIR /build
COPY . .

RUN make grpc-go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /Q-n-A -ldflags '-s -w'

FROM alpine:3.15.0 AS runner
EXPOSE 9000
EXPOSE 9001

COPY --from=builder /Q-n-A .

HEALTHCHECK --interval=60s --timeout=3s --retries=5 CMD ./Q-n-A healthcheck || exit 1
ENTRYPOINT ./Q-n-A serve
