package main

import (
	"fmt"
	"sixshop/apilog/configuration"
	"sixshop/apilog/consumer"
	"sixshop/apilog/db"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// 유틸성이 좋은 viper를 통해 configuration을 초기화 합니다.
func setConf(profile string) {
	viper.AddConfigPath(".")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&configuration.Conf)
	if err != nil {
		panic(err)
	}

	// kafka.yaml 파일의 수정이 발생할 경우 Conf 재 설정합니다.
	// 새로운 API가 생겨날 때 재 시작하지 않고 바로 수정을 하도록 합니다.
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&configuration.Conf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(configuration.Conf)

	})
	viper.WatchConfig()

}

func setConsumer() *consumer.ApiDataConsumer {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": configuration.Conf.Kafka.BootstrapServers,
		"group.id":          configuration.Conf.Kafka.GroupId,
		"auto.offset.reset": configuration.Conf.Kafka.AutoOffsetReset,
	}
	c := &consumer.ApiDataConsumer{
		Consumer: consumer.KafKaConsumer(cfg),
	}
	return c
}

func main() {
	profile := "conf"
	setConf(profile)
	db.InitDB()
	consumer := setConsumer()
	consumer.Consume()
}
