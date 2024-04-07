start:
	cd src/logs && git clean -xdf
	echo y | docker system prune -a
#	docker network create api-network
	docker compose up tls
	docker compose up setup
	docker compose up -d

stop:
	docker stop $$(docker ps -q)

apagar-logs:
	cd src/logs && git clean -xdf

build:
	cd src/logs && git clean -xdf
	cd src && docker compose up --build
