package structs

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// URL define a struct to handle with the urls
type URL struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	LongURL   string        `json:"long_url" bson:"long_url"`
	ShortURL  string        `json:"short_url" bson:"short_url"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
