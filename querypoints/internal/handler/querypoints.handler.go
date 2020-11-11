package handler

import (
	"context"
	"fmt"
	"io"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoqp "github.com/lucasalmeron/microc3/querypoints/pkg/querypoints/proto"
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

func buildUserResponse(querypoint querypoint.QueryPoint) *protoqp.ResponseQueryPoint {
	return &protoqp.ResponseQueryPoint{
		Id:         querypoint.ID,
		Name:       querypoint.Name,
		Phone:      querypoint.Phone,
		Address:    querypoint.Address,
		District:   querypoint.District,
		Department: querypoint.Department,
		Actions:    querypoint.Actions,
		CreatedAt:  querypoint.CreatedAt,
		ModifiedAt: querypoint.ModifiedAt,
		DeletedAt:  querypoint.DeletedAt,
	}
}

type QueryPointsHandler struct{}

func (e *QueryPointsHandler) GetList(ctx context.Context, req *empty.Empty, res *protoqp.ResponseQueryPointsArray) error {
	log.Info("Received QueryPoints.GetList request")
	reqQueryPoint := new(querypoint.QueryPoint)
	queryPoints, err := reqQueryPoint.GetList()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	var response []*protoqp.ResponseQueryPoint
	for _, u := range queryPoints {
		response = append(response, buildUserResponse(u))
	}

	res.QueryPoints = response

	return nil
}

func (e *QueryPointsHandler) GetByID(ctx context.Context, req *protoqp.RequestQueryPointID, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received QueryPoints.GetByID request")
	reqQueryPoint := new(querypoint.QueryPoint)
	foundQueryPoint, err := reqQueryPoint.GetbyID(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE+
	res.Id = foundQueryPoint.ID
	res.Name = foundQueryPoint.Name
	res.Phone = foundQueryPoint.Phone
	res.Address = foundQueryPoint.Address
	res.District = foundQueryPoint.District
	res.Department = foundQueryPoint.Department
	res.Actions = foundQueryPoint.Actions
	res.CreatedAt = foundQueryPoint.CreatedAt
	res.ModifiedAt = foundQueryPoint.ModifiedAt
	res.DeletedAt = foundQueryPoint.DeletedAt

	return nil
}

func (e *QueryPointsHandler) GetByName(ctx context.Context, req *protoqp.RequestQueryPointQuery, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received QueryPoints.GetByName request")
	reqQueryPoint := new(querypoint.QueryPoint)
	foundQueryPoint, err := reqQueryPoint.GetbyName(req.Query)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE+
	res.Id = foundQueryPoint.ID
	res.Name = foundQueryPoint.Name
	res.Phone = foundQueryPoint.Phone
	res.Address = foundQueryPoint.Address
	res.District = foundQueryPoint.District
	res.Department = foundQueryPoint.Department
	res.Actions = foundQueryPoint.Actions
	res.CreatedAt = foundQueryPoint.CreatedAt
	res.ModifiedAt = foundQueryPoint.ModifiedAt
	res.DeletedAt = foundQueryPoint.DeletedAt

	return nil
}

func (e *QueryPointsHandler) GetByIDs(ctx context.Context, stream protoqp.QueryPoints_GetByIDsStream) error {
	log.Info("Received QueryPoints.GetByIDs request")
	reqQueryPoint := new(querypoint.QueryPoint)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error(err)
			return status.Error(codes.Internal, err.Error())
		}
		foundQueryPoint, err := reqQueryPoint.GetbyID(req.Id)
		if err != nil {
			log.Error(err)
			return status.Error(codes.Internal, err.Error())
		}

		err = stream.Send(&protoqp.ResponseQueryPoint{
			Id:         foundQueryPoint.ID,
			Name:       foundQueryPoint.Name,
			Phone:      foundQueryPoint.Phone,
			Address:    foundQueryPoint.Address,
			District:   foundQueryPoint.District,
			Department: foundQueryPoint.Department,
			Actions:    foundQueryPoint.Actions,
			CreatedAt:  foundQueryPoint.CreatedAt,
			ModifiedAt: foundQueryPoint.ModifiedAt,
			DeletedAt:  foundQueryPoint.DeletedAt,
		})
		if err != nil {
			log.Error(err)
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

func (e *QueryPointsHandler) GetPaginated(ctx context.Context, req *protoqp.RequestPageOptions, res *protoqp.ResponsePage) error {
	log.Info("Received QueryPoint.GetPaginated request")
	pageOptions := new(user.PageOptions)
	pageOptions.PageNumber = req.PageNumber
	pageOptions.RegistersNumber = req.RegistersNumber
	pageOptions.OrderBy.Field = req.OrderBy.Field
	pageOptions.OrderBy.Value = req.OrderBy.Value
	for _, filter := range req.Filters {
		pageOptions.Filters = append(pageOptions.Filters, user.Filter{filter.Field, filter.Value})
	}
	err := pageOptions.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

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
	log.Info("Received QueryPoint.Create request")

	reqQueryPoint := &querypoint.QueryPoint{
		Name:       req.Name,
		Phone:      req.Phone,
		Address:    req.Address,
		District:   req.District,
		Department: req.Department,
		Actions:    req.Actions,
	}

	err := reqQueryPoint.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	createdQueryPoint, err := reqQueryPoint.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = createdQueryPoint.ID
	res.Name = createdQueryPoint.Name
	res.Phone = createdQueryPoint.Phone
	res.Address = createdQueryPoint.Address
	res.District = createdQueryPoint.District
	res.Department = createdQueryPoint.Department
	res.Actions = createdQueryPoint.Actions
	res.CreatedAt = createdQueryPoint.CreatedAt
	res.ModifiedAt = createdQueryPoint.ModifiedAt
	res.DeletedAt = createdQueryPoint.DeletedAt

	err = pubCreated.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *QueryPointsHandler) Update(ctx context.Context, req *protoqp.RequestUpdateQueryPoint, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received QueryPoints.Update request")
	reqQueryPoint := &querypoint.QueryPoint{
		ID:         req.Id,
		Name:       req.Name,
		Phone:      req.Phone,
		Address:    req.Address,
		District:   req.District,
		Department: req.Department,
		Actions:    req.Actions,
	}

	err := reqQueryPoint.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	updatedQueryPoint, err := reqQueryPoint.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedQueryPoint.ID
	res.Name = updatedQueryPoint.Name
	res.Phone = updatedQueryPoint.Phone
	res.Address = updatedQueryPoint.Address
	res.District = updatedQueryPoint.District
	res.Department = updatedQueryPoint.Department
	res.Actions = updatedQueryPoint.Actions
	res.CreatedAt = updatedQueryPoint.CreatedAt
	res.ModifiedAt = updatedQueryPoint.ModifiedAt
	res.DeletedAt = updatedQueryPoint.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *QueryPointsHandler) Delete(ctx context.Context, req *protoqp.RequestQueryPointID, res *protoqp.ResponseQueryPoint) error {
	log.Info("Received QueryPoint.Delete request")
	reqQueryPoint := new(querypoint.QueryPoint)
	deletedQueryPoint, err := reqQueryPoint.Delete(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = deletedQueryPoint.ID
	res.Name = deletedQueryPoint.Name
	res.Phone = deletedQueryPoint.Phone
	res.Address = deletedQueryPoint.Address
	res.District = deletedQueryPoint.District
	res.Department = deletedQueryPoint.Department
	res.Actions = deletedQueryPoint.Actions
	res.CreatedAt = deletedQueryPoint.CreatedAt
	res.ModifiedAt = deletedQueryPoint.ModifiedAt
	res.DeletedAt = deletedQueryPoint.DeletedAt

	err = pubDeleted.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}
