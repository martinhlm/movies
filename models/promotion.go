package models

// Promotion model
type Promotion struct {
	Title    string `bson:"title,omitempty"`
	Name     string `bson:"name"`
	Category string `bson:"category"`
	Image    string `bson:"image"`
	URL      string `bson:"url"`
}
