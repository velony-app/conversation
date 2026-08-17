package server

import (
	v1 "github.com/velony-app/conversation/api/velony/conversation/v1"
	"github.com/velony-app/conversation/internal/conf"
	"github.com/velony-app/conversation/internal/presentation/api"

	"github.com/go-kratos/kratos/v3/transport/grpc"
)

func NewGRPCServer(
	c *conf.Server,
	service *api.Service,
	metrics *ServerMetrics,
) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(
			serverMiddleware(metrics)...,
		),
	}

	if c.Grpc.Network != "" {
		opts = append(
			opts,
			grpc.Network(c.Grpc.Network),
		)
	}

	if c.Grpc.Addr != "" {
		opts = append(
			opts,
			grpc.Address(c.Grpc.Addr),
		)
	}

	if c.Grpc.Timeout != nil {
		opts = append(
			opts,
			grpc.Timeout(c.Grpc.Timeout.AsDuration()),
		)
	}

	srv := grpc.NewServer(opts...)

	v1.RegisterConversationServiceServer(
		srv,
		service,
	)

	return srv
}
