package bounty

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"gorm.io/gorm"

	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/model/tag"
)

// TODO: bounty model

const (
	AccessIn  = 1
	AccessOut = 2

	PaymentModeStage  = 1
	PaymentModePeriod = 2

	BountyPaymentTermsStatusUnpaid = 1
	BountyPaymentTermsStatusPaid   = 2

	BountyStatusReadyToWork = 1
	BountyStatusWordStarted = 2
	BountyStatusCompleted   = 3 // close 后设置为完成状态
	BountyStatusExpired     = 4

	ApplicantStatusPending    = 0 // 发起申请，待批准
	ApplicantStatusApplied    = 1
	ApplicantStatusRefunded   = 2
	ApplicantStatusWithdraw   = 3
	ApplicantStatusRefused    = 4
	ApplicantStatusApproved   = 5 // 批准
	ApplicantStatusUnApproved = 6 // 取消批准，申请者不能再上锁

	FOUNDER   = 1
	APPLICANT = 2
	OTHERS    = 3
)

func CreateBounty(db *gorm.DB, bounty *Bounty) (uint64, error) {
	if err := db.Create(&bounty).Error; err != nil {
		return 0, err
	}
	return bounty.ID, nil
}

func CreateContact(db *gorm.DB, contact *BountyContact) error {
	return db.Create(&contact).Error
}

func CreateDeposit(db *gorm.DB, deposit *BountyDeposit) error {
	return db.Create(&deposit).Error
}

func CreatePaymentTerms(db *gorm.DB, paymentTerm *BountyPaymentTerms) error {
	return db.Create(&paymentTerm).Error
}

func CreatePaymentPeriod(db *gorm.DB, paymentPeriod *BountyPaymentPeriod) error {
	return db.Create(&paymentPeriod).Error
}

func UpdateBountyDepositAmount(db *gorm.DB, bountyID uint64, tokenAmount int) error {
	var sql = fmt.Sprintf("UPDATE bounty SET founder_deposit = founder_deposit + %d WHERE id = %d", tokenAmount, bountyID)
	dbTmp := db.Exec(sql)
	err := dbTmp.Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBountyDepositContract(db *gorm.DB, bountyID uint64, depositContract string) error {
	return db.Model(&Bounty{}).Where("id = ?", bountyID).Update("deposit_contract", depositContract).Error
}

func UpdateBountyDepositStatus(db *gorm.DB, bountyID uint64, status int) error {
	return db.Model(&BountyDeposit{}).Where("bounty_id = ?", bountyID).Update("status", status).Error
}

func UpdateBountyDetailDepositStatus(db *gorm.DB, bountyID, founderComerID uint64, status int) error {
	return db.Model(&BountyDeposit{}).Where("bounty_id = ? and comer_id = ?", bountyID, founderComerID).Update("status", status).Error
}
func GetAndUpdateTagID(db *gorm.DB, name string) (tagID uint64, err error) {
	err = db.Table("tag").Select("id").Where("name = ? and 'category' = 'comerSkill' ", name).Find(&tagID).Error
	if err != nil {
		return 0, err
	}

	if tagID == 0 {
		var isIndex bool
		if len(name) > 2 && name[0:1] == "#" {
			isIndex = true
		}
		skill := tag.Tag{
			Name:     name,
			IsIndex:  isIndex,
			Category: tag.Bounty,
		}
		tag.FirstOrCreateTag(db, &skill)
		return skill.ID, nil
	}
	return tagID, nil
}

func CreateTagTargetRel(db *gorm.DB, tagTargetRel *tag.TagTargetRel) error {
	return db.Create(&tagTargetRel).Error
}

// GetPaymentTermsByBountyId get payment_terms list
func GetPaymentTermsByBountyId(db *gorm.DB, bountyId uint64, termList *[]BountyPaymentTerms) error {
	return db.Model(&BountyPaymentTerms{}).Where("bounty_id = ? ", bountyId).Find(termList).Error
}

func GetPaymentPeriodsByBountyId(db *gorm.DB, bountyId uint64, termList *[]BountyPaymentPeriod) error {
	return db.Model(&BountyPaymentPeriod{}).Where("bounty_id = ? ", bountyId).Find(termList).Error
}

func GetBountyTagNames(db *gorm.DB, bountyId uint64) (tagNames []string, err error) {
	var tagIds []uint64
	if err := db.Model(&tag.TagTargetRel{}).Where("target= ? and target_id = ?", "bounty", bountyId).Select("tag_id").Find(&tagIds).Error; err != nil {
		return nil, err
	}
	if len(tagIds) >= 0 {
		if err := db.Model(&tag.Tag{}).Where("id in ?", tagIds).Select("name").Find(&tagNames).Error; err != nil {
			return nil, err
		}
	}
	return
}

func GetApplicantCountOfBounty(db *gorm.DB, bountyId uint64) (cnt int64, err error) {
	if err := db.Model(&BountyApplicant{}).Where("bounty_id = ?", bountyId).Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func GetApplicantByBountyAndComer(db *gorm.DB, bountyId uint64, comerId uint64) (app BountyApplicant, err error) {
	if err := db.Model(&BountyApplicant{}).Where("bounty_id = ? and comer_id = ?", bountyId).Find(&app).Error; err != nil {
		return BountyApplicant{}, err
	}
	return app, nil
}

func GetBountyDepositByBountyAndComer(db *gorm.DB, bountyID uint64, crtComerId uint64) (bd BountyDeposit, err error) {
	if err := db.Model(&BountyDeposit{}).Where("bounty_id = ? and comer_id = ?", bountyID, crtComerId).Find(&bd).Error; err != nil {
		return BountyDeposit{}, err
	}
	return bd, nil
}
func PageSelectOnChainBounties(db *gorm.DB, pagination *model.Pagination) (bounties []Bounty, er error) {
	var cnt int64
	if err := db.Model(Bounty{}).Count(&cnt).Error; err != nil {
		return bounties, err
	}

	tx := db.Model(Bounty{}).Where("deposit_contract <> ?", "").Order(pagination.Sort).Offset(pagination.GetOffset()).Limit(pagination.Limit)
	if pagination.Keyword != "" {
		tx = tx.Where("title like ?", fmt.Sprintf("%%%v%%", pagination.Keyword))
	}

	if pagination.Mode != 0 {
		tx = tx.Where("status = ?", pagination.Mode)
	}

	if err := tx.Find(&bounties).Error; err != nil {
		return bounties, err
	}

	pagination.Rows = bounties
	pagination.TotalRows = cnt
	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return bounties, nil
}

func PageSelectBountiesByStartupId(db *gorm.DB, startupId uint64, pagination *model.Pagination) (bounties []Bounty, err error) {
	cntSql := fmt.Sprintf("select count(b.id) from bounty b left join transaction bd on b.id=bd.source_id  where bd.status=1 and bd.source_type=1 and b.startup_id=%d", startupId)
	var cnt int64
	if err := db.Raw(cntSql).Scan(&cnt).Error; err != nil {
		return bounties, err
	}
	pageSql := fmt.Sprintf("select b.* from bounty b left join transaction bd on b.id=bd.source_id  where bd.status=1 and bd.source_type=1 and b.startup_id=%d order by %s limit %d,%d", startupId, pagination.Sort, pagination.GetOffset(), pagination.GetLimit())

	if err := db.Raw(pageSql).Scan(&bounties).Error; err != nil {
		return bounties, err
	}
	pagination.Rows = bounties
	pagination.TotalRows = cnt

	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return bounties, nil
}

func PageSelectPostedBounties(db *gorm.DB, comerId uint64, pagination *model.Pagination) (bounties []Bounty, er error) {

	if err := db.Scopes(model.Paginate(&Bounty{}, pagination, db)).Where("comer_id = ?", comerId).Find(&bounties).Error; err != nil {
		return bounties, err
	}
	pagination.Rows = bounties
	return bounties, nil
}

func PageSelectParticipatedBounties(db *gorm.DB, comerId uint64, pagination *model.Pagination) (bounties []Bounty, er error) {
	db = db.Table("bounty").
		// Select("bounty.*").
		Joins("INNER JOIN bounty_applicant ON bounty_applicant.bounty_id = bounty.id").
		Where("bounty_applicant.comer_id = ? and bounty.is_deleted = false", comerId).
		Group("bounty.id")
	var cnt int64
	if err := db.Count(&cnt).Error; err != nil {
		return nil, err
	}

	if err := db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("bounty.created_at desc").Find(&bounties).Error; err != nil {
		return nil, err
	}

	pagination.Rows = bounties
	pagination.TotalRows = cnt
	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return
}

func GetDetailByBountyID(db *gorm.DB, bountyID uint64) (*DetailResponse, error) {
	var detailResponse DetailResponse
	var detailBounty DetailBounty
	var sql = fmt.Sprintf("select title, status, discussion_link, apply_cutoff_date, applicant_deposit, description, deposit_contract, deposit_token_symbol, chain_id, created_at from bounty where id = %d", bountyID)
	err := db.Raw(sql).Scan(&detailBounty).Error
	if err != nil {
		return nil, err
	}
	var tagIds []uint64
	err = db.Table("tag_target_rel").Select("tag_id").Where("target_id = ? and target = ?", bountyID, tag.Bounty).Find(&tagIds).Error
	if err != nil {
		return nil, err
	}
	var skillNames []string
	var skillName string
	for _, tagId := range tagIds {
		db.Table("tag").Select("name").Where("id = ?", tagId).Find(&skillName)
		skillNames = append(skillNames, skillName)
	}
	var contacts []Contact
	db.Table("bounty_contact").Select("contact_type, contact_address").Where("bounty_id = ?", bountyID).Find(&contacts)
	detailResponse.ApplicantSkills = skillNames
	detailResponse.Contacts = contacts
	detailResponse.DetailBounty = detailBounty
	return &detailResponse, nil
}

func GetPaymentByBountyID(db *gorm.DB, bountyID uint64, comerID uint64) (*PaymentResponse, error) {
	var founderComerID uint64
	err := db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&founderComerID).Error
	if err != nil {
		return nil, err
	}
	var bountyPaymentInfo BountyPaymentInfo
	err = db.Table("bounty").Select("payment_mode").Where("id = ? and status != 0", bountyID).Find(&bountyPaymentInfo).Error
	if err != nil {
		return nil, err
	}

	var paymentResponse PaymentResponse
	if bountyPaymentInfo.PaymentMode == PaymentModeStage {
		var StagePayments []StageTerm
		db.Table("bounty_payment_terms").Select("seq_num, status, token1_symbol,token1_amount, token2_symbol, token2_amount, terms").Where("bounty_id = ?", bountyID).Order("seq_num asc").Find(&StagePayments)
		for _, stagePayments := range StagePayments {
			paymentResponse.Rewards.Token1Symbol = stagePayments.Token1Symbol
			paymentResponse.Rewards.Token2Symbol = stagePayments.Token2Symbol
			paymentResponse.Rewards.Token1Amount += stagePayments.Token1Amount
			paymentResponse.Rewards.Token2Amount += stagePayments.Token2Amount
		}
		paymentResponse.BountyPaymentInfo = bountyPaymentInfo
		paymentResponse.StageTerms = StagePayments
		return &paymentResponse, nil
	}
	if bountyPaymentInfo.PaymentMode == PaymentModePeriod {
		var periodModes []PeriodMode
		db.Table("bounty_payment_terms").Select("seq_num, status, token1_symbol,token1_amount, token2_symbol, token2_amount").Where("bounty_id = ?", bountyID).Order("seq_num asc").Find(&periodModes)
		// var terms string
		// db.Table("bounty_payment_period").Select("target").Where("bounty_id = ?", bountyID).Order("seq_num asc").Last(&terms)
		for _, periodMode := range periodModes {
			paymentResponse.Rewards.Token1Symbol = periodMode.Token1Symbol
			paymentResponse.Rewards.Token2Symbol = periodMode.Token2Symbol
			paymentResponse.Rewards.Token1Amount += periodMode.Token1Amount
			paymentResponse.Rewards.Token2Amount += periodMode.Token2Amount
		}
		paymentResponse.BountyPaymentInfo = bountyPaymentInfo
		paymentResponse.PeriodTerms = new(PeriodTerms)
		paymentResponse.PeriodTerms.PeriodModes = periodModes

		var periodInfo PeriodInfo
		db.Table("bounty_payment_period").Select("hours_per_day, period_type, target").Where("bounty_id = ?", bountyID).Find(&periodInfo)
		paymentResponse.HoursPerDay = periodInfo.HoursPerDay
		paymentResponse.PeriodType = periodInfo.PeriodType
		paymentResponse.PeriodTerms.Terms = periodInfo.Target

		return &paymentResponse, nil
	}
	return nil, nil
}

type GetBountyStateResponse struct {
	BountyStatus              int8   `json:"bountyStatus"  gorm:"column:status"`
	ApplicantCount            int64  `json:"applicantCount"`
	DepositBalance            int64  `json:"depositBalance"`
	DepositTokenSymbol        string `json:"depositTokenSymbol"`
	FounderDepositAmount      int64  `json:"founderDepositAmount" gorm:"column:founder_deposit"`
	ApplicantDepositAmount    int64  `json:"applicantDepositAmount" gorm:"column:applicant_deposit"`
	ApplicantDepositMinAmount int64  `json:"applicantDepositMinAmount"`
	ApprovedStatus            int8   `json:"approvedStatus"`
	DepositLock               bool   `json:"depositLock"`
	TimeLock                  int64  `json:"timeLock"`
	MyRole                    int8   `json:"myRole"`
	MyDepositAmount           int64  `json:"myDepositAmount"`
	MyStatus                  int8   `json:"myStatus"`
}

func GetBountyState(db *gorm.DB, bountyID uint64, comerID uint64) (*GetBountyStateResponse, error) {
	var founderComerID uint64
	err := db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&founderComerID).Error
	if err != nil {
		return nil, err
	}

	var stateResponse GetBountyStateResponse

	err = db.Raw("select founder_deposit, deposit_token_symbol, applicant_deposit,status from bounty where id = ?", bountyID).Row().Scan(&stateResponse.FounderDepositAmount, &stateResponse.DepositTokenSymbol, &stateResponse.ApplicantDepositMinAmount, &stateResponse.BountyStatus)
	if err != nil {
		return nil, err
	}

	var (
		accessInResults, accessOutResults                               []int64
		applicantsTotalAccessInDeposit, applicantsTotalAccessOutDeposit int64
	)

	db.Table("bounty_deposit").Select("token_amount").Where("bounty_id = ? and comer_id != ? and access = 1", bountyID, founderComerID).Find(&accessInResults)
	for _, result := range accessInResults {
		applicantsTotalAccessInDeposit += result
	}
	db.Table("bounty_deposit").Select("token_amount").Where("bounty_id = ? and comer_id != ? and access = 2", bountyID, founderComerID).Find(&accessOutResults)
	for _, result := range accessOutResults {
		applicantsTotalAccessOutDeposit += result
	}
	var bountyDepositStatus uint64
	db.Table("bounty_deposit").Select("status").Where("bounty_id = ? and comer_id = ?", bountyID, comerID).Find(&bountyDepositStatus)
	var approvedStatus uint64
	db.Table("bounty_applicant").Select("status").Where("bounty_id = ? and status in ?", bountyID, [2]int{ApplicantStatusApproved, ApplicantStatusUnApproved}).Find(&approvedStatus)

	if bountyDepositStatus == 3 {
		stateResponse.DepositLock = true
	}

	stateResponse.ApplicantDepositAmount = applicantsTotalAccessInDeposit - applicantsTotalAccessOutDeposit
	if stateResponse.ApplicantDepositAmount < 0 {
		return nil, errors.New("applicant deposit less then 0")
	}
	// db.Table("bounty_deposit").Select("status").Where("bounty_id = ? and comer_id = ?", bountyID, comerID).Order("id desc").Find(&stateResponse.BountyStatus)

	var (
		applicantApplyStatus int64
	)
	if founderComerID == comerID {
		stateResponse.MyRole = FOUNDER
	} else {
		err := db.Raw("select status from bounty_applicant where bounty_id = ? and comer_id = ?", bountyID, comerID).Row().Scan(&applicantApplyStatus)
		if err == sql.ErrNoRows {
			stateResponse.MyRole = OTHERS
		} else {
			stateResponse.MyRole = APPLICANT
		}
	}
	stateResponse.MyStatus = int8(applicantApplyStatus)
	stateResponse.ApprovedStatus = int8(approvedStatus)
	// 申请者数量
	stateResponse.ApplicantCount, _ = GetApplicantCountOfBounty(db, bountyID)
	stateResponse.DepositBalance = stateResponse.ApplicantDepositAmount + stateResponse.FounderDepositAmount
	return &stateResponse, nil
}

func UpdateBountyCloseStatusByID(db *gorm.DB, bountyID uint64) (string, error) {
	var tokenAmount int
	err := db.Table("bounty").Select("founder_deposit").Where("id = ?", bountyID).Find(&tokenAmount).Error
	if err != nil {
		return "", err
	}
	if tokenAmount != 0 {
		// return "the deposit is not null", nil
	}
	err = db.Table("bounty").Where("id = ?", bountyID).Update("status", BountyStatusCompleted).Error
	if err != nil {
		return "", err
	}
	return "close bounty", nil
}

func UpdatePaidByBountyID(db *gorm.DB, bountyID uint64, request *PaidRequest) error {

	paidInfo := ""
	if len(request.PaidInfo) != 0 {
		b, _ := json.Marshal(request.PaidInfo)
		paidInfo = string(b)
	}
	return db.Table("bounty_payment_terms").Where("bounty_id = ? and seq_num = ?", bountyID, request.SeqNum).Updates(map[string]interface{}{"status": BountyPaymentTermsStatusPaid, "paid_info": string(paidInfo)}).Error
}

func CreateApplicants(db *gorm.DB, request *BountyApplicantForBounty) error {
	var count int64
	err := db.Table("bounty_applicant").Where("bounty_id = ? and comer_id = ?", request.BountyID, request.ComerID).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return db.Table("bounty_applicant").Where("bounty_id = ? and comer_id = ?", request.BountyID, request.ComerID).Updates(request).Error
	} else {
		return db.Create(&request).Error
	}
}

func GetActivitiesByBountyID(db *gorm.DB, bountyID uint64) (*[]ActivitiesResponse, error) {
	var sql = fmt.Sprintf("SELECT content, source_type, timestamp, cp.name, cp.avatar, pu.comer_id from post_update pu left join comer_profile cp on cp.comer_id = pu.comer_id where pu.source_id = %d", bountyID)
	var activitiesTotal []ActivitiesResponse
	err := db.Raw(sql).Find(&activitiesTotal).Error
	if err != nil {
		return nil, err
	}
	return &activitiesTotal, nil
}

func GetApplicants(db *gorm.DB, bountyID uint64) (*[]Applicant, error) {
	var applicantComerIDs []uint64
	err := db.Table("bounty_applicant").Select("distinct comer_id").Where("bounty_id = ?", bountyID).Find(&applicantComerIDs).Error
	if err != nil {
		return nil, err
	}
	var comerInfo ComerInfo
	var bountyApplicant BountyApplicant
	var applicant Applicant
	var comer account.Comer
	var applicants []Applicant
	for _, applicantComerID := range applicantComerIDs {
		db.Table("comer_profile").Select("comer_id, name, avatar").Where("comer_id = ?", applicantComerID).Find(&comerInfo)
		db.Table("bounty_applicant").Select("description, status, apply_at").Where("comer_id = ? and bounty_id = ?", applicantComerID, bountyID).Last(&bountyApplicant)
		db.Table("comer").Select("address").Where("id = ?", applicantComerID).Last(&comer)

		applicant.Name = comerInfo.Name
		applicant.Image = comerInfo.Avatar
		applicant.Description = bountyApplicant.Description
		applicant.Applyat = bountyApplicant.ApplyAt
		applicant.Status = bountyApplicant.Status
		if comer.Address != nil {
			applicant.Address = *comer.Address
		}

		applicant.ComerID = strconv.Itoa(int(applicantComerID))
		applicants = append(applicants, applicant)
	}
	return &applicants, nil
}

func GetFounderByBountyID(db *gorm.DB, bountyID uint64) (*FounderResponse, error) {
	var comerID uint64
	var comerInfo account.ComerProfile
	var tagIds []uint64
	var skillNames []string
	var skillName string
	var founderInfo FounderResponse
	db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&comerID)

	db.Table("comer_profile").Select("name, avatar, time_zone, email, location").Where("comer_id = ?", comerID).Find(&comerInfo)
	db.Table("tag_target_rel").Select("tag_id").Where("target = ? and target_id = ?", tag.ComerSkill, comerID).Find(&tagIds)
	for _, tagId := range tagIds {
		db.Table("tag").Select("name").Where("id = ?", tagId).Find(&skillName)
		skillNames = append(skillNames, skillName)
	}
	founderInfo.ComerID = comerID
	founderInfo.Name = comerInfo.Name
	founderInfo.Image = comerInfo.Avatar
	founderInfo.TimeZone = comerInfo.TimeZone
	founderInfo.ApplicantsSkills = skillNames
	founderInfo.Location = comerInfo.Location
	founderInfo.Email = comerInfo.Email
	return &founderInfo, nil
}

func GetApprovedApplicantByBountyID(db *gorm.DB, bountyID uint64) (*ApprovedResponse, error) {
	var comerID uint64
	var bountyStauts uint64
	db.Table("bounty").Select("status").Where("id = ?", bountyID).Find(&bountyStauts)
	if bountyStauts == BountyStatusReadyToWork {
		return nil, nil
	}
	db.Table("bounty_applicant").Select("comer_id").Where("bounty_id = ? and status = ?", bountyID, ApplicantStatusApproved).Find(&comerID)

	var comerInfo account.ComerProfile
	var tagIds []uint64
	var skillNames []string
	var skillName string
	var approvedInfo ApprovedResponse
	var Address string
	db.Table("comer_profile").Select("name, avatar, time_zone").Where("comer_id = ?", comerID).Find(&comerInfo)
	db.Table("tag_target_rel").Select("tag_id").Where("target_id = ? and target = ?", comerID, tag.ComerSkill).Find(&tagIds)
	for _, tagId := range tagIds {
		db.Table("tag").Select("name").Where("id = ?", tagId).Find(&skillName)
		skillNames = append(skillNames, skillName)
	}
	db.Table("comer").Select("address").Where("id = ?", comerID).Find(&Address)
	approvedInfo.ComerID = comerID
	approvedInfo.Name = comerInfo.Name
	approvedInfo.Image = comerInfo.Avatar
	approvedInfo.Address = Address
	approvedInfo.ApplicantsSkills = skillNames
	return &approvedInfo, nil
}

func GetDepositRecordsByBountyID(db *gorm.DB, bountyID uint64) (*[]DepositRecord, error) {
	var depositRecords []DepositRecord
	var sql = fmt.Sprintf("SELECT token_amount, access, bd.created_at, name, avatar, cp.comer_id FROM `bounty_deposit` bd LEFT JOIN comer_profile cp on bd.comer_id = cp.comer_id WHERE bounty_id = %d", bountyID)
	err := db.Raw(sql).Find(&depositRecords).Error
	if err != nil {
		return nil, err
	}

	return &depositRecords, nil
}

func UpdateApplicantStatus(db *gorm.DB, bountyID, comerID uint64, status int) (err error) {
	// fmt.Println("UpdateApplicantStatus > ", status)
	return db.Table("bounty_applicant").Where("bounty_id = ? and comer_id = ? and status != ?", bountyID, comerID, ApplicantStatusApproved).Updates(BountyApplicant{Status: status}).Error
}

func UpdateApplicantApprovedStatus(db *gorm.DB, bountyID, comerID uint64, status int) (err error) {
	err = db.Table("bounty_applicant").Where("bounty_id = ? and comer_id = ?", bountyID, comerID).Updates(BountyApplicant{Status: status, ApproveAt: time.Now()}).Error
	if err != nil {
		return err
	}
	// update other comer reject
	err = db.Table("bounty_applicant").Where("bounty_id = ? and comer_id != ? and status in (?, ?)", bountyID, comerID, ApplicantStatusPending, ApplicantStatusApplied).Updates(BountyApplicant{Status: ApplicantStatusRefused}).Error
	return err
}

type RejectedDepositsApplicants struct {
	BountyId    uint64
	ComerID     uint64
	TokenSymbol string
	TokenAmount int
}

func GetApplicantsRejectedDeposits(db *gorm.DB, bountyID, comerID uint64) (rsp []*RejectedDepositsApplicants, err error) {
	sqlStr := `
		SELECT bounty_id, comer_id, token_symbol
			, SUM(IF(access = 1, token_amount, 0)) - SUM(IF(access = 2, token_amount, 0)) AS deposit
		FROM bounty_deposit
		WHERE bounty_id = ? and comer_id != ?
		GROUP BY bounty_id, comer_id, token_symbol
		HAVING deposit > 0
	`
	rows, _ := db.Raw(sqlStr, bountyID, comerID).Rows()

	for rows.Next() {
		r := &RejectedDepositsApplicants{}
		err = rows.Scan(&r.BountyId, &r.ComerID, &r.TokenSymbol, &r.TokenAmount)
		if err != nil {
			return nil, err
		}
		rsp = append(rsp, r)
	}

	err = rows.Close()
	return
}

func GetApplicantsReleaseDeposits(db *gorm.DB, bountyID, comerID uint64) (rsp *RejectedDepositsApplicants, err error) {
	rsp = &RejectedDepositsApplicants{}

	sqlStr := `
		SELECT bounty_id, comer_id, token_symbol
			, SUM(IF(access = 1, token_amount, 0)) - SUM(IF(access = 2, token_amount, 0)) AS deposit
		FROM bounty_deposit
		WHERE bounty_id = ? and comer_id = ?
		GROUP BY bounty_id, comer_id, token_symbol
		HAVING deposit >= 0
	`
	err = db.Raw(sqlStr, bountyID, comerID).Row().Scan(&rsp.BountyId, &rsp.ComerID, &rsp.TokenSymbol, &rsp.TokenAmount)
	return
}

func GetAllApplicantsReleaseDeposits(db *gorm.DB, bountyID uint64) (rsp []*RejectedDepositsApplicants, err error) {

	sqlStr := `
		SELECT bounty_id, comer_id, token_symbol
			, SUM(IF(access = 1, token_amount, 0)) - SUM(IF(access = 2, token_amount, 0)) AS deposit
		FROM bounty_deposit
		WHERE bounty_id = ?
		GROUP BY bounty_id, comer_id, token_symbol
		HAVING deposit > 0
	`
	rows, _ := db.Raw(sqlStr, bountyID).Rows()

	for rows.Next() {
		r := &RejectedDepositsApplicants{}
		err = rows.Scan(&r.BountyId, &r.ComerID, &r.TokenSymbol, &r.TokenAmount)
		if err != nil {
			return nil, err
		}
		rsp = append(rsp, r)
	}

	err = rows.Close()
	return
}

func GetStartupByBountyID(db *gorm.DB, bountyID uint64) (*StartupListResponse, error) {
	var startupID uint64
	err := db.Table("bounty").Select("startup_id").Where("id = ?", bountyID).Find(&startupID).Error
	if err != nil {
		return nil, err
	}
	var startupListResponse StartupListResponse
	err = db.Table("startup").Select("id, name, mode, logo, chain_id, tx_hash, contract_audit, website, discord, twitter, telegram, docs, mission").Where("id = ?", startupID).First(&startupListResponse).Error
	if err != nil {
		return nil, err
	}

	var tagIDs []uint64
	err = db.Table("tag_target_rel").Select("tag_id").Where("target_id = ? and target = ?", startupID, tag.Startup).Find(&tagIDs).Error
	if err != nil {
		return nil, err
	}
	var tagName string
	var tagNames []string
	for _, tagID := range tagIDs {
		err = db.Table("tag").Select("name").Where("id = ?", tagID).Find(&tagName).Error
		if err != nil {
			return nil, err
		}
		tagNames = append(tagNames, tagName)
	}
	startupListResponse.Tag = tagNames
	return &startupListResponse, nil
}

func GetBountyRoleByComerID(db *gorm.DB, bountyID, comerID uint64) int {
	var comerIDtmp uint64
	db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&comerIDtmp)
	if comerIDtmp == comerID {
		return FOUNDER
	} else {
		var applicantStatus int
		db.Table("bounty_applicant").Select("status").Where("comer_id = ?").Find(&applicantStatus)
		if applicantStatus == 2 {
			return APPLICANT
		} else {
			return OTHERS
		}
	}
}

func UpdateBountyStatus(db *gorm.DB, bountyID uint64, status int) error {
	return db.Table("bounty").Where("id = ?", bountyID).Update("status", status).Error
}

func UpdateApplicantRejectStatus(db *gorm.DB, bountyID, applicantComerID uint64) error {
	return db.Table("bounty_applicant").Where("bounty_id = ? and comer_id = ?", bountyID, applicantComerID).Update("status", ApplicantStatusRefused).Error
}

func UpdateApplicantApproveStatus(db *gorm.DB, comerID, bountyID uint64) error {
	var approvedApplicantComerIDs []uint64
	err := db.Table("bounty_deposit").Select("comer_id").Where("bounty_id = ? and comer_id != ?", bountyID, comerID).Find(&approvedApplicantComerIDs).Error
	if err != nil {
		return err
	}
	for _, approvedApplicantComerID := range approvedApplicantComerIDs {
		db.Table("bounty_applicant").Where("comer_id = ?", approvedApplicantComerID).Update("status", ApplicantStatusApplied)
		db.Table("bounty_deposit").Where("comer_id = ?", approvedApplicantComerID).Update("access", AccessIn)
	}
	return nil
}

func GetApplicantLockStatus(db *gorm.DB, bountyID, comerID uint64) (int, error) {
	var applicantStatus int
	err := db.Table("bounty_deposit").Select("status").Where("bounty_id = ? and comer_id = ?", bountyID, comerID).Find(&applicantStatus).Error
	if err != nil {
		return 0, err
	}
	return applicantStatus, nil
}

func UpdateApplicantDepositLockStatus(db *gorm.DB, bountyID, comerID uint64, depositStatus int) error {
	var founderComerID uint64
	err := db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&founderComerID).Error
	if err != nil {
		return err
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		err := db.Table("bounty_deposit").Where("bounty_id = ? and comer_id = ?", bountyID, comerID).Update("status", depositStatus).Error
		if err != nil {
			return err
		}
		err = db.Table("bounty_deposit").Where("bounty_id = ? and comer_id = ?", bountyID, founderComerID).Update("status", depositStatus).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func GetComerIDByBountyID(db *gorm.DB, bountyID uint64) (founderComerID uint64, err error) {
	err = db.Table("bounty").Select("comer_id").Where("id = ?", bountyID).Find(&founderComerID).Error
	if err != nil {
		return 0, nil
	}
	return
}

func DecrApplicantDepositByBountyID(db *gorm.DB, bountyID uint64, num int) error {
	return db.Table("bounty").Where("id = ?", bountyID).Update("applicant_deposit", gorm.Expr("applicant_deposit - ?", num)).Error
}

func IncrApplicantDepositByBountyID(db *gorm.DB, bountyID uint64, num int) error {
	return db.Table("bounty").Where("id = ?", bountyID).Update("applicant_deposit", gorm.Expr("applicant_deposit + ?", num)).Error
}

func UpdateApplicantDepositByBountyID(db *gorm.DB, bountyID uint64) error {
	return db.Table("bounty").Where("id = ?", bountyID).Update("applicant_deposit", 0).Error
}

func UpdateFounderDepositByBountyID(db *gorm.DB, bountyID uint64) error {
	return db.Table("bounty").Where("id = ?", bountyID).Update("founder_deposit", 0).Error
}

func UpdateApplicantRevokeTimeByComerID(db *gorm.DB, applicantComerID uint64) error {
	return db.Table("bounty_applicant").Where("comer_id = ?", applicantComerID).Updates(BountyApplicant{Status: ApplicantStatusWithdraw, RevokeAt: time.Now()}).Error
}

func CountBountiesPostedByComer(db *gorm.DB, targetComerID uint64) (cnt int64, er error) {
	er = db.Model(&Bounty{}).Where("is_deleted = false and comer_id = ?", targetComerID).Count(&cnt).Error
	return
}
func UpdateDepositStatus(db *gorm.DB, txHash string, status int) error {
	return db.Table("bounty_deposit").Where("tx_hash = ?", txHash).Update("status", status).Error
}
