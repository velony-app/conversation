package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/velony-app/conversation/internal/application/command"
	"github.com/velony-app/conversation/internal/application/common"
	"github.com/velony-app/conversation/internal/application/query"
	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/repository"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type ConversationUsecase struct {
	conversationRepository     repository.ConversationRepository
	userRepository             repository.UserRepository
	conversationUserRepository repository.ConversationUserRepository
	unitOfWork                 repository.UnitOfWork
}

func NewConversationUsecase(
	conversationRepository repository.ConversationRepository,
	userRepository repository.UserRepository,
	conversationUserRepository repository.ConversationUserRepository,
	unitOfWork repository.UnitOfWork,
) *ConversationUsecase {
	return &ConversationUsecase{
		conversationRepository:     conversationRepository,
		userRepository:             userRepository,
		conversationUserRepository: conversationUserRepository,
		unitOfWork:                 unitOfWork,
	}
}

func (s *ConversationUsecase) GetConversation(
	ctx context.Context,
	qry *query.GetConversationQuery,
) (*query.GetConversationQueryResult, error) {
	return nil, nil
}

func (s *ConversationUsecase) CreateConversation(
	ctx context.Context,
	cmd *command.CreateConversationCommand,
) (*command.CreateConversationCommandResult, error) {
	rawActorID, err := uuid.Parse(cmd.ActorID)
	if err != nil {
		return nil, err
	}
	actorID := value_object.NewUserID(rawActorID)
	title, err := value_object.NewConversationTitle(cmd.Title)
	if err != nil {
		return nil, err
	}
	avatar, err := value_object.NewResourceName(cmd.Avatar)
	if err != nil {
		return nil, err
	}

	var createdConversation *entity.Conversation

	if err := s.unitOfWork.Do(ctx, func(ctx context.Context) error {
		user, err := s.userRepository.Find(ctx, actorID)
		if err != nil {
			return err
		}

		conversation, conversationUser := user.CreateConversation(title, avatar)

		if err := s.conversationRepository.Save(ctx, conversation); err != nil {
			return err
		}
		if err := s.conversationUserRepository.Save(ctx, conversationUser); err != nil {
			return err
		}

		createdConversation = conversation

		return nil
	}); err != nil {
		return nil, err
	}

	return &command.CreateConversationCommandResult{
		Conversation: &common.ConversationResult{
			ID:         createdConversation.ID.String(),
			Title:      createdConversation.Title.String(),
			Avatar:     createdConversation.AvatarImage.String(),
			CreateTime: createdConversation.CreateTime,
		},
	}, nil
}

func (s *ConversationUsecase) UpdateConversation(
	ctx context.Context,
	cmd *command.UpdateConversationCommand,
) (*command.UpdateConversationCommandResult, error) {
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

	var updatedConversation *entity.Conversation

	if err := s.unitOfWork.Do(ctx, func(ctx context.Context) error {
		conversation, err := s.conversationRepository.Find(ctx, conversationID)
		if err != nil {
			return err
		}
		conversationUser, err := s.conversationUserRepository.FindByConversationAndUser(ctx, conversationID, actorID)
		if err != nil {
			return err
		}

		if cmd.Title != nil {
			title, err := value_object.NewConversationTitle(*cmd.Title)
			if err != nil {
				return err
			}
			if err := conversation.ChangeTitle(conversationUser, title); err != nil {
				return err
			}
		}

		if cmd.Title != nil {
			if err := s.conversationRepository.Save(ctx, conversation); err != nil {
				return err
			}
		}

		updatedConversation = conversation

		return nil
	}); err != nil {
		return nil, err
	}

	return &command.UpdateConversationCommandResult{
		Conversation: &common.ConversationResult{
			ID:         updatedConversation.ID.String(),
			Title:      updatedConversation.Title.String(),
			Avatar:     updatedConversation.AvatarImage.String(),
			CreateTime: updatedConversation.CreateTime,
		},
	}, nil
}

func (s *ConversationUsecase) DeleteConversation(
	ctx context.Context,
	cmd *command.DeleteConversationCommand,
) (*command.DeleteConversationCommandResult, error) {
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
		conversation, err := s.conversationRepository.Find(ctx, conversationID)
		if err != nil {
			return err
		}
		conversationUser, err := s.conversationUserRepository.FindByConversationAndUser(ctx, conversationID, actorID)
		if err != nil {
			return err
		}

		if err := conversation.Delete(conversationUser); err != nil {
			return err
		}

		if err := s.conversationRepository.Save(ctx, conversation); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &command.DeleteConversationCommandResult{}, nil
}
