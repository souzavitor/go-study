package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/souzavitor/go-study/pkg/api/helpers"

	mgo "gopkg.in/mgo.v2"
)

// Create a new collection and ensure its indexes
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. ", err)
	}

	session := helpers.CreateDatabaseConnection()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(os.Getenv("DB_DATABASE")).C("urls")
	index := mgo.Index{
		Key:    []string{"short_url"},
		Unique: true,
	}

	err = collection.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
