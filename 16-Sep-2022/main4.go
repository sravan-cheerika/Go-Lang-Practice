package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	Name     string 
	LastName string
	Age      uint8
}

func sendResponse(response http.ResponseWriter, request *http.Request) {
	person := person{Name: "Sravan", LastName: "Reddy", Age: 25}
	jsonResponse, jsonError := json.Marshal(person)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	fmt.Println(string(jsonResponse))
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/person", sendResponse)
	http.ListenAndServe(":8000", route)
}
