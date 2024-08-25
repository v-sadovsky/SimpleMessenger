package chats

import (
	"context"
	"errors"

	"github.com/v_sadovsky/simple_messenger/homework_6/internal/app/models"
)

func (us *usecase) SendMessage(ctx context.Context, messageInfo *SendMessageInfoDTO) (*models.Message, error) {

	return nil, errors.New("not implemented")
}
