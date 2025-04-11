package startup

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
	"database/sql"
	"encoding/json"
	"math"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

// GetStartup  get startup
func GetStartup(db *gorm.DB, startupID uint64, startup *Startup) error {
	return db.Debug().Where("is_deleted = false AND id = ?", startupID).
		Preload("Wallets").
		Preload("HashTags", "category = ?", tag.Startup).
		Preload("Members").Preload("Members.Comer").
		Preload("Members.ComerProfile").
		Preload("Follows").
		Find(&startup).Error
}

func GetStartupById(db *gorm.DB, startup uint64) (st Startup, err error) {
	st = Startup{}
	err = db.Model(&Startup{}).Where("is_deleted = false and id = ?", startup).Find(&st).Error
	return
}

// CreateStartup  create startup
func CreateStartup(db *gorm.DB, startup *Startup) (err error) {
	return db.Create(startup).Error
}

// StartupOnTheChain Satrtup On-chain operation
func StartupOnChain(db *gorm.DB, txHash string, chainID uint64, comerID uint64) (err error) {
	return db.Where(Startup{
		TxHash:  txHash,
		ChainID: chainID,
		ComerID: comerID,
		OnChain: false,
	}).Select("on_chain").Updates(Startup{
		OnChain: true,
	}).Error
}

// CreateStartupWallet  create startup wallet
func CreateStartupWallet(db *gorm.DB, wallets []Wallet) (err error) {
	return db.Create(&wallets).Error
}

// BatchUpdateStartupWallet  batch update startup wallets
func BatchUpdateStartupWallet(db *gorm.DB, wallets []Wallet) (err error) {
	return db.Save(&wallets).Error
}

// FirstOrCreateWallet first or create wallet
func FirstOrCreateWallet(db *gorm.DB, wallet *Wallet) error {
	return db.Where("startup_id = ? AND wallet_name = ? ", wallet.StartupID, wallet.WalletName).FirstOrCreate(&wallet).Error
}

// DeleteStartupWallet delete startup wallet where not used
func DeleteStartupWallet(db *gorm.DB, startupID uint64, walletIds []uint64) error {
	return db.Delete(&Wallet{}, "startup_id = ? AND id NOT IN ?", startupID, walletIds).Error
}

// ListStartups  list startups
func ListStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false")
	if comerID != 0 {
		db = db.Where("comer_id = ?", comerID)
	}
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").Preload("Members.ComerProfile").Preload("Follows").Find(startups).Error
	return
}

// CreateStartupFollowRel create comer relation for startup and comer
func CreateStartupFollowRel(db *gorm.DB, comerID, startupID uint64) error {
	return db.Create(&FollowRelation{ComerID: comerID, StartupID: startupID}).Error
}

// DeleteStartupFollowRel delete comer relation for startup and comer
func DeleteStartupFollowRel(db *gorm.DB, input *FollowRelation) error {
	return db.Where("comer_id = ? AND startup_id = ?", input.ComerID, input.StartupID).Delete(input).Error
}

// ListFollowedStartups  list followed startups
func ListFollowedStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false").Joins("INNER JOIN startup_follow_rel ON startup_follow_rel.comer_id = ? AND startup_id = startup.id", comerID)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").Preload("Members.ComerProfile").Preload("Follows").Find(startups).Error
	return
}

// StartupNameIsExist check startup's  name is existed
func StartupNameIsExist(db *gorm.DB, name string) (isExit bool, err error) {
	var count int64
	err = db.Table("startup").Where("name = ?", name).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExit = false
	} else {
		isExit = true
	}
	return
}

// StartupTokenContractIsExist check startup's  token contract is existed
func StartupTokenContractIsExist(db *gorm.DB, tokenContract string) (isExit bool, err error) {
	var count int64
	err = db.Table("startup").Where("token_contract_address = ?", tokenContract).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExit = false
	} else {
		isExit = true
	}
	return
}

// UpdateStartupBasicSetting  update startup security and social setting
func UpdateStartupBasicSetting(db *gorm.DB, startupID uint64, input *BasicSetting) (err error) {
	return db.Table("startup").Where("id = ?", startupID).Select("kyc", "contract_audit", "website", "discord", "twitter", "telegram", "docs").Updates(input).Error
}

// UpdateStartupFinanceSetting  update startup finance setting
func UpdateStartupFinanceSetting(db *gorm.DB, startupID uint64, input *FinanceSetting) (err error) {
	var fieldMap map[string]interface{}
	fieldMap = make(map[string]interface{})
	fieldMap["token_contract_address"] = input.TokenContractAddress
	fieldMap["launch_network"] = input.LaunchNetwork
	fieldMap["token_name"] = input.TokenName
	fieldMap["token_symbol"] = input.TokenSymbol
	fieldMap["total_supply"] = input.TotalSupply
	if input.PresaleStart.IsZero() {
		fieldMap["presale_start"] = sql.NullTime{}
	} else {
		fieldMap["presale_start"] = input.PresaleStart
	}
	if input.PresaleEnd.IsZero() {
		fieldMap["presale_end"] = sql.NullTime{}
	} else {
		fieldMap["presale_end"] = input.PresaleEnd
	}
	if input.LaunchDate.IsZero() {
		fieldMap["launch_date"] = sql.NullTime{}
	} else {
		fieldMap["launch_date"] = input.LaunchDate
	}

	return db.Table("startup").Where("id = ?", startupID).Updates(fieldMap).Error
}

// ListParticipatedStartups  list participated startups
func ListParticipatedStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false").Joins("INNER JOIN startup_team_member_rel ON startup_team_member_rel.comer_id = ? AND startup_id = startup.id AND startup.comer_id != startup_team_member_rel.comer_id", comerID)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").Preload("Members.ComerProfile").Preload("Follows").Find(startups).Error
	return
}

// ListBeMemberStartups  list I am a member of startups
func ListBeMemberStartups(db *gorm.DB, comerID uint64, input *ListStartupRequest, startups *[]Startup) (total int64, err error) {
	db = db.Where("is_deleted = false").Joins("INNER JOIN startup_team_member_rel ON startup_team_member_rel.comer_id = ? AND startup_id = startup.id", comerID)
	if input.Keyword != "" {
		db = db.Where("name like ?", "%"+input.Keyword+"%")
	}
	if input.Mode != 0 {
		db = db.Where("mode = ?", input.Mode)
	}
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Wallets").Preload("HashTags", "category = ?", tag.Startup).Preload("Members").Preload("Members.Comer").Preload("Members.ComerProfile").Preload("Follows").Find(startups).Error
	return
}

// StartupFollowIsExist check startup and comer is existed
func StartupFollowIsExist(db *gorm.DB, startupID, comerID uint64) (isExist bool, err error) {
	var count int64
	err = db.Table("startup_follow_rel").Where("startup_id = ? AND comer_id = ?", startupID, comerID).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		isExist = false
	} else {
		isExist = true
	}
	return
}

// GetComerStartups  list comer all startups
func ListComerStartups(db *gorm.DB, comerID uint64, startups []*ListComerStartup) ([]*ListComerStartup, error) {
	err := db.Table("startup").Select("id, name, on_chain, comer_id").Where("comer_id = ? and is_deleted = 0", comerID).Order("convert(name using gbk)").Find(&startups).Error
	if err != nil {
		return nil, err
	}
	return startups, nil
}

func UpdateStartupCover(db *gorm.DB, startupId, comerId uint64, request UpdateStartupCoverRequest) error {
	return db.Model(&Startup{}).Where("id = ? and comer_id = ?", startupId, comerId).Update("cover", request.Image).Error
}

func UpdateStartupSecurity(db *gorm.DB, startupId, comerId uint64, request UpdateStartupSecurityRequest) error {
	return db.Model(&Startup{}).Where("id = ? and comer_id = ?", startupId, comerId).
		Updates(map[string]interface{}{"kyc": request.KYC, "contract_audit": request.ContractAudit}).Error
}

func CountStartupsPostedByComer(db *gorm.DB, targetComerID uint64) (cnt int64, err error) {
	err = db.Model(&Startup{}).Where("is_deleted = false and comer_id = ?", targetComerID).Count(&cnt).Error
	return
}

func UpdateStartupTabSequence(db *gorm.DB, startupId uint64, request UpdateStartupTabSequenceRequest) error {
	bytes, _ := json.Marshal(request.Tabs)
	return db.Model(&Startup{}).Where("is_deleted = false and id = ?", startupId).Update("tab_sequence", datatypes.JSON(bytes)).Error
}

func ExistStartupGroupByName(db *gorm.DB, startupId uint64, name string) (exist bool, err error) {
	var cnt int64
	err = db.Model(&StartupGroup{}).Where("startup_id = ? and name = ?", startupId, name).Count(&cnt).Error
	if cnt > 0 {
		exist = true
	}
	return
}

func CreateStartupGroup(db *gorm.DB, group *StartupGroup) error {
	return db.Create(group).Error
}

func GetStartupGroupById(db *gorm.DB, groupId uint64) (group StartupGroup, err error) {
	err = db.Model(&StartupGroup{}).Where("id = ?", groupId).Find(&group).Error
	return
}

func DeleteStartupGroup(db *gorm.DB, groupID uint64) error {
	return db.Delete(&StartupGroup{}, "id = ?", groupID).Error
}
func DeleteStartupGroupMemberRelsByGroupId(db *gorm.DB, groupId uint64) error {
	return db.Delete(&StartupGroupMemberRel{}, "group_id = ?", groupId).Error
}

func DeleteStartupGroupMemberRelsByComerIdAndStartupId(db *gorm.DB, comerID, startupId uint64) error {
	return db.Delete(&StartupGroupMemberRel{}, "comer_id = ? and startup_id = ?", comerID, startupId).Error
}

func UpdateStartupGroup(db *gorm.DB, groupID uint64, name string) error {
	return db.Model(&StartupGroup{}).Where("id = ?", groupID).Update("name", name).Error
}

func SelectStartupGroupsByStartupId(db *gorm.DB, startupId uint64) (groups []*StartupGroup, err error) {
	err = db.Model(&StartupGroup{}).Where("startup_id = ?", startupId).Find(&groups).Error
	return
}

func GetGroupMemberRelByComerIdAndStartupId(db *gorm.DB, comerID, startupID uint64) (rel StartupGroupMemberRel, err error) {
	err = db.Model(&StartupGroupMemberRel{}).Where("comer_id = ? and startup_id = ?", comerID, startupID).Find(&rel).Error
	return
}

func CreateGroupMemberRel(db *gorm.DB, rel *StartupGroupMemberRel) error {
	return db.Create(rel).Error
}

func UpdateGroupMemberRel(db *gorm.DB, rel StartupGroupMemberRel) error {
	return db.Model(&StartupGroupMemberRel{RelationBase: model.RelationBase{ID: rel.ID}}).Updates(rel).Error
}

func GetFollowedStartupsOfComer(db *gorm.DB, comerID uint64, page *model.Pagination) (startups []Startup, err error) {
	db = db.Where("is_deleted = false").
		Joins("INNER JOIN startup_follow_rel ON startup_follow_rel.comer_id = ? AND startup_id = startup.id", comerID)

	var total int64
	if err = db.Table("startup").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").
		Limit(page.GetLimit()).
		Offset(page.GetOffset()).
		Find(&startups).Error
	page.Rows = startups
	page.TotalRows = total
	page.TotalPages = int(math.Ceil(float64(total) / float64(page.Limit)))
	return
}

func ProfileStartupModuleDataInfo(db *gorm.DB, startupId uint64) (info StartupModuleDataInfo, err error) {
	var (
		bountyCnt,
		crowdfundingCnt,
		proposalCnt int64
	)

	if err = db.Table("bounty").
		Where("startup_id = ? and is_deleted = false and status in ?", startupId, []int{1, 2, 3, 4}).
		Count(&bountyCnt).Error; err != nil {
		return
	}

	if err = db.Table("crowdfunding").
		Where("startup_id = ? and is_deleted = false", startupId).
		Count(&crowdfundingCnt).Error; err != nil {
		return
	}
	if err = db.Table("governance_proposal").
		Where("startup_id = ? and is_deleted = false", startupId).
		Count(&proposalCnt).Error; err != nil {
		return
	}
	info = StartupModuleDataInfo{
		BountyCnt:       bountyCnt,
		CrowdfundingCnt: crowdfundingCnt,
		ProposalCnt:     proposalCnt,
		OtherDappCnt:    0,
	}
	return
}

func SelectStartupFans(db *gorm.DB, startupID uint64, pagination *model.Pagination) (fans StartupFans, err error) {
	db = db.Where("startup_follow_rel.startup_id = ?", startupID).
		Select("comer_profile.comer_id, comer_profile.avatar, comer_profile.name").
		Joins("left join comer_profile on comer_profile.comer_id = startup_follow_rel.comer_id").
		Where("comer_profile.is_deleted = false")

	var cnt int64
	if err = db.Table("startup_follow_rel").Count(&cnt).Error; err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}
	pagination.TotalRows = cnt
	if err = db.
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Table("startup_follow_rel").
		Scan(&fans).Error; err != nil {
		return nil, err
	}

	pagination.Rows = fans
	pagination.TotalPages = int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	return
}
