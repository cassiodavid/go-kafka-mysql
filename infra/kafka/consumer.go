package kafka

import (
	kafika "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *kafika.ConfigMap
	Topics    []string
}

func NewConsumer2(configMap *kafika.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c Consumer) Consume(msgChan chan *kafika.Message) error {
	consumer, err := kafika.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		//Aqui quebra qualquer um kkkkkkkkk == a NULL
		if err == nil {
			msgChan <- msg
		}
	}
}
