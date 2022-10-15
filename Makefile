build:
	go fmt ./...
	go build -o bin/tochknet cmd/tochknet/*.go

build-arm:
	go fmt ./...
	GOOS=linux GOARCH=arm64 go build -o bin/tochknet cmd/tochknet/*.go

generate:
	go generate ./...

run: generate
	go run cmd/tochknet/main.go -log=debug -cache

migrate:
	go run cmd/migrations/main.go up ./migrations

test:
	go test ./...

lint:
	golangci-lint run

all: generate build