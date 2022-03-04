FROM golang:1.17.7 AS builder
ARG TARGETARCH
WORKDIR /temp

RUN apt-get update && apt-get install jq unzip -y --no-install-recommends

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN if [ "$TARGETARCH" = "amd64" ]; then PROTOC_ARCH="x86_64"; elif [[ "$TARGETARCH" = "arm"* ]]; then PROTOC_ARCH="aarch_64"; else exit 1; fi && \
  URI=$(wget -O - -q https://api.github.com/repos/protocolbuffers/protobuf/releases | \
  jq -r --arg arch "linux-$PROTOC_ARCH" '.[0].assets[] | select(.name | test($arch)) | .browser_download_url') && \
  wget --progress=dot:giga "$URI" -O "protobuf.zip" && \
  unzip -o protobuf.zip -d protobuf && \
  chmod -R 755 protobuf/*
ENV PATH $PATH:/temp/protobuf/bin

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.* .
RUN go mod download

COPY . .
RUN make grpc-go && \
  CGO_ENABLED=0 GOOS=linux GOARCH="$TARGETARCH" go build -o /Q-n-A -ldflags '-s -w'

FROM gcr.io/distroless/static AS runner

COPY --chown=nonroot:nonroot --from=builder /Q-n-A /
USER nonroot

EXPOSE 9000
EXPOSE 9001

HEALTHCHECK --interval=60s --timeout=3s --retries=5 CMD /Q-n-A healthcheck || exit 1
ENTRYPOINT ["/Q-n-A", "serve"]
