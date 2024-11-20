package tests

import (
	"im/chat"
	"testing"
)

func TestPutUserToOnline(t *testing.T) {
	chat.PutUserToOnline(int64(1), nil)
}

func TestDeleteUserFromOnline(t *testing.T) {
	chat.DeleteUserFromOnline(nil)
}
