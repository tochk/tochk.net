build:
	go fmt ./...
	GOOS=linux go build -o bin/tochkru cmd/tochkru/main.go

generate:
	go generate ./...

run:
	go run cmd/tochkru/main.go -log=debug

all: generate build