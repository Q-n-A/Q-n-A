FROM node:16.14.0 AS builder
WORKDIR /temp

RUN apt-get update && apt-get install jq unzip -y --no-install-recommends

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN URI=$(wget -O - -q https://api.github.com/repos/protocolbuffers/protobuf/releases | \
  jq -r '.[0].assets[] | select(.name | test("linux-x86_64.zip")) | .browser_download_url') && \
  wget --progress=dot:giga "$URI" -O "protobuf.zip" && \
  unzip -o protobuf.zip -d protobuf && \
  chmod -R 755 protobuf/*
ENV PATH $PATH:/temp/protobuf/bin

COPY ./Q-n-A_UI/package*.json .
RUN npm ci --unsafe-perm

COPY ./Q-n-A_UI .
RUN npm run gen && \
  npm run build

FROM caddy:2.4.6-alpine AS runner
EXPOSE 80

COPY --from=builder /temp/dist /usr/share/caddy
