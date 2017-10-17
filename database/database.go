package database

import (
	"log"
	"movies/models"

	"gopkg.in/mgo.v2"
)

type m map[string]interface{}

var mongodbSession *mgo.Session

// Connect to MongoDB database
func Connect() {
	session := dbSession()
	var err error

	/*
		err = insertPromotions(session)
		if err != nil {
			fmt.Println("Insert success")
		}
	*/

	err = updatePromotions(session)
	if err != nil {
		log.Printf("Error update promotions, go error: %v\n", err)
	}
}

func dbSession() *mgo.Session {
	var session *mgo.Session
	var err error
	if mongodbSession == nil {
		session, err = mgo.Dial("localhost")
		if err != nil {
			log.Fatalf("Can't connect to mongo, go error: %v\n", err)
		}

	}
	return session
}

func insertPromotions(session *mgo.Session) error {
	promotions := session.DB("local").C("promotions")
	var promotionList = []models.Promotion{
		{"promo titulo", "promo+titulo", "month", "some_image", "some_url"},
		{"Promotion title", "Promotion+title", "dynamic", "some_url_image", "other_url"},
		{"Promotion of dynamic", "Promotion+of+dynamic", "dynamic", "some_url_image", "other_url"},
	}

	for _, promotion := range promotionList {
		err := promotions.Insert(promotion)
		if err != nil {
			return err
		}
	}
	return nil
}

func updatePromotions(session *mgo.Session) error {
	promotions := session.DB("local").C("promotions")
	change := m{"$set": models.Promotion{
		Title: "Otro título", Name: "Otro+título",
	}}

	/*
		promo := models.Promotion{
			Title:    "promo titulo",
			Name:     "promo+titulo",
			Category: "month",
			Image:    "some_image",
			URL:      "some_url",
		}
	*/

	err := promotions.Update(models.Promotion{Title: "Un titulo"}, change)
	if err != nil {
		return err
	}
	return nil
}
