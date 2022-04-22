package main

import (
	"sixshop/apilog/configuration"
	kafkasvc "sixshop/apilog/kafka"
)

func main() {
	p := &kafkasvc.Consumer{
		Consumer: configuration.KafKaConsumer(),
	}
	p.Consume()
}
