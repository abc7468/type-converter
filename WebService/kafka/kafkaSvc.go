package kafkasvc

import "sixshop/web/data"

type KafkaSvc interface {
	Print()
	Produce(data *data.Data)
}
