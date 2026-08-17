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

type MessageUsecase struct {
	conversationUserRepository repository.ConversationUserRepository
	messageRepository          repository.MessageRepository
	unitOfWork                 repository.UnitOfWork
}

func NewMessageUsecase(
	conversationUserRepository repository.ConversationUserRepository,
	messageRepository repository.MessageRepository,
	unitOfWork repository.UnitOfWork,
) *MessageUsecase {
	return &MessageUsecase{
		conversationUserRepository: conversationUserRepository,
		messageRepository:          messageRepository,
		unitOfWork:                 unitOfWork,
	}
}

func (s *MessageUsecase) SendMessage(
	ctx context.Context,
	cmd *command.SendMessageCommand,
) (*command.SendMessageCommandResult, error) {
	rawActorID, err := uuid.Parse(cmd.ActorID)
	if err != nil {
		return nil, err
	}
	rawConversationID, err := uuid.Parse(cmd.ConversationID)
	if err != nil {
		return nil, err
	}

	content, err := value_object.NewMessageContent(cmd.Content)
	if err != nil {
		return nil, err
	}

	actorID := value_object.NewUserID(rawActorID)
	conversationID := value_object.NewConversationID(rawConversationID)

	var createdMessage *entity.Message

	if err := s.unitOfWork.Do(ctx, func(ctx context.Context) error {
		conversationUser, err := s.conversationUserRepository.FindByConversationAndUser(ctx, conversationID, actorID)
		if err != nil {
			return err
		}

		message, err := entity.NewMessage(conversationUser, content)
		if err != nil {
			return err
		}

		if err := s.messageRepository.Save(ctx, message); err != nil {
			return err
		}

		createdMessage = message

		return nil
	}); err != nil {
		return nil, err
	}

	return &command.SendMessageCommandResult{
		Message: &common.MessageResult{
			ID:             createdMessage.ID.String(),
			ConversationID: createdMessage.ConversationID.String(),
			UserID:         createdMessage.UserID.String(),
			Content:        createdMessage.Content.String(),
			CreateTime:     createdMessage.CreateTime,
			UpdateTime:     createdMessage.UpdateTime,
		},
	}, nil
}

func (s *MessageUsecase) EditMessage(
	ctx context.Context,
	cmd *command.EditMessageCommand,
) (*command.EditMessageCommandResult, error) {
	rawActorID, err := uuid.Parse(cmd.ActorID)
	if err != nil {
		return nil, err
	}
	rawMessageID, err := uuid.Parse(cmd.MessageID)
	if err != nil {
		return nil, err
	}
	content, err := value_object.NewMessageContent(cmd.Content)
	if err != nil {
		return nil, err
	}

	actorID := value_object.NewUserID(rawActorID)
	messageID := value_object.NewMessageID(rawMessageID)

	var updatedMessage *entity.Message

	if err := s.unitOfWork.Do(ctx, func(ctx context.Context) error {
		message, err := s.messageRepository.Find(ctx, messageID)
		if err != nil {
			return err
		}
		conversationUser, err := s.conversationUserRepository.FindByConversationAndUser(ctx, message.ConversationID, actorID)
		if err != nil {
			return err
		}

		if err := message.Edit(conversationUser, content); err != nil {
			return err
		}

		if err := s.messageRepository.Save(ctx, message); err != nil {
			return err
		}

		updatedMessage = message

		return nil
	}); err != nil {
		return nil, err
	}

	return &command.EditMessageCommandResult{
		Message: &common.MessageResult{
			ID:             updatedMessage.ID.String(),
			ConversationID: updatedMessage.ConversationID.String(),
			UserID:         updatedMessage.UserID.String(),
			Content:        updatedMessage.Content.String(),
			CreateTime:     updatedMessage.CreateTime,
			UpdateTime:     updatedMessage.UpdateTime,
		},
	}, nil
}

func (s *MessageUsecase) DeleteMessage(
	ctx context.Context,
	cmd *command.DeleteMessageCommand,
) (*command.DeleteMessageCommandResult, error) {
	rawActorID, err := uuid.Parse(cmd.ActorID)
	if err != nil {
		return nil, err
	}
	rawMessageID, err := uuid.Parse(cmd.MessageID)
	if err != nil {
		return nil, err
	}

	actorID := value_object.NewUserID(rawActorID)
	messageID := value_object.NewMessageID(rawMessageID)

	if err := s.unitOfWork.Do(ctx, func(ctx context.Context) error {
		message, err := s.messageRepository.Find(ctx, messageID)
		if err != nil {
			return err
		}
		conversationUser, err := s.conversationUserRepository.FindByConversationAndUser(ctx, message.ConversationID, actorID)
		if err != nil {
			return err
		}

		if err := message.Delete(conversationUser); err != nil {
			return err
		}

		if err := s.messageRepository.Save(ctx, message); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &command.DeleteMessageCommandResult{}, nil
}
