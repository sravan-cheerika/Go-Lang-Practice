package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Global
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
	Message string `json:"message"`
}

var totalNumOfConnections int

func HandleClients(w http.ResponseWriter, r *http.Request) {
	go broadcastMessagesToClients()

	websocket, err := upgrader.Upgrade(w, r, nil) // returning *websocket.Conn
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket :: ", err)
	}
	defer websocket.Close()
	//Save all  Clients data base
	clients[websocket] = true

	for i := range clients {
		
		fmt.Println(i.RemoteAddr(), " ", i.LocalAddr(), "  ", websocket.RemoteAddr(), " ")
		fmt.Println("time stamp : ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("total Connections : ", totalNumOfConnections)
	}
	for {
		var message Message
		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			delete(clients, websocket)
			break
		}
		broadcast <- message
	}
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/echo", HandleClients)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
func broadcastMessagesToClients() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("error occurred while writing message to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
