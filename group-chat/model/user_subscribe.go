package model

// UserSubscribe 用户订阅主题
type UserSubscribe struct {
	Id    int64
	Topic string
}

func NewUserScribe(id int64, topic string) *UserSubscribe {
	return &UserSubscribe{
		Id:    id,
		Topic: topic,
	}
}
