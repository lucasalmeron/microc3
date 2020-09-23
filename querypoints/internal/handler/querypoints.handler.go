package handler

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoqp "github.com/lucasalmeron/microc3/querypoints/internal/proto"
	user "github.com/lucasalmeron/microc3/users/pkg/users"

	querypoint "github.com/lucasalmeron/microc3/querypoints/pkg/querypoints"
)

var (
	pubCreated  micro.Event
	pubMofidied micro.Event
	pubDeleted  micro.Event
)

func InitEvents(c client.Client) {
	pubCreated = micro.NewEvent("go.micro.querypoints.created", c)
	pubMofidied = micro.NewEvent("go.micro.querypoints.modified", c)
	pubDeleted = micro.NewEvent("go.micro.querypoints.deleted", c)
}

func buildUserResponse(user querypoint.QueryPoint) *protoqp.ResponseQueryPoint {
	return &protoqp.ResponseQueryPoint{
		Id:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		DocumentNumber: user.DocumentNumber,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		GDEUser:        user.GDEUser,
		Position:       user.Position,
		CreatedAt:      user.CreatedAt,
		ModifiedAt:     user.ModifiedAt,
		DeletedAt:      user.DeletedAt,
	}
}

type QueryPointsHandler struct{}

func (e *QueryPointsHandler) GetList(ctx context.Context, req *empty.Empty, res *protoqp.ResponseQueryPointsArray) error {
	log.Info("Received Users.GetUsers request")
	reqQueryPoint := new(querypoint.QueryPoint)
	users, err := reqQueryPoint.GetList()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	var response []*protoqp.ResponseQueryPoint
	for _, u := range users {
		response = append(response, buildUserResponse(u))
	}

	res.QueryPoints = response

	return nil
}

func (e *QueryPointsHandler) GetByID(ctx context.Context, req *protoqp.RequestQueryPointID, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received Users.GetUser request")
	reqQueryPoint := new(querypoint.QueryPoint)
	foundQueryPoint, err := reqQueryPoint.GetbyID(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE+
	res.Id = foundQueryPoint.ID
	res.FirstName = foundQueryPoint.FirstName
	res.LastName = foundQueryPoint.LastName
	res.DocumentNumber = foundQueryPoint.DocumentNumber
	res.Email = foundQueryPoint.Email
	res.PhoneNumber = foundQueryPoint.PhoneNumber
	res.GDEUser = foundQueryPoint.GDEUser
	res.Position = foundQueryPoint.Position
	res.CreatedAt = foundQueryPoint.CreatedAt
	res.ModifiedAt = foundQueryPoint.ModifiedAt
	res.DeletedAt = foundQueryPoint.DeletedAt

	return nil
}

func (e *QueryPointsHandler) GetPaginated(ctx context.Context, req *protoqp.RequestPageOptions, res *protoqp.ResponsePage) error {
	log.Info("Received Users.GetPaginatedUsers request")
	pageOptions := new(user.PageOptions)
	pageOptions.PageNumber = req.PageNumber
	pageOptions.RegistersNumber = req.RegistersNumber
	pageOptions.OrderBy.Field = req.OrderBy.Field
	pageOptions.OrderBy.Value = req.OrderBy.Value
	for _, filter := range req.Filters {
		pageOptions.Filters = append(pageOptions.Filters, user.Filter{filter.Field, filter.Value})
	}
	pageOptions.Validate()

	reqQueryPoint := new(querypoint.QueryPoint)

	fmt.Println(pageOptions)

	paginated, err := reqQueryPoint.GetPaginated(pageOptions)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	paginated.CalcNumberOfPages(pageOptions)
	//RESPONSE
	res.Length = paginated.Length
	res.Data = make([]*protoqp.ResponseQueryPoint, len(paginated.Data))
	res.PageNumber = paginated.PageNumber
	res.NumberOfPages = paginated.NumberOfPages

	for i, u := range paginated.Data {
		res.Data[i] = buildUserResponse(u)
	}

	return nil
}

func (e *QueryPointsHandler) Create(ctx context.Context, req *protoqp.RequestCreateQueryPoint, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received Users.CreateUser request")

	reqUser := &user.User{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DocumentNumber: req.DocumentNumber,
		Password:       req.Password,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		GDEUser:        req.GDEUser,
		Position:       req.Position,
	}

	err := reqUser.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	createdUser, err := reqUser.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = createdUser.ID
	res.FirstName = createdUser.FirstName
	res.LastName = createdUser.LastName
	res.DocumentNumber = createdUser.DocumentNumber
	res.Email = createdUser.Email
	res.PhoneNumber = createdUser.PhoneNumber
	res.GDEUser = createdUser.GDEUser
	res.Position = createdUser.Position
	res.CreatedAt = createdUser.CreatedAt
	res.ModifiedAt = createdUser.ModifiedAt
	res.DeletedAt = createdUser.DeletedAt

	err = pubCreated.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *QueryPointsHandler) Update(ctx context.Context, req *protoqp.RequestUpdateQueryPoint, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received Users.UpdateUser request")
	reqUser := &user.User{
		ID:             req.Id,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DocumentNumber: req.DocumentNumber,
		Password:       req.Password,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		GDEUser:        req.GDEUser,
		Position:       req.Position,
	}

	updatedUser, err := reqUser.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedUser.ID
	res.FirstName = updatedUser.FirstName
	res.LastName = updatedUser.LastName
	res.DocumentNumber = updatedUser.DocumentNumber
	res.Email = updatedUser.Email
	res.PhoneNumber = updatedUser.PhoneNumber
	res.GDEUser = updatedUser.GDEUser
	res.Position = updatedUser.Position
	res.CreatedAt = updatedUser.CreatedAt
	res.ModifiedAt = updatedUser.ModifiedAt
	res.DeletedAt = updatedUser.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *QueryPointsHandler) Delete(ctx context.Context, req *protoqp.RequestQueryPointID, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received Users.DeleteUser request")
	reqUser := new(user.User)
	deletedUser, err := reqUser.Delete(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = deletedUser.ID
	res.FirstName = deletedUser.FirstName
	res.LastName = deletedUser.LastName
	res.DocumentNumber = deletedUser.DocumentNumber
	res.Email = deletedUser.Email
	res.PhoneNumber = deletedUser.PhoneNumber
	res.GDEUser = deletedUser.GDEUser
	res.Position = deletedUser.Position
	res.CreatedAt = deletedUser.CreatedAt
	res.ModifiedAt = deletedUser.ModifiedAt
	res.DeletedAt = deletedUser.DeletedAt

	err = pubDeleted.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}
