package main

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"

	"github.com/cyril-ui-developer/codersnippet/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){


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
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
    
	// Create some variables holding test data. 
    title := "Welcome to Cary"
    content := "Welcome to Cary city"
    expires := "30"

	id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }
    http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}