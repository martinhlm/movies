package models

// Promotion model
type Promotion struct {
	Title    string   `bson:"title,omitempty"`
	Name     string   `bson:"name,omitempty"`
	Category string   `bson:"category,omitempty"`
	Image    string   `bson:"image,omitempty"`
	URL      string   `bson:"url,omitempty"`
	Tags     []string `bson:"tags,omitempty"`
}
