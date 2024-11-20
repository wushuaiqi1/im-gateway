package controller

import (
	"im/group-chat/service"
	"net/http"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController() UserController {
	return UserController{
		userService: service.NewUserService(),
	}
}

func (receiver UserController) add(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	receiver.userService.Insert(name)
}
