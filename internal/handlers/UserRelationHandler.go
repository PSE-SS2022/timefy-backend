package handlers

import . "github.com/PSE-SS2022/timefy-backend/internal/models"

func GetFriends(userId string) []User {
	var result []User
	return result
}

func SendFriendRequest(userId string, userToAddUids string) {

}

func RemoveFriend(userId string, removeUserUID string) {

}

func SearchUser(searchString string) []User {
	var result []User
	return result
}

func FriendRequestResponse(userId, userToRespondId string, respons bool) {

}

func BlockUser(userId, userToBlockUID string) {

}

func GetUserRelation(userId, userToCheckUID string) UserRelation {
	var result UserRelation
	return result
}

func HasSpaceForNewEvent(user User) bool {
	return false
}

func TryDeletingOldestOccuredEvent(user User) bool {
	return false
}

func IsBlockedByUser(blockedUser, blockingUser User) bool {
	return false
}
