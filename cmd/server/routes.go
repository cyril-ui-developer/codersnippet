package main

import (
	"net/http"
	"github.com/bmizerany/pat" 
)


func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))
	mux.Post("/snippet/add", http.HandlerFunc(app.addSnippet))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}