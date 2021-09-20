package main

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"

	"github.com/cyril-ui-developer/codersnippet/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/"{
		app.notFound(w)
		return
	}

    panic("oops! something went wrong") 
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
   
	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}
	w.Write([]byte("Welcome to Coder Snippet"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	
	s, err := app.snippets.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }
	fmt.Fprintf(w, "%v...", s)
}

func (app *application) addSnippet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header()["Date"] = nil
	w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
    
	// Create some variables holding test data. 
    title := "Hello world"
    content := "Welcome to Hello World"
    expires := "30"

	id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }
    http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}