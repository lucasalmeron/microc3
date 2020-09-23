module github.com/lucasalmeron/microc3/querypoints

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/lucasalmeron/microc3/users v0.0.0-20200923231205-cdef43fcb4be
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.9.1
	go.mongodb.org/mongo-driver v1.3.4
	google.golang.org/grpc v1.28.0
	google.golang.org/protobuf v1.25.0
)
