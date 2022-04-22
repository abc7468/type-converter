package configuration

import (
	kafkasvc "sixshop/apilog/kafka"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func KafKaConsumer() *kafka.Consumer {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": "172.24.25.218:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
	consumer := kafkasvc.APIDataConsumer(cfg)
	return consumer
}
