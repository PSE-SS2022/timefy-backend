package handlers

import . "github.com/PSE-SS2022/timefy-backend/internal/models"

type EventHandler struct {
}

func (eventHandler EventHandler) CreateEvent(userId, title, description string, planningType PlanningAlgorithmType, questions map[string][]string) bool {
	return false
}

func (eventHandler EventHandler) JoinEvent(userId, eventId string, answers map[int]int) bool {
	return false
}

func (eventHandler EventHandler) LeaveEvent(userId, eventId string) {

}

func (eventHandler EventHandler) GetEventById(userId string, EventId string) Event {
	var result Event
	return result
}

func (eventHandler EventHandler) GetEventByNameSearch(userId, name string) []Event {
	var result []Event
	return result
}

func (eventHandler EventHandler) GetEventFeed(userId string, filter EventFilter) []Event {
	var result []Event
	return result
}

func (eventHandler EventHandler) GetPlannedEvents(userId string) []Event {
	var result []Event
	return result
}

func (eventHandler EventHandler) EditEvent(editorId, eventId, title, description string) bool {
	return false
}

func (eventHandler EventHandler) CancelEvent(userId, eventId string) {

}
