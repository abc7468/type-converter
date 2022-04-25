package consumer

import (
	"fmt"
	"sixshop/apilog/data"
	"sixshop/apilog/utils"
	"sync"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ApiDataConsumer struct {
	Consumer   *kafka.Consumer
	ticker     *time.Ticker
	c          chan data.Data
	mu         sync.RWMutex
	msgBuf     []*data.Data
	maxBufSize int
}

func (ac *ApiDataConsumer) flushBuffer() {
	if len(ac.msgBuf) > 0 {
		ac.msgBuf = make([]*data.Data, 0, ac.maxBufSize)
	}
}

func (ac *ApiDataConsumer) insertMessage(msg *data.Data) {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	ac.msgBuf = append(ac.msgBuf, msg)
	if len(ac.msgBuf) >= ac.maxBufSize {
		ac.flushBuffer()
	}
}

func (ac *ApiDataConsumer) Consume() {
	topic := "api"
	ac.c = make(chan data.Data, 0)
	ac.ticker = time.NewTicker(time.Duration(60) * time.Second)
	go func() {
		for {
			select {
			case message, ok := <-ac.c:
				if ok {
					fmt.Println(message)
				} else {
					continue
				}
			case <-ac.ticker.C:
				ac.mu.Lock()
				ac.flushBuffer()
				ac.mu.Unlock()
			}

		}
	}()

	ac.Consumer.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := ac.Consumer.ReadMessage(-1)
		if err == nil {
			d := data.Data{}
			utils.FromBytes(&d, msg.Value)

			d.MakeData()
			ac.c <- d
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
func KafKaConsumer(cfg *kafka.ConfigMap) *kafka.Consumer {
	consumer := APIDataConsumer(cfg)
	return consumer
}

func APIDataConsumer(cfg *kafka.ConfigMap) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		panic(err)
	}
	return consumer
}
