package kafkasvc

import (
	"fmt"
	"sixshop/apilog/data"
	"sixshop/apilog/utils"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Consumer struct {
	Consumer *kafka.Consumer
}

func (c *Consumer) Consume() {
	topic := "myTopic"
	c.Consumer.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := c.Consumer.ReadMessage(-1)
		if err == nil {
			d := data.Data{}
			utils.FromBytes(&d, msg.Value)
			fmt.Println(string(msg.Value))

			d.MakeProductInfo()

		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func APIDataConsumer(cfg *kafka.ConfigMap) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		panic(err)
	}
	return consumer
}
