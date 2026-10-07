package main

import (
	"chat-go/database"
	"chat-go/handler"
	"chat-go/internal/user"
	"chat-go/router"
	"log"

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

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	repo := user.NewRepository(db)
	service := user.NewService(repo)
	userHandler := handler.NewUserHandler(service)

	router.Init(userHandler)

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
