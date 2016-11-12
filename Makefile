build:
	docker build -t resin.io-go-server -f Dockerfile.local .

run:
	docker run --name go-server resin.io-go-server
