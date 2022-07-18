package planning

import (
	"github.com/PSE-SS2022/timefy-backend/internal/database"
	"github.com/PSE-SS2022/timefy-backend/internal/models"
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
	var maxParticipants float64 = 0

	for _, timeSlot := range timeSlots {

		var currentPossibleParticipants float64 = 0

		for _, attendant := range attendants {
			user, result := database.UserRepositoryInstance.GetUserById(attendant.GetUserId())

			if attendant.CanParticipate(timeSlot) && result && !userHasPlannedEventAtTime(user, timeSlot) {
				// get amount of possible timeslots from registrations that collide with given on
				var possibleEvents float64 = float64(planner.getAmountOfPotentialEventsAtTime(user, timeSlot))
				if possibleEvents <= 0 {
					possibleEvents = 1
				}

				currentPossibleParticipants += (1 / possibleEvents)
			}
		}

		if currentPossibleParticipants > maxParticipants {
			maxParticipants = currentPossibleParticipants
			result = timeSlot
		}
	}

	return result
}

func (planner *ComplexPlanner) Notify(eventId string) {

}

func (planner *ComplexPlanner) getAmountOfPotentialEventsAtTime(user User, timeSlot TimeSlot) int {
	var amountOfPotentialEvents int = 0

	events := planner.getRegisteredEventsOfUser(user)

	for _, event := range events {
		// get Attendant for user
		attendant := planner.getUserAttendantDataOfEvent(user, event)

		// if we have a timeslot collision we are less likely to attend
		for _, timeSlots := range attendant.PossibleTimes {
			if timeSlots.Collides(timeSlot) {
				amountOfPotentialEvents++
			}
		}
	}

	return amountOfPotentialEvents
}

func (planner *ComplexPlanner) getRegisteredEventsOfUser(user User) []Event {
	events := database.EventRepositoryInstance.GetEventsOfUser(user)
	return events
}

func (planner *ComplexPlanner) getUserAttendantDataOfEvent(user User, event Event) EventAttendant {
	var attendant models.EventAttendant

	for _, attendant := range event.GetAttendants() {
		if attendant.GetUserId() == user.GetID() {
			return attendant
		}
	}

	return attendant
}

func userHasPlannedEventAtTime(user User, slot TimeSlot) bool {

	for _, event := range user.GetScheduledEvents() {

		for _, timeSlot := range event.GetPossibleTimes() {
			if timeSlot.Collides(slot) {
				return true
			}
		}
	}

	return false
}
