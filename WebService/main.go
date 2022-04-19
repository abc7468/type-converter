package main

import (
	"net/http"
	"sixshop/web/app"
)

const port string = ":8080"

func main() {
	app := app.MakeRouter()

	http.ListenAndServe(port, app)
}
