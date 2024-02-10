FROM golang:1.23.3-alpine AS certificates

RUN --mount=type=cache,target=/var/cache/apk/ \
      apk update && \
      apk add ca-certificates

RUN --mount=type=cache,target=/etc/ssl/certs/ \
    update-ca-certificates

FROM golang:1.23.3-alpine AS suppress

WORKDIR /application/

RUN --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x

RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 GOOS=linux go build -o /suppress/ .

FROM scratch

LABEL maintainer="Konstantinos Chatsatourian <kchatsatourian@gmail.com>"

COPY --from=certificates --link /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=suppress --link /suppress/suppress /suppress/

ENTRYPOINT [ "/suppress/suppress" ]
