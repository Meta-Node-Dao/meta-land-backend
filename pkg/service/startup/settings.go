package startup

import (
	"ceres/pkg/initialization/mysql"
	model2 "ceres/pkg/model"
	model "ceres/pkg/model/startup"
	"ceres/pkg/model/tag"
	"ceres/pkg/router"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/qiniu/x/log"
)

// UpdateStartupBasicSetting update startup security and social setting
func UpdateStartupBasicSetting(startupID uint64, request *model.UpdateStartupBasicSettingRequest) (err error) {
	//get startup
	var startup model.Startup
	if err = model.GetStartup(mysql.DB, startupID, &startup); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if startup.ID == 0 {
		return router.ErrBadRequest.WithMsg("startup does not exist")
	}
	var tagIds []uint64
	var tagRelList []tag.TagTargetRel
	startupBasicSetting := model.BasicSetting{
		KYC:           startup.KYC,
		ContractAudit: startup.ContractAudit,
		Website:       startup.Website,
		Discord:       startup.Discord,
		Twitter:       startup.Twitter,
		Telegram:      startup.Telegram,
		Docs:          startup.Docs,
	}
	if request.KYC != nil {
		startupBasicSetting.KYC = *request.KYC
	}
	if request.ContractAudit != nil {
		startupBasicSetting.ContractAudit = *request.ContractAudit
	}
	if request.Website != nil {
		startupBasicSetting.Website = *request.Website
	}
	if request.Discord != nil {
		startupBasicSetting.Discord = *request.Discord
	}
	if request.Twitter != nil {
		startupBasicSetting.Twitter = *request.Twitter
	}
	if request.Telegram != nil {
		startupBasicSetting.Telegram = *request.Telegram
	}
	if request.Docs != nil {
		startupBasicSetting.Docs = *request.Docs
	}
	if err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
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
				TargetID: startupID,
			})
			tagIds = append(tagIds, hashTag.ID)
		}
		if len(tagIds) > 0 {
			//delete not used hashtags
			if er = tag.DeleteTagRel(tx, startupID, tag.Startup, tagIds); er != nil {
				return er
			}
		}
		if len(tagRelList) > 0 {
			//batch create startup hashtag rel
			if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
				return er
			}
		}
		//update startup basic setting
		if er = model.UpdateStartupBasicSetting(tx, startupID, &startupBasicSetting); er != nil {
			return er
		}
		return er
	}); err != nil {
		log.Warn(err)
		return
	}
	return
}

func UpdateStartupFinanceSetting(startupID, comerID uint64, request *model.UpdateStartupFinanceSettingRequest) (err error) {
	//get startup
	var startup model.Startup
	if err = model.GetStartup(mysql.DB, startupID, &startup); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if startup.ID == 0 {
		return router.ErrBadRequest.WithMsg("startup does not exist")
	}
	var walletIds []uint64
	var walletList []model.Wallet
	startupFinanceSetting := model.FinanceSetting{
		TokenContractAddress: *request.TokenContractAddress,
		LaunchNetwork:        *request.LaunchNetwork,
		TokenName:            *request.TokenName,
		TokenSymbol:          *request.TokenSymbol,
		TotalSupply:          *request.TotalSupply,
		PresaleStart:         ConverToDatetime(*request.PresaleStart),
		PresaleEnd:           ConverToDatetime(*request.PresaleEnd),
		LaunchDate:           ConverToDatetime(*request.LaunchDate),
	}

	if err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		for _, v := range request.Wallets {
			wallet := model.Wallet{
				ComerID:       comerID,
				StartupID:     startup.ID,
				WalletName:    v.WalletName,
				WalletAddress: v.WalletAddress,
			}
			if er = model.FirstOrCreateWallet(tx, &wallet); er != nil {
				return er
			}
			wallet.WalletName = v.WalletName
			wallet.WalletAddress = v.WalletAddress
			walletList = append(walletList, wallet)
			walletIds = append(walletIds, wallet.ID)
		}
		//batch update startup wallet
		if er = model.BatchUpdateStartupWallet(tx, walletList); er != nil {
			return er
		}
		//delete not used startup wallet
		if er = model.DeleteStartupWallet(tx, startupID, walletIds); er != nil {
			return er
		}
		//update startup finance setting
		if er = model.UpdateStartupFinanceSetting(tx, startupID, &startupFinanceSetting); er != nil {
			return er
		}
		return er
	}); err != nil {
		log.Warn(err)
		return
	}
	return
}

func ConverToDatetime(strTime string) (t time.Time) {
	var err error
	const timeFormat = "2006-01-02T15:04:05Z"
	if t, err = time.ParseInLocation(timeFormat, strTime, time.UTC); err != nil {
		if t, err = time.ParseInLocation("2006-01-02", strTime, time.UTC); err != nil {
			t = time.Time{}
			return
		}
	}
	return t
}

func UpdateStartupCover(startupId, comerId uint64, request model.UpdateStartupCoverRequest) error {
	if strings.TrimSpace(request.Image) == "" {
		return router.ErrBadRequest.WithMsg("cover image is empty")
	}
	return model.UpdateStartupCover(mysql.DB, startupId, comerId, request)
}

func UpdateStartupSecurity(startupId, comerId uint64, request model.UpdateStartupSecurityRequest) error {
	return model.UpdateStartupSecurity(mysql.DB, startupId, comerId, request)
}

func UpdateStartupTabSequence(startupId, comerId uint64, request model.UpdateStartupTabSequenceRequest) error {
	st, err := model.GetStartupById(mysql.DB, startupId)
	if err != nil {
		return err
	}
	if st.ComerID != comerId {
		return errors.New("can not update tab sequence, current comer is not founder of startup")
	}
	if request.Tabs == nil || len(request.Tabs) == 0 {
		request.Tabs = model2.DefaultModules
	} else if len(request.Tabs) < len(model2.DefaultModules) {
		return errors.New("invalid tab sequence param")
	}
	return model.UpdateStartupTabSequence(mysql.DB, startupId, request)
}

func ResetStartupTabSequence(startupId, comerId uint64) error {
	return UpdateStartupTabSequence(startupId, comerId, model.UpdateStartupTabSequenceRequest{Tabs: nil})
}

func UpdateSocialsAndTags(startupId, comerId uint64, request model.UpdateStartupSocialsAndTagsRequest) error {

	st, err := model.GetStartupById(mysql.DB, startupId)
	if err != nil {
		return err
	}
	if st.ComerID != comerId {
		return errors.New("can not update social, current comer is not founder of startup")
	}
	var tagIds []uint64
	var tagRelList []tag.TagTargetRel
	return mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
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
				TargetID: startupId,
			})
			tagIds = append(tagIds, hashTag.ID)
		}
		if len(tagIds) > 0 {
			//delete not used hashtags
			if er = tag.DeleteTagRel(tx, startupId, tag.Startup, tagIds); er != nil {
				return er
			}
		}
		if len(tagRelList) > 0 {
			//batch create startup hashtag rel
			if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
				return er
			}
		}
		toBeUpdated := map[string]interface{}{}
		if len(request.Socials) > 0 {
			for _, social := range request.Socials {
				toBeUpdated[social.SocialType.String()] = social.SocialLink
			}
		}
		if len(request.DeletedSocials) > 0 {
			for _, social := range request.DeletedSocials {
				toBeUpdated[social.String()] = ""
			}
		}
		if len(toBeUpdated) > 0 {
			return tx.Model(&model.Startup{}).Where("id = ? and is_deleted = false", startupId).Updates(toBeUpdated).Error
		}

		return nil
	})

}

func UpdateStartupBasicSettingNew(startupId, comerId uint64, request model.UpdateStartupBasicSettingRequestNew) error {
	st, err := model.GetStartupById(mysql.DB, startupId)
	if err != nil {
		return err
	}
	if st.ComerID != comerId {
		return errors.New("can not update basic setting, current comer is not founder of startup")
	}
	update := model.Startup{
		Name:     request.Name,
		Mode:     request.Mode,
		Logo:     request.Logo,
		Mission:  request.Mission,
		Overview: request.Overview,
		Cover:    request.Cover,
		// TxHash:   request.TxHash,
		// ChainID: request.ChainID,
	}
	// 判断是否已上链 现在不确定is_set是个什么字段 没有使用过 也没有文档 暂时使用这个字段作为上链的依据
	// 已经将 is_set 字段修改为 on_chain 满足规范使用
	if !st.OnChain {
		update.TxHash = request.TxHash
		update.ChainID = request.ChainID
	}

	var tagIds []uint64
	var tagRelList []tag.TagTargetRel
	return mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		if er = tx.Model(&model.Startup{}).
			Where("id = ? and is_deleted = false", startupId).
			Updates(update).Error; er != nil {
			return er
		}
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
				TargetID: startupId,
			})
			tagIds = append(tagIds, hashTag.ID)
		}
		if len(tagIds) > 0 {
			//delete not used hashtags
			if er = tag.DeleteTagRel(tx, startupId, tag.Startup, tagIds); er != nil {
				return er
			}
		}
		if len(tagRelList) > 0 {
			//batch create startup hashtag rel
			if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
				return er
			}
		}
		return nil
	})
}
