package mongostorage

import (
	"context"
	"fmt"
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

/*
ID   string `json:"id" bson:"_id,omitempty"`
	User string `json:"user" bson:"user"`

	Admin       bool `json:"admin" bson:"admin"`
	Permissions []struct {
		Read        bool   `json:"read" bson:"read"`
		Write       bool   `json:"write" bson:"write"`
		Responsible bool   `json:"responsible" bson:"responsible"`
		Query       bool   `json:"query" bson:"query"`
		Health      bool   `json:"health" bson:"health"`
		QueryPoint  string `json:"queryPoint" bson:"queryPoint"`
	} `json:"permissions" bson:"permissions"`

	ModifierUser string `json:"modifierUser" bson:"modifierUser"`
	CreatedAt    int64  `json:"createdAt" bson:"createdAt"`
	ModifiedAt   int64  `json:"modifiedAt" bson:"modifiedAt"`
	DeletedAt    int64  `json:"deletedAt" bson:"deletedAt"`
*/

func (service *AuthService) buildBsonObject(auth auth.Auth) (bson.D, error) {
	permissions := bson.A{}
	for _, permission := range auth.Permissions {
		permissions = append(permissions, bson.D{
			{"_id", primitive.NewObjectID()},
			{"access", permission.Access},
		})
	}
	userId, err := primitive.ObjectIDFromHex(auth.User)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bson.D{
		{"user", userId},
		{"admin", auth.Admin},
		{"permissions", permissions},
	}, nil
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

func (service *AuthService) GetByUserID(ctx context.Context, userID string) (*auth.Auth, error) {

	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var auth auth.Auth
	err = service.collection.FindOne(ctx, bson.D{{"user", objectId}}).Decode(&auth)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return &auth, nil
		}
		log.Println(err)
		return nil, err
	}

	return &auth, nil
}

func (service *AuthService) Create(ctx context.Context, reqAuth auth.Auth) (*auth.Auth, error) {

	BSONObj, err := service.buildBsonObject(reqAuth)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	BSONObj = append(BSONObj, bson.E{"createdAt", time.Now().Unix()})

	newauthID, err := service.collection.InsertOne(
		ctx,
		BSONObj,
	)
	if err != nil {
		return nil, err
	}
	reqAuth.ID = newauthID.InsertedID.(primitive.ObjectID).Hex()
	return &reqAuth, nil
}

func (service *AuthService) Update(ctx context.Context, reqAuth auth.Auth) (*auth.Auth, error) {

	objectID, err := primitive.ObjectIDFromHex(reqAuth.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var auth auth.Auth

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	BSONObj, err := service.buildBsonObject(reqAuth)
	if err != nil {
		log.Println(err)
		return nil, err
	}
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

func (service *AuthService) PushPermission(ctx context.Context, userID string, permission auth.Permission) (*auth.Auth, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
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
		bson.D{{"user", objectID}},
		bson.M{
			"$push": bson.M{"permissions": bson.D{
				{"_id", primitive.NewObjectID()},
				{"access", permission.Access},
			}},
		},
		&opt,
	).Decode(&auth)
	if err != nil {
		return nil, err
	}
	fmt.Println(auth)

	return &auth, nil
}

func (service *AuthService) DeletePermission(ctx context.Context, userID string, permissionID string) (*auth.Auth, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	permissionObjectID, err := primitive.ObjectIDFromHex(permissionID)
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
		bson.D{{"user", userObjectID}},
		bson.M{
			"$pull": bson.M{"permissions": bson.D{
				{"_id", permissionObjectID},
			}},
		},
		&opt,
	).Decode(&auth)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
