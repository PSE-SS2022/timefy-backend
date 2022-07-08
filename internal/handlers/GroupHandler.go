package handlers

import . "github.com/PSE-SS2022/timefy-backend/internal/models"

func CreateGroup(userId, groupName, description string) {

}

func GetGroupsOfUser(userId string) []Group {
	var results []Group
	return results
}

func GetGroupById(userId string) Group {
	var result Group
	return result
}

func GetMembersOfGroup(userId, groupId string) []User {
	var result []User
	return result
}

func GroupInviteResponse(userId, groupId string, response bool) {

}

func LeaveGroup(userId, groupId string) {

}

func Invite(userId, userToInviteUID string) {

}

func RemoveMember(userId, userToRemoveUID string) {

}

func SetGroupMemberRole(userId, userToAssignUID string, role GroupRole) {

}

func DeleteGroup(userId, groupId string) {

}
