package database

import (
	"fmt"
	"movies/models"

	"gopkg.in/mgo.v2"
)

// Connect to MongoDB database
func Connect() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success")

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

	fmt.Println("Success")
}
