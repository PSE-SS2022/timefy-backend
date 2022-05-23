package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ReportedUserId string             `bson:"ReportedUserId" json:"ReportedUserId"`
	ReportDate     string             `bson:"ReportDate" json:"ReportDate"`
	EventId        string             `bson:"EventId" json:"EventId"`
}
