package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"` // to ignore the field "-"
}

type Post struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	Caption         string             `json:"caption" bson:"caption"`
	ImageUrl        string             `json:"imageUrl" bson:"imageUrl"`
	PostedTimestamp string             `json:"postedTimestamp" bson:"postedTimestamp"`
	CreatedBy       primitive.ObjectID `json:"createdBy" bson:"createdBy"`
}
