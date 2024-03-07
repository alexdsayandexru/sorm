package kafka

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IProducer interface {
	Send(data []byte) (bool, error)
}

type Producer struct {
	host  string
	port  int
	topic string
	time  time.Time
}

func GetProducer(host string, port int, topic string) *Producer {
	return &Producer{
		host:  host,
		port:  port,
		topic: topic,
		time:  time.Now(),
	}
}

func (p *Producer) Send(data []byte) (bool, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": fmt.Sprintf("%s:%d", p.host, p.port)})
	if err != nil {
		return false, err
	}
	defer producer.Close()
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
	return err == nil, err
}
