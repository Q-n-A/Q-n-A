# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.194.3/containers/ubuntu/.devcontainer/base.Dockerfile

FROM golang:1.17.5 AS builder

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN go get github.com/k1LoW/tbls

FROM golang:1.17.5 AS devContainer

ENV TZ=Asia/Tokyo

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get install curl sudo git vim mariadb-client clang-format -y --no-install-recommends

COPY --from=builder /go/bin/air /usr/local/bin/
COPY --from=builder /go/bin/wire /usr/local/bin/
COPY --from=builder /go/bin/tbls /usr/local/bin/
