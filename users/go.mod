module github.com/lucasalmeron/microc3/users

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/lucasalmeron/microc3/auth v0.0.0-20201102234245-b1986d0b8145
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.9.1
	go.mongodb.org/mongo-driver v1.4.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/tools v0.0.0-20200331025713-a30bf2db82d4 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)
