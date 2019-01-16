package database

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

var MongoConnection *mgo.Session

func ConnectMongoDatabase() {
	url := "localhost:27017"
	var err error
	MongoConnection, err = mgo.Dial(url)
	if err != nil {
		log.Fatal("Error connecting mongo database .... :( ", err)
		os.Exit(0)
	}
}
