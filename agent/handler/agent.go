package handler

import (
	context "context"
	containerd "github.com/containerd/containerd/v2/client"
	"proto"
)

type Agent struct {
	Client *containerd.Client
}

func (a Agent) CreatePod(ctx context.Context, in *proto.Pod, out *proto.Pod) error {
	a.Client.Pull(ctx, in.Image)
	return nil
}
