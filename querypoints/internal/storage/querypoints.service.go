package mongostorage

import (
	"context"
	"log"
	"time"

	querypoint "github.com/lucasalmeron/microc3/querypoints/pkg/querypoints"
	user "github.com/lucasalmeron/microc3/users/pkg/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryPointService struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewQueryPointService(db *mongo.Database) querypoint.Repository {
	return &QueryPointService{db, db.Collection("querypoints")}
}

func (service *QueryPointService) buildBsonObject(querypoint querypoint.QueryPoint) bson.D {
	return bson.D{
		{"name", querypoint.Name},
		{"phone", querypoint.Phone},
		{"address", querypoint.Address},
		{"district", querypoint.District},
		{"department", querypoint.Department},
		{"actions", querypoint.Actions},
	}

}

func (service *QueryPointService) GetList(ctx context.Context) ([]querypoint.QueryPoint, error) {

	cursor, err := service.collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []querypoint.QueryPoint
	if err = cursor.All(ctx, &results); err != nil {
		log.Println(err)
		return nil, err
	}

	return results, nil
}

func (service *QueryPointService) GetByID(ctx context.Context, queryPointID string) (*querypoint.QueryPoint, error) {

	objectId, err := primitive.ObjectIDFromHex(queryPointID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var queryPoint querypoint.QueryPoint
	err = service.collection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&queryPoint)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return &queryPoint, nil
		}
		log.Println(err)
		return nil, err
	}

	return &queryPoint, nil
}

func (service *QueryPointService) GetPaginated(ctx context.Context, pageOptions *user.PageOptions) (*querypoint.Page, error) {

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
	var results []querypoint.Page
	if err = cursor.All(ctx, &results); err != nil {
		log.Println(err)
		return nil, err
	}

	if len(results) == 0 {
		return &querypoint.Page{}, nil
	}

	return &results[0], nil
}

func (service *QueryPointService) Create(ctx context.Context, reqQueryPoint querypoint.QueryPoint) (*querypoint.QueryPoint, error) {

	BSONObj := service.buildBsonObject(reqQueryPoint)
	BSONObj = append(BSONObj, bson.E{"createdAt", time.Now().Unix()})

	newQueryPointID, err := service.collection.InsertOne(
		ctx,
		BSONObj,
	)
	if err != nil {
		return nil, err
	}
	reqQueryPoint.ID = newQueryPointID.InsertedID.(primitive.ObjectID).Hex()
	return &reqQueryPoint, nil
}

func (service *QueryPointService) Update(ctx context.Context, reqQueryPoint querypoint.QueryPoint) (*querypoint.QueryPoint, error) {

	objectID, err := primitive.ObjectIDFromHex(reqQueryPoint.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var querypoint querypoint.QueryPoint

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	BSONObj := service.buildBsonObject(reqQueryPoint)
	BSONObj = append(BSONObj, bson.E{Key: "modifiedAt", Value: time.Now().Unix()})

	err = service.collection.FindOneAndUpdate(
		ctx,
		bson.D{{"_id", objectID}},
		bson.M{
			"$set": BSONObj,
		},
		&opt,
	).Decode(&querypoint)
	if err != nil {
		return nil, err
	}

	return &querypoint, nil
}

func (service *QueryPointService) Delete(ctx context.Context, querypointID string) (*querypoint.QueryPoint, error) {
	objectID, err := primitive.ObjectIDFromHex(querypointID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var querypoint querypoint.QueryPoint

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
	).Decode(&querypoint)
	if err != nil {
		return nil, err
	}

	return &querypoint, nil
}
