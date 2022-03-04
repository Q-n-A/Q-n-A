FROM golang:1.17.7 AS protoc
WORKDIR /temp

RUN apt-get update && apt-get install jq unzip -y --no-install-recommends

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN URI=$(wget -O - -q https://api.github.com/repos/protocolbuffers/protobuf/releases | \
  jq -r '.[0].assets[] | select(.name | test("linux-x86_64.zip")) | .browser_download_url') && \
  wget --progress=dot:giga "$URI" -O "protobuf.zip" && \
  unzip -o protobuf.zip -d protobuf && \
  chmod -R 755 protobuf/*

FROM golang:1.17.7 AS builder
WORKDIR /temp

COPY --from=protoc /temp/protobuf /temp/protobuf/
ENV PATH $PATH:/temp/protobuf/bin

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.* .
RUN go mod download

COPY . .
RUN make grpc-go && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /Q-n-A -ldflags '-s -w'

FROM gcr.io/distroless/static AS runner

EXPOSE 9000
EXPOSE 9001

COPY --chown=nonroot:nonroot --from=builder /Q-n-A /
USER nonroot

HEALTHCHECK --interval=60s --timeout=3s --retries=5 CMD /Q-n-A healthcheck || exit 1
ENTRYPOINT ["/Q-n-A", "serve"]
