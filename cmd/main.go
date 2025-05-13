package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "NOT FOUND 404")
		return
	}
	fmt.Fprintf(w, "home route")
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "add two numbers")
}
func handleSubtract(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "subtract two numbers")
}
func handleMultiply(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "multiply two numbers")
}
func handleDivide(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "divide two numbers")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/add", handleAdd)
	mux.HandleFunc("/subtract", handleSubtract)
	mux.HandleFunc("/multply", handleMultiply)
	mux.HandleFunc("/divide", handleDivide)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Server is running at PORT", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
