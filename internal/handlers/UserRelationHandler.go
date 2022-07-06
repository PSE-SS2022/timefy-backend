package handlers

import . "github.com/PSE-SS2022/timefy-backend/internal/models"

type UserRelationHandler struct {
}

func (userRelationHandler UserRelationHandler) GetFriends(userId string) []User {
	var result []User
	return result
}

func (userRelationHandler UserRelationHandler) SendFriendRequest(userId string, userToAddUids string) {

}

func (userRelationHandler UserRelationHandler) RemoveFriend(userId string, removeUserUID string) {

}

func (userRelationHandler UserRelationHandler) SearchUser(searchString string) []User {
	var result []User
	return result
}

func (userRelationHandler UserRelationHandler) FriendRequestResponse(userId, userToRespondId string, respons bool) {

}

func (userRelationHandler UserRelationHandler) BlockUser(userId, userToBlockUID string) {

}

func (userRelationHandler UserRelationHandler) GetUserRelation(userId, userToCheckUID string) UserRelation {
	var result UserRelation
	return result
}

func (userRelationHandler UserRelationHandler) HasSpaceForNewEvent(user User) bool {
	return false
}

func (userRelationHandler UserRelationHandler) TryDeletingOldestOccuredEvent(user User) bool {
	return false
}

func (userRelationHandler UserRelationHandler) IsBlockedByUser(blockedUser, blockingUser User) bool {
	return false
}
