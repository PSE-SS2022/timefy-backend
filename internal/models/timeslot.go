package models

import "time"

type TimeSlot struct {
	StartTime time.Time `bson:"StartTime" json:"StartTime"`
	EndTime   time.Time `bson:"EndTime" json:"EndTime"`
}

func (timeSlot TimeSlot) StartToString() string {
	return timeSlot.StartTime.GoString()
}

func (timeSlot TimeSlot) EndToString() string {
	return timeSlot.StartTime.GoString()
}

func (timeSlot TimeSlot) Equals(compare TimeSlot) bool {
	return (timeSlot.StartTime.Equal(compare.StartTime)) && (timeSlot.EndTime.Equal(compare.EndTime))
}
