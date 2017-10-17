package models

// Promotion model
type Promotion struct {
	Title    string `bson:title`
	Name     string `bson:name`
	Category string `bson:category`
	image    string `bson:image`
	url      string `bson:url`
}
