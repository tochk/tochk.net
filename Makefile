build:
	go fmt ./...
	GOOS=linux go build -o bin/tochkru cmd/tochkru/main.go

generate:
	go generate ./...

run:
	go run cmd/tochkru/main.go -log=debug

migrate:
	go run cmd/migrations/main.go up ./migrations

migrate_down:
	go run cmd/migrations/main.go down ./migrations


all: generate build