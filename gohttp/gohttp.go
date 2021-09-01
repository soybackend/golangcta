package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola CTA %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Ir a http://localhost:8080/loquesea")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
