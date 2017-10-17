package database

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Connect to MongoDB database
func Connect() {
	_, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success")
}
