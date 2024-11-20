package tests

import (
	group_chat "im/group-chat/controller"
	"log"
	"net/http"
	"testing"
)

func TestGroupChat(t *testing.T) {
	//创建控制器
	controller := group_chat.NewGroupChatController()
	http.HandleFunc("/chat/group", controller.GroupChat)
	err := http.ListenAndServe("localhost:8002", nil)
	if err != nil {
		log.Println(err)
	}
}
