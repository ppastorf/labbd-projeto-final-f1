# Projeto final Lab BD - Grupo 14

## Integrantes

Pedro Pastorello Fernandes - 10262502
Vinicius Leite Ribeiro - 
Rafael Polakieviz - 

## Dependências

1. [Docker](https://docs.docker.com/engine/install/) + [docker-compose](https://docs.docker.com/compose/install/)

### Backend
1. [Go]()
1. Dependências: comando `go get` na pasta `backend/`

### Frontend
1. [NodeJs](https://nodejs.org/en/download/)
1. Dependências: comando `npm install` na pasta `frontend/`

## Componentes

### Postgres + PgAdmin
Containers do Postgres e PgAdmin. No arquivo do compose você encontra sobre as portas abertas por cada um dos containers e os logins/senha de acesso usadas para se conectar a eles.

Na pasta `db/` encontra-se o script de inicialização do banco e na pasta `db/data/` encontram-se os arquivos de dados que são usados para popular o banco.

### Backend
Servidor escrito em Go, encontra se na pasta `backend/`.
Para compilar o backend use o comando `docker-compose build` na pasta raíz do repositório.

### Frontend
Site escrito em JavaScript com a framework Vue.js, seu código encontra-se na pasta `frontend/`.
Para buildar o frontend use o comando `npm run build` na pasta do frontend. Após buildar, serão gerados os arquivos estáticos do site na pasta `frontend/dist`. Essa pasta é montada pelo `docker-compose` para dentro do container do backend, que os serve para seus clientes.

## Instruções para deploy
Para subir o ambiente completo da aplicação rode os seguintes comandos em ordem:

1. `npm i npm-run-all` para instalar os comandos requeridos no package.json
1. `npm run build` na pasta `frontend/` para buildar os arquivos do frontend
1. `docker-compose build` para montar o container do backend
1. `docker-compose up -d` para subir o ambiente
1. Logar no container do pgAdmin e adicionar o servidor do Postgres (hostname: nome do container no `docker-compose.yaml`)
1. Depois de adicionado, você deve rodar o script de inicialização do banco. Para isso, no pgAdmin clique com o botão direito em cima do banco do Postgres que você adicionou e vá em CREATE script. Na aba que abrir, apague o código e cole o conteúdo do arquivo `data/script.sql` e execute.
1. Agora você pode acessar o sistema batendo em `localhost:8080` (porta exposta pelo container do backend)
