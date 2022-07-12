package models

import (
	"testing"
	"time"
)

func TestCanParticipateSuccess1(t *testing.T) {
	timeSlotStart := time.Date(2009, 11, 17, 20, 30, 00, 651387237, time.UTC)
	timeSlotEnd := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)

	slot := TimeSlot{timeSlotStart, timeSlotEnd}
	slots := []TimeSlot{slot}

	attendant := EventAttendant{"", slots, EventRoleAttendant}
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

	slot1 := TimeSlot{timeSlot1Start, timeSlot1End}
	slot2 := TimeSlot{timeSlot2Start, timeSlot2End}
	slots1 := []TimeSlot{slot1}

	attendant := EventAttendant{"", slots1, EventRoleAttendant}
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

	slot1 := TimeSlot{timeSlot1Start, timeSlot1End}
	slot2 := TimeSlot{timeSlot2Start, timeSlot2End}

	slots1 := []TimeSlot{slot1}

	attendant := EventAttendant{"", slots1, EventRoleAttendant}

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

	slot1 := TimeSlot{timeSlot1Start, timeSlot1End}
	slot2 := TimeSlot{timeSlot2Start, timeSlot2End}

	slots1 := []TimeSlot{slot1}

	attendant := EventAttendant{"", slots1, EventRoleAttendant}

	canParticipate := attendant.CanParticipate(slot2)

	if !canParticipate {
		t.Logf("Participating test success")
	} else {
		t.Errorf("Participating test failed, got true but expected false")
	}
}
