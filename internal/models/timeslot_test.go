package models

import (
	"testing"
	"time"
)

func TestEqualsSuccess1(t *testing.T) {
	start := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)
	end := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	slot := TimeSlot{start, end}

	equals := slot.Equals(slot)

	if !equals {
		t.Errorf("Expected result to be equal")
	}
}

func TestEqualsFail1(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)
	end1 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	start2 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	end2 := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start2, end2}

	equals := slot1.Equals(slot2)

	if equals {
		t.Errorf("Didn't expect result to be equal")
	}
}

func TestEqualsFail2(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)
	end1 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	start2 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start2, end1}

	equals := slot1.Equals(slot2)

	if equals {
		t.Errorf("Didn't expect result to be equal")
	}
}

func TestEqualsFail3(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)
	end1 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	end2 := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start1, end2}

	equals := slot1.Equals(slot2)

	if equals {
		t.Errorf("Didn't expect result to be equal")
	}
}

func TestCollides1(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 20, 00, 00, 0, time.UTC)
	end1 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	start2 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	end2 := time.Date(2009, 11, 17, 24, 00, 00, 0, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start2, end2}

	collides := slot1.collides(slot2)

	if collides {
		t.Errorf("Didn't expect timeslots to collide")
	}
}

func TestCollides2(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 20, 00, 00, 0, time.UTC)
	end1 := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)
	start2 := time.Date(2009, 11, 17, 22, 00, 00, 0, time.UTC)
	end2 := time.Date(2009, 11, 17, 24, 00, 00, 0, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start2, end2}

	collides := slot1.collides(slot2)

	if collides {
		t.Errorf("Expected timeslots not to collide")
	}
}

func TestCollides3(t *testing.T) {
	start1 := time.Date(2009, 11, 17, 20, 00, 00, 0, time.UTC)
	end1 := time.Date(2009, 11, 17, 23, 00, 00, 0, time.UTC)
	start2 := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)
	end2 := time.Date(2009, 11, 17, 24, 00, 00, 0, time.UTC)

	slot1 := TimeSlot{start1, end1}
	slot2 := TimeSlot{start2, end2}

	collides := slot2.collides(slot1)

	if !collides {
		t.Errorf("Expected timeslots to collide")
	}
}

func TestCollides4(t *testing.T) {
	start := time.Date(2009, 11, 17, 20, 00, 00, 0, time.UTC)
	end := time.Date(2009, 11, 17, 21, 00, 00, 0, time.UTC)

	slot := TimeSlot{start, end}

	collides := slot.collides(slot)

	if !collides {
		t.Errorf("Expected timeslots to collide")
	}
}
