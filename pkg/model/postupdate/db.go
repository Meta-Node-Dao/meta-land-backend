/**
 * @Author: Sun
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2022/7/3 10:46
 */

package postupdate

import "gorm.io/gorm"

func CreatePostUpdate(db *gorm.DB, postUpdate *PostUpdate) error {
	return db.Create(&postUpdate).Error
}
