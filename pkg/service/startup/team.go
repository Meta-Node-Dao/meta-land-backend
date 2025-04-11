package startup

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	"ceres/pkg/model/startup"
	model "ceres/pkg/model/startup_team"
	"gorm.io/gorm"

	"github.com/qiniu/x/log"
)

// ListStartupTeamMembers get startup team members comer
func ListStartupTeamMembers(startupID, crtComerId uint64, request *model.ListStartupTeamMemberRequest, response *model.ListStartupTeamMemberResponse) (err error) {
	total, err := model.ListStartupTeamMembers(mysql.DB, startupID, request, &response.List)
	if err != nil {
		log.Warn(err)
		return
	}
	if total == 0 {
		response.List = make([]*model.StartupTeamMember, 0)
		return
	}
	for _, member := range response.List {
		var f *bool
		if crtComerId != member.ComerID {
			if followed, err := account.ComerFollowIsExist(mysql.DB, crtComerId, member.ComerID); err != nil {
				return err
			} else {
				f = &followed
			}
		} else {
			f = nil
		}
		member.FollowedByMe = f
	}
	response.Total = total
	return
}

// CreateStartupTeamMember create startup team member
func CreateStartupTeamMember(startupID, comerID uint64, request *model.CreateStartupTeamMemberRequest) error {
	startupTeam := model.StartupTeamMember{
		StartupID: startupID,
		ComerID:   comerID,
		Position:  request.Position,
	}

	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.CreateStartupTeamMembers(mysql.DB, &startupTeam); err != nil {
			log.Warn(err)
		}
		return AddComer2Group(startupID, request.GroupId, comerID)
	})

}

// UpdateStartupTeamMember update startup team member
func UpdateStartupTeamMember(startupID, comerID uint64, request *model.UpdateStartupTeamMemberRequest) error {
	startupTeam := model.StartupTeamMember{
		StartupID: startupID,
		ComerID:   comerID,
		Position:  request.Position,
	}
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.UpdateStartupTeamMember(mysql.DB, &startupTeam); err != nil {
			log.Warn(err)
			return err
		}
		return AddComer2Group(startupID, request.GroupId, comerID)
	})
}

// DeleteStartupTeamMember delete startup team member
func DeleteStartupTeamMember(startupID, comerID uint64) (err error) {
	startupTeam := model.StartupTeamMember{
		StartupID: startupID,
		ComerID:   comerID,
	}
	if err = model.DeleteStartupTeamMember(mysql.DB, &startupTeam); err != nil {
		log.Warn(err)
	}
	if err = startup.DeleteStartupGroupMemberRelsByComerIdAndStartupId(mysql.DB, comerID, startupID); err != nil {
		log.Warn(err)
	}
	return
}
