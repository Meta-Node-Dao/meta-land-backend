package image

import (
	"gorm.io/gorm"
)

func GetImageList(db *gorm.DB, input ListRequest, tags *[]Image) (count int64, err error) {
	if input.Category != "" {
		db = db.Where("category = ?", input.Category)
	}
	err = db.Find(tags).Count(&count).Limit(input.Limit).Offset(input.Offset).Error
	return
}
