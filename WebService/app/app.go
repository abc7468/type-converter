package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sixshop/web/model"

	"github.com/gorilla/mux"
)

func inputHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/input.html")
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	body := &model.Data{}
	json.NewDecoder(r.Body).Decode(body)
	fmt.Println(body)
}

func MakeRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", inputHandler).Methods("GET")
	r.HandleFunc("/send", sendHandler).Methods("POST")
	return r
}
