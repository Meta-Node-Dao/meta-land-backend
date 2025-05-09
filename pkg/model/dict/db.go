package dict

import "gorm.io/gorm"

func SelectByDictType(db *gorm.DB, tp string) (list []DictData, err error) {
	err = db.Model(&DictData{}).Where("status = ? and dict_type = ?", Enabled, tp).Find(&list).Error
	return
}

func SelectByDictTypeAndLabel(db *gorm.DB, tp, value string) (dict DictData, err error) {
	err = db.Model(&DictData{}).
		Where("status = ? and dict_type = ? and dict_label = ?", Enabled, tp, value).
		Find(&dict).Error
	return
}
