package main

import (
	"net/http"
	"sixshop/web/app"
	"sixshop/web/configuration"
	kafkasvc "sixshop/web/kafka"
)

const port string = ":8081"

func main() {
	p := &kafkasvc.Producer{
		Producer: configuration.KafKaProducer(),
	}

	app := app.App{}
	app.P = p
	r := app.MakeRouter()

	http.ListenAndServe(port, r)
}
