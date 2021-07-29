.PHONY: build 

run: 
	go run main.go
build:
	go build main.go -o bakery
protos:
	buf generate
#	buf protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/bakery/v1/bakery.proto