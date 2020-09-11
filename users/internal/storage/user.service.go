package mongostorage

import (
	"context"
	"log"

	user "github.com/lucasalmeron/microc3/users/pkg/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) user.Repository {
	return &UserService{db, db.Collection("users")}
}

func (service *UserService) buildBsonObject(user user.User) bson.M {
	return bson.M{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	}

}

func (service *UserService) GetUsers(ctx context.Context) ([]user.User, error) {

	cursor, err := service.collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []user.User
	if err = cursor.All(ctx, &results); err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func (service *UserService) GetUser(ctx context.Context, userID string) (*user.User, error) {

	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var user user.User
	err = service.collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (service *UserService) CreateUser(ctx context.Context, reqUser user.User) (*user.User, error) {

	newUserID, err := service.collection.InsertOne(
		ctx,
		service.buildBsonObject(reqUser),
	)
	if err != nil {
		return nil, err
	}
	reqUser.ID = newUserID.InsertedID.(primitive.ObjectID).Hex()
	return &reqUser, nil
}

func (service *UserService) UpdateUser(ctx context.Context, reqUser user.User) (*user.User, error) {

	objectID, err := primitive.ObjectIDFromHex(reqUser.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var user user.User

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err = service.collection.FindOneAndUpdate(
		ctx,
		bson.D{{"_id", objectID}},
		service.buildBsonObject(reqUser),
		&opt,
	).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) DeleteUser(ctx context.Context, userID string) error {
	return nil
}
