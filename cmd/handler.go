package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "NOT FOUND 404", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "home route")
}

func (app *Application) handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Failed to read requestbody", err)
	}

	fmt.Fprintf(w, "add of two numbers is : %s", reqBody)
}
func (app *Application) handleSubtract(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "subtract two numbers")
}
func (app *Application) handleMultiply(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "multiply two numbers")
}
func (app *Application) handleDivide(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "divide two numbers")
}
