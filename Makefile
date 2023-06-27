proto:
	@if [ ! -d "../server-contract/protos" ]; then \
		echo "Make sure the server-contract project exists with the proto files."; \
	else \
		echo "Copying proto files..."; \
		cp -v ../server-contract/protos/types.proto ./protos/; \
		sed -i '' 's/^option\ go_package.*/option\ go_package\ =\ "sv.player\/protos\/pb";/g' ./protos/types.proto; \
		cp -v ../server-contract/protos/player.proto ./protos/service.proto; \
		sed -i '' 's/^option\ go_package.*/option\ go_package\ =\ "sv.player\/protos\/pb";/g' ./protos/service.proto; \
		echo "Generating proto..."; \
		protoc -I=${PROTO_SRC_DIR} --go-grpc_opt=require_unimplemented_servers=false --go_out=${PROTO_SRC_DIR} --go_opt=paths=source_relative \
			--go-grpc_out=${PROTO_SRC_DIR} --go-grpc_opt=paths=source_relative \
			${PROTO_SRC_DIR}/*.proto; \
		echo "Removing proto files..."; \
		rm -v ./protos/service.proto ./protos/types.proto; \
	fi

build:
	go build -o myapp

lint:
	golangci-lint run

run:
	go run main.go

test:
	go test -v ./...

.PHONY: proto build lint run test
