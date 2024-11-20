package service

type IUserSubscribeService interface {
	SubscribeTopic(id int64, topic string) error
	CancelSubscribeTopic(id int64, topic string) error
	GetSubscribeTopicsByUser(id int64) error
	GetUserListByTopic(topic string) error
}

type UserSubscribeService struct {
}

func NewUserSubscribeService() IUserSubscribeService {
	return UserSubscribeService{}
}

func (receiver UserSubscribeService) SubscribeTopic(id int64, topic string) error {
	return nil
}

func (receiver UserSubscribeService) CancelSubscribeTopic(id int64, topic string) error {
	return nil
}
func (receiver UserSubscribeService) GetSubscribeTopicsByUser(id int64) error {
	return nil
}

func (receiver UserSubscribeService) GetUserListByTopic(topic string) error {
	return nil
}
