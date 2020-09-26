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

	protousers "github.com/lucasalmeron/microc3/users/pkg/users/proto"

	user "github.com/lucasalmeron/microc3/users/pkg/users"
)

var (
	pubCreated  micro.Event
	pubMofidied micro.Event
	pubDeleted  micro.Event
)

func InitEvents(c client.Client) {
	pubCreated = micro.NewEvent("go.micro.users.created", c)
	pubMofidied = micro.NewEvent("go.micro.users.modified", c)
	pubDeleted = micro.NewEvent("go.micro.users.deleted", c)
}

func buildUserResponse(user user.User) *protousers.ResponseUser {
	return &protousers.ResponseUser{
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

type UsersHandler struct{}

func (e *UsersHandler) GetUsers(ctx context.Context, req *empty.Empty, res *protousers.ResponseUsersArray) error {
	log.Info("Received Users.GetUsers request")
	reqUser := new(user.User)
	users, err := reqUser.GetList()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	var response []*protousers.ResponseUser
	for _, u := range users {
		response = append(response, buildUserResponse(u))
	}

	res.Users = response

	return nil
}

func (e *UsersHandler) GetUserByID(ctx context.Context, req *protousers.RequestUserID, res *protousers.ResponseUser) error {
	log.Info("Received Users.GetUser request")
	fmt.Println(req.Id)
	reqUser := new(user.User)
	foundUser, err := reqUser.GetbyID(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
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

func (e *UsersHandler) GetUserByEmail(ctx context.Context, req *protousers.RequestUserEmail, res *protousers.ResponseUser) error {
	log.Info("Received Users.GetUser request")
	reqUser := new(user.User)
	foundUser, err := reqUser.GetbyEmail(req.Email)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
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

func (e *UsersHandler) GetPaginatedUsers(ctx context.Context, req *protousers.RequestPageOptions, res *protousers.ResponsePage) error {
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

	reqUser := new(user.User)

	fmt.Println(pageOptions)

	paginatedUsers, err := reqUser.GetPaginated(pageOptions)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	paginatedUsers.CalcNumberOfPages(pageOptions)
	//RESPONSE
	res.Length = paginatedUsers.Length
	res.Data = make([]*protousers.ResponseUser, len(paginatedUsers.Data))
	res.PageNumber = paginatedUsers.PageNumber
	res.NumberOfPages = paginatedUsers.NumberOfPages

	for i, u := range paginatedUsers.Data {
		res.Data[i] = buildUserResponse(u)
	}

	return nil
}

func (e *UsersHandler) CreateUser(ctx context.Context, req *protousers.RequestCreateUser, res *protousers.ResponseUser) error {
	log.Info("Received Users.CreateUser request")
	fmt.Println(req)
	if req.Password != req.Repassword {
		return status.Error(codes.InvalidArgument, "Passwords do not match")
	}

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

	fmt.Println(reqUser)

	err := reqUser.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := reqUser.GetbyEmail(reqUser.Email)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if exist.ID != "" {
		return status.Error(codes.InvalidArgument, "User already exist")
	}
	err = reqUser.EncryptPassword()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
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

func (e *UsersHandler) DeleteUser(ctx context.Context, req *protousers.RequestUserID, res *protousers.ResponseUser) error {
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
