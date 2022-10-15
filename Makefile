build:
	go fmt ./...
	go build -o bin/tochknet cmd/tochknet/*.go

build-arm:
	go fmt ./...
	GOOS=linux GOARCH=arm64 go build -o bin/tochknet cmd/tochknet/*.go

generate:
	go generate ./...

run: generate
	go run cmd/tochknet/main.go -log=debug

migrate:
	go run cmd/migrations/main.go up ./migrations

migrate_down:
	go run cmd/migrations/main.go down ./migrations

test:
	go test ./...

all: generate build