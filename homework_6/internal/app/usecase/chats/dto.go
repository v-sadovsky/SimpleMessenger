package chats

import "github.com/v_sadovsky/simple_messenger/homework_6/internal/app/models"

type SendMessageInfoDTO struct {
	Owner   models.User
	ChatID  models.ChatID
	Message models.Message
}
