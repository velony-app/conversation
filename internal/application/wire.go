package application

import (
	"github.com/velony-app/conversation/internal/application/usecase"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	usecase.NewConversationUsecase,
	usecase.NewConversationUserUsecase,
	usecase.NewMessageUsecase,
)
