package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are // encountered during parsing the application will be terminated.
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorlog,
		infoLog:  infolog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.route(),
	}

	infolog.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()
	errorlog.Fatal(err)
}
