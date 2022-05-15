// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/checkSMS/checkSMS.proto

package go_micro_service_sendSMS

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

// Api Endpoints for CheckSMS service

func NewCheckSMSEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CheckSMS service

type CheckSMSService interface {
	Check(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type checkSMSService struct {
	c    client.Client
	name string
}

func NewCheckSMSService(name string, c client.Client) CheckSMSService {
	return &checkSMSService{
		c:    c,
		name: name,
	}
}

func (c *checkSMSService) Check(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CheckSMS.check", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CheckSMS service

type CheckSMSHandler interface {
	Check(context.Context, *Request, *Response) error
}

func RegisterCheckSMSHandler(s server.Server, hdlr CheckSMSHandler, opts ...server.HandlerOption) error {
	type checkSMS interface {
		Check(ctx context.Context, in *Request, out *Response) error
	}
	type CheckSMS struct {
		checkSMS
	}
	h := &checkSMSHandler{hdlr}
	return s.Handle(s.NewHandler(&CheckSMS{h}, opts...))
}

type checkSMSHandler struct {
	CheckSMSHandler
}

func (h *checkSMSHandler) Check(ctx context.Context, in *Request, out *Response) error {
	return h.CheckSMSHandler.Check(ctx, in, out)
}
