package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL  string             `json:"imageurl" bson:"imageurl,omitempty"`
	Timestamp string             `json:"creation"`
	UserId    string             `json:"userId" bson:"userid"`
}

