package handlers

import . "github.com/PSE-SS2022/timefy-backend/internal/models"

type GroupHandler struct {
}

func (groupHandler GroupHandler) CreateGroup(userId, groupName, description string) {

}

func (groupHandler GroupHandler) GetGroupsOfUser(userId string) []Group {
	var results []Group
	return results
}

func (groupHandler GroupHandler) GetGroupById(userId string) Group {
	var result Group
	return result
}

func (groupHandler GroupHandler) GetMembersOfGroup(userId, groupId string) []User {
	var result []User
	return result
}

func (groupHandler GroupHandler) GroupInviteResponse(userId, groupId string, response bool) {

}

func (groupHandler GroupHandler) LeaveGroup(userId, groupId string) {

}

func (groupHandler GroupHandler) Invite(userId, userToInviteUID string) {

}

func (groupHandler GroupHandler) RemoveMember(userId, userToRemoveUID string) {

}

func (groupHandler GroupHandler) SetGroupMemberRole(userId, userToAssignUID string, role GroupRole) {

}

func (groupHandler GroupHandler) DeleteGroup(userId, groupId string) {

}
