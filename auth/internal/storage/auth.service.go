package mongostorage

import (
	"context"
	"log"
	"time"

	auth "github.com/lucasalmeron/microc3/auth/pkg/auth"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthService struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewAuthService(db *mongo.Database) auth.Repository {
	return &AuthService{db, db.Collection("auth")}
}

func (service *AuthService) buildBsonObject(auth auth.Auth) bson.D {
	return bson.D{
		{"name", auth.Name},
		{"phone", auth.Phone},
		{"address", auth.Address},
		{"district", auth.District},
		{"department", auth.Department},
		{"responsibles", auth.Responsibles},
		{"actions", auth.Actions},
	}

}

func (service *AuthService) GetList(ctx context.Context) ([]auth.Auth, error) {

	cursor, err := service.collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []auth.Auth
	if err = cursor.All(ctx, &results); err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func (service *AuthService) GetByID(ctx context.Context, authID string) (*auth.Auth, error) {

	objectId, err := primitive.ObjectIDFromHex(authID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var auth auth.Auth
	err = service.collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&auth)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return &auth, nil
		}
		log.Println(err)
		return nil, err
	}

	return &auth, nil
}

func (service *AuthService) Create(ctx context.Context, reqauth auth.Auth) (*auth.Auth, error) {

	BSONObj := service.buildBsonObject(reqauth)
	BSONObj = append(BSONObj, bson.E{"createdAt", time.Now().Unix()})

	newauthID, err := service.collection.InsertOne(
		ctx,
		BSONObj,
	)
	if err != nil {
		return nil, err
	}
	reqauth.ID = newauthID.InsertedID.(primitive.ObjectID).Hex()
	return &reqauth, nil
}

func (service *AuthService) Update(ctx context.Context, reqauth auth.Auth) (*auth.Auth, error) {

	objectID, err := primitive.ObjectIDFromHex(reqauth.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var auth auth.Auth

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	BSONObj := service.buildBsonObject(reqauth)
	BSONObj = append(BSONObj, bson.E{Key: "modifiedAt", Value: time.Now().Unix()})

	err = service.collection.FindOneAndUpdate(
		ctx,
		bson.D{{"_id", objectID}},
		bson.M{
			"$set": BSONObj,
		},
		&opt,
	).Decode(&auth)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (service *AuthService) Delete(ctx context.Context, authID string) (*auth.Auth, error) {
	objectID, err := primitive.ObjectIDFromHex(authID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var auth auth.Auth

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	err = service.collection.FindOneAndUpdate(
		ctx,
		bson.D{{"_id", objectID}},
		bson.M{
			"$set": bson.M{"deletedAt": time.Now().Unix()},
		},
		&opt,
	).Decode(&auth)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
