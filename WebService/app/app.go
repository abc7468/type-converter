package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func inputHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/input.html")
}

func MakeRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", inputHandler).Methods("GET")
	return r
}
