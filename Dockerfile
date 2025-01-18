FROM golang:1.24.5-alpine AS certificates

RUN --mount=type=cache,target=/var/cache/apk/ \
      apk update && \
      apk add ca-certificates

RUN --mount=type=cache,target=/etc/ssl/certs/ \
    update-ca-certificates

FROM golang:1.24.5-alpine AS suppress

WORKDIR /application/

RUN --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x

RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 GOOS=linux go build -o /suppress/ .

FROM scratch

LABEL \
    org.opencontainers.image.description="Suppress is a simple RSS application for Telegram written in Go." \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.source="https://github.com/kchatsatourian/suppress" \
    org.opencontainers.image.title="suppress" \
    org.opencontainers.image.vendor="Konstantinos Chatsatourian"

COPY --from=certificates --link /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=suppress --link /suppress/suppress /suppress/

ENTRYPOINT [ "/suppress/suppress" ]
