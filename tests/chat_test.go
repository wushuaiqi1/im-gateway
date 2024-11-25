package tests

import (
	"fmt"
	"im/chat"
	"log"
	"net/http"
	"testing"
)

func TestPeerToPeerChat(t *testing.T) {
	fmt.Println("TestPeerToPeerChat")
	http.HandleFunc("/chat1", chat.PeerToPeerChat)
	err := http.ListenAndServe("localhost:8002", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
