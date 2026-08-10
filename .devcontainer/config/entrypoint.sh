#!/usr/bin/env bash

set -e

mkdir -p /home/devuser/.ssh

# Limpa o authorized_keys antes de repopular, evitando acúmulo de chaves antigas
> /home/devuser/.ssh/authorized_keys

# Copia TODAS as chaves públicas (.pub) encontradas em .devcontainer/
shopt -s nullglob
for pubkey in /workspace/.devcontainer/*.pub; do
    cat "$pubkey" >> /home/devuser/.ssh/authorized_keys
done
shopt -u nullglob

chmod 700 /home/devuser/.ssh

if [ -f /home/devuser/.ssh/authorized_keys ]; then
    chmod 600 /home/devuser/.ssh/authorized_keys
fi

chown -R devuser:devuser /home/devuser/.ssh

exec "$@"