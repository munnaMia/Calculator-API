package main

import "net/http"

func handleAdd(w http.ResponseWriter, r *http.Request){}
func handleSubtract(w http.ResponseWriter, r *http.Request){}
func handleMultiply(w http.ResponseWriter, r *http.Request){}
func handleDivide(w http.ResponseWriter, r *http.Request){}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/add", handleAdd)
	mux.HandleFunc("/subtract", handleSubtract)
	mux.HandleFunc("/multply", handleMultiply)
	mux.HandleFunc("/divide", handleDivide)
}