package main

import (
	"net/http"
	"github.com/bmizerany/pat" 
)


func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/list-snippets", app.session.Enable(http.HandlerFunc(app.home)))
	mux.Get("/snippet/:id", app.session.Enable(http.HandlerFunc(app.showSnippet)))
	mux.Post("/snippet/add", app.session.Enable(http.HandlerFunc(app.addSnippet)))
	mux.Post("/login", app.session.Enable(http.HandlerFunc(app.loginUser)))
	mux.Post("/register", app.session.Enable(http.HandlerFunc(app.registerUser)))
	mux.Post("/logout", app.session.Enable(http.HandlerFunc(app.logoutUser)))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}