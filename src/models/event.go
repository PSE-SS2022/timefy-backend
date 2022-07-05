package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `bson:"Title" json:"Title"`
	Description  string             `bson:"Description" json:"Description"`
	CreationDate string             `bson:"CreationDate" json:"CreationDate"`
	CreatorId    string             `bson:"CreatorId" json:"CreatorId"`
}
