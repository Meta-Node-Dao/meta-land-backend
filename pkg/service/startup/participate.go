package startup

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/startup"
	"github.com/qiniu/x/log"
)

func ListParticipateStartups(ComerID uint64, request *model.ListStartupRequest, response *model.ListStartupsResponse) (err error) {
	startups := make([]model.Startup, 0)
	total, err := model.ListParticipatedStartups(mysql.DB, ComerID, request, &startups)
	if err != nil {
		log.Warn(err)
		return
	}
	response.Total = total
	response.List = startups
	for i, startup := range response.List {
		response.List[i].MemberCount = len(startup.Members)
		response.List[i].FollowCount = len(startup.Follows)
	}
	return
}

func ListBeMemberStartups(ComerID uint64, request *model.ListStartupRequest, response *model.ListStartupsResponse) (err error) {
	startups := make([]model.Startup, 0)
	total, err := model.ListBeMemberStartups(mysql.DB, ComerID, request, &startups)
	if err != nil {
		log.Warn(err)
		return
	}
	response.Total = total
	response.List = startups
	for i, startup := range response.List {
		response.List[i].MemberCount = len(startup.Members)
		response.List[i].FollowCount = len(startup.Follows)
	}
	return
}
