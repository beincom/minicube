package main

import (
	"agent/handler"
	"context"
	containerd "github.com/containerd/containerd/v2/client"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"proto"
)

func main() {

	srv := micro.NewService()
	srv.Init(
		micro.Name(proto.AgentServiceName),
	)

	// os.Getenv("CONTAINERD_ADDRESS")
	client, err := containerd.New("/run/containerd/containerd.sock")
	defer client.Close()

	if err != nil {
		logger.Fatal(err)
	}

	err = handler.Agent{
		Client: client,
	}.CreatePod(context.Background(), &proto.Pod{
		Image: "alpine",
		Name:  "test",
	}, &proto.Pod{})

	if err != nil {
		logger.Fatal(err)
	}

	// Register handler
	err = proto.RegisterAgentHandler(srv.Server(), &handler.Agent{
		Client: client,
	})

	if err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

	// call to manger service

	// Create service
	//managerService := proto.NewManagerService(proto.ManagerServiceName, grpc.NewClient())
	//
	//node, err := managerService.RegisterNode(context.Background(), &proto.Node{
	//	Name: "Name1",
	//})
	//
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//
	//logger.Info(node.Name)
}
