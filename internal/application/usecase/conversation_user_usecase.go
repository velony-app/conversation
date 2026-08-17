package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/velony-app/conversation/internal/application/command"
	"github.com/velony-app/conversation/internal/application/common"
	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/repository"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type ConversationUserUsecase struct {
	conversationRepository     repository.ConversationRepository
	userRepository             repository.UserRepository
	conversationUserRepository repository.ConversationUserRepository
	unitOfWork                 repository.UnitOfWork
}

func NewConversationUserUsecase(
	conversationRepository repository.ConversationRepository,
	userRepository repository.UserRepository,
	conversationUserRepository repository.ConversationUserRepository,
	unitOfWork repository.UnitOfWork,
) *ConversationUserUsecase {
	return &ConversationUserUsecase{
		conversationRepository:     conversationRepository,
		userRepository:             userRepository,
		conversationUserRepository: conversationUserRepository,
		unitOfWork:                 unitOfWork,
	}
}

func (s *ConversationUserUsecase) Join(
	ctx context.Context,
	cmd *command.JoinConversationCommand,
) (*command.JoinConversationCommandResult, error) {
	rawActorID, err := uuid.Parse(cmd.ActorID)
	if err != nil {
		return nil, err
	}
	rawConversationID, err := uuid.Parse(cmd.ConversationID)
	if err != nil {
		return nil, err
	}

	actorID := value_object.NewUserID(rawActorID)
	conversationID := value_object.NewConversationID(rawConversationID)

	var createdConversationUser *entity.ConversationUser

	if err := s.unitOfWork.Do(ctx, func(ctx context.Context) error {
		user, err := s.userRepository.Find(ctx, actorID)
		if err != nil {
			return err
		}
		conversation, err := s.conversationRepository.Find(ctx, conversationID)
		if err != nil {
			return err
		}

		conversationUser, err := user.JoinConversation(conversation)
		if err != nil {
			return err
		}

		createdConversationUser = conversationUser

		return s.conversationUserRepository.Save(ctx, conversationUser)
	}); err != nil {
		return nil, err
	}

	return &command.JoinConversationCommandResult{
		ConversationUser: &common.ConversationUserResult{
			ID:       createdConversationUser.ID.UserID().String(),
			Role:     createdConversationUser.Role.String(),
			JoinTime: createdConversationUser.JoinTime,
		},
	}, nil
}

func (s *ConversationUserUsecase) Leave(
	ctx context.Context,
	cmd *command.LeaveConversationCommand,
) (*command.LeaveConversationCommandResult, error) {
	rawActorID, err := uuid.Parse(cmd.ActorID)
	if err != nil {
		return nil, err
	}
	rawConversationID, err := uuid.Parse(cmd.ConversationID)
	if err != nil {
		return nil, err
	}

	actorID := value_object.NewUserID(rawActorID)
	conversationID := value_object.NewConversationID(rawConversationID)

	if err := s.unitOfWork.Do(ctx, func(ctx context.Context) error {
		conversationUser, err := s.conversationUserRepository.FindByConversationAndUser(ctx, conversationID, actorID)
		if err != nil {
			return err
		}

		if err := conversationUser.Leave(); err != nil {
			return err
		}

		return s.conversationUserRepository.Save(ctx, conversationUser)
	}); err != nil {
		return nil, err
	}

	return &command.LeaveConversationCommandResult{}, nil
}
