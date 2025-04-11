package account

import (
	"ceres/pkg/initialization/mysql"
	model2 "ceres/pkg/model"
	model "ceres/pkg/model/account"
	"ceres/pkg/model/startup"
	startup2 "ceres/pkg/service/startup"
	"errors"
	"github.com/qiniu/x/log"
)

func FollowComer(comerID, targetComerID uint64) (err error) {
	if comerID == targetComerID {
		return errors.New("can not follow yourself")
	}
	return model.CreateComerFollowRel(mysql.DB, comerID, targetComerID)
}

func UnfollowComer(comerID, targetComerID uint64) (err error) {
	followRel := model.FollowRelation{
		ComerID:       comerID,
		TargetComerID: targetComerID,
	}
	if err = model.DeleteComerFollowRel(mysql.DB, &followRel); err != nil {
		log.Warn(err)
	}
	return
}

func FollowedByComer(comerID, targetComerID uint64) (isFollowed bool, err error) {
	isFollowed, err = model.ComerFollowIsExist(mysql.DB, comerID, targetComerID)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}

// GetConnectorsOfComer fans
func GetConnectorsOfComer(currentComerId, targetComerId uint64, pagination *model2.Pagination) error {
	profiles, err := model.GetFollowersOfComer(mysql.DB, targetComerId, pagination)
	if err != nil {
		return err
	}
	var followerInfos []model.ComerFollowerInfo
	if len(profiles) > 0 {
		followerInfos, err = packComerFollowerInfo(currentComerId, profiles)
		if err != nil {
			return err
		}
	}
	pagination.Rows = followerInfos
	return nil
}

func packComerFollowerInfo(currentComerId uint64, profiles []model.ComerProfile) (followerInfos []model.ComerFollowerInfo, err error) {
	for _, profile := range profiles {
		var f *bool
		if currentComerId != profile.ComerID {
			if followed, err := FollowedByComer(currentComerId, profile.ComerID); err != nil {
				return followerInfos, err
			} else {
				f = &followed
			}
		}
		followerInfos = append(followerInfos, model.ComerFollowerInfo{
			ComerId:      profile.ComerID,
			ComerAvatar:  profile.Avatar,
			ComerName:    profile.Name,
			FollowedByMe: f,
		})
	}
	return
}

func GetComersFollowedByComer(currentComerId, targetComerId uint64, pagination *model2.Pagination) error {
	profiles, err := model.GetFollowedByComer(mysql.DB, targetComerId, pagination)
	if err != nil {
		return err
	}
	var followerInfos []model.ComerFollowerInfo
	if len(profiles) > 0 {
		followerInfos, err = packComerFollowerInfo(currentComerId, profiles)
		if err != nil {
			return err
		}
	}
	pagination.Rows = followerInfos
	return nil
}

func GetComerFollowedStartups(currentComerId, targetComerId uint64, pagination *model2.Pagination) error {
	startups, err := startup.GetFollowedStartupsOfComer(mysql.DB, targetComerId, pagination)
	if err != nil {
		return err
	}
	var followerInfos []model.StartupFollowerInfo
	if len(startups) > 0 {
		for _, st := range startups {
			followed, err := startup2.StartupFollowedByComer(st.ID, currentComerId)
			// swallow
			if err != nil {
				return err
			}
			followerInfos = append(followerInfos, model.StartupFollowerInfo{
				StartupId:    st.ID,
				StartupLogo:  st.Logo,
				StartupName:  st.Name,
				FollowedByMe: followed,
			})
		}
	}
	pagination.Rows = followerInfos
	return nil
}
