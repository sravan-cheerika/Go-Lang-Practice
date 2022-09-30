package main

import (
	"log"
	"net/http"
	"sum_test/subDirectory"
)

// PlayerServer currently returns Hello, world given _any_ request.
func main() {

	//http.HandlerFunc(subDirectory.PlayerServer)
	http.HandleFunc("/hello", subDirectory.MethodHello)
	http.HandleFunc("/headers", subDirectory.MethodHeader)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
