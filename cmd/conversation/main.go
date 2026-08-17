package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/velony-app/conversation/internal/conf"

	"github.com/go-kratos/kratos/contrib/otel/v3/tracing"
	"github.com/go-kratos/kratos/v3"
	"github.com/go-kratos/kratos/v3/config"
	"github.com/go-kratos/kratos/v3/config/file"
	"github.com/go-kratos/kratos/v3/log"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"

	_ "go.uber.org/automaxprocs"
)

// Set at build time:
//
//	go build \
//	  -ldflags "-X main.Name=conversation -X main.Version=x.y.z"
var (
	Name    = "conversation"
	Version = "dev"

	flagconf string
)

func init() {
	flag.StringVar(
		&flagconf,
		"conf",
		"../../configs",
		"config path, eg: -conf ./configs",
	)
}

func newApp(
	logger *slog.Logger,
	grpcServer *grpc.Server,
	httpServer *http.Server,
) *kratos.App {
	id, err := os.Hostname()
	if err != nil {
		id = Name
	}

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger),
		kratos.Server(
			grpcServer,
			httpServer,
		),
	)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		slog.Error(
			"application terminated",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
}

func run() error {
	id, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("get hostname: %w", err)
	}

	logger := newLogger(id)
	log.SetDefault(logger)

	bootstrap, cleanupConfig, err := loadBootstrap()
	if err != nil {
		return err
	}
	defer cleanupConfig()

	app, cleanup, err := wireApp(
		bootstrap.Server,
		bootstrap.Data,
		logger,
	)
	if err != nil {
		return fmt.Errorf("initialize application: %w", err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		return fmt.Errorf("run application: %w", err)
	}

	return nil
}

func newLogger(id string) *slog.Logger {
	return log.NewLogger(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			},
		),
		log.WithExtractor(tracing.TraceAttrs),
	).With(
		slog.String("service.id", id),
		slog.String("service.name", Name),
		slog.String("service.version", Version),
	)
}

func loadBootstrap() (
	*conf.Bootstrap,
	func(),
	error,
) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)

	cleanup := func() {
		_ = c.Close()
	}

	if err := c.Load(); err != nil {
		cleanup()

		return nil, func() {}, fmt.Errorf(
			"load configuration: %w",
			err,
		)
	}

	var bootstrap conf.Bootstrap

	if err := c.Scan(&bootstrap); err != nil {
		cleanup()

		return nil, func() {}, fmt.Errorf(
			"scan configuration: %w",
			err,
		)
	}

	return &bootstrap, cleanup, nil
}
