/**
 * @Author: Sun
 * @Description:
 * @File:  proto
 * @Version: 1.0.0
 * @Date: 2022/7/3 10:47
 */

package postupdate

import (
	"ceres/pkg/model"
	"time"
)

type PostUpdate struct {
	model.RelationBase
	SourceType int       `gorm:"sourceType"`
	SourceID   uint64    `gorm:"sourceID"`
	ComerID    uint64    `gorm:"comerID"`
	Content    string    `gorm:"column:content"`
	TimeStamp  time.Time `gorm:"column:timestamp"` // post time
}

// TableName the PostUpdate table name for gorm
func (PostUpdate) TableName() string {
	return "post_update"
}
