FROM golang:1.25.6

LABEL \
    org.opencontainers.image.description="Suppress is a simple RSS application for Telegram written in Go." \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.source="https://github.com/kchatsatourian/suppress" \
    org.opencontainers.image.title="suppress" \
    org.opencontainers.image.vendor="Konstantinos Chatsatourian"

RUN <<-EOF
    set -o xtrace
    apt-get update
    apt-get upgrade --assume-yes
    apt-get dist-upgrade --assume-yes
    apt-get install --assume-yes --no-install-recommends make
    apt-get autoremove --assume-yes
    apt-get autoclean
    apt-get clean
    rm --force --recursive /var/lib/apt/lists/*
EOF
