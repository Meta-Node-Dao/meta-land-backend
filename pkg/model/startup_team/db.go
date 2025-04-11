package startup_team

import (
	"ceres/pkg/model"
	"ceres/pkg/model/startup_group"
	"ceres/pkg/model/tag"
	"gorm.io/gorm"
	"math"
)

// ListStartupTeamMembers  get startup team members
func ListStartupTeamMembers(db *gorm.DB, startupID uint64, input *ListStartupTeamMemberRequest, output *[]*StartupTeamMember) (total int64, err error) {
	if startupID != 0 {
		db = db.Where("startup_id = ?", startupID)
	}
	if err = db.Table("startup_team_member_rel").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at ASC").Limit(input.Limit).Offset(input.Offset).Preload("Comer").Preload("ComerProfile").Preload("ComerProfile.Skills", "category = ?", tag.ComerSkill).Find(output).Error
	return
}
func SelectStartupMembers(db *gorm.DB, startupID uint64, pagination *model.Pagination) (err error) {
	var total int64
	db = db.
		Where("startup_team_member_rel.startup_id = ?", startupID).
		Select("startup_group.id as group_id, startup_group.name as group_name, startup_team_member_rel.position, comer_profile.comer_id, comer_profile.name as comer_name, comer_profile.avatar as comer_avatar, startup.id as startup_id, startup_group_member_rel.created_at").
		Joins("LEFT JOIN startup on startup.id = startup_team_member_rel.startup_id").
		Joins("LEFT JOIN startup_group_member_rel on startup_group_member_rel.comer_id = startup_team_member_rel.comer_id and startup_group_member_rel.startup_id = startup_team_member_rel.startup_id").
		Joins("LEFT JOIN startup_group on startup_group.id = startup_group_member_rel.group_id").
		Joins("LEFT JOIN comer_profile on startup_team_member_rel.comer_id =comer_profile.comer_id").
		Where("startup.is_deleted = false and comer_profile.is_deleted = false")
	err = db.Table("startup_team_member_rel").Count(&total).Error
	if err != nil {
		return
	}
	var groups []startup_group.GroupMember

	err = db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Table("startup_team_member_rel").Scan(&groups).Error
	if err != nil {
		return
	}
	pagination.TotalRows = total
	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	pagination.Rows = groups
	return
}
func ComerIsTeamMemberOfStartup(db *gorm.DB, startupID, comerID uint64) (is bool, err error) {
	var cnt int64
	err = db.Model(&StartupTeamMember{}).Where("startup_id = ? and comer_id = ?", startupID, comerID).Count(&cnt).Error
	if err != nil {
		return
	}
	if cnt == 0 {
		var ci uint64
		err = db.Select("comer_id").Where("id = ? and is_deleted = false", startupID).Table("startup").Scan(&ci).Error
		if err != nil {
			return
		}
		if ci == comerID {
			return true, nil
		}
		return
	}
	return true, nil
}

// CreateStartupTeamMembers  add startup team members
func CreateStartupTeamMembers(db *gorm.DB, input *StartupTeamMember) (err error) {
	return db.Create(input).Error
}

// UpdateStartupTeamMember  update startup team member title
func UpdateStartupTeamMember(db *gorm.DB, input *StartupTeamMember) (err error) {
	return db.Table("startup_team_member_rel").Where("comer_id = ? AND startup_id = ?", input.ComerID, input.StartupID).Update("position", input.Position).Error
}

// DeleteStartupTeamMember delete startup team member
func DeleteStartupTeamMember(db *gorm.DB, input *StartupTeamMember) (err error) {
	return db.Where("comer_id = ? AND startup_id = ?", input.ComerID, input.StartupID).Delete(input).Error
}
