package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ReportedUserId string             `bson:"ReportedUserId" json:"ReportedUserId"`
	ReportDate     string             `bson:"ReportDate" json:"ReportDate"`
	EventId        string             `bson:"EventId" json:"EventId"`
	AmountReports  int                `bson:"AmountReports" json:"AmountReports"`
}

type ExtendedReport struct {
	ID             string `bson:"ID" json:"ID"`
	ReportedUserId string `bson:"ReportedUserId" json:"ReportedUserId"`
	FirstName      string `bson:"FirstName" json:"FirstName"`
	LastName       string `bson:"LastName" json:"LastName"`
	ReportDate     string `bson:"ReportDate" json:"ReportDate"`
	EventId        string `bson:"EventId" json:"EventId"`
	Title          string `bson:"Title" json:"Title"`
	Description    string `bson:"Description" json:"Description"`
}
