package service

import (
	"im/group-chat/model"
	"log"
)

type IUserService interface {
	Insert(name string)
}

type UserService struct {
	// 定义map模拟数据库用户存储
	userMap map[int]model.User
}

func NewUserService() IUserService {
	return &UserService{
		userMap: make(map[int]model.User),
	}
}

func (receiver *UserService) Insert(name string) {
	user := model.NewUser(name)
	receiver.userMap[user.Id] = user
	log.Println("UserService insert success...", receiver.userMap)
}
