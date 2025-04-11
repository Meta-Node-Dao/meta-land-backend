package startup_group

import (
	"ceres/pkg/model"
	"gorm.io/gorm"
	"math"
)

func SelectStartupGroupsMembers(db *gorm.DB, groupId uint64, pagination *model.Pagination) (err error) {
	var groups []GroupMember
	db = db.Where("startup_group_member_rel.group_id = ?", groupId).
		Select("startup_group.id as group_id, startup_group.name as group_name,startup_team_member_rel.position, comer_profile.comer_id, comer_profile.name as comer_name, comer_profile.avatar as comer_avatar, startup.id as startup_id, startup_group_member_rel.created_at").
		Joins("LEFT JOIN comer_profile on comer_profile.comer_id = startup_group_member_rel.comer_id").
		Joins("LEFT JOIN startup_group on startup_group.id = startup_group_member_rel.group_id").
		Joins("LEFT JOIN startup on startup.id = startup_group_member_rel.startup_id").
		Joins("LEFT JOIN startup_team_member_rel on startup_team_member_rel.startup_id = startup_group_member_rel.startup_id and startup_team_member_rel.comer_id=startup_group_member_rel.comer_id").
		Where("startup.is_deleted = false")
	var total int64
	if err := db.Table("startup_group_member_rel").Count(&total).Error; err != nil {
		return err
	}
	err = db.
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Table("startup_group_member_rel").
		Scan(&groups).
		Error
	pagination.TotalRows = total
	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	pagination.Rows = groups
	return
}
