package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/munnaMia/Calculator-API/internals/utils"
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

	reqBody := utils.MUST(io.ReadAll(r.Body)) // Read Request body

	err := json.Unmarshal(reqBody, &app.data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "add of two numbers is : %d", app.data["a"]+app.data["b"])
}

func (app *Application) handleSubtract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	reqBody := utils.MUST(io.ReadAll(r.Body)) // Read Request body

	err := json.Unmarshal(reqBody, &app.data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Substract of two numbers is : %d", app.data["a"]-app.data["b"])
}

func (app *Application) handleMultiply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	reqBody := utils.MUST(io.ReadAll(r.Body)) // Read Request body

	err := json.Unmarshal(reqBody, &app.data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Multiply of two numbers is : %d", app.data["a"]*app.data["b"])
}

func (app *Application) handleDivide(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		return
	}

	reqBody := utils.MUST(io.ReadAll(r.Body)) // Read Request body

	err := json.Unmarshal(reqBody, &app.data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Division of two numbers is : %d", app.data["a"]/app.data["b"])
}
