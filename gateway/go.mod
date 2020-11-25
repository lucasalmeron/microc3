module github.com/lucasalmeron/microc3/gateway

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/lucasalmeron/microc3/auth v0.0.0-20201111210833-6b5aaaab6010
	github.com/lucasalmeron/microc3/querypoints v0.0.0-20201111220146-9d2aaa4f288f
	github.com/lucasalmeron/microc3/users v0.0.0-20201111210833-6b5aaaab6010
	github.com/micro/go-micro/v2 v2.9.1
)
