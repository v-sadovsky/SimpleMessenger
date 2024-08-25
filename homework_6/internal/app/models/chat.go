package models

type ChatID uint64

type Chat struct {
	ID           ChatID
	Participants []User
	Messages     []Message
}
