# Para subir o projevo via make


make new 

make build

O make new vai criar os containers de kafka, mongodb e o elk

O command make build vai criar 5 containers do projeto em go

Se o projeto já foi criado e precisa apenas subir make start


# Para subir o Docker

docker network create api-network

docker compose up tls

docker compose up setup


docker compose up -d


cd src 

docker compose up --build

# Primeiros passos

