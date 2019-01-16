package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitiateSocketServer(router *gin.Engine) {
	// Configure websocket route
	router.GET("/ws", func(c *gin.Context) {
		handleConnections(c.Writer, c.Request)
	})
	// go handleMessages()
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// Save the socket client in some manner. so that we could get back
	// database.RedisClient.Set("socket", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// to the client for sending the messages
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		// Based on the message type we could publish the message to the
		// individual receiver or to a group
		// Send the newly received message to the broadcast channel
		// broadcast <- msg
	}
}
