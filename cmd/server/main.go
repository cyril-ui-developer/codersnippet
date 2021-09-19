package main

import (
	"flag"
	"net/http"
	"os"
	"log"
)
type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

func main(){
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


    // Initialize a new instance of application containing the dependencies.
    app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
    }


    
	srv := &http.Server{
        Addr:     *addr,
        ErrorLog: errorLog,
        Handler:  app.routes(),
    }

    infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}