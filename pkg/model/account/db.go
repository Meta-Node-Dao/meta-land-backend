package account

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
	"gorm.io/gorm"
	"math"
)

// GetComerByAddress  get comer entity by comer's address
func GetComerByAddress(db *gorm.DB, address string, comer *Comer) error {
	return db.Model(&Comer{}).Where("address = ? AND is_deleted = false", address).Find(comer).Error
}

// GetComerByID  get comer entity by comer's ID
func GetComerByID(db *gorm.DB, comerID uint64, comer *Comer) (err error) {
	return db.Model(&Comer{}).Where("id = ? AND is_deleted = false", comerID).Find(comer).Error
}

// CreateComer create a comer
func CreateComer(db *gorm.DB, comer *Comer) (err error) {
	return db.Create(comer).Error
}

// UpdateComerAddress update the comer address
func UpdateComerAddress(db *gorm.DB, comerID uint64, address string) (err error) {
	return db.Model(&Comer{Base: model.Base{ID: comerID}}).Update("address", address).Error
}

func GetComerAccount(db *gorm.DB, accountType ComerAccountType, oin string, comerAccount *ComerAccount) error {
	return db.Model(&ComerAccount{}).Where("type = ? AND oin = ? AND is_deleted = false", accountType, oin).Find(comerAccount).Error
}

func GetComerAccountById(db *gorm.DB, accountId uint64, comerAccount *ComerAccount) error {
	return db.Model(&ComerAccount{}).Where("id=? AND is_deleted = false", accountId).Find(comerAccount).Error
}

func ListAccount(db *gorm.DB, comerID uint64, accountList *[]ComerAccount) (err error) {
	return db.Model(&ComerAccount{}).Where("comer_id = ? AND is_deleted = false", comerID).Find(accountList).Error
}

func CreateAccount(db *gorm.DB, comerAccount *ComerAccount) (err error) {
	return db.Create(comerAccount).Error
}

func DeleteAccount(db *gorm.DB, comerID, accountID uint64) error {
	return db.Where("comer_id = ? AND id = ? AND is_deleted = false", comerID, accountID).Delete(&ComerAccount{}).Error
}

// GetComerProfile update the comer address
func GetComerProfile(db *gorm.DB, comerID uint64, profile *ComerProfile) (err error) {
	return db.Where("comer_id = ? AND is_deleted = false", comerID).Preload("Skills", "category = ?", tag.ComerSkill).Find(profile).Error
}

// CreateComerProfile update the comer address
func CreateComerProfile(db *gorm.DB, comerProfile *ComerProfile) error {
	return db.Create(&comerProfile).Error
}

// UpdateComerProfile update the comer address
func UpdateComerProfile(db *gorm.DB, comerProfile *ComerProfile) error {
	return db.Where("comer_id = ? AND is_deleted = false", comerProfile.ComerID).Select("avatar", "name", "location", "time_zone", "website", "email", "twitter", "discord", "telegram", "medium", "bio").Updates(comerProfile).Error
}

func UpdateComerProfileLocation(db *gorm.DB, comerId uint64, location string) error {
	return db.Model(&ComerProfile{}).Where("comer_id = ? AND is_deleted = false", comerId).Update("location", location).Error
}

// CreateComerFollowRel create comer relation for comer and target comer
func CreateComerFollowRel(db *gorm.DB, comerID, targetComerID uint64) error {
	return db.Create(&FollowRelation{ComerID: comerID, TargetComerID: targetComerID}).Error
}

// DeleteComerFollowRel delete comer relation for comer and target comer
func DeleteComerFollowRel(db *gorm.DB, input *FollowRelation) error {
	return db.Where("comer_id = ? AND target_comer_id = ?", input.ComerID, input.TargetComerID).Delete(input).Error
}

// ComerFollowIsExist check startup and comer is existed
func ComerFollowIsExist(db *gorm.DB, comerID, targetComerID uint64) (isExist bool, err error) {
	isExist = false
	var count int64
	err = db.Table("comer_follow_rel").Where("comer_id = ? AND target_comer_id = ?", comerID, targetComerID).Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		isExist = true
	}
	return
}

func ListFollowComer(db *gorm.DB, comerID uint64, output *[]FollowComer) (total int64, err error) {
	if comerID != 0 {
		db = db.Where("comer_id = ?", comerID)
	}
	err = db.Order("created_at ASC").Preload("Comer").Preload("ComerProfile").Preload("ComerProfile.Skills").Find(output).Count(&total).Error
	return
}

func ListFollowedComer(db *gorm.DB, comerID uint64, output *[]FollowedComer) (total int64, err error) {
	if comerID != 0 {
		db = db.Where("target_comer_id = ?", comerID)
	}
	err = db.Order("created_at ASC").Preload("Comer").Preload("ComerProfile").Preload("ComerProfile.Skills").Find(output).Count(&total).Error
	return
}

// BindComerAccountToComerId bind comerAccount to comer
func BindComerAccountToComerId(db *gorm.DB, comerAccountId, comerID uint64) (err error) {
	var crtAccount ComerAccount
	db.First(&crtAccount, comerAccountId)
	if crtAccount.ComerID == comerID || crtAccount.ComerID == 0 {
		return db.Model(&ComerAccount{Base: model.Base{ID: comerAccountId}}).Updates(ComerAccount{ComerID: comerID, IsLinked: true}).Error
	}
	return db.Transaction(func(tx *gorm.DB) (err error) {
		var comer Comer
		tx.First(&comer, crtAccount.ComerID)
		var accounts []ComerAccount
		if err = tx.Model(&ComerAccount{}).Where("comer_id = ? and is_deleted = false", comer.ID).Find(&accounts).Error; err != nil {
			return
		}
		if accounts == nil || (len(accounts) == 1 && comer.AddressStr() == "") {
			if err = tx.Delete(&comer).Error; err != nil {
				return
			}
			if err = tx.Model(&ComerAccount{Base: model.Base{ID: comerAccountId}}).Updates(ComerAccount{ComerID: comerID, IsLinked: true}).Error; err != nil {
				return
			}
		}
		return nil
	})
}

func GetComerAccountsByComerId(db *gorm.DB, comerId uint64, accounts *[]ComerAccount) (err error) {
	return db.Model(&ComerAccount{}).Where("comer_id = ? and is_deleted = false", comerId).Find(accounts).Error
}

func UpdateComerSocial(db *gorm.DB, comerId uint64, request SocialModifyRequest) error {
	return db.Model(&ComerProfile{}).
		Where("comer_id = ? and is_deleted = false", comerId).
		Update(request.SocialType.String(), request.SocialLink).Error
}

func UpdateComerCover(db *gorm.DB, comerID uint64, request UpdateComerCoverRequest) error {
	return db.Model(&ComerProfile{}).
		Where("comer_id = ? and is_deleted = false", comerID).
		Update("cover", request.Image).Error
}

func GetFollowersOfComer(db *gorm.DB, targetComerId uint64, pagination *model.Pagination) (comerProfiles []ComerProfile, err error) {
	// ComerFollowerInfo
	db = db.Table("comer_profile").
		Joins("LEFT JOIN comer ON comer.id = comer_profile.comer_id").
		Joins("LEFT JOIN comer_follow_rel ON  comer_follow_rel.comer_id = comer.id").
		Where("comer_follow_rel.target_comer_id = ? and comer.is_deleted = false and comer_profile.is_deleted = false", targetComerId)

	var cnt int64
	if err = db.Count(&cnt).Error; err != nil {
		return
	}

	if err = db.Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("created_at desc").
		Find(&comerProfiles).Error; err != nil {
		return
	}
	pagination.Rows = comerProfiles
	pagination.TotalRows = cnt
	pagination.TotalPages = int(math.Ceil(float64(cnt) / float64(pagination.Limit)))

	return
}

func GetFollowedByComer(db *gorm.DB, targetComerId uint64, pagination *model.Pagination) (comerProfiles []ComerProfile, err error) {
	// ComerFollowerInfo
	db = db.Table("comer_profile").
		Joins("LEFT JOIN comer ON  comer.id = comer_profile.comer_id").
		Joins("LEFT JOIN comer_follow_rel ON  comer_follow_rel.target_comer_id = comer.id").
		Where("comer_follow_rel.comer_id = ? and comer.is_deleted = false and comer_profile.is_deleted=false", targetComerId)
	var cnt int64
	if err = db.Count(&cnt).Error; err != nil {
		return
	}

	if err = db.Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("created_at desc").
		Find(&comerProfiles).Error; err != nil {
		return
	}
	pagination.Rows = comerProfiles
	pagination.TotalRows = cnt
	pagination.TotalPages = int(math.Ceil(float64(cnt) / float64(pagination.Limit)))

	return
}

func UpdateComerBio(db *gorm.DB, comerId uint64, bio string) error {
	return db.
		Model(&ComerProfile{}).
		Where("is_deleted = false and comer_id = ?", comerId).
		Update("bio", bio).
		Error
}

func UpdateLanguageInfos(db *gorm.DB, comerId uint64, request UpdateLanguageInfosRequest) error {
	return db.Model(&ComerProfile{}).
		Where("comer_id = ?", comerId).
		Update("languages", request.Languages).
		Error
}

func UpdateEducationInfos(db *gorm.DB, comerId uint64, request UpdateEducationsRequest) error {
	return db.Model(&ComerProfile{}).
		Where("comer_id = ?", comerId).
		Update("educations", request.Educations).
		Error
}

func UpdateBasicInfo(db *gorm.DB, comerId uint64, request UpdateBasicInfoRequest) error {
	return db.Model(&ComerProfile{}).
		Where("comer_id = ?", comerId).
		Updates(ComerProfile{
			Name:     request.Name,
			Avatar:   request.Avatar,
			Cover:    request.Cover,
			Location: request.Location,
			TimeZone: request.TimeZone,
		}).Error
}

func ProfileComerConnectedInfo(db *gorm.DB, comerId uint64) (info ComerConnectedInfo, err error) {
	var followedStartupCnt int64
	if err = db.Table("startup").
		Select("startup.id").
		Joins("INNER JOIN startup_follow_rel ON startup_follow_rel.startup_id = startup.id").
		Where("startup_follow_rel.comer_id = ? and startup.is_deleted = false", comerId).
		Count(&followedStartupCnt).Error; err != nil {
		return
	}
	var followedComerCnt int64
	if err = db.Table("comer").
		Select("comer.id").
		Joins("INNER JOIN comer_follow_rel ON comer_follow_rel.target_comer_id = comer.id").
		Where("comer_follow_rel.comer_id = ? and comer.is_deleted = false", comerId).
		Count(&followedComerCnt).Error; err != nil {
		return
	}
	var comerFollowingCnt int64
	if err = db.Table("comer").
		Select("comer.id").
		Joins("INNER JOIN comer_follow_rel ON comer_follow_rel.comer_id = comer.id").
		Where("comer_follow_rel.target_comer_id = ? and comer.is_deleted = false", comerId).
		Count(&comerFollowingCnt).Error; err != nil {
		return
	}
	info = ComerConnectedInfo{
		StartupCnt:  followedStartupCnt,
		ComerCnt:    followedComerCnt,
		FollowerCnt: comerFollowingCnt,
	}
	return
}

func ProfileComerModuleDataInfo(db *gorm.DB, comerId uint64, dataType BusinessModuleDataType) (info ComerModuleDataInfo, err error) {
	var (
		startupCnt,
		bountyCnt,
		crowdfundingCnt,
		proposalCnt int64
	)
	if dataType == Posted {
		if err = db.Table("startup").
			Where("comer_id = ? and is_deleted = false", comerId).
			Count(&startupCnt).Error; err != nil {
			return
		}

		if err = db.Table("bounty").
			Where("comer_id = ? and is_deleted = false", comerId).
			Count(&bountyCnt).Error; err != nil {
			return
		}

		if err = db.Table("crowdfunding").
			Where("comer_id = ? and is_deleted = false", comerId).
			Count(&crowdfundingCnt).Error; err != nil {
			return
		}

		if err = db.Table("governance_proposal").
			Where("author_comer_id = ? and is_deleted = false", comerId).
			Count(&proposalCnt).Error; err != nil {
			return
		}
	} else {
		if err = db.
			Where("is_deleted = false").
			Joins("INNER JOIN startup_team_member_rel ON startup_team_member_rel.comer_id = ? AND startup_id = startup.id AND startup.comer_id != startup_team_member_rel.comer_id", comerId).
			Table("startup").Count(&startupCnt).Error; err != nil {
			return
		}

		if err = db.Table("bounty").
			Select("bounty.id").
			Joins("INNER JOIN bounty_applicant ON bounty_applicant.bounty_id = bounty.id").
			Where("bounty_applicant.comer_id = ? and bounty.is_deleted = false", comerId).
			Group("bounty.id").
			Count(&bountyCnt).Error; err != nil {
			return
		}

		if err = db.Table("crowdfunding").
			Select("crowdfunding.id").
			Joins("INNER JOIN crowdfunding_investor ON crowdfunding_investor.crowdfunding_id = crowdfunding.id").
			Where("crowdfunding_investor.comer_id = ? and crowdfunding.is_deleted = false", comerId).
			Group("crowdfunding.id").
			Count(&crowdfundingCnt).Error; err != nil {
			return
		}

		if err = db.Table("governance_proposal").
			Select("governance_proposal.id").
			Joins("LEFT JOIN  governance_vote ON governance_proposal.id = governance_vote.proposal_id").
			Where("governance_vote.voter_comer_id = ? and governance_proposal.is_deleted = false", comerId).
			Group("governance_proposal.id").
			Count(&proposalCnt).Error; err != nil {
			return
		}
	}
	info = ComerModuleDataInfo{
		Type:            dataType,
		StartupCnt:      startupCnt,
		BountyCnt:       bountyCnt,
		CrowdfundingCnt: crowdfundingCnt,
		ProposalCnt:     proposalCnt,
	}
	return
}
