package models

type TimeSlot struct {
	BeginTime Time `bson:"BeginTime" json:"BeginTime"`
	EndTime   Time `bson:"EndTime" json:"EndTime"`
}
