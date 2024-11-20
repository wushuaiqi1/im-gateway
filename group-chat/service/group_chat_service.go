package service

import (
	"github.com/gorilla/websocket"
)

type ITopicChatService interface {
	// SendMessage 发送到指定群聊消息
	SendMessage(topic string, message []byte)
	// SubscribeTopic 订阅主题
	SubscribeTopic(topic string, uid int64, conn *websocket.Conn)
	// CancelSubscribeTopic 用户离线取消订阅
	CancelSubscribeTopic(conn *websocket.Conn)
}

type TopicChatService struct {
}

func NewTopicChatService() ITopicChatService {
	return TopicChatService{}
}

func (t TopicChatService) SubscribeTopic(topic string, uid int64, conn *websocket.Conn) {

}
func (t TopicChatService) CancelSubscribeTopic(conn *websocket.Conn) {

}
func (t TopicChatService) SendMessage(topic string, message []byte) {

}
