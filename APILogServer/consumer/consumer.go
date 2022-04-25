package consumer

import (
	"fmt"
	"sixshop/apilog/configuration"
	"sixshop/apilog/data"
	"sixshop/apilog/db"
	"sixshop/apilog/utils"
	"strings"
	"sync"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ApiDataConsumer struct {
	Consumer *kafka.Consumer
	ticker   *time.Ticker
	c        chan data.Data
	mu       sync.RWMutex
	msgBuf   []*data.Data
}

func bulkInsert(datas []*data.Data) error {
	db := db.DB.D
	var (
		placeholders []string
		vals         []interface{}
	)

	for _, data := range datas {
		placeholders = append(placeholders, fmt.Sprintf("(?,?,?,?,?,?)"))
		vals = append(vals, data.ProductInfo.Id, data.ProductInfo.Name, data.ProductInfo.Price, data.Body, data.Json, data.Time)
	}

	txn, err := db.Begin()
	if err != nil {
		fmt.Println(err)

		return err
	}

	insertStatement := fmt.Sprintf("INSERT INTO api_info (product_id, product_name, product_price, original_api, json_api, send_time) VALUES %s", strings.Join(placeholders, ","))
	_, err = txn.Exec(insertStatement, vals...)
	if err != nil {
		txn.Rollback()
		fmt.Println(err)
		return err
	}

	if err := txn.Commit(); err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}

func (ac *ApiDataConsumer) flushBuffer() {
	fmt.Println("flush")

	if len(ac.msgBuf) > 0 {
		bulkInsert(ac.msgBuf)
		ac.msgBuf = make([]*data.Data, 0, configuration.Conf.Kafka.MaxBufSize)
	}
}

func (ac *ApiDataConsumer) insertMessage(msg *data.Data) {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	ac.msgBuf = append(ac.msgBuf, msg)
	if len(ac.msgBuf) >= configuration.Conf.Kafka.MaxBufSize {
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
					ac.insertMessage(&message)
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
