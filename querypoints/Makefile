
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=. --go_out=. --go_opt=paths=source_relative pkg/querypoints/proto/querypoints.proto
    

.PHONY: build
build: proto

	go build -o querypoints-service cmd/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t querypoints-service:latest

run:
	go run cmd/main.go
