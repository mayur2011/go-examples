package main

import (
	"go-examples/11_http_examples/04_middleware/router"
	"log"
	"net/http"
)

func main() {
	log.Printf("%s\n", "Starting the http server application")
	router := router.InitRoutes()
	http.ListenAndServe(":8092", router)
}
