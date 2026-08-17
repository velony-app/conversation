package server

import (
	v1 "github.com/velony-app/conversation/api/velony/conversation/v1"
	"github.com/velony-app/conversation/internal/conf"
	"github.com/velony-app/conversation/internal/presentation/api"

	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(
	c *conf.Server,
	service *api.Service,
	metrics *ServerMetrics,
) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			serverMiddleware(metrics)...,
		),
	}

	if c.Http.Network != "" {
		opts = append(
			opts,
			http.Network(c.Http.Network),
		)
	}

	if c.Http.Addr != "" {
		opts = append(
			opts,
			http.Address(c.Http.Addr),
		)
	}

	if c.Http.Timeout != nil {
		opts = append(
			opts,
			http.Timeout(c.Http.Timeout.AsDuration()),
		)
	}

	srv := http.NewServer(opts...)

	// Operational infrastructure endpoint.
	srv.Handle(
		"/metrics",
		promhttp.Handler(),
	)

	// Connect the HTTP transport to the API/presentation service.
	v1.RegisterConversationServiceHTTPServer(
		srv,
		service,
	)

	return srv
}
