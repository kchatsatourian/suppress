FROM golang:1.21.2 AS suppress

LABEL maintainer="Konstantinos Chatsatourian <kchatsatourian@gmail.com>"

WORKDIR /suppress/

COPY go.mod /suppress/
COPY go.sum /suppress/

RUN set -o xtrace && \
    go mod download

COPY suppress.go /suppress/

RUN set -o xtrace && \
    CGO_ENABLED=0 GOOS=linux go build -o /suppress/

FROM scratch

COPY --from=suppress /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=suppress /suppress/ /suppress/

CMD [ "/suppress/suppress" ]
