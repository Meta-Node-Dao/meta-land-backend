package startup

import (
	"ceres/pkg/initialization/mysql"
	model2 "ceres/pkg/model"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/startup"
	"github.com/qiniu/x/log"
)

func FollowStartup(ComerID, startupID uint64) (err error) {
	return model.CreateStartupFollowRel(mysql.DB, ComerID, startupID)
}

func UnfollowStartup(ComerID, startupID uint64) (err error) {
	followRel := model.FollowRelation{
		StartupID: startupID,
		ComerID:   ComerID,
	}
	if err = model.DeleteStartupFollowRel(mysql.DB, &followRel); err != nil {
		log.Warn(err)
	}
	return
}

func ListFollowStartups(ComerID uint64, request *model.ListStartupRequest, response *model.ListStartupsResponse) (err error) {
	startups := make([]model.Startup, 0)
	total, err := model.ListFollowedStartups(mysql.DB, ComerID, request, &startups)
	if err != nil {
		log.Warn(err)
		return
	}
	response.Total = total
	response.List = startups
	return
}

func GetComersFollowedThisStartup(crtComerId, startupID uint64, pagination *model2.Pagination) error {
	fans, err := model.SelectStartupFans(mysql.DB, startupID, pagination)
	if err != nil {
		return err
	}
	if len(fans) > 0 {
		for _, fan := range fans {
			var f *bool
			if crtComerId != fan.ComerId {
				if followed, err := account.ComerFollowIsExist(mysql.DB, crtComerId, fan.ComerId); err != nil {
					return err
				} else {
					f = &followed
				}
			} else {
				f = nil
			}
			fan.FollowedByMe = f
		}
	}
	pagination.Rows = fans
	return nil
}
