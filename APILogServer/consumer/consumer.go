package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"sixshop/apilog/data"
	"sixshop/apilog/utils"
	"time"

	"github.com/Shopify/sarama"
)

type ConsumerGroupHandler interface {
	sarama.ConsumerGroupHandler
	WaitReady()
	Reset()
}

type ConsumerGroup struct {
	cg sarama.ConsumerGroup
}

func NewConsumerGroup(broker string, topics []string, group string, handler ConsumerGroupHandler) (*ConsumerGroup, error) {
	ctx := context.Background()
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V0_10_2_0
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	client, err := sarama.NewConsumerGroup([]string{broker}, group, cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := client.Consume(ctx, topics, handler)
			if err != nil {
				if err == sarama.ErrClosedConsumerGroup {
					break
				} else {
					panic(err)
				}
			}
			if ctx.Err() != nil {
				return
			}
			handler.Reset()
		}
	}()

	handler.WaitReady() // Await till the consumer has been set up

	return &ConsumerGroup{
		cg: client,
	}, nil
}

func (c *ConsumerGroup) Close() error {
	return c.cg.Close()
}

type ConsumerSessionMessage struct {
	Session sarama.ConsumerGroupSession
	Message *sarama.ConsumerMessage
}

func decodeMessage(msg []byte) error {
	d := data.Data{}
	utils.FromBytes(&d, msg)
	err := json.Unmarshal(msg, &d)
	if err != nil {
		return err
	}
	return nil
}

func StartBatchConsumer(broker, topic string) (*ConsumerGroup, error) {
	var count int64
	var start = time.Now()

	handler := NewBatchConsumerGroupHandler(&BatchConsumerConfig{
		MaxBufSize: 10,
		Callback: func(messages []*ConsumerSessionMessage) error {
			for i := range messages {
				if err := decodeMessage(messages[i].Message.Value); err == nil {
					messages[i].Session.MarkMessage(messages[i].Message, "")
				}
			}
			count += int64(len(messages))
			if count%5000 == 0 {
				fmt.Printf("batch consumer consumed %d messages at speed %.2f/s\n", count, float64(count)/time.Since(start).Seconds())
			}
			return nil
		},
	})
	consumer, err := NewConsumerGroup(broker, []string{topic}, "batch-consumer-"+fmt.Sprintf("%d", time.Now().Unix()), handler)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
