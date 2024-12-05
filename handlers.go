package main

import "net/http"

func (app *application) add(w http.ResponseWriter, r *http.Request)      {}
func (app *application) subtract(w http.ResponseWriter, r *http.Request) {}
func (app *application) multiply(w http.ResponseWriter, r *http.Request) {}
func (app *application) divide(w http.ResponseWriter, r *http.Request)   {}
func (app *application) sum(w http.ResponseWriter, r *http.Request)      {}
