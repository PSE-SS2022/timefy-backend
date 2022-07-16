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

func (timeSlot TimeSlot) collides(toCheckTimeSlot TimeSlot) bool {
	var earliestEnd time.Time
	var SlotToCompare TimeSlot

	if timeSlot.StartTime.Before(toCheckTimeSlot.StartTime) {
		earliestEnd = timeSlot.EndTime
		SlotToCompare = toCheckTimeSlot
	} else {
		earliestEnd = toCheckTimeSlot.EndTime
		SlotToCompare = timeSlot
	}

	// if first timeslots ends before other one starts they don't collide
	if earliestEnd.Before(SlotToCompare.StartTime) || earliestEnd.Equal(SlotToCompare.StartTime) {
		return false
	}

	return true
}
