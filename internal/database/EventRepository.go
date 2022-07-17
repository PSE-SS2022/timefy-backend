package database

import (
	"github.com/PSE-SS2022/timefy-backend/internal/models"
)

var EventRepositoryInstance EventRepository

const EVENT_REPO = "users"

type EventRepository struct {
}

func (eventRepository EventRepository) GetEventsOfUser(user models.User) []models.Event {
	var events []models.Event

	eventCollection := databaseMgrInstance.getCollection((EVENT_REPO))
	if eventCollection != nil {
		return events
	}
	//eventCollection.Find(context.TODO(), bson.M{"id": id}).Decode(&user)
	if user.ID.IsZero() {
		return events
	}

	return events
}

func (eventRepository EventRepository) GetAttendantData(user models.User, event models.Event) models.EventAttendant {
	var result models.EventAttendant
	return result
}
