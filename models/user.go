package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"FirstName" json:"FirstName"`
	LastName  string             `bson:"LastName" json:"LastName"`
	Email     string             `bson:"Email" json:"Email"`
	Tags      []Tag              `bson:"Tags" json:"Tags"`
}
