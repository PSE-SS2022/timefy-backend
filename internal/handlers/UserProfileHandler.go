package handlers

import (
	"github.com/PSE-SS2022/timefy-backend/internal/database"
	"github.com/PSE-SS2022/timefy-backend/internal/models"
)

func RegisterUser(firstName, lastName, email string) {

}

func DeleteUser(userId string) {

}

func SetProfilePicture(userId string /*image */) {

}

func DeleteProfilePicture(userId string) {

}

func GetUserById(userId string) (models.User, bool) {
	return database.UserRepositoryInstance.GetUserById(userId)
}
