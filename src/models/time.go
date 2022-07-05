package models

// TODO: Enforce specific format with these
type Time struct {
	Time timeEntry `bson:"Time" json:"Time"`
	Day  dayEntry  `bson:"Day" json:"Day"`
}

// TODO: Maybe ditch below and enforce format in other way
type timeEntry struct {
}

type dayEntry struct {
}
