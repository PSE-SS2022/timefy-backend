package models

import (
	"context"

	"github.com/PSE-SS2022/timefy-backend/internal/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	USER_REPO = "users"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FirstName      string             `bson:"FirstName" json:"FirstName"`
	LastName       string             `bson:"LastName" json:"LastName"`
	Email          string             `bson:"Email" json:"Email"`
	Tags           []Tag              `bson:"Tags" json:"Tags"`
	Roles          map[string]string  `bson:"Roles" json:"Roles"`
	AmountWarnings int                `bson:"AmountWarnings" json:"AmountWarnings"`
}

func GetUserByID(id primitive.ObjectID) (User, bool) {
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

func GetUserByMail(mail string) (User, bool) {
	var user User
	usersCollection := repos.GetCollection((USER_REPO))
	if usersCollection != nil {
		return user, false
	}
	usersCollection.FindOne(context.TODO(), bson.M{"Email": mail}).Decode(&user)
	if user.ID.IsZero() {
		return user, false
	}
	return user, true
}
