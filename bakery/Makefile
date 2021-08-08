.PHONY: build 

run: 
	go run main.go
build:
	go build main.go -o bakery
protos:
	buf generate
	
BUF_VERSION:=0.48.0
SQLC_VERSION:=1.8.0

install:
	go get \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc 
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Darwin-x86_64" \
		-o "$(shell go env GOPATH)/bin/buf"
	chmod +x "$(shell go env GOPATH)/bin/buf"
	curl -sSL \
		"https://github.com/kyleconroy/sqlc/releases/download/v$(SQLC_VERSION)/sqlc-v$(SQLC_VERSION)-darwin-amd64.zip" \
		-o sqlc.zip
	unzip -o -d "$(shell go env GOPATH)/bin/" sqlc.zip
	chmod +x "$(shell go env GOPATH)/bin/sqlc"
	rm sqlc.zip

sqlc:
	cd storage && sqlc generate