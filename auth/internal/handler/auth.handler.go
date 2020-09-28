package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoauth "github.com/lucasalmeron/microc3/auth/pkg/auth/proto"

	auth "github.com/lucasalmeron/microc3/auth/pkg/auth"
)

var (
	pubCreated  micro.Event
	pubMofidied micro.Event
	pubDeleted  micro.Event
)

func InitEvents(c client.Client) {
	pubCreated = micro.NewEvent("go.micro.auth.created", c)
	pubMofidied = micro.NewEvent("go.micro.auth.modified", c)
	pubDeleted = micro.NewEvent("go.micro.auth.deleted", c)
}

func buildUserResponse(auth auth.Auth) *protoauth.Responseauth {
	return &protoauth.Responseauth{
		Id:         auth.ID,
		Name:       auth.Name,
		Phone:      auth.Phone,
		Address:    auth.Address,
		District:   auth.District,
		Department: auth.Department,
		Actions:    auth.Actions,
		CreatedAt:  auth.CreatedAt,
		ModifiedAt: auth.ModifiedAt,
		DeletedAt:  auth.DeletedAt,
	}
}

type AuthsHandler struct{}

func (e *AuthsHandler) GetList(ctx context.Context, req *empty.Empty, res *protoauth.ResponseauthsArray) error {
	log.Info("Received auths.GetList request")
	reqauth := new(auth.Auth)
	users, err := reqauth.GetList()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	var response []*protoauth.Responseauth
	for _, u := range users {
		response = append(response, buildUserResponse(u))
	}

	res.auths = response

	return nil
}

func (e *AuthsHandler) GetByID(ctx context.Context, req *protoauth.RequestauthID, res *protoauth.Responseauth) error {
	log.Info("Received auths.GetByID request")
	reqauth := new(auth.Auth)
	foundauth, err := reqauth.GetbyID(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE+
	res.Id = foundauth.ID
	res.Name = foundauth.Name
	res.Phone = foundauth.Phone
	res.Address = foundauth.Address
	res.District = foundauth.District
	res.Department = foundauth.Department
	res.Actions = foundauth.Actions
	res.CreatedAt = foundauth.CreatedAt
	res.ModifiedAt = foundauth.ModifiedAt
	res.DeletedAt = foundauth.DeletedAt

	return nil
}

func (e *AuthsHandler) Create(ctx context.Context, req *protoauth.RequestCreateauth, res *protoauth.Responseauth) error {
	log.Info("Received auth.Create request")

	reqauth := &auth.Auth{
		Name:       req.Name,
		Phone:      req.Phone,
		Address:    req.Address,
		District:   req.District,
		Department: req.Department,
		Actions:    req.Actions,
	}

	err := reqauth.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	createdauth, err := reqauth.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = createdauth.ID
	res.Name = createdauth.Name
	res.Phone = createdauth.Phone
	res.Address = createdauth.Address
	res.District = createdauth.District
	res.Department = createdauth.Department
	res.Actions = createdauth.Actions
	res.CreatedAt = createdauth.CreatedAt
	res.ModifiedAt = createdauth.ModifiedAt
	res.DeletedAt = createdauth.DeletedAt

	err = pubCreated.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthsHandler) Update(ctx context.Context, req *protoauth.RequestUpdateauth, res *protoauth.Responseauth) error {
	log.Info("Received auths.Update request")
	reqauth := &auth.Auth{
		ID:         req.Id,
		Name:       req.Name,
		Phone:      req.Phone,
		Address:    req.Address,
		District:   req.District,
		Department: req.Department,
		Actions:    req.Actions,
	}

	updatedauth, err := reqauth.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedauth.ID
	res.Name = updatedauth.Name
	res.Phone = updatedauth.Phone
	res.Address = updatedauth.Address
	res.District = updatedauth.District
	res.Department = updatedauth.Department
	res.Actions = updatedauth.Actions
	res.CreatedAt = updatedauth.CreatedAt
	res.ModifiedAt = updatedauth.ModifiedAt
	res.DeletedAt = updatedauth.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthsHandler) Delete(ctx context.Context, req *protoauth.RequestauthID, res *protoauth.Responseauth) error {
	log.Info("Received auth.Delete request")
	reqauth := new(auth.Auth)
	deletedauth, err := reqauth.Delete(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = deletedauth.ID
	res.Name = deletedauth.Name
	res.Phone = deletedauth.Phone
	res.Address = deletedauth.Address
	res.District = deletedauth.District
	res.Department = deletedauth.Department
	res.Actions = deletedauth.Actions
	res.CreatedAt = deletedauth.CreatedAt
	res.ModifiedAt = deletedauth.ModifiedAt
	res.DeletedAt = deletedauth.DeletedAt

	err = pubDeleted.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}
