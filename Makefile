proto:
	protoc --go_out=./pb --go-grpc_out=./pb ./protos/player.proto

build:
	go build -o myapp

lint:
	golangci-lint run

run:
	go run main.go

test:
	go test -v ./...

.PHONY: proto build lint run test
