package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	value := strings.ToUpper(r.URL.Path[1:])
	fmt.Fprintf(w, "Hello %s", value)
}

func main() {
	http.HandleFunc("/", sayHello)
	port := ":8089"
	log.Fatalln(http.ListenAndServe(port, nil))
}
