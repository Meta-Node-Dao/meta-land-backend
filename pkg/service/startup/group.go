package startup

import (
	"ceres/pkg/initialization/mysql"
	model2 "ceres/pkg/model"
	model "ceres/pkg/model/startup"
	"ceres/pkg/model/startup_group"
	"ceres/pkg/model/startup_team"
	"errors"
	"gorm.io/gorm"
	"strings"
)

func CreateStartupGroup(startupId, comerId uint64, request model.CreateOrUpdateStartupGroupRequest) (group model.StartupGroup, err error) {
	st, err := model.GetStartupById(mysql.DB, startupId)
	if err != nil {
		return
	}
	if st.ComerID != comerId {
		err = errors.New("can not create startup group, current comer is not founder of startup")
		return
	}
	if strings.TrimSpace(request.Name) == "" {
		err = errors.New("startup group name can not be empty")
		return
	}
	exist, err := model.ExistStartupGroupByName(mysql.DB, startupId, strings.TrimSpace(request.Name))
	if err != nil {
		return
	}
	if exist {
		err = errors.New("can not create startup group with name " + request.Name)
		return
	}
	group = model.StartupGroup{
		StartupId: startupId,
		ComerId:   comerId,
		Name:      request.Name,
	}
	if err := model.CreateStartupGroup(mysql.DB, &group); err != nil {
		return model.StartupGroup{}, err
	}
	return group, nil
}

func DeleteStartupGroup(groupID, comerID uint64) error {
	group, err := model.GetStartupGroupById(mysql.DB, groupID)
	if err != nil {
		return err
	}
	st, err := model.GetStartupById(mysql.DB, group.StartupId)
	if err != nil {
		return err
	}

	if st.ComerID != comerID {
		return errors.New("can not delete startup group")
	}
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.DeleteStartupGroup(tx, groupID); err != nil {
			return err
		}
		if err := model.DeleteStartupGroupMemberRelsByGroupId(tx, groupID); err != nil {
			return err
		}
		return nil
	})

}

func UpdateStartupGroup(groupId, comerId uint64, request model.CreateOrUpdateStartupGroupRequest) (err error) {
	st, err := model.GetStartupGroupById(mysql.DB, groupId)
	if err != nil {
		return
	}
	if st.ComerId != comerId {
		err = errors.New("can not create startup group, current comer is not founder of startup")
		return
	}
	if strings.TrimSpace(request.Name) == "" {
		err = errors.New("startup group name can not be empty")
		return
	}
	exist, err := model.ExistStartupGroupByName(mysql.DB, st.StartupId, strings.TrimSpace(request.Name))
	if err != nil {
		return
	}
	if exist {
		err = errors.New("can not update startup group with name " + request.Name)
		return
	}
	if err := model.UpdateStartupGroup(mysql.DB, groupId, request.Name); err != nil {
		return err
	}
	return nil
}

func GetStartupGroups(startupId uint64) (groups []*model.StartupGroup, er error) {
	return model.SelectStartupGroupsByStartupId(mysql.DB, startupId)
}

func GetStartupGroupMembers(startupId, groupId uint64, pagination *model2.Pagination) (err error) {
	if startupId == 0 {
		return
	}

	if groupId != 0 {
		group, err := model.GetStartupGroupById(mysql.DB, groupId)
		if err != nil || group.StartupId != startupId {
			err = errors.New("invalid startupId and groupId")
			return err
		}
		return startup_group.SelectStartupGroupsMembers(mysql.DB, groupId, pagination)
	}
	return startup_team.SelectStartupMembers(mysql.DB, startupId, pagination)
}

func AddComer2Group(startupId, groupID, comerID uint64) error {
	stop, err := beforeAddingComer2Group(startupId, groupID, comerID)
	if stop {
		return err
	}
	return createOrUpdateGroupMember(comerID, startupId, groupID)
}

func beforeAddingComer2Group(startupId, groupID, comerID uint64) (stop bool, er error) {
	if groupID == 0 {
		return true, model.DeleteStartupGroupMemberRelsByComerIdAndStartupId(mysql.DB, comerID, startupId)
	}
	group, err := model.GetStartupGroupById(mysql.DB, groupID)
	if err != nil {
		return true, err
	}
	if group.StartupId != startupId {
		return true, errors.New("invalid group id")
	}
	// check comer is member of startup!
	isTeamMemberOrFunder, err := startup_team.ComerIsTeamMemberOfStartup(mysql.DB, startupId, comerID)
	if err != nil {
		return true, err
	}
	if !isTeamMemberOrFunder {
		return true, errors.New("comer is not member of startup")
	}
	return false, nil
}

func ChangeComerGroupAndPosition(startupId, groupID, comerID uint64, request model.ModifyLocationRequest) error {
	stop, er := beforeAddingComer2Group(startupId, groupID, comerID)
	if stop {
		return er
	}

	return mysql.DB.Transaction(func(tx *gorm.DB) (e error) {
		if e = startup_team.UpdateStartupTeamMember(mysql.DB, &startup_team.StartupTeamMember{
			ComerID:   comerID,
			StartupID: startupId,
			Position:  request.Position,
		}); e != nil {
			return e
		}
		return createOrUpdateGroupMember(comerID, startupId, groupID)
	})

}

func createOrUpdateGroupMember(comerID uint64, startupId uint64, groupID uint64) (e error) {
	rel, e := model.GetGroupMemberRelByComerIdAndStartupId(mysql.DB, comerID, startupId)
	if (e != nil && errors.Is(e, gorm.ErrRecordNotFound)) || rel.ID == 0 {
		rel := model.StartupGroupMemberRel{
			StartupId: startupId,
			ComerId:   comerID,
			GroupId:   groupID,
		}
		return model.CreateGroupMemberRel(mysql.DB, &rel)
	}

	if rel.ID != 0 {
		rel.GroupId = groupID
		return model.UpdateGroupMemberRel(mysql.DB, rel)
	}
	if e != nil {
		return e
	}
	return nil
}

func GetComerJoinedOrFollowedStartups(comerId uint64) (list []model.SimpleStartupInfo, err error) {
	var joined []model.SimpleStartupInfo
	err = mysql.DB.Where("startup.is_deleted = false").
		Select("startup.id as id, startup.name as name, startup.logo as logo").
		Joins("left join startup_team_member_rel on  startup.id = startup_team_member_rel.startup_id").
		Where("startup_team_member_rel.comer_id = ?", comerId).
		Table("startup").
		Scan(&joined).Error
	if err != nil {
		return
	}
	var joinedIds []uint64
	if len(joined) > 0 {
		for _, info := range joined {
			joinedIds = append(joinedIds, info.StartupId)
		}
	}
	var followed []model.SimpleStartupInfo
	db := mysql.DB.Where("startup.is_deleted = false").
		Select("startup.id as id", "startup.name as name", "startup.logo as logo").
		Joins("left join startup_follow_rel on startup_follow_rel.startup_id = startup.id").
		Where("startup_follow_rel.comer_id = ?", comerId)
	if len(joinedIds) > 0 {
		db = db.Where("startup_follow_rel.startup_id not in ?", joinedIds)
	}
	err = db.Table("startup").Scan(&followed).Error
	if err != nil {
		return
	}
	list = append(list, joined...)
	list = append(list, followed...)
	return
}
