// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/manager.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Manager service

func NewManagerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Manager service

type ManagerService interface {
	RegisterNode(ctx context.Context, in *Node, opts ...client.CallOption) (*Node, error)
}

type managerService struct {
	c    client.Client
	name string
}

func NewManagerService(name string, c client.Client) ManagerService {
	return &managerService{
		c:    c,
		name: name,
	}
}

func (c *managerService) RegisterNode(ctx context.Context, in *Node, opts ...client.CallOption) (*Node, error) {
	req := c.c.NewRequest(c.name, "Manager.RegisterNode", in)
	out := new(Node)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Manager service

type ManagerHandler interface {
	RegisterNode(context.Context, *Node, *Node) error
}

func RegisterManagerHandler(s server.Server, hdlr ManagerHandler, opts ...server.HandlerOption) error {
	type manager interface {
		RegisterNode(ctx context.Context, in *Node, out *Node) error
	}
	type Manager struct {
		manager
	}
	h := &managerHandler{hdlr}
	return s.Handle(s.NewHandler(&Manager{h}, opts...))
}

type managerHandler struct {
	ManagerHandler
}

func (h *managerHandler) RegisterNode(ctx context.Context, in *Node, out *Node) error {
	return h.ManagerHandler.RegisterNode(ctx, in, out)
}
