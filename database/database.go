package database

import (
	"fmt"
	"log"
	"movies/models"

	"gopkg.in/mgo.v2"
)

type m map[string]interface{}

var mongodbSession *mgo.Session

// Connect to MongoDB database
func Connect() {
	session := dbSession()

	insertPromotions(session)
}

func dbSession() *mgo.Session {
	var session *mgo.Session
	var err error
	if mongodbSession == nil {
		session, err = mgo.Dial("localhost")
		if err != nil {
			log.Fatalf("Can't connect to mongo, go error %v\n", err)
		}

	}
	return session
}

func insertPromotions(session *mgo.Session) {
	promotions := session.DB("local").C("promotions")
	var promotionList = []models.Promotion{
		{"promo titulo", "promo+titulo", "month", "some_image", "some_url"},
		{"Promotion title", "Promotion+title", "dynamic", "some_url_image", "other_url"},
		{"Promotion of dynamic", "Promotion+of+dynamic", "dynamic", "some_url_image", "other_url"},
	}

	for _, promotion := range promotionList {
		err := promotions.Insert(promotion)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Insert success")
}
