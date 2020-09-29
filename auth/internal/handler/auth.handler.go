package handler

import (
	"context"

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

func buildUserResponse(auth auth.Auth) *protoauth.ResponseAuth {
	return &protoauth.ResponseAuth{
		Id: auth.ID,

		CreatedAt:  auth.CreatedAt,
		ModifiedAt: auth.ModifiedAt,
		DeletedAt:  auth.DeletedAt,
	}
}

type AuthHandler struct{}

func (e *AuthHandler) LogIn(ctx context.Context, req *protoauth.RequestAuthLogIn, res *protoauth.ResponseLogIn) error {
	log.Info("Received auth.LogIn request")

	return nil
}

func (e *AuthHandler) AuthPath(ctx context.Context, req *protoauth.RequestAuthPath, res *protoauth.ResponseAuthPath) error {
	log.Info("Received auth.AuthPath request")

	return nil
}

func (e *AuthHandler) GetByID(ctx context.Context, req *protoauth.RequestAuthID, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.GetByID request")
	reqauth := new(auth.Auth)
	foundauth, err := reqauth.GetbyID(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE+
	res.Id = foundauth.ID

	res.CreatedAt = foundauth.CreatedAt
	res.ModifiedAt = foundauth.ModifiedAt
	res.DeletedAt = foundauth.DeletedAt

	return nil
}

func (e *AuthHandler) Create(ctx context.Context, req *protoauth.RequestCreateAuth, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.Create request")

	reqauth := &auth.Auth{
		/*Name:       req.Name,
		Phone:      req.Phone,
		Address:    req.Address,
		District:   req.District,
		Department: req.Department,
		Actions:    req.Actions,*/
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

	res.CreatedAt = createdauth.CreatedAt
	res.ModifiedAt = createdauth.ModifiedAt
	res.DeletedAt = createdauth.DeletedAt

	err = pubCreated.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthHandler) Update(ctx context.Context, req *protoauth.RequestUpdateAuth, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.Update request")
	reqauth := &auth.Auth{
		ID: req.Id,
		//User:       req.,

	}

	updatedauth, err := reqauth.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedauth.ID

	res.CreatedAt = updatedauth.CreatedAt
	res.ModifiedAt = updatedauth.ModifiedAt
	res.DeletedAt = updatedauth.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthHandler) Delete(ctx context.Context, req *protoauth.RequestAuthID, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.Delete request")
	reqauth := new(auth.Auth)
	deletedauth, err := reqauth.Delete(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = deletedauth.ID

	res.CreatedAt = deletedauth.CreatedAt
	res.ModifiedAt = deletedauth.ModifiedAt
	res.DeletedAt = deletedauth.DeletedAt

	err = pubDeleted.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}
