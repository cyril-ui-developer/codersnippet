package main

import (
    "database/sql"
	"flag"
	"net/http"
	"os"
	"log"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/cyril-ui-developer/codersnippet/pkg/models/mysql"
    "github.com/golangcollege/sessions"
)
type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
    session  *sessions.Session
    snippets *mysql.SnippetModel
    users    *mysql.UserModel
}
func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}


func main(){
	addr := flag.String("addr", ":4000", "HTTP network address")
    dsn := flag.String("dsn", "test:test@/codersnippet?parseTime=true", "MySQL data source name")
    secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }

    defer db.Close()
    // Initialize the session with New()
    session := sessions.New([]byte(*secret))
    session.Lifetime = 12 * time.Hour
    session.Secure = true

    // Initialize a new instance of application containing the dependencies.
    app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
        session:  session,
        snippets: &mysql.SnippetModel{DB: db},
        users:    &mysql.UserModel{DB: db},
    }


    
	srv := &http.Server{
        Addr:     *addr,
        ErrorLog: errorLog,
        Handler:  app.routes(),
    }

    infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}