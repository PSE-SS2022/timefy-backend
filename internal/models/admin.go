package models

import (
	"context"

	"github.com/PSE-SS2022/timefy-backend/internal/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ADMIN_ROLE     = "admin"
	MODERATOR_ROLE = "moderator"
	ADMIN_REPO     = "admins"
)

type Admin struct {
	// TODO: change the id to proper type
	ID        bsontype.Type `bson:"uuid,omitempty"`
	FirstName string        `bson:"FirstName" json:"FirstName"`
	LastName  string        `bson:"LastName" json:"LastName"`
	Email     string        `bson:"Email" json:"Email"`
	Password  string        `bson:"Password" json:"Password"`
	Role      string        `bson:"Role" json:"Role"`
}

func GetAdminById(id primitive.ObjectID) (Admin, bool) {
	var admin Admin
	usersCollection := repos.GetCollection(ADMIN_REPO)
	usersCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&admin)
	if admin.ID {
		return admin, false
	}
	return admin, true
}

func GetAdminByMail(email string) (Admin, bool) {
	var admin Admin
	usersCollection := repos.GetCollection(ADMIN_REPO)
	err := usersCollection.FindOne(context.TODO(), bson.M{"Email": email}).Decode(&admin)
	if admin.ID.IsZero() {
		println("could not decode: " + err.Error())
		println(email)
		return admin, false
	}
	return admin, true
}
