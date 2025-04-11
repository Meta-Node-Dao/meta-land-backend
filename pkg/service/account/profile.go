package account

import (
	"ceres/pkg/initialization/mysql"
	model2 "ceres/pkg/model"
	model "ceres/pkg/model/account"
	bountyModel "ceres/pkg/model/bounty"
	crowdfundingModel "ceres/pkg/model/crowdfunding"
	startupModel "ceres/pkg/model/startup"
	"ceres/pkg/model/tag"
	"ceres/pkg/router"
	"errors"
	"strings"

	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

// GetComerProfile get current comer profile
func GetComerProfile(comerID uint64, response *model.ComerProfileResponse) (err error) {
	var profile model.ComerProfile
	if err = model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		log.Warn(err)
		return err
	}
	response.ComerProfile = &profile
	var accounts []model.ComerAccount
	if err = model.GetComerAccountsByComerId(mysql.DB, comerID, &accounts); err != nil {
		log.Warn(err)
		return err
	}
	log.Infof("comer accounts for %d : %v \n", comerID, accounts)
	var accountBindingInfos = []*model.OauthAccountBindingInfo{
		{Linked: false, AccountType: 1},
		{Linked: false, AccountType: 2},
	}
	if len(accounts) > 0 {
		mp := make(map[model.ComerAccountType]uint64)
		for _, account := range accounts {
			mp[account.Type] = account.ID
		}
		for _, info := range accountBindingInfos {
			if v, ok := mp[info.AccountType]; ok {
				info.AccountId = v
				info.Linked = true
			}
		}
	}
	log.Infof("comer accounts bidingInfos for %d : %v \n", comerID, accountBindingInfos)
	// profile may not exist!
	response.ComerID = comerID
	response.ComerAccounts = accountBindingInfos
	return
}

// CreateComerProfile  create a new profil for comer
// current comer should not be exists now
func CreateComerProfile(comerID uint64, post *model.CreateProfileRequest) (err error) {
	//get comer profile
	var profile model.ComerProfile
	if err = model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		log.Warn(err)
		return
	}
	if profile.ID != 0 {
		return router.ErrBadRequest.WithMsg("user profile already exists")
	}
	var tagRelList []tag.TagTargetRel
	profile = model.ComerProfile{
		ComerID:  comerID,
		Name:     post.Name,
		Avatar:   post.Avatar,
		Location: post.Location,
		TimeZone: post.TimeZone,
		Website:  post.Website,
		Email:    post.Email,
		Twitter:  post.Twitter,
		Discord:  post.Discord,
		Telegram: post.Telegram,
		Medium:   post.Medium,
		Facebook: post.Facebook,
		Linktree: post.Linktree,
		BIO:      post.BIO,
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		//create skill
		for _, skillName := range post.SKills {
			var isIndex bool
			if len(skillName) > 2 && skillName[0:1] == "#" {
				isIndex = true
			}
			skill := tag.Tag{
				Name:     skillName,
				IsIndex:  isIndex,
				Category: tag.ComerSkill,
			}
			if er = tag.FirstOrCreateTag(tx, &skill); err != nil {
				return er
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    skill.ID,
				Target:   tag.ComerSkill,
				TargetID: comerID,
			})
		}
		//batch create comer skill relation
		if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
			log.Warn(er)
			return
		}
		//create comer profile
		if er = model.CreateComerProfile(tx, &profile); er != nil {
			log.Warn(er)
			return
		}
		return nil
	})

	return err
}

// UpdateComerProfile update the comer profile
// if profile is not exists then will return the not exits error
func UpdateComerProfile(comerID uint64, post *model.UpdateProfileRequest) (err error) {
	//get comer profile
	var profile model.ComerProfile
	if err = model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if profile.ID == 0 {
		return router.ErrBadRequest.WithMsg("user profile does not exists")
	}
	profile = model.ComerProfile{
		ComerID:  comerID,
		Name:     post.Name,
		Avatar:   post.Avatar,
		Location: post.Location,
		TimeZone: post.TimeZone,
		Website:  post.Website,
		Email:    post.Email,
		Twitter:  post.Twitter,
		Discord:  post.Discord,
		Telegram: post.Telegram,
		Medium:   post.Medium,
		Facebook: post.Facebook,
		Linktree: post.Linktree,
		BIO:      post.BIO,
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) { //create skill
		if er = processSkills(tx, comerID, post.SKills); er != nil {
			return
		}
		//create profile
		if er = model.UpdateComerProfile(tx, &profile); er != nil {
			return er
		}
		return nil
	})
	return
}
func processSkills(tx *gorm.DB, comerId uint64, skills []string) (err error) {
	var tagIds []uint64
	var tagRelList []tag.TagTargetRel
	return tx.Transaction(func(tx *gorm.DB) (er error) { //create skill
		for _, skillName := range skills {
			var isIndex bool
			if len(skillName) > 1 && skillName[0:1] == "#" {
				isIndex = true
			}
			skill := tag.Tag{
				Name:     skillName,
				Category: tag.ComerSkill,
				IsIndex:  isIndex,
			}
			if er = tag.FirstOrCreateTag(tx, &skill); er != nil {
				return er
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    skill.ID,
				Target:   tag.ComerSkill,
				TargetID: comerId,
			})
			tagIds = append(tagIds, skill.ID)
		}
		//delete not used skills
		if er = tag.DeleteTagRel(tx, comerId, tag.ComerSkill, tagIds); er != nil {
			return er
		}
		//batch create comer skill rel
		if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
			return er
		}
		return nil
	})

}

func UpdateComerCover(comerID uint64, request model.UpdateComerCoverRequest) error {
	if strings.TrimSpace(request.Image) == "" {
		return router.ErrBadRequest.WithMsg("cover image is empty")
	}
	return model.UpdateComerCover(mysql.DB, comerID, request)
}

func GetComerModuleInfo(targetComerID uint64) (info []model.ModuleInfo, err error) {
	if err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		startupCnt, err := startupModel.CountStartupsPostedByComer(tx, targetComerID)
		if err != nil {
			return err
		}
		bountyCnt, err := bountyModel.CountBountiesPostedByComer(tx, targetComerID)
		if err != nil {
			return err
		}
		crowdfundingCnt, err := crowdfundingModel.CountCrowdfundingPostedByComer(tx, targetComerID)
		if err != nil {
			return err
		}
		// proposal add in the future
		//startupCnt, err := startupModel.CountStartupsPostedByComer(tx, targetComerID)
		//if err!=nil{
		//	return err
		//}
		// proposalCnt := 0
		info = append(info, []model.ModuleInfo{
			{Module: model2.ModuleStartup, HasCreated: startupCnt > 0},
			{Module: model2.ModuleBounty, HasCreated: bountyCnt > 0},
			{Module: model2.ModuleCrowdfunding, HasCreated: crowdfundingCnt > 0},
			{Module: model2.ModuleGovernance},
		}...)
		return nil
	}); err != nil {
		return
	}
	return
}

func UpdateComerSkill(comerId uint64, skills []string) error {
	return processSkills(mysql.DB, comerId, skills)
}

func UpdateComerBio(comerId uint64, bio string) error {
	return model.UpdateComerBio(mysql.DB, comerId, bio)
}

func UpdateLanguages(comerId uint64, request model.UpdateLanguageInfosRequest) error {
	if len(request.Languages) > 0 {
		for _, info := range request.Languages {
			if strings.TrimSpace(info.Language) == "" {
				return errors.New("language can not be empty")
			}
			if err := info.Level.Check(); err != nil {
				return err
			}
		}
	}
	return model.UpdateLanguageInfos(mysql.DB, comerId, request)
}

func UpdateEducations(comerId uint64, request model.UpdateEducationsRequest) error {
	return model.UpdateEducationInfos(mysql.DB, comerId, request)
}

func CreateOrUpdateBasic(comerId uint64, request model.UpdateBasicInfoRequest) error {
	var profile model.ComerProfile
	err := model.GetComerProfile(mysql.DB, comerId, &profile)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) || (profile.ID == 0) {
		profile = model.ComerProfile{
			ComerID:  comerId,
			Name:     request.Name,
			Avatar:   request.Avatar,
			Cover:    request.Cover,
			Location: request.Location,
			TimeZone: request.TimeZone,
		}
		if err = model.CreateComerProfile(mysql.DB, &profile); err != nil {
			return err
		}
		return nil
	}
	return model.UpdateBasicInfo(mysql.DB, comerId, request)
}
