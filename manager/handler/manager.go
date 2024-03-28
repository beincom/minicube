package handler

import (
	"context"
	"github.com/samber/do"
	"proto"
)

type Manager struct {
	Container *do.Injector
	Nodes     []*proto.Node
}

func (m Manager) RegisterNode(ctx context.Context, in *proto.Node, out *proto.Node) error {
	out.Name = in.Name + "_registered"
	out.IpAddress = "demo_ip"
	m.Nodes = append(m.Nodes, out)
	return nil
}
