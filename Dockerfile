FROM golang:1.21.5 AS suppress

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

COPY --from=suppress /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=suppress /suppress/ /suppress/

ENTRYPOINT [ "/suppress/suppress" ]
