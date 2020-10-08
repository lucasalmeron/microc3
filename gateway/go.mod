module github.com/lucasalmeron/microc3/gateway

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/lucasalmeron/microc3/auth v0.0.0-20201008005346-f7a5420ddf9b
	github.com/lucasalmeron/microc3/users v0.0.0-20200930205716-85524c0bddd5
	github.com/micro/go-micro/v2 v2.9.1
)
