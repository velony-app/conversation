package api

import (
	"context"

	v1 "github.com/velony-app/conversation/api/velony/conversation/v1"
	"github.com/velony-app/conversation/internal/application/command"
	"github.com/velony-app/conversation/internal/application/usecase"

	"go.einride.tech/aip/resourcename"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	conversationDefaultPageSize = 25
	conversationMaximumPageSize = 100
	messageDefaultPageSize      = 25
	messageMaximumPageSize      = 100

	// TODO: Replace with the authenticated actor ID.
	//
	// This must be a valid UUID because the application use cases
	// parse ActorID using uuid.Parse.
	temporaryActorID = "00000000-0000-0000-0000-000000000001"
)

const (
	userResourcePattern             = "users/{user}"
	conversationResourcePattern     = "conversations/{conversation}"
	conversationUserResourcePattern = "conversations/{conversation}/users/{userMembership}"
	messageResourcePattern          = "conversations/{conversation}/messages/{message}"
)

type Service struct {
	v1.UnimplementedConversationServiceServer

	conversationUsecase     *usecase.ConversationUsecase
	messageUsecase          *usecase.MessageUsecase
	conversationUserUsecase *usecase.ConversationUserUsecase
}

func NewService(
	conversationUsecase *usecase.ConversationUsecase,
	conversationUserUsecase *usecase.ConversationUserUsecase,
	messageUsecase *usecase.MessageUsecase,
) *Service {
	return &Service{
		conversationUsecase:     conversationUsecase,
		conversationUserUsecase: conversationUserUsecase,
		messageUsecase:          messageUsecase,
	}
}

func (s *Service) CreateConversation(
	ctx context.Context,
	req *v1.CreateConversationRequest,
) (*v1.Conversation, error) {
	if req.GetConversation() == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"conversation is required",
		)
	}

	result, err := s.conversationUsecase.CreateConversation(
		ctx,
		&command.CreateConversationCommand{
			ActorID: temporaryActorID,
			Title:   req.GetConversation().GetTitle(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &v1.Conversation{
		Name: resourcename.Sprint(
			conversationResourcePattern,
			result.Conversation.ID,
		),
		Title:       result.Conversation.Title,
		AvatarImage: result.Conversation.Avatar,
		CreateTime:  timestamppb.New(result.Conversation.CreateTime),
	}, nil
}

func (s *Service) UpdateConversation(
	ctx context.Context,
	req *v1.UpdateConversationRequest,
) (*v1.Conversation, error) {
	if req.GetConversation() == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"conversation is required",
		)
	}

	var conversationID string

	if err := resourcename.Sscan(
		req.GetConversation().GetName(),
		conversationResourcePattern,
		&conversationID,
	); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	updateTitle := req.GetUpdateMask() == nil &&
		req.GetConversation().GetTitle() != ""

	if req.GetUpdateMask() != nil {
		for _, path := range req.GetUpdateMask().GetPaths() {
			switch path {
			case "*":
				updateTitle = true

			case "title":
				updateTitle = true

			case "name":
				// Identifier.

			case "avatar":
				// Output only.

			case "create_time":
				// Output only.

			default:
				return nil, status.Errorf(
					codes.InvalidArgument,
					"invalid update mask path: %s",
					path,
				)
			}
		}
	}

	if req.GetUpdateMask() != nil &&
		updateTitle &&
		req.GetConversation().GetTitle() == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"title is required",
		)
	}

	var title *string

	if updateTitle {
		value := req.GetConversation().GetTitle()
		title = &value
	}

	result, err := s.conversationUsecase.UpdateConversation(
		ctx,
		&command.UpdateConversationCommand{
			ActorID:        temporaryActorID,
			ConversationID: conversationID,
			Title:          title,
		},
	)
	if err != nil {
		return nil, err
	}

	return &v1.Conversation{
		Name: resourcename.Sprint(
			conversationResourcePattern,
			result.Conversation.ID,
		),
		Title:       result.Conversation.Title,
		AvatarImage: result.Conversation.Avatar,
		CreateTime:  timestamppb.New(result.Conversation.CreateTime),
	}, nil
}

func (s *Service) DeleteConversation(
	ctx context.Context,
	req *v1.DeleteConversationRequest,
) (*emptypb.Empty, error) {
	var conversationID string

	if err := resourcename.Sscan(
		req.GetName(),
		conversationResourcePattern,
		&conversationID,
	); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	if _, err := s.conversationUsecase.DeleteConversation(
		ctx,
		&command.DeleteConversationCommand{
			ActorID:        temporaryActorID,
			ConversationID: conversationID,
		},
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) CreateMessage(
	ctx context.Context,
	req *v1.CreateMessageRequest,
) (*v1.Message, error) {
	if req.GetMessage() == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"message is required",
		)
	}

	var conversationID string

	if err := resourcename.Sscan(
		req.GetParent(),
		conversationResourcePattern,
		&conversationID,
	); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	if req.GetMessage().GetContent() == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"message.content is required",
		)
	}

	result, err := s.messageUsecase.SendMessage(
		ctx,
		&command.SendMessageCommand{
			ActorID:        temporaryActorID,
			ConversationID: conversationID,
			Content:        req.GetMessage().GetContent(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &v1.Message{
		Name: resourcename.Sprint(
			messageResourcePattern,
			result.Message.ConversationID,
			result.Message.ID,
		),
		User: resourcename.Sprint(
			userResourcePattern,
			result.Message.ConversationID,
			result.Message.UserID,
		),
		Content:    result.Message.Content,
		CreateTime: timestamppb.New(result.Message.CreateTime),
		UpdateTime: timestamppb.New(result.Message.UpdateTime),
	}, nil
}

func (s *Service) UpdateMessage(
	ctx context.Context,
	req *v1.UpdateMessageRequest,
) (*v1.Message, error) {
	if req.GetMessage() == nil {
		return nil, status.Error(
			codes.InvalidArgument,
			"message is required",
		)
	}

	var (
		conversationID string
		messageID      string
	)

	if err := resourcename.Sscan(
		req.GetMessage().GetName(),
		messageResourcePattern,
		&conversationID,
		&messageID,
	); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	updateContent := req.GetUpdateMask() == nil

	if req.GetUpdateMask() != nil {
		for _, path := range req.GetUpdateMask().GetPaths() {
			switch path {
			case "*":
				updateContent = true

			case "content":
				updateContent = true

			case "name":
				// Identifier.

			case "conversation":
				// Output only.

			case "sender":
				// Output only.

			case "create_time":
				// Output only.

			case "update_time":
				// Output only.

			default:
				return nil, status.Errorf(
					codes.InvalidArgument,
					"invalid update mask path: %s",
					path,
				)
			}
		}
	}

	if !updateContent {
		return nil, status.Error(
			codes.InvalidArgument,
			"update_mask must include content",
		)
	}

	if req.GetMessage().GetContent() == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"message.content is required",
		)
	}

	result, err := s.messageUsecase.EditMessage(
		ctx,
		&command.EditMessageCommand{
			ActorID:   temporaryActorID,
			MessageID: messageID,
			Content:   req.GetMessage().GetContent(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &v1.Message{
		Name: resourcename.Sprint(
			messageResourcePattern,
			result.Message.ConversationID,
			result.Message.ID,
		),
		User: resourcename.Sprint(
			userResourcePattern,
			result.Message.ConversationID,
			result.Message.UserID,
		),
		Content:    result.Message.Content,
		CreateTime: timestamppb.New(result.Message.CreateTime),
		UpdateTime: timestamppb.New(result.Message.UpdateTime),
	}, nil
}

func (s *Service) DeleteMessage(
	ctx context.Context,
	req *v1.DeleteMessageRequest,
) (*emptypb.Empty, error) {
	var (
		conversationID string
		messageID      string
	)

	if err := resourcename.Sscan(
		req.GetName(),
		messageResourcePattern,
		&conversationID,
		&messageID,
	); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	if _, err := s.messageUsecase.DeleteMessage(
		ctx,
		&command.DeleteMessageCommand{
			ActorID:   temporaryActorID,
			MessageID: messageID,
		},
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
