module github.com/lucasalmeron/microc3/gateway

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/lucasalmeron/microc3/auth v0.0.0-20201029213517-44d0822db567
	github.com/lucasalmeron/microc3/querypoints v0.0.0-20201029162505-cf8b84ac2f10
	github.com/lucasalmeron/microc3/users v0.0.0-20201029213517-44d0822db567
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc v1.32.0
)
