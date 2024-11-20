package model

import (
	"github.com/google/uuid"
	"im/constant"
	"time"
)

type Message struct {
	Id        string
	Type      int8
	Body      []byte
	FromId    int
	ToId      int
	TimeStamp int64
}

// BuildTextMessage 构建文本消息
func BuildTextMessage(from int, to int, body []byte) Message {
	return Message{
		Id:        uuid.NewString(),
		Type:      constant.TEXT,
		Body:      body,
		FromId:    from,
		ToId:      to,
		TimeStamp: time.Now().UnixMilli(),
	}
}
