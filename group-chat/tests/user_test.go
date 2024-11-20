package tests

import (
	"encoding/json"
	"fmt"
	"im/group-chat/model"
	"testing"
)

func TestNewUserToJSON(t *testing.T) {
	u := model.User{Id: 0, Name: "武帅祺"}
	marshal, err := json.Marshal(u)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
