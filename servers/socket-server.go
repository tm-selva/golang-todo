package server

import (
	"log"
	"net/http"

	socketservice "../services/socket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

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
	go socketservice.InitiateSocketWorker()
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Save the socket client in some manner. so that we could get back
	var userId uuid.UUID
	userId, _ = uuid.NewV4()
	socketClient := socketservice.Register(ws, userId.String())
	// database.RedisClient.Set("socket", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// listen for the incoming socket message
	go socketClient.IncomingSocketMessageListener()
}
