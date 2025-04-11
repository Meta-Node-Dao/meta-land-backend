package startup_team

import (
	"ceres/pkg/model"
	"ceres/pkg/model/account"
)

type StartupTeamMember struct {
	model.RelationBase
	ComerID      uint64               `gorm:"comer_id" json:"comerID"`
	StartupID    uint64               `gorm:"startup_id" json:"startupID"`
	Position     string               `gorm:"position" json:"position"`
	Comer        account.Comer        `gorm:"foreignkey:ID;references:ComerID" json:"comer"`
	ComerProfile account.ComerProfile `gorm:"foreignkey:ComerID;references:ComerID" json:"comerProfile"`
	FollowedByMe *bool                `gorm:"-" json:"followedByMe"`
}

// TableName Startup table name for gorm
func (StartupTeamMember) TableName() string {
	return "startup_team_member_rel"
}
