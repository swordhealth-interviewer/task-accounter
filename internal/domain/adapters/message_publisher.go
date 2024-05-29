package adapters

type MessagePublisherInterface interface {
	PublishMessages(message interface{}, queueName string) error
}
