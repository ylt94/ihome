// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/login/login.proto

package go_micro_service_register

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

// Api Endpoints for Login service

func NewLoginEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Login service

type LoginService interface {
	Login(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type loginService struct {
	c    client.Client
	name string
}

func NewLoginService(name string, c client.Client) LoginService {
	return &loginService{
		c:    c,
		name: name,
	}
}

func (c *loginService) Login(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Login.Login", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginHandler interface {
	Login(context.Context, *Request, *Response) error
}

func RegisterLoginHandler(s server.Server, hdlr LoginHandler, opts ...server.HandlerOption) error {
	type login interface {
		Login(ctx context.Context, in *Request, out *Response) error
	}
	type Login struct {
		login
	}
	h := &loginHandler{hdlr}
	return s.Handle(s.NewHandler(&Login{h}, opts...))
}

type loginHandler struct {
	LoginHandler
}

func (h *loginHandler) Login(ctx context.Context, in *Request, out *Response) error {
	return h.LoginHandler.Login(ctx, in, out)
}
