package chats_repository

import (
	"context"

	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/models"
)

func (r *Repository) CreateChat(ctx context.Context, chat *models.Chat) (models.ChatID, error) {
	return 0, models.ErrAlreadyExists
}
