package models

type MsgID int64

type Message struct {
	ID        MsgID
	Text      string
	Timestamp int64
}
