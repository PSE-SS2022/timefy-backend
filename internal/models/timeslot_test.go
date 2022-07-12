package models

import (
	"testing"
	"time"
)

func TestEqualsSuccess1(t *testing.T) {
	start := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	end := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)
	slot := TimeSlot{start, end}

	equals := slot.Equals(slot)

	if equals {
		t.Logf("")
	} else {
		t.Errorf("")
	}
}

func TestEqualsFail1(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	end1 := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)
	start2 := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)
	end2 := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start2, end2}

	equals := slot1.Equals(slot2)

	if !equals {
		t.Logf("")
	} else {
		t.Errorf("")
	}
}

func TestEqualsFail2(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	end1 := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)
	start2 := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start2, end1}

	equals := slot1.Equals(slot2)

	if !equals {
		t.Logf("")
	} else {
		t.Errorf("")
	}
}

func TestEqualsFail3(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)
	end1 := time.Date(2009, 11, 17, 22, 00, 00, 651387237, time.UTC)
	end2 := time.Date(2009, 11, 17, 21, 00, 00, 651387237, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start1, end2}

	equals := slot1.Equals(slot2)

	if !equals {
		t.Logf("")
	} else {
		t.Errorf("")
	}
}
