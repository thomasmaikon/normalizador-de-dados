# 📦 Desafio HUBLA

Desafio programacao full-stack para implementacao de plataforma normalizadora de transacoes encontrados em registros

## 🚀 Tecnologias Utilizadas

- Backend:
  - Golang (1.8)
  - Gin (Framework Web em Golang)
  - 🔐 JWT (JSON Web Tokens para autenticação)
  - 🧪 Teste de integração, unidade e E2E (com uso de Docker para o banco de dados)
  - 🗄️ Banco de dados PostgreSQL
- Frontend:
  - Next.js (v16)
  - Bootstrap 5 (Framework CSS)
  - 🍪 js-cookie (manipulação de cookies no lado do cliente)
  - 📢 react-toastify (exibição de notificações na interface)
- DevOps:
  - 🐳 Docker Compose (para orquestração do ambiente de desenvolvimento)
  - 🐳 Dockerfile (para construir as imagens do backend e frontend)

## ✅ Pré-requisitos

- 🐹 Golang (versão 1.8 ou superior) instalado
- 🟢 Node.js (versão 16 ou superior) e npm instalados
- 🐳 Docker e Docker Compose instalados

## ⚙️ Instalação e Execução

- Acesse o diretorio raiz do projeto e execute
``docker-compose up`` , adicione -d caso nao queria ver logs
- Acesse o endereco `` localhost:3000 `` do seu navegador
- Ao acessar cadastre um usuario, o login nao precisa ser genuinamente valido com @ e etc...

### Backend - Caso queria rodar manualmente o projeto
1. Acesse o diretório do backend: ``cd backend``
2. Instale as dependências do Golang: ``go mod download``
3. Execute o aplicativo: ``go run main.go``, lembre-se de trocar o host do banco para localhost no arquivo ``config.env``
4. Tenha o  banco de dados PostgreSQL rodando com as seguintes especificacoes: 
   1. port=5432
   2. user=postgres
   3. password=pgadmin
   4. dbname=desafio
   5. host=localhost

### Frontend - Caso queria rodar manualmente o projeto
1. Acesse o diretório do frontend: ``cd frontend``
2. Instale as dependências: ``npm install``
3. Execute o aplicativo: ``npm run dev``
4. O aplicativo estará disponível em: http://localhost:3000