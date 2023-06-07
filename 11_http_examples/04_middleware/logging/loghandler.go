package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome ..!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About ..!")
}

func loggingHandlerMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//before
		startTm := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
		//after
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(startTm))
	})
}

func main() {

	http.Handle("/index", loggingHandlerMW(http.HandlerFunc(index)))
	http.Handle("/about", loggingHandlerMW(http.HandlerFunc(about)))
	server := &http.Server{
		Addr: ":8089",
	}
	server.ListenAndServe()
}
