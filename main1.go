package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello 1234")
}

func HandlerWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome program")
}

func UserDetails(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Gorilla!\n"))

}

func main() {
	r := mux.NewRouter()
	//r.Host("www.localhost")
	r.HandleFunc("/hello", Handler)
	r.HandleFunc("/welcome", HandlerWelcome)
	r.HandleFunc("/userDetails", UserDetails)
	err := http.ListenAndServe(":9000", r)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
