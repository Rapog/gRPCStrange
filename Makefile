.PHONY: run

all: run

run:
	cd gRPCServer && docker-compose up
	go run client/main.go

server_run:
	cd gRPCServer && make

stop:
	docker ps -a -q | xargs docker stop

clear:
	docker ps -a -q | xargs docker rm