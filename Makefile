build:
	go fmt ./...
	GOOS=linux go build -o bin/tochkru cmd/tochkru/main.go

generate:
	go generate ./...

all: generate build