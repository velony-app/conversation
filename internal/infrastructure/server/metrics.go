package server

import (
	"fmt"

	"github.com/go-kratos/kratos/contrib/otel/v3/metrics"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const instrumentationName = "github.com/velony-app/conversation/internal/infrastructure/server"

type ServerMetrics struct {
	requests metric.Int64Counter
	seconds  metric.Float64Histogram
}

func NewServerMetrics() (*ServerMetrics, error) {
	meter := otel.Meter(instrumentationName)

	requests, err := metrics.DefaultRequestsCounter(
		meter,
		metrics.DefaultServerRequestsCounterName,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create server requests counter: %w",
			err,
		)
	}

	seconds, err := metrics.DefaultSecondsHistogram(
		meter,
		metrics.DefaultServerSecondsHistogramName,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create server duration histogram: %w",
			err,
		)
	}

	return &ServerMetrics{
		requests: requests,
		seconds:  seconds,
	}, nil
}
