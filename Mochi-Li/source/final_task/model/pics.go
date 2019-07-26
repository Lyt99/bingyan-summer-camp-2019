package model

type Pic struct {
	Url  string `json:"url" bson:"url"`
	Name string `json:"name" bson:"name"`
}
