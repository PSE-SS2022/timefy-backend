package models

import (
	"context"
	"time"

	"github.com/PSE-SS2022/timefy-backend/internal/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USER_REPO = "users"

type LimitedOperationType int

const (
	EventEditLimiation LimitedOperationType = iota
	EventCreationLimitation
	UserProfilePictureSetLimitation
	UserReportLimitation
)

var LimitedOperationAmounts []int = []int{1, 1, 1, 5}
var LimitedOperationDelay []time.Time = []time.Time{Seconds(30), Seconds(60), Seconds(300), Seconds(30)}

func Seconds(seconds int) time.Time {
	return time.Date(0, 0, 0, 0, 0, seconds, 0, time.UTC)
}

type InviteType int

const (
	InviteUserFriendRequest InviteType = iota
	InviteGroupRequest
)

type Invite struct {
	Id   string
	Type InviteType
}

type User struct {
	ID                 primitive.ObjectID                   `bson:"_id,omitempty"`
	UID                string                               `bson:"UID" json:"UID"`
	FirstName          string                               `bson:"FirstName" json:"FirstName"`
	LastName           string                               `bson:"LastName" json:"LastName"`
	Email              string                               `bson:"Email" json:"Email"`
	Roles              map[string]string                    `bson:"Roles" json:"Roles"`
	AmountWarnings     int                                  `bson:"AmountWarnings" json:"AmountWarnings"`
	BannedUntil        time.Time                            `bson:"BannedUntil" json:"BannedUntil"`
	NotificationTokens []string                             `bson:"NotificationTokens" json:"NotificationTokens"`
	ScheduledEvents    []ScheduledEvent                     `bson:"ScheduledEvents" json:"ScheduledEvents"`
	LimitedOperations  map[LimitedOperationType][]time.Time `bson:"LimitedOperations" json:"LimitedOperations"`
	BlockedUsers       []string                             `bson:"BlockedUsers" json:"BlockedUsers"`
	Invites            []Invite                             `bson:"Invites" json:"Invites"`
}

func GetUserByID(id string) (User, bool) {
	var user User
	usersCollection := repos.GetCollection((USER_REPO))
	if usersCollection != nil {
		return user, false
	}
	usersCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if user.ID.IsZero() {
		return user, false
	}
	return user, true
}

// this should be calculated in userhandler on registration/first login
func (user *User) SetUID(uid string) {
	user.UID = uid
}

func (user *User) SetFcmToken(fcmToken string) {

}

func (user User) GetResetTime(operationType LimitedOperationType) time.Time {
	return LimitedOperationDelay[operationType]
}

func (user User) GetAmount(operationType LimitedOperationType) int {
	return LimitedOperationAmounts[operationType]
}

func (user User) HasOperationAvailable(operationType LimitedOperationType) bool {
	operations := user.LimitedOperations[operationType]
	amount := user.GetAmount(operationType)
	resetTime := user.GetResetTime(operationType)

	for x := 0; x < amount; x++ {
		timeToCompare := time.Since(operations[x])
		if timeToCompare.Seconds() >= float64(resetTime.Second()) {
			return true
		}
	}

	return false
}

func (user *User) ConsumeOperation(operationType LimitedOperationType) {
	operations := user.LimitedOperations[operationType]
	amount := user.GetAmount(operationType)
	resetTime := user.GetResetTime(operationType)

	for x := 0; x < amount; x++ {
		timeToCompare := time.Since(operations[x])
		if timeToCompare.Seconds() >= float64(resetTime.Second()) {
			operations[x] = time.Now()
			return
		}
	}
}

func (user *User) HandleInvite(objectId string, inviteType InviteType, answer bool) {

}

func (user *User) Friendinvite(userId string, answer bool) {

}

func (user *User) GroupInvite(groupId string, answer bool) {

}

func (user *User) GetScheduledEvents() []Event {
	var result []Event
	return result
}

type ScheduledEvent struct {
	EventId string `bson:"EventId" json:"EventId"`
	Synced  bool   `bson:"Synced" json:"Synced"`
}

type FriendRelation struct {
	Id      string `bson:"Id" json:"Id"`
	UserAId string `bson:"UserAId" json:"UserAId"`
	UserBId string `bson:"UserBId" json:"UserBId"`
}

type UserRelation int

const (
	UserRelationFriends UserRelation = iota
	UserRelationBlocked
	UserRelationNotFriends
)
