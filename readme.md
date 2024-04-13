# Sobre o projeto

O projeto vai buscar os arquivos da CVM (https://dados.cvm.gov.br/dados/) da pasta FI e salvar no mongodb atualmente são 402 milhões de registros de balacente, cda, extrato e outras informações dos fundos de investimentos

O projeto é utilizado para estudo e muita coisa foi e vai ser alterado com o tempo, por isso não estou criando os testes por enquanto


# Requisitos

Primeiro passo é renomear o arquivo env.Example para .env

É necessário ter docker instalado na máquin e a duas maneiras de subir via arquivo makefile ou rodando os comandos manualemente logo abaixo tem às duas maneiras


# Para subir via makefile


make new 

make build

O make new vai criar os containers de kafka, mongodb e o elk

O command make build vai criar 5 containers do projeto em go


Se o projeto já foi criado e precisa apenas subir make start e um make build


# Para subir via Docker


Subindo o ELK, mongodb e kafka

docker network create api-network

docker compose up tls

docker compose up setup

docker compose up -d

Agora é necessário subir o projeto 

cd src 

docker compose up --build

# Primeiros passos

Depois de subir o projeto precisa acessar a rota localhost:8080/api/v1/fundos/sincronizar como post

Vai ser ser criado um banco no mongo db como api_fundos e as filas vão começar a rodar, a primeira fila é a de fazer os download do arquivo  e depois vai ser executado a de processar o arquivo se no .env estiver como PERSISTENCIA estiver como local vai ser salvo todos os dados na hora caso contrário vai ser jogado para uma nova fila de persistência dos dados

