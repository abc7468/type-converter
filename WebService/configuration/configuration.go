package configuration

import (
	kafkasvc "sixshop/web/kafka"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func KafKaProducer() *kafka.Producer {
	cfg := kafka.ConfigMap{"bootstrap.servers": "172.24.25.218:29092"}
	producer := kafkasvc.APIDataProducer(cfg)
	return producer
}
