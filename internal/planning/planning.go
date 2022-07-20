package planning

import (
	"time"

	"github.com/PSE-SS2022/timefy-backend/internal/database"
	. "github.com/PSE-SS2022/timefy-backend/internal/models"
	"github.com/jasonlvhit/gocron"
)

const INTERVAL_IN_MINUTES_SCHEDULE uint64 = 10

type CronJob struct {
}

func (cronJob *CronJob) SetupCronJob() {
	if cronJob != nil {
		gocron.Every(INTERVAL_IN_MINUTES_SCHEDULE).Minutes().Do(cronJob.ScheduleDueEvents)
	}
}

func (cronJob *CronJob) ScheduleDueEvents() {
	for _, event := range database.EventRepositoryInstance.GetEvents() {
		if !event.GetIsScheduled() && time.Now().After(event.GetDeadline()) {
			// get planner based on type used for event creation
			planner := GetPlanner(event)
			timeSlot := planner.Evaluate(event.GetAttendants(), event.GetPossibleTimes())

			// set results in event
			event.SetPlannedTimeSlot(timeSlot)
			event.SetIsScheduled(true)
			// update event in database
			database.EventRepositoryInstance.UpdateEvent(event)
		}
	}
}

func (cronJob *CronJob) DeleteOccuredEvents() {

}

func GetPlanner(event Event) PlanningAlgorithm {
	var planner PlanningAlgorithm

	switch event.GetPlanningAlgorithmType() {
	case PlanningAlgorithmTypeSimple:
		return SimplePlanner{}
	case PlanningAlgorithmTypeComplex:
		return ComplexPlanner{}
	default:
	}

	return planner
}

type PlanningAlgorithm interface {
	Evaluate(attendants []EventAttendant, timeSlots []TimeSlot) TimeSlot
	Notify(eventId string)
}

type SimplePlanner struct {
}

func (planner SimplePlanner) Evaluate(attendants []EventAttendant, timeSlots []TimeSlot) TimeSlot {
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

func (planner SimplePlanner) Notify(eventId string) {

}

type ComplexPlanner struct {
}

func (planner ComplexPlanner) Evaluate(attendants []EventAttendant, timeSlots []TimeSlot) TimeSlot {
	var result TimeSlot
	var maxParticipants float64 = 0

	for _, timeSlot := range timeSlots {

		var currentPossibleParticipants float64 = 0

		for _, attendant := range attendants {
			user, result := database.UserRepositoryInstance.GetUserById(attendant.GetUserId())

			if attendant.CanParticipate(timeSlot) && result && !planner.userHasPlannedEventAtTime(user, timeSlot) {
				// get amount of possible timeslots from registrations that collide with given one
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

func (planner ComplexPlanner) Notify(eventId string) {

}

// todo divide through amount of timeslots of event
func (planner ComplexPlanner) getAmountOfPotentialEventsAtTime(user User, timeSlot TimeSlot) int {
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

func (planner ComplexPlanner) getRegisteredEventsOfUser(user User) []Event {
	events := database.EventRepositoryInstance.GetEventsOfUser(user)
	return events
}

func (planner ComplexPlanner) getUserAttendantDataOfEvent(user User, event Event) EventAttendant {
	var attendant EventAttendant

	for _, attendant := range event.GetAttendants() {
		if attendant.GetUserId() == user.GetID() {
			return attendant
		}
	}

	return attendant
}

func (planner ComplexPlanner) userHasPlannedEventAtTime(user User, slot TimeSlot) bool {
	for _, scheduled := range user.GetScheduledEvents() {

		event, result := database.EventRepositoryInstance.GetEventById(scheduled.GetEventId())

		if result {
			continue
		}

		for _, timeSlot := range event.GetPossibleTimes() {
			if timeSlot.Collides(slot) {
				return true
			}
		}
	}

	return false
}
