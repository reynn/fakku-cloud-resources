FROM golang:1.17 as BUILDER

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o app main.go

FROM debian as CLI_TOOLS

COPY bootstrap.sh /tmp/bootstrap.sh
ADD https://github.com/TomWright/dasel/releases/download/v1.22.1/dasel_linux_amd64 /tmp/dasel
ADD https://releases.hashicorp.com/vault/1.9.3/vault_1.9.3_linux_amd64.zip /tmp
RUN apt update \
    && apt install unzip \
    && unzip -d /tmp /tmp/vault_*_linux_amd64.zip \
    && chmod +x /tmp/*

FROM debian

COPY --from=BUILDER /app/app /usr/bin/app
COPY --from=CLI_TOOLS /tmp/vault /usr/bin/vault
COPY --from=CLI_TOOLS /tmp/dasel /usr/bin/dasel
COPY --from=CLI_TOOLS /tmp/bootstrap.sh /usr/bin/bootstrap.sh

CMD "/usr/bin/bootstrap.sh"
