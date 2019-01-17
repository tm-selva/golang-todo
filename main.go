package main

import (
	"log"

	"github.com/gin-gonic/gin"

	database "./database"
	authRoute "./routes/auth"
	server "./servers"
)

// require the config file based on the environment dev/prod
func main() {

	// Initiate the database and create connection
	database.InitiateAllDatabase()

	// Set the router as the default
	router := gin.Default()

	// Initiate the websocket server
	server.InitiateSocketServer(router)

	// Serve frontend static files

	router.Static("/index", "public")

	authRoute.Routes(router)

	err := router.Run(":3000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
