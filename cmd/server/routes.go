package main

import (
	"net/http"
	"github.com/bmizerany/pat" 
)


func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/list-snippets", app.session.Enable(http.HandlerFunc(app.listSnippets)))
	mux.Get("/snippet/:id", app.session.Enable(http.HandlerFunc(app.getSnippet)))
	mux.Post("/snippet/add", app.session.Enable(app.requireAuthentication(http.HandlerFunc(app.addSnippet))))
	mux.Post("/login", app.session.Enable(http.HandlerFunc(app.loginUser)))
	mux.Post("/register", app.session.Enable(http.HandlerFunc(app.registerUser)))
	mux.Post("/logout", app.session.Enable(app.requireAuthentication(http.HandlerFunc(app.logoutUser))))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}