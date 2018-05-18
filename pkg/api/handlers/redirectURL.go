package handlers

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/context"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/souzavitor/go-study/pkg/structs"
)

type MessageResponse struct {
	Error string `json:"error"`
}

// Redirect short URL to long URL
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var url structs.URL
	db := context.Get(r, "database").(*mgo.Database)
	collection := db.C("urls")
	err := collection.Find(bson.M{"short_url" : vars["shortURL"]}).One(&url)
	if err != nil {
		var message MessageResponse
		message.Error = err.Error();

		log.Print(message.Error)
		w.WriteHeader(http.StatusNotFound)

		response, _ := json.Marshal(message)
		fmt.Fprint(w, string(response))
		return
	}

	http.Redirect(w, r, url.LongURL, 301)
}
