all: build run logs

build:
	docker build -t resin.io-go-server -f Dockerfile.local .

stop:
	docker stop go-server

remove:
	docker rm go-server

run:
	docker run -d -p 80:80 --name go-server -e DOWNLOAD_URL=http://ipv4.download.thinkbroadband.com/5MB.zip -e SAVE_FILE=/go/src/github.com/jooldesign/resin.io-speed-test/server/static/data.csv resin.io-go-server

logs:
	docker logs -f go-server