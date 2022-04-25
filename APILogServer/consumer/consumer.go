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

// msgbuf에 모인 데이터를 bulk 방식을 통해 RDB에 INSERT하는 함수입니다.
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

// 조건이 충족되면 msgBuf에 담긴 값을 처리하는 함수입니다.
func (ac *ApiDataConsumer) flushBuffer() {
	fmt.Println("flush")

	if len(ac.msgBuf) > 0 {
		bulkInsert(ac.msgBuf)
		ac.msgBuf = make([]*data.Data, 0, configuration.Conf.Kafka.MaxBufSize)
	}
}

// msgBuf에 값을 채워넣는 함수입니다.
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
	ac.ticker = time.NewTicker(time.Duration(configuration.Conf.Kafka.FlushTickerSec) * time.Second)
	// ApiDataConsumer가 가진 chan에 값이 전달되면 작동하는 고루틴입니다.
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

	// kafka의 producer가 발행한 데이터를 받아오는 부분입니다.
	ac.Consumer.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := ac.Consumer.ReadMessage(-1)
		if err == nil {
			// 데이터를 data.Data 구조에 파싱하여 chan을 통해 전달합니다.
			d := data.Data{}
			utils.FromBytes(&d, msg.Value)
			d.MakeData()
			ac.c <- d
		} else {
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
