package helpers

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// CreateDatabaseConnection create a new connection
func CreateDatabaseConnection() (session *mgo.Session) {
	connectionString := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Fatal("could not dial mongo. ", err)
	}
	return
}
