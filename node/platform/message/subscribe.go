package message

type Subscribed struct {
	*Message
}

func CreateSubscribed(topicName string) *Subscribed {
	return &Subscribed{
		Message: &Message{
			Type:  MessageType.Subscribed,
			Value: topicName,
		},
	}
}
