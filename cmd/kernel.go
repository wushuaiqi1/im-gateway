package main

import (
	"fmt"
	"im/chat"
	"im/higo"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/higo", higo.Hello)
	fmt.Println("TestPeerToPeerChat")
	http.HandleFunc("/chat", chat.PeerToPeerChat)
	err := http.ListenAndServe("localhost:7999", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
