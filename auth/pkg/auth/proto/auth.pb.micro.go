// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pkg/auth/proto/auth.proto

package go_micro_service_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for Auth service

func NewAuthEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Auth service

type AuthService interface {
	LogIn(ctx context.Context, in *RequestAuthLogIn, opts ...client.CallOption) (*ResponseLogIn, error)
	AuthPath(ctx context.Context, in *RequestAuthPath, opts ...client.CallOption) (*ResponseAuthPath, error)
	GetByID(ctx context.Context, in *RequestAuthID, opts ...client.CallOption) (*ResponseAuth, error)
	GetByUserID(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseAuth, error)
	GetQueryPointsByUserID(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseAuthWithQueryPoint, error)
	Create(ctx context.Context, in *RequestCreateAuth, opts ...client.CallOption) (*ResponseAuth, error)
	Update(ctx context.Context, in *RequestUpdateAuth, opts ...client.CallOption) (*ResponseAuth, error)
	Delete(ctx context.Context, in *RequestAuthID, opts ...client.CallOption) (*ResponseAuth, error)
	PushPermission(ctx context.Context, in *RequestPushPermission, opts ...client.CallOption) (*ResponseAuth, error)
	DeletePermission(ctx context.Context, in *RequestDeletePermission, opts ...client.CallOption) (*ResponseAuth, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) LogIn(ctx context.Context, in *RequestAuthLogIn, opts ...client.CallOption) (*ResponseLogIn, error) {
	req := c.c.NewRequest(c.name, "Auth.LogIn", in)
	out := new(ResponseLogIn)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) AuthPath(ctx context.Context, in *RequestAuthPath, opts ...client.CallOption) (*ResponseAuthPath, error) {
	req := c.c.NewRequest(c.name, "Auth.AuthPath", in)
	out := new(ResponseAuthPath)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetByID(ctx context.Context, in *RequestAuthID, opts ...client.CallOption) (*ResponseAuth, error) {
	req := c.c.NewRequest(c.name, "Auth.GetByID", in)
	out := new(ResponseAuth)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetByUserID(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseAuth, error) {
	req := c.c.NewRequest(c.name, "Auth.GetByUserID", in)
	out := new(ResponseAuth)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetQueryPointsByUserID(ctx context.Context, in *RequestUserID, opts ...client.CallOption) (*ResponseAuthWithQueryPoint, error) {
	req := c.c.NewRequest(c.name, "Auth.GetQueryPointsByUserID", in)
	out := new(ResponseAuthWithQueryPoint)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Create(ctx context.Context, in *RequestCreateAuth, opts ...client.CallOption) (*ResponseAuth, error) {
	req := c.c.NewRequest(c.name, "Auth.Create", in)
	out := new(ResponseAuth)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Update(ctx context.Context, in *RequestUpdateAuth, opts ...client.CallOption) (*ResponseAuth, error) {
	req := c.c.NewRequest(c.name, "Auth.Update", in)
	out := new(ResponseAuth)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Delete(ctx context.Context, in *RequestAuthID, opts ...client.CallOption) (*ResponseAuth, error) {
	req := c.c.NewRequest(c.name, "Auth.Delete", in)
	out := new(ResponseAuth)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) PushPermission(ctx context.Context, in *RequestPushPermission, opts ...client.CallOption) (*ResponseAuth, error) {
	req := c.c.NewRequest(c.name, "Auth.PushPermission", in)
	out := new(ResponseAuth)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DeletePermission(ctx context.Context, in *RequestDeletePermission, opts ...client.CallOption) (*ResponseAuth, error) {
	req := c.c.NewRequest(c.name, "Auth.DeletePermission", in)
	out := new(ResponseAuth)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	LogIn(context.Context, *RequestAuthLogIn, *ResponseLogIn) error
	AuthPath(context.Context, *RequestAuthPath, *ResponseAuthPath) error
	GetByID(context.Context, *RequestAuthID, *ResponseAuth) error
	GetByUserID(context.Context, *RequestUserID, *ResponseAuth) error
	GetQueryPointsByUserID(context.Context, *RequestUserID, *ResponseAuthWithQueryPoint) error
	Create(context.Context, *RequestCreateAuth, *ResponseAuth) error
	Update(context.Context, *RequestUpdateAuth, *ResponseAuth) error
	Delete(context.Context, *RequestAuthID, *ResponseAuth) error
	PushPermission(context.Context, *RequestPushPermission, *ResponseAuth) error
	DeletePermission(context.Context, *RequestDeletePermission, *ResponseAuth) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		LogIn(ctx context.Context, in *RequestAuthLogIn, out *ResponseLogIn) error
		AuthPath(ctx context.Context, in *RequestAuthPath, out *ResponseAuthPath) error
		GetByID(ctx context.Context, in *RequestAuthID, out *ResponseAuth) error
		GetByUserID(ctx context.Context, in *RequestUserID, out *ResponseAuth) error
		GetQueryPointsByUserID(ctx context.Context, in *RequestUserID, out *ResponseAuthWithQueryPoint) error
		Create(ctx context.Context, in *RequestCreateAuth, out *ResponseAuth) error
		Update(ctx context.Context, in *RequestUpdateAuth, out *ResponseAuth) error
		Delete(ctx context.Context, in *RequestAuthID, out *ResponseAuth) error
		PushPermission(ctx context.Context, in *RequestPushPermission, out *ResponseAuth) error
		DeletePermission(ctx context.Context, in *RequestDeletePermission, out *ResponseAuth) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) LogIn(ctx context.Context, in *RequestAuthLogIn, out *ResponseLogIn) error {
	return h.AuthHandler.LogIn(ctx, in, out)
}

func (h *authHandler) AuthPath(ctx context.Context, in *RequestAuthPath, out *ResponseAuthPath) error {
	return h.AuthHandler.AuthPath(ctx, in, out)
}

func (h *authHandler) GetByID(ctx context.Context, in *RequestAuthID, out *ResponseAuth) error {
	return h.AuthHandler.GetByID(ctx, in, out)
}

func (h *authHandler) GetByUserID(ctx context.Context, in *RequestUserID, out *ResponseAuth) error {
	return h.AuthHandler.GetByUserID(ctx, in, out)
}

func (h *authHandler) GetQueryPointsByUserID(ctx context.Context, in *RequestUserID, out *ResponseAuthWithQueryPoint) error {
	return h.AuthHandler.GetQueryPointsByUserID(ctx, in, out)
}

func (h *authHandler) Create(ctx context.Context, in *RequestCreateAuth, out *ResponseAuth) error {
	return h.AuthHandler.Create(ctx, in, out)
}

func (h *authHandler) Update(ctx context.Context, in *RequestUpdateAuth, out *ResponseAuth) error {
	return h.AuthHandler.Update(ctx, in, out)
}

func (h *authHandler) Delete(ctx context.Context, in *RequestAuthID, out *ResponseAuth) error {
	return h.AuthHandler.Delete(ctx, in, out)
}

func (h *authHandler) PushPermission(ctx context.Context, in *RequestPushPermission, out *ResponseAuth) error {
	return h.AuthHandler.PushPermission(ctx, in, out)
}

func (h *authHandler) DeletePermission(ctx context.Context, in *RequestDeletePermission, out *ResponseAuth) error {
	return h.AuthHandler.DeletePermission(ctx, in, out)
}
