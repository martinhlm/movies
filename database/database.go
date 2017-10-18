package database

import (
	"fmt"
	"log"
	"movies/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type m map[string]interface{}

var mongodbSession *mgo.Session

// Connect to MongoDB database
func Connect() {
	session := dbSession()
	var err error

	//err = insertPromotions(session)
	//err = updatePromotions(session)
	//err = findPromotion(session)
	//err = iteratePromotions(session)
	//err = insertNestingPromotions(session)
	//err = indexingPromotions(session)
	err = concurrentPromotions(session)

	if err != nil {
		log.Printf("Error, go error: %v\n", err)
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
	var arrays []string //:= ["one", "two", "three"]
	arrays[0] = "one"

	var promotionList = []models.Promotion{
		{"promo titulo", "promo+titulo", "month", "some_image", "some_url",
			arrays, models.Author{"", ""}},
		{"Promotion title", "Promotion+title", "dynamic", "some_url_image",
			"other_url", arrays, models.Author{"", ""}},
		{"Promotion of dynamic", "Promotion+of+dynamic", "dynamic",
			"some_url_image", "other_url", arrays, models.Author{"", ""}},
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

func findPromotion(session *mgo.Session) error {
	promotions := session.DB("local").C("promotions")

	var promo models.Promotion
	err := promotions.Find(models.Promotion{Title: "Otro título"}).One(&promo)
	if err != nil {
		return err
	}

	fmt.Printf("Promotion %v\n", promo)
	return nil
}

func iteratePromotions(session *mgo.Session) error {
	promotions := session.DB("local").C("promotions")

	iter := promotions.Find(nil).Iter()

	var promo models.Promotion
	for iter.Next(&promo) {
		fmt.Printf("Promotion: %v\n", promo)
	}

	return nil
}

func insertNestingPromotions(session *mgo.Session) error {
	promotions := session.DB("local").C("promotions")
	m := map[string]interface{}{
		"title": "nesting title",
		"name":  "nesting+title",
		"tags":  []string{"face", "skin"},
		"author": bson.M{
			"name":  "Martin",
			"email": "martin@fadermex",
		},
	}

	err := promotions.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func indexingPromotions(session *mgo.Session) error {
	promotions := session.DB("local").C("promotions")
	// root field
	err := promotions.EnsureIndexKey("title")

	// nested field
	err = promotions.EnsureIndexKey("author.name")

	return err
}

func concurrentPromotions(session *mgo.Session) error {
	promotions := session.DB("local").C("promotions")

	done := make(chan error)

	go f(promotions, "nesting title", done)
	go f(promotions, "Promotion title", done)

	//if err = firstError(2, done); err != nil {
	//return err
	//}
	return nil
}

func f(promotions *mgo.Collection, title string, done chan error) {
	var promotion models.Promotion
	err := promotions.Find(models.Promotion{Title: title}).One(&promotion)
	if err != nil {
		fmt.Printf("Promotion: %v\n", promotion)
	}

	done <- err
}
