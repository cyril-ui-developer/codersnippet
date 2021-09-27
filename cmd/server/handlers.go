package main

import (
	//"fmt"
	"net/http"
	"strconv"
	"errors"
	"encoding/json"
	"io/ioutil"
	"strings"
	"unicode/utf8"
	

	"github.com/cyril-ui-developer/codersnippet/pkg/models"
)


func (app *application) home(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	js, errs := json.Marshal(s)
	if errs != nil {
        app.errorLog.Println(errs)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
	js, errs := json.Marshal(s)
	if errs != nil {
        app.errorLog.Println(errs)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) addSnippet(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	requestBody, _ := ioutil.ReadAll(r.Body)
	 var sp  models.Snippet
	 json.Unmarshal(requestBody, &sp)

	// Create some variables holding test data. 
    title :=  sp.Title
    content := sp.Content
    expires :=  sp.Expires
 
	// Initialize a map to hold any validation errors.
	errors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
        errors["title"] = "This field cannot be blank"
    } else if utf8.RuneCountInString(title) > 75 {
        errors["title"] = "This field is too long (maximum is 75 characters)"
    }
    
	// Check that the Content field isn't blank.
    if strings.TrimSpace(content) == "" {
        errors["content"] = "This field cannot be blank"
    }else if utf8.RuneCountInString(title) > 250 {
        errors["content"] = "This field is too long (maximum is 250 characters)"
    }

	js, errs := json.Marshal(errors)
	if errs != nil {
        app.errorLog.Println(errs)
        http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
        return
    }
	if len(errors) > 0 {
		w.Write(js)
        return
    }

	id, erro := app.snippets.Insert(title, content, expires)
    if erro != nil {
        app.serverError(w, erro)
        return
    }
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}