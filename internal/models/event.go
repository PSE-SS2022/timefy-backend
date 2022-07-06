package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID                    primitive.ObjectID    `bson:"_id,omitempty"`
	Title                 string                `bson:"Title" json:"Title"`
	Description           string                `bson:"Description" json:"Description"`
	CreationDate          string                `bson:"CreationDate" json:"CreationDate"`
	CreatorId             string                `bson:"CreatorId" json:"CreatorId"`
	PossibleTimes         []TimeSlot            `bson:"PossibleTimes" json:"PossibleTimes"`
	MultipleChoice        MCQuestionair         `bson:"MultipleChoice" json:"MultipleChoice"`
	Attendants            []EventAttendant      `bson:"Attendants" json:"Attendants"`
	Algorithm             PlanningAlgorithmType `bson:"Algorithm" json:"Algorithm"`
	MaxAmountOfAttendants int                   `bson:"MaxAmountOfAttendants" json:"MaxAmountOfAttendants"`
	Category              EventCategory         `bson:"Category" json:"Category"`
	LastEditorId          string                `bson:"LastEditorId" json:"LastEditorId"`
	Visibility            EventVisibility       `bson:"Visibility" json:"Visibility"`
	JoinDeadline          time.Time             `bson:"JoinDeadline" json:"JoinDeadline"`
	IsScheduled           bool                  `bson:"IsScheduled" json:"IsScheduled"`
	EventDate             time.Time             `bson:"EventDate" json:"EventDate"`
}

func (event Event) IsFull() bool {
	return false
}

func (event Event) DeadlineArrived() bool {
	return false
}

type UserQuestionairInput map[int](map[int]bool)

func (event *Event) ParseUserQuestionaire(answers UserQuestionairInput) {

}

type MCAnswer struct {
	Answer          string   `bson:"Answer" json:"Answer"`
	SelectedUserIds []string `bson:"SelectedUserIds" json:"SelectedUserIds"`
}

func (answer *MCAnswer) AddUser(userId string) {

}

type MCQuestion struct {
	Question string           `bson:"Question" json:"Question"`
	Answers  map[int]MCAnswer `bson:"Answers" json:"Answers"`
}

func (answer MCQuestion) GetAnswers() map[int]MCAnswer {
	var result map[int]MCAnswer
	return result
}

func (answer MCQuestion) GetUsersForAnswer(index int) string {
	return ""
}

type MCQuestionair struct {
	Question map[int]MCQuestion `bson:"Question" json:"Question"`
}

func (questionair MCQuestionair) GetQuestions() map[int]MCQuestion {
	var result map[int]MCQuestion
	return result
}

func (questionair MCQuestionair) GetUsersForAnswers(index int) []string {
	var result []string
	return result
}

type EventRole int

const (
	EventRoleOwner = iota
	EventRoleOrganizer
	EventRoleHelper
	EventRoleAttendant
)

type EventAttendant struct {
	UserId        string    `bson:"UserId" json:"UserId"`
	PossibleTimes TimeSlot  `bson:"PossibleTimes" json:"PossibleTimes"`
	Role          EventRole `bson:"Role" json:"Role"`
}

func (attendant EventAttendant) CanParticipate(slot TimeSlot) bool {
	return false
}

func (attendant EventAttendant) GetUser() User {
	var result User
	return result
}

type EventCategory int

const (
	EventCategorySports EventCategory = iota
	EventCategoryFestival
	EventCategoryStudy
	EventCategoryReading
)

type EventVisibility int

const (
	EventVisibilityPublic EventVisibility = iota
	EventVisibilityGroup
	EventVisibilityFriendsOnly
	EventVisibilityPrivate
)

type EventFilter interface {
	IsValidEvent(event Event) bool
}

type EventFilterCategory struct {
	Category EventCategory `bson:"Category" json:"Category"`
}

func (eventFilterCategory EventFilterCategory) IsValidEvent(event Event) bool {
	return false
}
