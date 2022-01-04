# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.194.3/containers/ubuntu/.devcontainer/base.Dockerfile

FROM logica0419/protoc-go:1.17.5 AS builder

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN go get github.com/k1LoW/tbls
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

FROM logica0419/protoc-go:1.17.5 AS devcontainer

ENV TZ=Asia/Tokyo

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get install curl sudo git vim mariadb-client clang-format -y --no-install-recommends

COPY --from=builder /go/bin/air /usr/local/bin/
COPY --from=builder /go/bin/wire /usr/local/bin/
COPY --from=builder /go/bin/tbls /usr/local/bin/
COPY --from=builder /go/bin/grpcurl /usr/local/bin/

EXPOSE 9090
EXPOSE 9001
