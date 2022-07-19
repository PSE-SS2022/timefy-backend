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
	TimeSlot              TimeSlot              `bson:"TimeSlot" json:"TimeSlot"`
}

func (event *Event) SetPlannedTimeSlot(timeSlot TimeSlot) {
	event.TimeSlot = timeSlot
}

func (event Event) GetIsScheduled() bool {
	return event.IsScheduled
}

func (event *Event) SetIsScheduled(isScheduled bool) {
	event.IsScheduled = isScheduled
}

func (event Event) GetDeadline() time.Time {
	return event.JoinDeadline
}

func (event Event) GetPlanningAlgorithmType() PlanningAlgorithmType {
	return event.Algorithm
}

func (event Event) GetAttendants() []EventAttendant {
	return event.Attendants
}

func (event Event) GetPossibleTimes() []TimeSlot {
	return event.PossibleTimes
}

func (event Event) ContainsUser(userId string) bool {
	for _, attendant := range event.Attendants {
		if attendant.GetUserId() == userId {
			return true
		}
	}

	return false
}

func (event Event) IsFull() bool {
	return len(event.Attendants) == event.MaxAmountOfAttendants
}

func (event Event) DeadlineArrived() bool {
	return time.Now().After(event.JoinDeadline)
}

type UserQuestionairInput map[int](map[int]bool)

func (event *Event) ParseUserQuestionaire(answers UserQuestionairInput) {

}

type MCAnswer struct {
	Answer          string   `bson:"Answer" json:"Answer"`
	SelectedUserIds []string `bson:"SelectedUserIds" json:"SelectedUserIds"`
}

func (answer *MCAnswer) AddUser(userId string) {
	answer.SelectedUserIds = append(answer.SelectedUserIds, userId)
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
	EventRoleOwner EventRole = iota
	EventRoleOrganizer
	EventRoleHelper
	EventRoleAttendant
)

type EventAttendant struct {
	UserId        string     `bson:"UserId" json:"UserId"`
	PossibleTimes []TimeSlot `bson:"PossibleTimes" json:"PossibleTimes"`
	Role          EventRole  `bson:"Role" json:"Role"`
}

func (attendant EventAttendant) CanParticipate(slot TimeSlot) bool {

	for _, possibleTime := range attendant.PossibleTimes {
		startIsBefore := possibleTime.StartTime.Before(slot.StartTime) || possibleTime.StartTime.Equal(slot.StartTime)
		endIsAfter := possibleTime.EndTime.After(slot.EndTime) || possibleTime.EndTime.Equal(slot.EndTime)
		if startIsBefore && endIsAfter {
			return true
		}
	}

	return false
}

func (attendant EventAttendant) GetUserId() string {
	return attendant.UserId
}

type EventCategory int

const (
	EventCategorySports EventCategory = iota
	EventCategoryFestival
	EventCategoryStudy
	EventCategoryReading
	EventCategoryEating
	EventCategorySocial
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
	return eventFilterCategory.Category == event.Category
}
