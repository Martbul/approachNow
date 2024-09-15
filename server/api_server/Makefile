swagger:
	@swagger generate spec -o swagger.yaml --scan-models

build:
	go build -o bin/api_server

run: build
	./bin/api_server
