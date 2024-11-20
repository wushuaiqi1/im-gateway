package chat

import "time"

type Message struct {
	Session   string
	Body      string
	From      int64
	To        int64
	Timestamp int64
}

// NewMessage 创建一条新消息
func NewMessage(Session string, Body string, From int64, To int64) *Message {
	return &Message{
		Session:   Session,
		Body:      Body,
		From:      From,
		To:        To,
		Timestamp: time.Now().UnixMilli(),
	}
}
