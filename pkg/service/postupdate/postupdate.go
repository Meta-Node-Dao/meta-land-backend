/**
 * @Author: Sun
 * @Description:
 * @File:  postupdate
 * @Version: 1.0.0
 * @Date: 2022/7/3 10:37
 */

package postupdate

import (
	"ceres/pkg/model/bounty"
	model "ceres/pkg/model/postupdate"
	"gorm.io/gorm"
	"time"
)

const Bounty = 1

func CreatePostUpdate(tx *gorm.DB, bountyID uint64, request *bounty.BountyRequest) error {
	postUpdate := &model.PostUpdate{
		SourceType: Bounty, //1 bounty
		SourceID:   bountyID,
		ComerID:    request.ComerID,
		Content:    request.Description,
		TimeStamp:  time.Now(),
	}
	err := model.CreatePostUpdate(tx, postUpdate)
	if err != nil {
		return err
	}
	return nil
}
