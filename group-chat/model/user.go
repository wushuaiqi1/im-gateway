package model

type User struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// 定义全局自增id
var id int = 0

func NewUser(name string) User {
	id++
	return User{
		Name: name,
		Id:   id,
	}
}
