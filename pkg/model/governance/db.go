package governance

import (
	"ceres/pkg/model"
	"gorm.io/gorm"
	"math"
	"time"
)

func CreateGovernanceSetting(db *gorm.DB, setting *GovernanceSetting) error {
	return db.Create(setting).Error
}

func UpdateGovernanceSetting(db *gorm.DB, settingId uint64, request *GovernanceSetting) error {
	return db.Model(&GovernanceSetting{}).Where("id = ?", settingId).Updates(request).Error
}

func DeleteGovernanceSetting(db *gorm.DB, settingId uint64) error {
	return db.Where("id = ?", settingId).Delete(&GovernanceSetting{}).Error
}

func CreateGovernanceStrategies(db *gorm.DB, strategies []*GovernanceStrategy) error {
	return db.Create(&strategies).Error
}

func DeleteStrategiesBySettingId(db *gorm.DB, settingId uint64) error {
	return db.Where("setting_id = ?", settingId).Delete(&GovernanceStrategy{}).Error
}

func CreateGovernanceAdmins(db *gorm.DB, admins []*GovernanceAdmin) error {
	return db.Create(&admins).Error
}
func DeleteAdminsBySettingId(db *gorm.DB, settingId uint64) error {
	return db.Where("setting_id = ?", settingId).Delete(&GovernanceAdmin{}).Error
}

func CreateProposal(db *gorm.DB, proposal *GovernanceProposal) error {
	return db.Create(proposal).Error
}

func DeleteProposal(db *gorm.DB, comerId, proposalId uint64) error {
	return db.Model(&GovernanceProposal{}).Where("id = ?", proposalId).Update("is_deleted", true).Error
}

func GetProposalById(db *gorm.DB, proposalId uint64) (proposal GovernanceProposal, err error) {
	err = db.Model(&GovernanceProposal{}).Where("id = ? and is_deleted = false", proposalId).Find(&proposal).Error
	return
}

func GetProposalPublicInfo(db *gorm.DB, proposalId uint64) (proposal ProposalPublicInfo, err error) {
	err = db.
		Where("governance_proposal.id = ?", proposalId).
		Select("governance_proposal.*,governance_setting.*,startup.logo as startup_logo, startup.name as startup_name, comer_profile.avatar as author_comer_avatar, comer_profile.name as author_comer_name").
		Joins("left join startup on startup.id = governance_proposal.startup_id").
		Joins("left join governance_setting on governance_setting.startup_id = governance_proposal.startup_id").
		Joins("left join comer_profile on comer_profile.comer_id = governance_proposal.author_comer_id").
		Table("governance_proposal").Scan(&proposal).Error
	return
}

func CreateProposalChoices(db *gorm.DB, choices []*GovernanceChoice) error {
	return db.Create(&choices).Error
}

func DeleteProposalChoices(db *gorm.DB, proposalId uint64) error {
	return db.Model(&GovernanceChoice{}).Where("proposal_id = ?", proposalId).Update("is_deleted", true).Error
}

func DeleteVoteByProposalIdAndVoterComer(db *gorm.DB, proposalId, voterComerId uint64) error {
	return db.Where("proposal_id = ? and voter_comer_id = ?", proposalId, voterComerId).Delete(&GovernanceVote{}).Error
}

func GetProposalChoices(db *gorm.DB, proposalId uint64) (choices []*GovernanceChoice, err error) {
	err = db.Model(&GovernanceChoice{}).Where("proposal_id = ?", proposalId).Order("seq_num asc").Find(&choices).Error
	return
}

func CreateProposalVote(db *gorm.DB, vote *GovernanceVote) error {
	return db.Create(vote).Error
}

func ReVoteChoiceOfProposal(db *gorm.DB, voteID uint64, request VoteRequest) error {
	return db.Model(&GovernanceVote{}).Where("id = ?", voteID).
		Updates(map[string]interface{}{"votes": request.Votes, "ipfs_hash": request.IpfsHash}).
		Error
}

func GetChoiceByProposalIdAndChoiceId(db *gorm.DB, proposalId, choiceItemId uint64) (choice GovernanceChoice, err error) {
	err = db.Model(&GovernanceChoice{}).
		Where("proposal_id = ? and id = ?", proposalId, choiceItemId).
		Find(&choice).
		Error
	return
}

func GetGovernanceSetting(db *gorm.DB, startupId uint64) (setting GovernanceSetting, err error) {
	err = db.Model(&GovernanceSetting{}).Where("startup_id = ?", startupId).Find(&setting).Error
	return
}

func GetGovernanceStrategies(db *gorm.DB, settingId uint64) (strategies GovernanceStrategies, err error) {
	err = db.Model(&GovernanceStrategy{}).Where("setting_id = ?", settingId).Find(&strategies).Error
	return
}

func GetGovernanceStrategiesByAuthorComerID(db *gorm.DB, startupId uint64) (strategies GovernanceStrategies, err error) {
	err = db.Table("governance_strategy").
		Joins("left join governance_setting on governance_setting.id = governance_strategy.setting_id").
		Where("governance_setting.startup_id = ?", startupId).
		Scan(&strategies).Error
	return
}

func GetGovernanceAdmins(db *gorm.DB, settingId uint64) (admins GovernanceAdmins, err error) {
	err = db.Model(&GovernanceAdmin{}).Where("setting_id = ?", settingId).Find(&admins).Error
	return
}

func GetGovernanceAdminsByStartupId(db *gorm.DB, startupId uint64) (admins GovernanceAdmins, err error) {
	err = db.Table("governance_admin").
		Joins("left join governance_setting on governance_setting.id = governance_admin.setting_id").
		Where("governance_setting.startup_id = ?", startupId).
		Scan(&admins).Error
	return
}

func GetVoteRecordByProposalIdAndComerId(db *gorm.DB, proposalId, comerId uint64) (vote GovernanceVote, err error) {
	err = db.Model(&GovernanceVote{}).
		Where("proposal_id = ? and voter_comer_id = ?", proposalId, comerId).
		Find(&vote).Error
	return
}

func GetVoteByProposalIdAndComerIdAndChoiceId(db *gorm.DB, proposalId, choiceItemId, comerId uint64) (vote GovernanceVote, err error) {
	err = db.Model(&GovernanceVote{}).
		Where("proposal_id = ? and choice_item_id = ? and voter_comer_id = ?", proposalId, choiceItemId, comerId).
		Find(&vote).Error
	return
}

func GetVotesOfProposal(db *gorm.DB, proposalId uint64) (records VoteRecords, err error) {
	err = db.Where("governance_vote.proposal_id = ?", proposalId).Find(&records).Error
	return
}
func GetVoteRecordsByProposalId(db *gorm.DB, proposalId uint64, pagination *model.Pagination) (records VoteRecords, err error) {
	db = db.Where("governance_vote.proposal_id = ?", proposalId).
		Select("governance_vote.*, comer_profile.avatar as voter_comer_avatar, comer_profile.name as voter_comer_name, governance_choice.item_name as choice_item_name").
		Joins("left join comer_profile on comer_profile.comer_id = governance_vote.voter_comer_id").
		Joins("left join governance_choice on governance_choice.id = governance_vote.choice_item_id")
	var cnt int64
	if err = db.Table("governance_vote").Count(&cnt).Error; err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}
	if err = db.
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("governance_vote.created_at desc").
		Table("governance_vote").
		Scan(&records).Error; err != nil {
		return nil, err
	}
	pagination.TotalRows = cnt
	pagination.TotalPages = int(float64(cnt) / float64(pagination.GetLimit()))
	pagination.Rows = records
	return
}

func SelectProposalList(db *gorm.DB, request *ProposalListRequest) (proposals []ProposalPublicInfo, err error) {
	db = db.Where("governance_proposal.is_deleted = false")
	if len(request.States) != 0 {
		db = db.Where("governance_proposal.status in ?", request.States)
	}
	db = db.Select("governance_proposal.*, startup.logo as startup_logo, startup.name as startup_name, comer_profile.avatar as author_comer_avatar, comer_profile.name as author_comer_name").
		Joins("left join startup on startup.id = governance_proposal.startup_id").
		Joins("left join comer_profile on comer_profile.comer_id = governance_proposal.author_comer_id")
	var cnt int64
	if err := db.Table("governance_proposal").Count(&cnt).Error; err != nil {
		return proposals, err
	}
	if cnt == 0 {
		return proposals, nil
	}
	if err := db.
		Offset(request.GetOffset()).
		Limit(request.GetLimit()).
		Order("governance_proposal.created_at desc").
		Table("governance_proposal").Scan(&proposals).Error; err != nil {
		return proposals, err
	}
	request.TotalRows = cnt
	request.TotalPages = int(math.Ceil(float64(cnt) / float64(request.Limit)))
	request.Rows = proposals
	return
}

func SelectProposalListByStartupId(db *gorm.DB, startupId uint64, request *model.Pagination) (proposals []ProposalPublicInfo, err error) {
	db = db.Where("governance_proposal.startup_id = ? and governance_proposal.is_deleted = false", startupId).
		Select("governance_proposal.*, startup.logo as startup_logo, startup.name as startup_name, comer_profile.avatar as author_comer_avatar, comer_profile.name as author_comer_name").
		Joins("left join startup on startup.id = governance_proposal.startup_id").
		Joins("left join comer_profile on comer_profile.comer_id = governance_proposal.author_comer_id")
	var cnt int64
	if err := db.Table("governance_proposal").Count(&cnt).Error; err != nil {
		return proposals, err
	}
	if cnt == 0 {
		return proposals, nil
	}
	if err := db.
		Offset(request.GetOffset()).
		Limit(request.GetLimit()).
		Order("governance_proposal.created_at desc").
		Table("governance_proposal").Scan(&proposals).Error; err != nil {
		return proposals, err
	}
	request.TotalRows = cnt
	request.TotalPages = int(math.Ceil(float64(cnt) / float64(request.Limit)))
	request.Rows = proposals
	return
}

func SelectProposalListByComerPosted(db *gorm.DB, comerId uint64, request *model.Pagination) (proposals []ProposalPublicInfo, err error) {
	db = db.Where("governance_proposal.author_comer_id = ? and governance_proposal.is_deleted = false", comerId).
		Select("governance_proposal.*, startup.logo as startup_logo, startup.name as startup_name, comer_profile.avatar as author_comer_avatar, comer_profile.name as author_comer_name").
		Joins("left join startup on startup.id = governance_proposal.startup_id").
		Joins("left join comer_profile on comer_profile.comer_id = governance_proposal.author_comer_id")
	var cnt int64
	if err := db.Table("governance_proposal").Count(&cnt).Error; err != nil {
		return proposals, err
	}
	if cnt == 0 {
		return proposals, nil
	}
	if err := db.
		Offset(request.GetOffset()).
		Limit(request.GetLimit()).
		Order("governance_proposal.created_at desc").
		Table("governance_proposal").Scan(&proposals).Error; err != nil {
		return proposals, err
	}
	request.TotalRows = cnt
	request.TotalPages = int(math.Ceil(float64(cnt) / float64(request.Limit)))
	request.Rows = proposals
	return
}

func SelectProposalListByComerParticipate(db *gorm.DB, comerId uint64, request *model.Pagination) (proposals []ProposalPublicInfo, err error) {
	db = db.
		Where("governance_proposal.is_deleted = false").
		Select("governance_proposal.*, startup.logo as startup_logo, startup.name as startup_name, comer_profile.avatar as author_comer_avatar, comer_profile.name as author_comer_name").
		Joins("left join governance_vote on governance_vote.proposal_id = governance_proposal.id").
		Joins("left join startup on startup.id = governance_proposal.startup_id").
		Joins("left join comer_profile on comer_profile.comer_id = governance_proposal.author_comer_id").
		Where("governance_vote.voter_comer_id = ?", comerId).
		Group("governance_proposal.id")
	var cnt int64
	if err := db.
		Table("governance_proposal").Count(&cnt).Error; err != nil {
		return proposals, err
	}
	if cnt == 0 {
		return proposals, nil
	}
	if err := db.
		Offset(request.GetOffset()).
		Limit(request.GetLimit()).
		Order("governance_proposal.created_at desc").
		Table("governance_proposal").Scan(&proposals).Error; err != nil {
		return proposals, err
	}
	request.TotalRows = cnt
	request.TotalPages = int(math.Ceil(float64(cnt) / float64(request.Limit)))
	request.Rows = proposals
	return
}

func UpdateProposalStatus(db *gorm.DB, proposalId uint64, status ProposalStatus) error {
	return db.Model(&GovernanceProposal{}).Where("id = ?", proposalId).Update("status", status).Error
}
func SelectToBeStartedProposalListWithin1Min(db *gorm.DB) (list []GovernanceProposal, err error) {
	now := time.Now()
	err = db.Model(&GovernanceProposal{}).
		Where("status = ? and (start_time between ? and ? or start_time < ? and end_time > ?)",
			ProposalUpcoming,
			now.Add(-30*time.Second),
			now.Add(30*time.Second),
			now,
			now).
		Find(&list).Error
	return
}

func SelectToEndedProposalListWithin1Min(db *gorm.DB) (list []GovernanceProposal, err error) {
	err = db.Model(&GovernanceProposal{}).
		Where("is_deleted = false and status not in  ? and end_time <= ?",
			[]ProposalStatus{ProposalInvalid, ProposalEnded, ProposalUpcoming, ProposalPending},
			time.Now().Add(time.Second*30)).
		Find(&list).Error
	return
}
