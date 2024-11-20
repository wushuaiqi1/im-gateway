package tests

import (
	"fmt"
	"github.com/google/uuid"
	"im/chat"
	"testing"
)

func TestNewMessage(t *testing.T) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return
	}
	msg := chat.NewMessage(newUUID.String(), "您好，我是卡卡", 001, 002)
	fmt.Println(msg)
}
