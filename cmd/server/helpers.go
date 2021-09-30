package main

import (
    "fmt"
    "net/http"
    "runtime/debug"
    "encoding/json"

    "github.com/cyril-ui-developer/codersnippet/pkg/models"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    app.errorLog.Output(2, trace)

    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *application) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
    app.clientError(w, http.StatusNotFound)
}


func (app *application) customServerMessage(w http.ResponseWriter, status , message string) {
   
    successMsg :=  models.Message{status, message}

   jsMsg, msgErr := json.Marshal(successMsg)
   if msgErr != nil{
      app.serverError(w, msgErr)
   }
  w.Write(jsMsg)
}

func (app *application) isAuthenticated(r *http.Request) bool {
    return app.session.Exists(r, "authenticatedUserID")
}