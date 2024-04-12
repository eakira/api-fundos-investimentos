# Para subir o projevo via make

make new 

make build

O make new vai criar os containers de kafka, mongodb e o elk

O command make build vai criar 5 containers do projeto em go


Se o projeto já foi criado e precisa apenas subir make start e um make build


# Para subir o Docker

docker network create api-network

docker compose up tls

docker compose up setup


docker compose up -d


cd src 

docker compose up --build

# Primeiros passos

Depois de subir o projeto precisa acessar a rota localhost:8080/api/v1/fundos/sincronizar como post

Vai ser ser criado um banco no mongo db como api_fundos e lá vai trazer todos os arquivos