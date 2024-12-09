FROM golang:1.23.4

LABEL maintainer="Konstantinos Chatsatourian <kchatsatourian@gmail.com>"

RUN set -o xtrace && \
    apt-get update && \
    apt-get upgrade --assume-yes && \
    apt-get dist-upgrade --assume-yes && \
    apt-get install --assume-yes --no-install-recommends make && \
    apt-get autoremove --assume-yes && \
    apt-get autoclean && \
    apt-get clean && \
    rm --force --recursive /var/lib/apt/lists/*
