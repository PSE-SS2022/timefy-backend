package models

import (
	"testing"
	"time"
)

func TestCanParticipateSuccess1(t *testing.T) {
	timeSlotStart := time.Date(2009, 11, 17, 20, 30, 00, 651387237, time.UTC)
	timeSlotEnd := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)

	slot := TimeSlot{StartTime: timeSlotStart, EndTime: timeSlotEnd}
	slots := []TimeSlot{slot}

	attendant := EventAttendant{UserId: "", PossibleTimes: slots, Role: EventRoleAttendant}
	canParticipate := attendant.CanParticipate(slot)

	if canParticipate {
		t.Logf("Participating test success")
	} else {
		t.Errorf("Participating test failed, got false but expected true")
	}
}

func TestCanParticipateSuccess2(t *testing.T) {
	timeSlot1Start := time.Date(2009, 11, 17, 20, 00, 00, 651387237, time.UTC)
	timeSlot1End := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)
	timeSlot2Start := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	timeSlot2End := time.Date(2009, 11, 17, 21, 30, 00, 651387237, time.UTC)

	slot1 := TimeSlot{StartTime: timeSlot1Start, EndTime: timeSlot1End}
	slot2 := TimeSlot{StartTime: timeSlot2Start, EndTime: timeSlot2End}
	slots1 := []TimeSlot{slot1}

	attendant := EventAttendant{UserId: "", PossibleTimes: slots1, Role: EventRoleAttendant}
	canParticipate := attendant.CanParticipate(slot2)

	if canParticipate {
		t.Logf("Participating test success")
	} else {
		t.Errorf("Participating test failed, got false but expected true")
	}
}

func TestCanParticipateFail1(t *testing.T) {
	timeSlot1Start := time.Date(2009, 11, 17, 21, 30, 00, 651387237, time.UTC)
	timeSlot1End := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)
	timeSlot2Start := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	timeSlot2End := time.Date(2009, 11, 17, 21, 30, 00, 651387237, time.UTC)

	slot1 := TimeSlot{StartTime: timeSlot1Start, EndTime: timeSlot1End}
	slot2 := TimeSlot{StartTime: timeSlot2Start, EndTime: timeSlot2End}

	slots1 := []TimeSlot{slot1}

	attendant := EventAttendant{UserId: "", PossibleTimes: slots1, Role: EventRoleAttendant}

	canParticipate := attendant.CanParticipate(slot2)

	if !canParticipate {
		t.Logf("Participating test success")
	} else {
		t.Errorf("Participating test failed, got true but expected false")
	}
}

func TestCanParticipateFail2(t *testing.T) {
	timeSlot1Start := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	timeSlot1End := time.Date(2009, 11, 17, 21, 30, 00, 651387237, time.UTC)

	timeSlot2Start := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	timeSlot2End := time.Date(2009, 11, 17, 22, 30, 00, 651387237, time.UTC)

	slot1 := TimeSlot{StartTime: timeSlot1Start, EndTime: timeSlot1End}
	slot2 := TimeSlot{StartTime: timeSlot2Start, EndTime: timeSlot2End}

	slots1 := []TimeSlot{slot1}

	attendant := EventAttendant{UserId: "", PossibleTimes: slots1, Role: EventRoleAttendant}

	canParticipate := attendant.CanParticipate(slot2)

	if !canParticipate {
		t.Logf("Participating test success")
	} else {
		t.Errorf("Participating test failed, got true but expected false")
	}
}

func TestIsFull1(t *testing.T) {
	attendants := []EventAttendant{EventAttendant{}, EventAttendant{}}
	event := Event{Attendants: attendants, MaxAmountOfAttendants: 2}

	isFull := event.IsFull()

	if !isFull {
		t.Errorf("Expected event to be full")
	}
}

func TestIsFull2(t *testing.T) {
	attendants := []EventAttendant{EventAttendant{}}
	event := Event{Attendants: attendants, MaxAmountOfAttendants: 2}

	isFull := event.IsFull()

	if isFull {
		t.Errorf("Expected event not to be full")
	}
}
