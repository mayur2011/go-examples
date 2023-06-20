package main

import (
	"go-examples/11_http_examples/04_middleware/router"
	"log"
	"net/http"
)

func main() {
	port := ":8092"
	log.Printf("Starting the http server on port: %s", port)
	router := router.InitRoutes()
	http.ListenAndServe(port, router)
}
