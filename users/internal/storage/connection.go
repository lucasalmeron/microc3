package mongostorage

import (
	"context"
	"fmt"

	user "github.com/lucasalmeron/microc3/users/pkg/users"
	log "github.com/micro/go-micro/v2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgo *MongoDB

type MongoDB struct {
	connection *mongo.Database
	mongoURI   string
	database   string
}

func NewMongoDBConnection(mongoURI string, database string) error {

	mgo = new(MongoDB)

	mgo.mongoURI = mongoURI
	mgo.database = database

	if mongoURI == "" {
		mgo.mongoURI = fmt.Sprintf("mongodb://localhost:27017")
	}
	if database == "" {
		mgo.database = "cus3users"
	}

	err := mgo.connect()
	if err != nil {
		return err
	}
	//SET REPOSITORIES
	user.SetRepository(NewUserService(mgo.connection))

	return nil
}

func GetMongoDBConnection() *MongoDB {
	return mgo
}

func (mongoDB *MongoDB) connect() error {

	clientOpts := options.Client().ApplyURI(mgo.mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	mongoDB.connection = client.Database(mgo.database)

	log.Info("MongoDB connection success")

	return nil
}

func (mongoDB *MongoDB) CloseConnection() error {
	return mongoDB.connection.Client().Disconnect(context.TODO())
}
