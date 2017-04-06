package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/souzavitor/go-study/pkg/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var router = api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
