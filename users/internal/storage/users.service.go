package mongostorage

import (
	"context"
	"log"
	"time"

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

func (service *UserService) buildBsonObject(user user.User) bson.D {
	return bson.D{
		{"firstName", user.FirstName},
		{"lastName", user.LastName},
		{"documentNumber", user.DocumentNumber},
		{"password", user.Password},
		{"email", user.Email},
		{"phoneNumber", user.PhoneNumber},
		{"GDEUser", user.GDEUser},
		{"position", user.Position},
	}

}

func (service *UserService) GetList(ctx context.Context) ([]user.User, error) {

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

func (service *UserService) GetByID(ctx context.Context, userID string) (*user.User, error) {

	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var user user.User
	err = service.collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&user)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return &user, nil
		}
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (service *UserService) GetByEmail(ctx context.Context, email string) (*user.User, error) {

	var user user.User
	err := service.collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return &user, nil
		}
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (service *UserService) GetPaginated(ctx context.Context, pageOptions *user.PageOptions) (*user.Page, error) {

	queryMap := make(map[string]primitive.A)
	queryMap["id"] = bson.A{}
	queryMap["firstName"] = bson.A{}
	queryMap["lastName"] = bson.A{}
	queryMap["documentNumber"] = bson.A{}
	queryMap["email"] = bson.A{}
	queryMap["position"] = bson.A{}
	queryAnd := bson.A{}

	if len(pageOptions.Filters) > 0 {
		for _, filter := range pageOptions.Filters {
			switch filter.Field {
			case "id":
				objectId, err := primitive.ObjectIDFromHex(filter.Value)
				if err != nil {
					log.Println(err)
					return nil, err
				}
				queryMap["id"] = append(queryMap["id"], bson.D{{"_id", objectId}})
			case "firstName":
				queryMap["firstName"] = append(queryMap["firstName"], bson.D{{"firstName", filter.Value}})
			case "lastName":
				queryMap["lastName"] = append(queryMap["lastName"], bson.D{{"lastName", filter.Value}})
			case "documentNumber":
				queryMap["documentNumber"] = append(queryMap["documentNumber"], bson.D{{"documentNumber", filter.Value}})
			case "email":
				queryMap["email"] = append(queryMap["email"], bson.D{{"email", filter.Value}})
			case "position":
				queryMap["position"] = append(queryMap["position"], bson.D{{"position", filter.Value}})
			case "queryPoint":

			}
		}
		for _, value := range queryMap {
			if len(value) > 0 {
				queryAnd = append(queryAnd, bson.D{{"$or", value}})
			}
		}
	}

	buildBSONOrderBy := func(field string, order string) primitive.D {
		if order == "asc" {
			return bson.D{{"$sort", bson.D{{field, 1}}}}
		}
		return bson.D{{"$sort", bson.D{{field, -1}}}}
	}

	var sortStage primitive.D
	switch pageOptions.OrderBy.Field {
	case "firstName":
		sortStage = buildBSONOrderBy("firstName", pageOptions.OrderBy.Value)
	case "lastName":
		sortStage = buildBSONOrderBy("lastName", pageOptions.OrderBy.Value)
	case "documentNumber":
		sortStage = buildBSONOrderBy("documentNumber", pageOptions.OrderBy.Value)
	case "email":
		sortStage = buildBSONOrderBy("email", pageOptions.OrderBy.Value)
	case "createdAt":
		sortStage = buildBSONOrderBy("createdAt", pageOptions.OrderBy.Value)
	case "modifiedAt":
		sortStage = buildBSONOrderBy("modifiedAt", pageOptions.OrderBy.Value)
	default:
		sortStage = buildBSONOrderBy("createdAt", "desc")
	}

	mongoPipeline := mongo.Pipeline{}

	projectStage1 := bson.D{{"$project", bson.M{
		"_id":            "$_id",
		"firstName":      "$firstName",
		"lastName":       "$lastName",
		"documentNumber": "$documentNumber",
		"password":       "$password",
		"email":          "$email",
		"phoneNumber":    "$phoneNumber",
		"GDEUser":        "$GDEUser",
		"position":       "$position",
		"createdAt":      "$createdAt",
		"modifiedAt":     "$modifiedAt",
		"deletedAt":      "$deletedAt",
	}}}
	mongoPipeline = append(mongoPipeline, projectStage1)

	if len(pageOptions.Filters) > 0 {
		matchStage := bson.D{{"$match", bson.D{{"$and", queryAnd}}}}
		mongoPipeline = append(mongoPipeline, matchStage)
	}

	mongoPipeline = append(mongoPipeline, sortStage)

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", nil},
			{"count", bson.D{{"$sum", 1}}},
			{"users", bson.D{{"$push",
				bson.M{
					"_id":            "$_id",
					"firstName":      "$firstName",
					"lastName":       "$lastName",
					"documentNumber": "$documentNumber",
					"password":       "$password",
					"email":          "$email",
					"phoneNumber":    "$phoneNumber",
					"GDEUser":        "$GDEUser",
					"position":       "$position",
					"createdAt":      "$createdAt",
					"modifiedAt":     "$modifiedAt",
					"deletedAt":      "$deletedAt",
				},
			}},
			}},
		}}
	mongoPipeline = append(mongoPipeline, groupStage)

	projectStage2 := bson.D{{"$project", bson.D{
		{"data", bson.D{{"$slice", bson.A{"$users", pageOptions.RegistersNumber * (pageOptions.PageNumber - 1), pageOptions.RegistersNumber}}}},
		{"length", "$count"},
	}}}

	mongoPipeline = append(mongoPipeline, projectStage2)

	cursor, err := service.collection.Aggregate(ctx, mongoPipeline)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []user.Page
	if err = cursor.All(ctx, &results); err != nil {
		log.Println(err)
		return nil, err
	}

	if len(results) == 0 {
		return &user.Page{}, nil
	}

	return &results[0], nil
}

func (service *UserService) Create(ctx context.Context, reqUser user.User) (*user.User, error) {

	BSONObj := service.buildBsonObject(reqUser)
	BSONObj = append(BSONObj, bson.E{"createdAt", time.Now().Unix()})

	newUserID, err := service.collection.InsertOne(
		ctx,
		BSONObj,
	)

	if err != nil {
		return nil, err
	}
	reqUser.ID = newUserID.InsertedID.(primitive.ObjectID).Hex()
	return &reqUser, nil
}

func (service *UserService) Update(ctx context.Context, reqUser user.User) (*user.User, error) {

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

	BSONObj := service.buildBsonObject(reqUser)
	BSONObj = append(BSONObj, bson.E{Key: "modifiedAt", Value: time.Now().Unix()})

	err = service.collection.FindOneAndUpdate(
		ctx,
		bson.D{{"_id", objectID}},
		bson.M{
			"$set": BSONObj,
		},
		&opt,
	).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) Delete(ctx context.Context, userID string) (*user.User, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
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
		bson.M{
			"$set": bson.M{"deletedAt": time.Now().Unix()},
		},
		&opt,
	).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
