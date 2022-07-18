package database

import (
	"context"
	"log"

	"github.com/PSE-SS2022/timefy-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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

	cur, err := eventCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// iterate through all objects
	for cur.Next(context.TODO()) {
		var event models.Event
		err := cur.Decode(&event)
		if err != nil {
			log.Fatal(err)
		}

		// check if event contains user
		for _, attendant := range event.GetAttendants() {
			if attendant.GetUserId() == user.GetID() {
				events = append(events, event)
				break
			}
		}
	}

	return events
}
