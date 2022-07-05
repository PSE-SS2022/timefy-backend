package models

type Tag struct {
	TagName   string     `bson:"TagName" json:"TagName"`
	TimeSlots []TimeSlot `bson:"TimeSlots" json:"TimeSlots"` // Associated timeslots with that tag
}
