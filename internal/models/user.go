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

type InviteType int

const (
	InviteUserFriendRequest InviteType = iota
	InviteGroupRequest
)

type User struct {
	ID                 primitive.ObjectID                   `bson:"_id,omitempty"`
	UID                string                               `bson:"UID" json:"UID"`
	FirstName          string                               `bson:"FirstName" json:"FirstName"`
	LastName           string                               `bson:"LastName" json:"LastName"`
	Email              string                               `bson:"Email" json:"Email"`
	Roles              []string                             `bson:"Roles" json:"Roles"` // ?
	AmountWarnings     int                                  `bson:"AmountWarnings" json:"AmountWarnings"`
	BannedUntil        time.Time                            `bson:"BannedUntil" json:"BannedUntil"`
	NotificationTokens []string                             `bson:"NotificationTokens" json:"NotificationTokens"`
	ScheduledEvents    []ScheduledEvent                     `bson:"ScheduledEvents" json:"ScheduledEvents"`
	LimitedOperations  map[LimitedOperationType][]time.Time `bson:"LimitedOperations" json:"LimitedOperations"`
	BlockedUsers       []string                             `bson:"BlockedUsers" json:"BlockedUsers"`
	//Invites []
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

func (user User) CalculateUID(lastName string) {

}

func (user User) SetFcmToken(fcmToken string) {

}

func (user User) GetResetTime(operationType LimitedOperationType) time.Duration {
	return 0
}

func (user User) GetAmount(operationType LimitedOperationType) int {
	return 0
}

func (user User) HasOperationAvailable(operationType LimitedOperationType) bool {
	return false
}

func (user *User) ConsumeOperation(operationType LimitedOperationType) {

}

func (user *User) HandleInvite(objectId string, inviteType InviteType, answer bool) {

}

func (user *User) Friendinvite(userId string, answer bool) {

}

func (user *User) GroupInvite(groupId string, answer bool) {

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
	Friends UserRelation = iota
	Blocked
	NotFriends
)
