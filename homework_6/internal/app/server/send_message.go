package server

import (
	"context"
	"time"

	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/models"
	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/usecase/chats"
	pb "github.com/v_sadovsky/simple_messenger/homework_6/pkg/api/chat_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	// 1. validation
	if err := validateSendMessageRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 2. convert delivery models to domain models/DTO
	messageInfo := newMessageFromPbSendMessageRequest(req)

	// 3. call usecase
	message, err := s.ChatsUsecase.SendMessage(ctx, messageInfo)
	if err != nil {
		return nil, err // обработается на уровне middleware
	}

	// 4. convert domain models/DTO to delivery models
	response := newSendMessageResponseFromDTO(message)

	// 5. return result
	return response, nil
}

func validateSendMessageRequest(_ *pb.SendMessageRequest) error {
	//
	return nil
}

func newMessageFromPbSendMessageRequest(req *pb.SendMessageRequest) *chats.SendMessageInfoDTO {
	return &chats.SendMessageInfoDTO{
		ChatID: models.ChatID(req.GetChatId()),
		Owner: models.User{
			ID: models.UserID(req.GetUserId()),
		},
		Message: models.Message{
			Text: req.GetMessage(),
		},
	}
}

func newSendMessageResponseFromDTO(message *models.Message) *pb.SendMessageResponse {
	t := time.Unix(message.Timestamp, 0)
	return &pb.SendMessageResponse{
		Timestamp: timestamppb.New(t),
	}
}
