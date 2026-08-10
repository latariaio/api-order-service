# Go Dev Environment (Docker + SSH) 🐳🔑

Ambiente de desenvolvimento Go em container, acessível via SSH pelo VSCode
(Remote-SSH) ou pelo Zed (Remote Development)

✅ Postgres.

## 1. Gere um par de chaves SSH dedicado pra esse ambiente

⚠️ **Não reuse sua chave pessoal do GitHub — crie uma só pra containers de dev:**

Gere um par de chaves:
```bash
ssh-keygen -t ed25519 -f ./.devcontainer/id_ed25519 -C "go-dev-container" -N ""
```

Isso gera `id_ed25519` (privada, fica de fora do build) e `id_ed25519.pub`
(pública, é copiada pra dentro do container no Dockerfile).

Ou copie sua chave de containers local:
```bash 
cp ~/.ssh/sua_chave.pub .devcontainer/sua_chave.pub
```

⚠️ O `.gitignore` já ignora qualquer chave `.pub` que estiver em `.devcontainer`.

## 2. Suba os Containers

```bash
docker compose up -d --build
```

Isso sobe:
- `dev-container` — Container principal, porta `8080` da API
- `dev-ssh-container` — Container SSH na porta `2222`
- `postgres-container` — Postgres 18, porta `5432`, banco `app`

## 3. Configure seu `~/.ssh/config` (Opcional)

```bash
Host go-dev-container
    HostName localhost
    Port 2222
    User devuser
    IdentityFile /caminho/completo/para/sua/chave/privada
```

## 4. Conecte

**VSCode:**
1. Instale a extensão *Remote - SSH*
2. `Ctrl+Shift+P` → `Remote-SSH: Connect to Host` → `go-dev-container`
3. Abra a pasta `/workspace`

**Zed:**
1. `Ctrl+Shift+P` → `zed: remote` (ou o menu de Remote Development)
2. Adicione o host `go-dev-container` (ele lê do seu `~/.ssh/config`)
3. Abra a pasta `/workspace`

## 5. Teste a conexão manualmente (opcional, pra debugar)

```bash
ssh go-dev-container
```

Se cair no bash do `devuser` dentro do container, tá tudo certo ✅

## Notas

- `go mod` cache fica num volume nomeado (`go-mod-cache`), então não se
  perde entre rebuilds da imagem.
- Prometheus espera um endpoint `/metrics` na sua API rodando na porta
  `8080` — ajuste `.devcontainer/prometheus.yml` se sua API expõe métricas
  em outra porta/rota.
- Pra reconstruir do zero: `docker compose down -v && docker compose up -d --build`
