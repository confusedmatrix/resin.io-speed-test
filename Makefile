all: build run

build:
	docker build -t resin.io-go-server -f Dockerfile.local .

stop:
	docker stop go-server

remove:
	docker rm go-server

run:
	docker run -d -p 8080:8080 --name go-server resin.io-go-server
