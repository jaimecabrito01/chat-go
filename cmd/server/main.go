package main

import (
	"chat-go/database"
	"chat-go/internal/user"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool)

func handlerConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true
	for {
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	repo := user.NewRepository(db)
	service := user.NewService(repo)

}
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				delete(clients, client)
				client.Close()
			}
		}
	}
}
