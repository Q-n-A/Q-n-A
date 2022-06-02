# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.194.3/containers/ubuntu/.devcontainer/base.Dockerfile

FROM golang:1.18.3 AS builder
WORKDIR /temp

RUN apt-get update && apt-get install jq unzip -y --no-install-recommends

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN URI=$(wget -O - -q https://api.github.com/repos/protocolbuffers/protobuf/releases | \
  jq -r '.[0].assets[] | select(.name | test("linux-x86_64.zip")) | .browser_download_url') && \
  wget --progress=dot:giga "$URI" -O "protobuf.zip" && \
  unzip -o protobuf.zip -d protobuf && \
  chmod -R 755 protobuf/*

RUN URI=$(wget -O - -q https://api.github.com/repos/hadolint/hadolint/releases | \
  jq -r '.[0].assets[] | select(.name | test("Linux-x86_64")) | .browser_download_url') && \
  wget --progress=dot:giga "$URI" -O "hadolint" && \
  chmod -R 755 hadolint

RUN  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
  go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest && \
  go install github.com/cosmtrek/air@latest && \
  go install github.com/google/wire/cmd/wire@latest && \
  go install github.com/k1LoW/tbls@main && \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

FROM golang:1.18.3 AS devcontainer

ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get install git mariadb-client clang-format gnupg2 -y --no-install-recommends && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

COPY --from=builder /temp/protobuf/bin /usr/local/bin/
COPY --from=builder /temp/protobuf/include /usr/local/include/
COPY --from=builder /temp/hadolint /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-go /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-go-grpc /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-doc /usr/local/bin/
COPY --from=builder /go/bin/air /usr/local/bin/
COPY --from=builder /go/bin/wire /usr/local/bin/
COPY --from=builder /go/bin/tbls /usr/local/bin/
COPY --from=builder /go/bin/grpcurl /usr/local/bin/

EXPOSE 9090
EXPOSE 9001
EXPOSE 6060
