package models

import "time"

type TimeSlot struct {
	StartTime time.Time `bson:"StartTime" json:"StartTime"`
	EndTime   time.Time `bson:"EndTime" json:"EndTime"`
}

func (t TimeSlot) StartToString() string {
	return t.StartTime.GoString()
}

func (t TimeSlot) EndToString() string {
	return t.StartTime.GoString()
}
