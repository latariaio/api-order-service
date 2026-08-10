FROM golang:1.26-bookworm

ARG UID=1000
ARG GID=1000

RUN apt-get update && apt-get install -y \
    bash \
    curl \
    git \
    openssh-server \
    postgresql-client \
    sudo \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Cria o usuário de desenvolvimento
RUN groupadd -g ${GID} devuser \
    && useradd -u ${UID} -g ${GID} -ms /bin/bash devuser \
    && echo "devuser ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && usermod -p '*' devuser

# Diretório necessário para o sshd
RUN mkdir -p /var/run/sshd

# Ferramentas Go
RUN go install golang.org/x/tools/gopls@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

ENV PATH="/go/bin:${PATH}"

COPY .devcontainer/config/sshd_config /etc/ssh/sshd_config
COPY .devcontainer/config/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

WORKDIR /workspace

RUN echo 'export PATH="/usr/local/go/bin:/go/bin:${PATH}"' > /etc/profile.d/golang.sh \
    && chmod +x /etc/profile.d/golang.sh

ENTRYPOINT ["/entrypoint.sh"]