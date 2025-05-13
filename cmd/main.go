package main

import (
	"log"
	"net/http"
)

type Application struct{
	data map[string]int
	help string
}

func main() {
	app := &Application{}
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/add", app.handleAdd)
	mux.HandleFunc("/subtract", app.handleSubtract)
	mux.HandleFunc("/multiply", app.handleMultiply)
	mux.HandleFunc("/divide", app.handleDivide)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Server is running at PORT", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
