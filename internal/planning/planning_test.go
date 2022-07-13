package planning

import (
	"testing"
	"time"

	. "github.com/PSE-SS2022/timefy-backend/internal/models"
)

func TestSimpleEvaluate(t *testing.T) {
	planner := SimplePlanner{}
	timeSlot1Start := time.Date(2009, 11, 17, 20, 30, 00, 651387237, time.UTC)
	timeSlot1End := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)

	timeSlot2Start := time.Date(2009, 11, 18, 20, 30, 00, 651387237, time.UTC)
	timeSlot2End := time.Date(2009, 11, 18, 21, 00, 00, 651387237, time.UTC)

	slot1 := TimeSlot{StartTime: timeSlot1Start, EndTime: timeSlot1End}
	slot2 := TimeSlot{StartTime: timeSlot2Start, EndTime: timeSlot2End}

	slots1 := []TimeSlot{slot1}
	slots2 := []TimeSlot{slot2}
	slotsEvent := []TimeSlot{slot1, slot2}

	attendant1 := EventAttendant{UserId: "", PossibleTimes: slots1, Role: EventRoleAttendant}
	attendant2 := EventAttendant{UserId: "", PossibleTimes: slots2, Role: EventRoleAttendant}
	attendant3 := EventAttendant{UserId: "", PossibleTimes: slots2, Role: EventRoleAttendant}

	attendants := []EventAttendant{attendant1, attendant2, attendant3}
	result := planner.Evaluate(attendants, slotsEvent)

	if result.Equals(slot2) {
		t.Logf("")
	} else {
		t.Errorf("")
	}
}
