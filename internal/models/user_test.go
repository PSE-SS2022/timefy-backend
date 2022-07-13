package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHasOperation(t *testing.T) {
	var testObjectId primitive.ObjectID
	var testMap map[string]string
	user := NewUser(testObjectId, "", "", "", "", testMap)

	if !user.HasOperationAvailable(EventEditLimiation) {
		t.Fatalf("Expected user to have an edit operation")
	}
}

func TestConsumerOperation1(t *testing.T) {
	var testObjectId primitive.ObjectID
	var testMap map[string]string
	user := NewUser(testObjectId, "", "", "", "", testMap)

	user.ConsumeOperation(EventEditLimiation)

	if user.HasOperationAvailable(EventEditLimiation) {
		t.Fatalf("Expected user not to have an edit operation")
	}
}

func TestConsumerOperation2(t *testing.T) {
	var testObjectId primitive.ObjectID
	var testMap map[string]string
	user := NewUser(testObjectId, "", "", "", "", testMap)

	user.ConsumeOperation(UserReportLimitation)

	if !user.HasOperationAvailable(UserReportLimitation) {
		t.Fatalf("Expected user to have a report operation")
	}
}

func TestConsumerOperation3(t *testing.T) {
	var testObjectId primitive.ObjectID
	var testMap map[string]string
	user := NewUser(testObjectId, "", "", "", "", testMap)

	for x := 0; x < 3; x++ {
		user.ConsumeOperation(UserReportLimitation)
	}

	if !user.HasOperationAvailable(UserReportLimitation) {
		t.Fatalf("Expected user to have a report operation")
	}
}

func TestConsumerOperation4(t *testing.T) {
	var testObjectId primitive.ObjectID
	var testMap map[string]string
	user := NewUser(testObjectId, "", "", "", "", testMap)

	for x := 0; x < 5; x++ {
		user.ConsumeOperation(UserReportLimitation)
	}

	if user.HasOperationAvailable(UserReportLimitation) {
		t.Fatalf("Expected user not to have a report operation")
	}
}

func TestConsumerOperation5(t *testing.T) {
	var testObjectId primitive.ObjectID
	var testMap map[string]string
	user := NewUser(testObjectId, "", "", "", "", testMap)

	user.ConsumeOperation(EventEditLimiation)

	if !user.HasOperationAvailable(EventCreationLimitation) {
		t.Fatalf("Expected user to have an event creation operation")
	}
}
