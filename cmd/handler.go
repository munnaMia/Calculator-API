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

	app.help = `
	Hey try this URLs  and send some POST requests to them

	no 1 : http://localhost:8080/add --------> to add two number
	no 2 : http://localhost:8080/subtract ---> to subtract two number
	no 3 : http://localhost:8080/multiply ---> to multiply two number
	no 4 : http://localhost:8080/divide -----> to divide two number

	follow this architecture to send json data : '{"a":int, "b":int}'

	!use a & b as the key.
	!use int values only.
	`
	fmt.Fprintf(w, "%s", app.help)
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
