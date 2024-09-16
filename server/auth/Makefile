build:
	go build -o bin/auth

run: build
	./bin/auth

.PHONY: protos

protos:
		protoc -I protos/ --go_out=. --go-grpc_out=. protos/auth.proto


