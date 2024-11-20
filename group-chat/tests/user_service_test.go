package tests

import (
	"im/group-chat/service"
	"testing"
)

func TestInsert(t *testing.T) {
	us := service.NewUserService()
	us.Insert("张三")
	us.Insert("李四")
}
