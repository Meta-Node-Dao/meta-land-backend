package tag

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//GetTagList get tag list tag ids
func GetTagList(db *gorm.DB, input ListRequest, tags *[]Tag) (total int64, err error) {
	db = db.Where("is_index = ? AND is_deleted = false", input.IsIndex)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Category != "" {
		db = db.Where("category = ?", input.Category)
	}
	if err = db.Table("tag").Order("created_at DESC").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Find(tags).Error
	return
}

//FirstOrCreateTag first or create tags
func FirstOrCreateTag(db *gorm.DB, tag *Tag) error {
	return db.Where("name = ?", tag.Name).FirstOrCreate(&tag).Error
}

//DeleteTagRel delete comer skill relation where not used
func DeleteTagRel(db *gorm.DB, comerID uint64, target Category, skillIds []uint64) error {
	return db.Delete(&TagTargetRel{}, "target_id = ? AND target = ? AND tag_id NOT IN ?", comerID, target, skillIds).Error
}

//BatchCreateTagRel delete comer skill relation where not used
func BatchCreateTagRel(db *gorm.DB, tagTargetRelList []TagTargetRel) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&tagTargetRelList).Error
}
