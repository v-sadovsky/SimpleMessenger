package chats_repository

import (
	"context"

	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/models"
)

func (r *Repository) SaveMessage(ctx context.Context, chatID models.ChatID, msg *models.Message) (models.ChatID, error) {
	return 0, models.ErrNotFound
}