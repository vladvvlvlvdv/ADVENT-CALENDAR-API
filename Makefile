# Makefile

dev: swag
	MODE="dev" nodemon --exec "make swag && go run cmd/main.go" --signal SIGTERM

build:
	MODE="build" go build cmd/main.go && ./main

swag:
	swag init -g cmd/main.go

test:
	MODE="test" go test ./... -v