
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/lucasalmeron/microc3/querypoints/pkg/querypoints/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. pkg/querypoints/proto/querypoints.proto

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
