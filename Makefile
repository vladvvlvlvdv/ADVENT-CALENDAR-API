# Makefile

dev: swag
	MODE="dev" nodemon --exec "make swag && go run cmd/main.go" --signal SIGTERM

build:
	go build cmd/main.go && ./main

swag:
	swag init -g cmd/main.go
