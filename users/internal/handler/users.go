package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protousers "github.com/lucasalmeron/microc3/users/internal/proto"

	user "github.com/lucasalmeron/microc3/users/pkg/users"
)

var (
	pub micro.Publisher
)

func NewPublisher(c client.Client) {
	pub = micro.NewEvent("go.micro.users.created", c)
}

type UsersHandler struct{}

func (e *UsersHandler) GetUsers(ctx context.Context, req *empty.Empty, res *protousers.ResponseUsersArray) error {
	log.Info("Received Users.GetUsers request")
	reqUser := new(user.User)
	users, err := reqUser.GetUsers()
	if err != nil {
		log.Error(err)
		status.Error(codes.Internal, err.Error())
	}

	var response []*protousers.ResponseUser
	for _, u := range users {
		response = append(response, &protousers.ResponseUser{
			Id:             u.ID,
			FirstName:      u.FirstName,
			LastName:       u.LastName,
			DocumentNumber: u.DocumentNumber,
			Email:          u.Email,
			PhoneNumber:    u.PhoneNumber,
			GDEUser:        u.GDEUser,
			Position:       u.Position,
			CreatedAt:      u.CreatedAt,
			ModifiedAt:     u.ModifiedAt,
			DeletedAt:      u.DeletedAt,
		})
	}

	res.Users = response

	return nil
}

func (e *UsersHandler) GetUser(ctx context.Context, req *protousers.RequestUserID, res *protousers.ResponseUser) error {
	log.Info("Received Users.GetUser request")
	reqUser := new(user.User)
	foundUser, err := reqUser.GetUser(req.Id)
	if err != nil {
		log.Error(err)
		status.Error(codes.Internal, err.Error())
	}

	//RESPONSE+
	res.Id = foundUser.ID
	res.FirstName = foundUser.FirstName
	res.LastName = foundUser.LastName
	res.DocumentNumber = foundUser.DocumentNumber
	res.Email = foundUser.Email
	res.PhoneNumber = foundUser.PhoneNumber
	res.GDEUser = foundUser.GDEUser
	res.Position = foundUser.Position
	res.CreatedAt = foundUser.CreatedAt
	res.ModifiedAt = foundUser.ModifiedAt
	res.DeletedAt = foundUser.DeletedAt

	return nil
}

func (e *UsersHandler) CreateUser(ctx context.Context, req *protousers.RequestCreateUser, res *protousers.ResponseUser) error {
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

	createdUser, err := reqUser.Save()
	if err != nil {
		log.Error(err)
		status.Error(codes.Internal, err.Error())
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

	return nil
}

func (e *UsersHandler) UpdateUser(ctx context.Context, req *protousers.RequestUpdateUser, res *protousers.ResponseUser) error {
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
		status.Error(codes.Internal, err.Error())
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

	return nil
}

func (e *UsersHandler) DeleteUser(ctx context.Context, req *protousers.RequestUserID, res *protousers.ResponseUser) error {
	log.Info("Received Users.DeleteUser request")
	reqUser := new(user.User)
	deletedUser, err := reqUser.Delete(req.Id)
	if err != nil {
		log.Error(err)
		status.Error(codes.Internal, err.Error())
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

	return nil
}
