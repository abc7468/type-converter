package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sixshop/web/data"
	kafkasvc "sixshop/web/kafka"

	"github.com/gorilla/mux"
)

type App struct {
	P kafkasvc.KafkaSvc
}

func (a App) inputHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/input.html")
}

func (a App) sendHandler(w http.ResponseWriter, r *http.Request) {
	body := &data.Data{}
	json.NewDecoder(r.Body).Decode(body)
	fmt.Println(body)
	a.P.Produce(body)

}

func (a App) MakeRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", a.inputHandler).Methods("GET")
	r.HandleFunc("/send", a.sendHandler).Methods("POST")
	return r
}
