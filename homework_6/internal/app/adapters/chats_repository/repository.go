package chats_repository

import "github.com/v_sadovsky/simple_messenger/homework_6/internal/app/usecase/chats"

type Repository struct {
}

var (
	_ chats.MessagesRepository = (*Repository)(nil)
)

func NewRepository() *Repository {
	return &Repository{}
}
