package main

import "net/http"

func (app *application) routes() {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", app.add)
	mux.HandleFunc("/subtract", app.subtract)
	mux.HandleFunc("/multiply", app.multiply)
	mux.HandleFunc("/divide", app.divide)
	mux.HandleFunc("/sum", app.sum)
}
