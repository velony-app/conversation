package server

import (
	"github.com/go-kratos/kratos/contrib/otel/v3/metrics"
	"github.com/go-kratos/kratos/contrib/otel/v3/tracing"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/middleware/validate"

	"go.einride.tech/aip/fieldbehavior"
	"google.golang.org/protobuf/proto"
)

func serverMiddleware(
	serverMetrics *ServerMetrics,
) []middleware.Middleware {
	return []middleware.Middleware{
		// Outermost middleware.
		//
		// A panic from anything below this layer should become
		// a transport error rather than killing the process.
		recovery.Recovery(),

		// Tracing and metrics wrap validation too, meaning invalid
		// requests can still be observed.
		tracing.Server(),

		metrics.Server(
			metrics.WithSeconds(serverMetrics.seconds),
			metrics.WithRequests(serverMetrics.requests),
		),

		// API-contract validation.
		validate.Validator(
			validateRequiredFields,
		),
	}
}

func validateRequiredFields(req any) error {
	message, ok := req.(proto.Message)
	if !ok {
		return nil
	}

	return fieldbehavior.ValidateRequiredFields(message)
}
