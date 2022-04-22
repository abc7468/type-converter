package kafkasvc

import (
	"fmt"
	"sixshop/web/data"
	"sixshop/web/utils"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Producer struct {
	Producer *kafka.Producer
}

func (p *Producer) Print() {
	go func() {
		for e := range p.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

func (p *Producer) Produce(data *data.Data) {
	topic := "myTopic"
	dataByte := utils.ToBytes(data)
	fmt.Println(&p.Producer)
	err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          dataByte,
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
	p.Producer.Flush(15 * 1000)
}

func APIDataProducer(cfg kafka.ConfigMap) *kafka.Producer {

	p, err := kafka.NewProducer(&cfg)
	if err != nil {
		panic(err)
	}
	return p
}

// func Producer1() {

// 	p := FormatDataProducer()
// 	defer p.Close()

// 	// Delivery report handler for produced messages
// 	go func() {
// 		for e := range p.Events() {
// 			switch ev := e.(type) {
// 			case *kafka.Message:
// 				if ev.TopicPartition.Error != nil {
// 					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
// 				} else {
// 					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
// 				}
// 			}
// 		}
// 	}()

// 	// Produce messages to topic (asynchronously)
// topic := "myTopic"

// for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
// 	_ = p.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          []byte(word),
// 	}, nil)
// }

// 	// Wait for message deliveries before shutting down
// 	p.Flush(15 * 1000)

// }
