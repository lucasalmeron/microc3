// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: internal/proto/users.proto

package go_micro_service_users

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Users service

func NewUsersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Users service

type UsersService interface {
	GetUsers(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseUsersArray, error)
	GetUser(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseUser, error)
	CreateUser(ctx context.Context, in *RequestCreateUser, opts ...client.CallOption) (*ResponseUser, error)
	UpdateUser(ctx context.Context, in *RequestUpdateUser, opts ...client.CallOption) (*ResponseUser, error)
	DeleteUser(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseUser, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) GetUsers(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseUsersArray, error) {
	req := c.c.NewRequest(c.name, "Users.GetUsers", in)
	out := new(ResponseUsersArray)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) GetUser(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Users.GetUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) CreateUser(ctx context.Context, in *RequestCreateUser, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Users.CreateUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) UpdateUser(ctx context.Context, in *RequestUpdateUser, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Users.UpdateUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) DeleteUser(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "Users.DeleteUser", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	GetUsers(context.Context, *empty.Empty, *ResponseUsersArray) error
	GetUser(context.Context, *RequestUserID, *ResponseUser) error
	CreateUser(context.Context, *RequestCreateUser, *ResponseUser) error
	UpdateUser(context.Context, *RequestUpdateUser, *ResponseUser) error
	DeleteUser(context.Context, *RequestUserID, *ResponseUser) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		GetUsers(ctx context.Context, in *empty.Empty, out *ResponseUsersArray) error
		GetUser(ctx context.Context, in *RequestUserID, out *ResponseUser) error
		CreateUser(ctx context.Context, in *RequestCreateUser, out *ResponseUser) error
		UpdateUser(ctx context.Context, in *RequestUpdateUser, out *ResponseUser) error
		DeleteUser(ctx context.Context, in *RequestUserID, out *ResponseUser) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) GetUsers(ctx context.Context, in *empty.Empty, out *ResponseUsersArray) error {
	return h.UsersHandler.GetUsers(ctx, in, out)
}

func (h *usersHandler) GetUser(ctx context.Context, in *RequestUserID, out *ResponseUser) error {
	return h.UsersHandler.GetUser(ctx, in, out)
}

func (h *usersHandler) CreateUser(ctx context.Context, in *RequestCreateUser, out *ResponseUser) error {
	return h.UsersHandler.CreateUser(ctx, in, out)
}

func (h *usersHandler) UpdateUser(ctx context.Context, in *RequestUpdateUser, out *ResponseUser) error {
	return h.UsersHandler.UpdateUser(ctx, in, out)
}

func (h *usersHandler) DeleteUser(ctx context.Context, in *RequestUserID, out *ResponseUser) error {
	return h.UsersHandler.DeleteUser(ctx, in, out)
}
