package handler

import (
	"context"
	"errors"
	"fmt"
	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/cio"
	"github.com/containerd/containerd/v2/pkg/namespaces"
	"github.com/containerd/containerd/v2/pkg/oci"
	"github.com/google/uuid"
	"proto"
)

type Agent struct {
	Client *containerd.Client
}

func (a Agent) CreatePod(ctx context.Context, in *proto.Pod, out *proto.Pod) error {

	if in.Image == "" {
		return errors.New("Image is required")
	}

	if in.Name == "" {
		return errors.New("Name is required")
	}

	ctx = namespaces.WithNamespace(ctx, "default")
	image, err := a.Client.Pull(ctx, in.Image, containerd.WithPullUnpack)
	if err != nil {
		return errors.New("Failed to pull image: " + err.Error())
	}

	id := generateNewID(in.Name)

	options := []oci.SpecOpts{
		oci.WithImageConfig(image),
	}

	if in.Limit != nil && in.Limit.Memory > 0 {
		options = append(options, oci.WithMemoryLimit(in.Limit.Memory))
	}

	container, err := a.Client.NewContainer(
		ctx,
		id,
		containerd.WithNewSpec(options...),
		containerd.WithContainerLabels(in.Labels),
		//containerd.WithNewSpec(oci.WithMemoryLimit(in.Limit.Memory)),
	)
	if err != nil {
	}

	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))

	if err != nil {
		return errors.New("Failed to start task: " + err.Error())
	}

	if err := task.Start(ctx); err != nil {
		return errors.New("Failed to start task: " + err.Error())
	}

	out.Id = id
	out.Name = in.Name
	//

	return nil
}

func generateNewID(name string) string {
	id := uuid.New()

	return fmt.Sprintf("%s-%s", name, id)
}
