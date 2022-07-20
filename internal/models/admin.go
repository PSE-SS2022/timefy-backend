package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ADMIN_ROLE     = "admin"
	MODERATOR_ROLE = "moderator"
)

type Admin struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"FirstName" json:"FirstName"`
	LastName  string             `bson:"LastName" json:"LastName"`
	Email     string             `bson:"Email" json:"Email"`
	Password  string             `bson:"Password" json:"Password"`
	Role      string             `bson:"Role" json:"Role"`
}
