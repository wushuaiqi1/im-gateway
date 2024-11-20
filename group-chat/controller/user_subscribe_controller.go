package controller

import "im/group-chat/service"

type UserSubscribeController struct {
	userSubscribeService service.IUserSubscribeService
}

func NewUserSubscribeController() *UserSubscribeController {
	return &UserSubscribeController{
		userSubscribeService: service.NewUserSubscribeService(),
	}
}
