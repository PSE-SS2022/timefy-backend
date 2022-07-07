package models

type Group struct {
	ID          string `bson:"_id,omitempty"`
	GID         string `bson:"GID" json:"GID"`
	Title       string `bson:"Title" json:Title"`
	Description string `bson:"Description" json:"Description"`
	Events      string `bson:"Events" json:"Events"`
}

func (group Group) CalculateGID() {

}

type GroupRelation struct {
	Id      string    `bson:"_id,omitempty"`
	GroupId string    `bson:"GroupId" json:"GroupId"`
	UserId  string    `bson:"UserId" json:"UserId"`
	Role    GroupRole `bson:"Role" json:"Role"`
}

type GroupRole int

const (
	GroupRoleOwner GroupRole = iota
	GroupRoleModerator
	GroupRoleMember
)
