package app

import (
	"github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
	"manager/app/factory"
	"manager/handler"
	"proto"
)

var (
	version = "latest"
)

func GRPC(options ...micro.Option) micro.Service {
	container, err := factory.Init()
	if err != nil {
		logger.Fatal("Error initializing registry: ", err)
	}

	srv := micro.NewService(micro.Server(grpc.NewServer(
		server.Name(proto.ManagerServiceName),
	)))

	err = proto.RegisterManagerHandler(srv.Server(), &handler.Manager{
		Container: container,
	})

	if err != nil {
		logger.Fatal("Error registering grpc handler: ", err)
	}

	return srv
}
