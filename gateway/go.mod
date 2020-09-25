module github.com/lucasalmeron/microc3/gateway

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gorilla/mux v1.7.3
	github.com/lucasalmeron/microc3/users v0.0.0-20200924000454-78f61c39f971
	github.com/micro/go-micro/v2 v2.9.1
)
