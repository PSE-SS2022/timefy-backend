package planning

import (
	. "github.com/PSE-SS2022/timefy-backend/internal/models"
)

const INTERVAL_IN_MINUTES int = 10

type CronJob struct {
}

func (cronJob *CronJob) SetupCronJob() {

}

func (cronJob *CronJob) ScheduleDueEvents() {

}

func (cronJob *CronJob) DeleteOccuredEvents() {

}

type PlanningAlgorithm interface {
	Evaluate(attendants []EventAttendant, timeSlots []TimeSlot) TimeSlot
	Notify(eventId string)
}

type SimplePlanner struct {
}

func (planner *SimplePlanner) Evaluate(attendants []EventAttendant, timeSlots []TimeSlot) TimeSlot {
	var result TimeSlot
	var maxParticipants int = 0

	for _, timeSlot := range timeSlots {

		var currentPossibleParticipants int = 0

		for _, attendant := range attendants {
			if attendant.CanParticipate(timeSlot) {
				currentPossibleParticipants++
			}
		}

		if currentPossibleParticipants > maxParticipants {
			maxParticipants = currentPossibleParticipants
			result = timeSlot
		}
	}

	return result
}

func (planner *SimplePlanner) Notify(eventId string) {

}

type ComplexPlanner struct {
}

func (planner *ComplexPlanner) Evaluate(attendants []EventAttendant, timeSlots []TimeSlot) TimeSlot {
	var result TimeSlot
	return result
}

func (planner *ComplexPlanner) Notify(eventId string) {

}

func (planner *ComplexPlanner) getRegisteredEventsOfAttendant(attendant EventAttendant) []Event {
	var result []Event
	return result
}

func (planner *ComplexPlanner) getTimeSlotsForEvent(event Event) []TimeSlot {
	var result []TimeSlot
	return result
}

func (planner *ComplexPlanner) fetchAvailableTimesForAttendant(attendant EventAttendant) {
}
