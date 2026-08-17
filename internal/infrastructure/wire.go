package infrastructure

import (
	"github.com/velony-app/conversation/internal/infrastructure/db/mysql"
	"github.com/velony-app/conversation/internal/infrastructure/server"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	mysql.NewConnection,
	mysql.NewUnitOfWork,
	mysql.NewConversationRepository,
	mysql.NewUserRepository,
	mysql.NewConversationUserRepository,
	mysql.NewMessageRepository,
	server.NewGRPCServer,
	server.NewHTTPServer,
	server.NewServerMetrics,
)
