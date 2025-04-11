package startup

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/startup"
	team "ceres/pkg/model/startup_team"
	"ceres/pkg/model/tag"
	"errors"
	"fmt"
	"strings"

	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

func CreateStartup(comerID uint64, request *model.CreateStartupRequest) (err error) {
	if strings.TrimSpace(request.Name) == "" {
		err = errors.New("startup name can not be empty")
		return
	}
	if len(request.HashTags) <= 0 {
		err = errors.New("must be select tags")
		return
	}
	if request.ChainID <= 0 {
		err = errors.New("please select the correct chain")
		return
	}

	isExist, err := StartupNameIsExist(request.Name)
	fmt.Println(request.Name, isExist)
	if err != nil {
		return
	}
	if isExist {
		err = errors.New("startup name is exist")
		return
	}

	if err := mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		startup := &model.Startup{
			Logo:     request.Logo,
			Mode:     request.Mode,
			Name:     request.Name,
			Mission:  request.Mission,
			TxHash:   request.TxHash,
			ChainID:  request.ChainID,
			Overview: request.Overview,
			ComerID:  comerID,
		}
		er = model.CreateStartup(tx, startup)
		if er != nil {
			return
		}
		// create default team member
		teamMember := team.StartupTeamMember{
			StartupID: startup.ID,
			ComerID:   comerID,
			Position:  "founder",
		}
		if er = team.CreateStartupTeamMembers(tx, &teamMember); er != nil {
			return
		}

		var tagRelList []tag.TagTargetRel
		for _, tagName := range request.HashTags {
			var isIndex bool
			if len(tagName) > 2 && tagName[0:1] == "#" {
				isIndex = true
			}
			hashTag := tag.Tag{
				Name:     tagName,
				Category: tag.Startup,
				IsIndex:  isIndex,
			}
			if er = tag.FirstOrCreateTag(tx, &hashTag); er != nil {
				return er
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    hashTag.ID,
				Target:   tag.Startup,
				TargetID: startup.ID,
			})
		}
		if len(tagRelList) > 0 {
			//batch create startup hashtag rel
			if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
				return er
			}
		}

		return er
	}); err != nil {
		log.Warn(err)
		return err
	}
	return
}

// ListStartups get current comer accounts
func ListStartups(comerID uint64, request *model.ListStartupRequest, response *model.ListStartupsResponse) (err error) {
	total, err := model.ListStartups(mysql.DB, comerID, request, &response.List)
	if err != nil {
		log.Warn(err)
		return
	}
	if total == 0 {
		response.List = make([]model.Startup, 0)
		return
	}
	response.Total = total
	for i, startup := range response.List {
		response.List[i].MemberCount = len(startup.Members)
		response.List[i].FollowCount = len(startup.Follows)
	}
	return
}

func GetStartup(startupID uint64, response *model.GetStartupResponse) (err error) {
	if err = model.GetStartup(mysql.DB, startupID, &response.Startup); err != nil {
		log.Warn(err)
		return
	}
	response.MemberCount = len(response.Members)
	response.FollowCount = len(response.Follows)
	return
}

func StartupNameIsExist(name string) (isExist bool, err error) {
	isExist, err = model.StartupNameIsExist(mysql.DB, name)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}

func StartupTokenContractIsExist(tokenContract string) (isExist bool, err error) {
	isExist, err = model.StartupTokenContractIsExist(mysql.DB, tokenContract)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}

func StartupFollowedByComer(startupID, comerID uint64) (isFollowed bool, err error) {
	isFollowed, err = model.StartupFollowIsExist(mysql.DB, startupID, comerID)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}

func GetStartupsByComerID(targetComerID, currentComerId uint64) ([]*model.ListComerStartup, error) {
	var startups []*model.ListComerStartup
	startupsResponse, err := model.ListComerStartups(mysql.DB, targetComerID, startups)
	if err != nil {
		return nil, err
	}
	if len(startupsResponse) > 0 {
		for _, st := range startupsResponse {
			if st.ComerId == currentComerId {
				st.IsFollowed = true
			} else {
				isFollowed, err := StartupFollowedByComer(st.StartupID, currentComerId)
				if err != nil {
					return nil, err
				}
				st.IsFollowed = isFollowed
			}
		}
	}
	return startupsResponse, nil
}
