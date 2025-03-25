# 🛒 Ecommerce Fullstack - Go + Vue 3

Este projeto é um boilerplate de e-commerce com backend em **Golang (Gin)** e frontend em **Vue 3 (Vite)**, utilizando **Docker** para orquestração.

---

## ✅ Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🚀 Como rodar o projeto localmente

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/ecommerce.git
cd ecommerce
```

### 2. Copie o arquivo `.env` do backend

Antes de rodar, copie o arquivo de exemplo com as variáveis de ambiente:

```bash
cp backend/.env.example backend/.env
```

⚠️ Verifique se os valores no `.env` estão corretos (como `DATABASE_URL`, JWT_SECRET, etc).

---

### 3. Suba os containers com Docker

Na raiz do projeto, rode:

```bash
docker-compose up --build
```

Esse comando vai:

- Subir o banco PostgreSQL
- Iniciar o backend em Go (porta `8080`)
- Iniciar o frontend em Vue 3 (porta `5173`)

---

### 4. Execute as migrations do banco

Após os containers estarem rodando, execute a migration manualmente:

```bash
docker exec -it ecommerce-backend sh -c "go run cmd/migration/main.go"
```

Você verá algo como:

```
✅ Migrations aplicadas com sucesso!
```

---

## 🌐 Acessos rápidos

- Frontend: [http://localhost:5173](http://localhost:5173)
- Backend (API): [http://localhost:8080](http://localhost:8080)

---

## 🛠 Estrutura do projeto

```
ecommerce/
├── backend/               # Backend em Go (Gin)
│   ├── cmd/               # Entrypoints: API e Migrations
│   ├── internal/          # Domínio (hexagonal)
│   ├── pkg/               # Conexão com banco de dados e Scripts SQL versionados
|   |   |── migrations/    # Arquivos das migrações
|   |   |── postgres.go    # Modulo de conexão com banco de dados
│   └── .env               # Configurações do backend
├── frontend/              # Frontend em Vue 3 (Vite)
├── docker-compose.yml
└── README.md
```

---

## 🧪 Comandos úteis

Rodar novamente as migrations:

```bash
docker exec -it ecommerce-backend sh -c "go run cmd/migration/main.go"
```

Derrubar e remover todos os containers/volumes:

```bash
docker-compose down -v
```

---

## 📦 Tecnologias utilizadas

- ⚙️ Golang 1.24 + Gin
- 🧱 PostgreSQL 17
- 🌐 Vue 3 + Vite
- 🐳 Docker + Docker Compose
- 🔄 Air (hot reload do backend)
- 🔐 JWT + Cookies HttpOnly
- 📂 Arquitetura Hexagonal

---

## 🤝 Contribuindo

Sinta-se à vontade para abrir issues, PRs ou sugestões.

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
