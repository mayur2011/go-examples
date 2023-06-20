package main

import (
	"go-examples/11_http_examples/04_middleware/router"
	"log"
	"net/http"
)

func main() {
	port := ":8092"
	log.Printf("%s\n", "Starting the http server on port: %d", port)
	router := router.InitRoutes()
	http.ListenAndServe(port, router)
}
