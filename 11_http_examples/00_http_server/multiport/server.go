package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// logging middleware code
func loggingMWHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//before executing the handler
		//start := time.Now()
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		//after executing the handler
		//log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func generateJWTToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Generating JWT Token..!")
}

func main() {
	router := mux.NewRouter()
	router.Use(loggingMWHandler)
	router.HandleFunc("/GenerateToken", generateJWTToken).Methods("POST")

	server := &http.Server{
		Addr:    ":8091",
		Handler: router,
	}
	server.ListenAndServe()
}
