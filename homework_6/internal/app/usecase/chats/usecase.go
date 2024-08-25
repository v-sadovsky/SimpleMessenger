package chats

import (
	"context"

	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/models"
)

// Usecase - primary port to service transport/delivery level
type Usecase interface {
	// SendMessage - send message from a user to his friend in a separate chat
	SendMessage(ctx context.Context, messageInfo *SendMessageInfoDTO) (*models.Message, error)
}

type (
	// MessagesRepository - port (secondary)
	MessagesRepository interface {
		CreateChat(ctx context.Context, chat *models.Chat) (models.ChatID, error)
		FindChat(ctx context.Context, chatID models.ChatID) (*models.Chat, error)
		SaveMessage(ctx context.Context, chatID models.ChatID, msg *models.Message) (models.ChatID, error)
	}
)

// Deps -
type Deps struct {
	// Adapters
	MessagesRepository MessagesRepository
}

type usecase struct {
	Deps
}

func NewUsecase(d Deps) Usecase {
	return &usecase{d}
}
