package database

import (
	"context"

	"github.com/PSE-SS2022/timefy-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

var UserRepositoryInstance userRepository

const USER_REPO = "users"

type userRepository struct {
}

func (userRepository userRepository) GetUserById(id string) (models.User, bool) {
	var user models.User
	usersCollection := databaseMgrInstance.getCollection((USER_REPO))

	if usersCollection != nil {
		return user, false
	}

	userResult := usersCollection.FindOne(context.TODO(), bson.M{"id": id})
	userResult.Decode(&user)

	if user.ID.IsZero() {
		return user, false
	}

	return user, true
}
