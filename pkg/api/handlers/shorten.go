package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"os"

	"github.com/gorilla/context"
	"github.com/souzavitor/go-study/pkg/api/helpers"
	"github.com/souzavitor/go-study/pkg/structs"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Shorten create a short URL and save in the database
func Shorten(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "database").(*mgo.Database)

	var url structs.URL

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url.ID = bson.NewObjectId()
	url.ShortURL = helpers.CreateShortURL(db)
	url.CreatedAt = time.Now()

	if err := db.C("urls").Insert(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url.ShortURL = os.Getenv("SHORTENER_SCHEME") + "://" + os.Getenv("SHORTENER_HOST") + "/" + url.ShortURL

	response, _ := json.Marshal(url)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(response))
}
