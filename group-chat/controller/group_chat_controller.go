package controller

import (
	"github.com/gorilla/websocket"
	"im/group-chat/service"
	"log"
	"net/http"
	"strconv"
)

type GroupChatController struct {
	tc      service.ITopicChatService
	upgrade websocket.Upgrader
}

func NewGroupChatController() *GroupChatController {
	return &GroupChatController{
		tc: service.NewTopicChatService(),
		upgrade: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

// GroupChat 实现群聊
func (cc *GroupChatController) GroupChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("GroupChat request:")
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		return
	}
	_, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return
	}
	topic := r.URL.Query().Get("topic")
	if topic == "" {
		return
	}
	//协议升级
	conn, err := cc.upgrade.Upgrade(w, r, nil)
	defer cc.CloseConn(conn)

}

func (cc *GroupChatController) CloseConn(conn *websocket.Conn) {

}
