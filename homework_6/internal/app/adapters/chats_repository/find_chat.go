package chats_repository

import (
	"context"

	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/models"
)

func (r *Repository) FindChat(ctx context.Context, chatID models.ChatID) (*models.Chat, error) {
	return nil, models.ErrNotFound
}
