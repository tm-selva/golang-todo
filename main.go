package main

import (
	"log"
	"net/http"

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
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	authRoute.Routes(router)

	// setting up server
	// err := http.ListenAndServe(":3000", nil)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
