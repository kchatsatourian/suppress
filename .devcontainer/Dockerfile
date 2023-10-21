FROM golang:1.21.3

LABEL maintainer="Konstantinos Chatsatourian <kchatsatourian@gmail.com>"

RUN set -o xtrace && \
    apt-get update && \
    apt-get upgrade --assume-yes && \
    apt-get dist-upgrade --assume-yes && \
    apt-get install --assume-yes make && \
    apt-get autoremove --assume-yes && \
    apt-get autoclean && \
    apt-get clean
