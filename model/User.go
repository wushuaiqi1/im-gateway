package model

var id = 0

type User struct {
	Id   int
	Name string
}

func NewUser(Name string) *User {
	// 并发场景存在消息id重复问题～
	id++
	return &User{
		Id:   id,
		Name: Name,
	}
}
