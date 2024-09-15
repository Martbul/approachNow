build:
	go build -o bin/near_users

run: build
	./bin/near_users

.PHONY: protos

protos:
		protoc -I protos/ --go_out=. --go-grpc_out=. protos/near_users.proto


