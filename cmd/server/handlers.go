package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/"{
		app.notFound(w)
		return
	}

	w.Write([]byte("Welcome to Coder Snippet"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	
	fmt.Fprintf(w, "Show a specific snippet with ID %d...", id)
}

func (app *application) addSnippet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header()["Date"] = nil
	w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
	
	
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
//	w.Write([]byte("Add a new snippet"))
w.Write([]byte(`[{"name":"Cyril"}]`))
}