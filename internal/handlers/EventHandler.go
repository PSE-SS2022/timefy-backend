package handlers

import (
	"github.com/PSE-SS2022/timefy-backend/internal/database"
	. "github.com/PSE-SS2022/timefy-backend/internal/models"
)

func CreateEvent(userId, title, description string, planningType PlanningAlgorithmType, questions map[string][]string) bool {
	return false
}

func JoinEvent(userId, eventId string, answers map[int]int) bool {
	return false
}

func LeaveEvent(userId, eventId string) {

}

func GetEventById(userId, eventId string) (Event, bool) {
	event, result := database.EventRepositoryInstance.GetEventById(eventId)

	if !result {
		return event, false
	}

	if UserCanSeeEvent(userId, eventId) {
		return event, true
	}

	return event, false
}

func UserCanSeeEvent(userId, eventId string) bool {
	return false
}

func GetEventByNameSearch(userId, name string) []Event {
	var result []Event
	return result
}

func GetEventFeed(userId string, filter EventFilter) []Event {
	var result []Event
	return result
}

func GetPlannedEvents(userId string) []Event {
	var result []Event
	return result
}

func EditEvent(editorId, eventId, title, description string) bool {
	return false
}

func CancelEvent(userId, eventId string) {

}
