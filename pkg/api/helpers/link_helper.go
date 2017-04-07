package helpers

import (
	"math"
	"strings"

	"github.com/souzavitor/go-study/pkg/structs"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateShortURL create a new short URL code
func CreateShortURL(session *mgo.Database) (shortURLCode string) {
	var link structs.URL
	var baseValue string
	var err error
	var num int

	collection := session.C("urls")

	err = collection.Find(nil).Sort("-created_at").One(&link)

	if err == mgo.ErrNotFound {
		num = 0
	} else {
		num = toBase10(link.ShortURL, 0)
		num++
	}

	baseValue = toBase(num, 0)

	err = collection.Find(bson.M{"short_url": baseValue}).One(&link)
	for err != mgo.ErrNotFound {
		baseValue = toBase(num, 0)
		err = collection.Find(bson.M{"short_url": baseValue}).One(&link)
		num++
	}
	shortURLCode = baseValue
	return
}

func toBase(num int, b int) string {
	if b == 0 {
		b = 62
	}
	base := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r := int(num % b)
	res := string([]rune(base)[r])
	q := int(math.Floor(float64(num) / float64(b)))
	for q > 0 {
		r = q % b
		q = int(math.Floor(float64(q) / float64(b)))
		res = string([]rune(base)[r]) + res
	}
	return res
}

func toBase10(num string, b int) int {
	if b == 0 {
		b = 62
	}
	var pos int

	base := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	limit := len(num)

	res := strings.Index(base, string([]rune(num)[0]))

	for pos = 1; pos < limit; pos++ {
		res = (b * res) + strings.Index(base, string([]rune(num)[pos]))
	}
	return res
}
