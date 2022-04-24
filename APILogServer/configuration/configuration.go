package configuration

import (
	"sixshop/apilog/mq"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var Conf = Configuration{}

type Configuration struct {
	Field FieldList `yaml:"field"`
	Kafka KafkaConf `yaml:"kafka"`
}

type FieldList struct {
	Product []string `yaml:"product"`
	Ids     []string `yaml:"ids"`
	Names   []string `yaml:"names"`
	Prices  []string `yaml:"prices"`
}

type KafkaConf struct {
	BootstrapServers string `yaml:"bootStrapServers"`
	GroupId          string `yaml:"groupId"`
	AutoOffsetReset  string `yaml:"autoOffsetReset"`
}

func KafKaConsumer(cfg *kafka.ConfigMap) *kafka.Consumer {
	consumer := mq.APIDataConsumer(cfg)
	return consumer
}
