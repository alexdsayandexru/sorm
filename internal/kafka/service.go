package kafka

type IKafkaService interface {
	Send(data string) (bool, error)
}
