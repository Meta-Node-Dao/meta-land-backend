/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/29 10:04
 */

package bounty

import (
	"ceres/pkg/initialization/mysql"
	model2 "ceres/pkg/model"
	model "ceres/pkg/model/bounty"
	model3 "ceres/pkg/model/postupdate"
	"ceres/pkg/model/startup"
	"ceres/pkg/model/tag"
	model4 "ceres/pkg/model/transaction"
	"ceres/pkg/service/transaction"
	"ceres/pkg/utility/tool"
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

// CreateComerBounty create bounty
func CreateComerBounty(request *model.BountyRequest) error {

	err := mysql.DB.Transaction(func(tx *gorm.DB) (ere error) {
		paymentMode, totalRewardToken := handlePayDetail(request.PayDetail)

		bountyID, err := createBounty(tx, paymentMode, totalRewardToken, request)
		if err != nil {
			log.Warn(err)
			return err
		}
		if bountyID == 0 {
			return errors.New(fmt.Sprintf("create bounty err: %d", bountyID))
		}

		err = transaction.CreateTransaction(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return err
		}

		// 创建bounty不默认创建一个post update
		// err = postupdate.CreatePostUpdate(tx, bountyID, request)
		// if err != nil {
		// 	log.Warn(err)
		// 	return err
		// }

		err = createDeposit(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return err
		}

		errorsLog := createPaymentTerms(tx, bountyID, request)
		if len(errorsLog) > 0 {
			return errors.New(fmt.Sprintf("create payment_terms err:%v", errorsLog))
		}

		err = creatPaymentPeriod(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return err
		}

		errorsLog = createContact(tx, bountyID, request)
		if len(errorsLog) > 0 {
			return errors.New(fmt.Sprintf("create contact address err:%v", errorsLog))
		}

		err = createApplicantsSkills(tx, bountyID, request)
		if err != nil {
			log.Warn(err)
			return err
		}

		getContract(request.ChainID, request.TxHash, bountyID)

		return nil
	})

	return err
}

func createBounty(tx *gorm.DB, paymentMode, totalRewardToken int, request *model.BountyRequest) (uint64, error) {
	bounty := &model.Bounty{
		StartupID:          request.StartupID,
		ComerID:            request.ComerID,
		ChainID:            request.ChainID,
		TxHash:             request.TxHash,
		Title:              request.Title,
		ApplyCutoffDate:    tool.ParseTimeString2Time(request.ExpiresIn),
		DiscussionLink:     request.BountyDetail.DiscussionLink,
		DepositTokenSymbol: request.Deposit.TokenSymbol,
		ApplicantDeposit:   request.ApplicantsDeposit,
		FounderDeposit:     request.Deposit.TokenAmount,
		Description:        request.Description,
		PaymentMode:        paymentMode,
		Status:             model.BountyStatusReadyToWork,
		TotalRewardToken:   totalRewardToken,
	}

	bountyID, err := model.CreateBounty(tx, bounty)
	if err != nil || bountyID == 0 {
		return 0, errors.New(fmt.Sprintf("created bounty err:%s", err))
	}
	return bountyID, nil
}

func createDeposit(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	deposit := &model.BountyDeposit{
		ChainID:     request.ChainID,
		TxHash:      request.TxHash,
		Status:      transaction.Pending,
		BountyID:    bountyID,
		ComerID:     request.ComerID,
		Access:      model.AccessIn,
		TokenSymbol: request.Deposit.TokenSymbol,
		TokenAmount: request.Deposit.TokenAmount,
		Timestamp:   time.Now(),
	}
	err := model.CreateDeposit(tx, deposit)
	if err != nil {
		return err
	}
	return nil
}

func createContact(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) []string {
	var errorLog []string
	for _, contact := range request.Contacts {
		contactModel := &model.BountyContact{
			BountyID:       bountyID,
			ContactType:    contact.ContactType,
			ContactAddress: contact.ContactAddress,
		}
		err := model.CreateContact(tx, contactModel)
		if err != nil {
			errorLog = append(errorLog, fmt.Sprintf("create contactAddress:%s err:%v", contact.ContactAddress, err))
			continue
		}
	}
	return errorLog
}

func createPaymentTerms(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) []string {
	paymentMode, _ := handlePayDetail(request.PayDetail)
	var errorLog []string
	if paymentMode == model.PaymentModeStage {
		for _, stage := range request.PayDetail.Stages {
			paymentTerms := &model.BountyPaymentTerms{
				BountyID:     bountyID,
				PaymentMode:  int8(paymentMode),
				Token1Symbol: stage.Token1Symbol,
				Token1Amount: stage.Token1Amount,
				Token2Symbol: stage.Token2Symbol,
				Token2Amount: stage.Token2Amount,
				Terms:        stage.Terms,
				SeqNum:       stage.SeqNum,
				Status:       model.BountyPaymentTermsStatusUnpaid,
			}
			err := model.CreatePaymentTerms(tx, paymentTerms)
			if err != nil {
				errorLog = append(errorLog, fmt.Sprintf("create stage %v err:%v", stage, err))
				continue
			}
		}
	} else {
		for i := 1; i <= request.Period.PeriodAmount; i++ {
			paymentTerms := &model.BountyPaymentTerms{
				BountyID:     bountyID,
				PaymentMode:  int8(paymentMode),
				Token1Symbol: request.Period.Token1Symbol,
				Token1Amount: request.Period.Token1Amount,
				Token2Symbol: request.Period.Token2Symbol,
				Token2Amount: request.Period.Token2Amount,
				Terms:        request.Period.Target,
				SeqNum:       i,
				Status:       model.BountyPaymentTermsStatusUnpaid,
			}
			err := model.CreatePaymentTerms(tx, paymentTerms)
			if err != nil {
				errorLog = append(errorLog, fmt.Sprintf("create period err:%v", err))
				return errorLog
			}
		}
	}

	return errorLog
}

func creatPaymentPeriod(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	paymentMode, _ := handlePayDetail(request.PayDetail)
	if paymentMode == model.PaymentModePeriod {
		periodAmount := int64(request.Period.Token1Amount + request.Period.Token2Amount)
		paymentPeriod := &model.BountyPaymentPeriod{
			BountyID:     bountyID,
			PeriodType:   request.Period.PeriodType,
			PeriodAmount: periodAmount,
			HoursPerDay:  request.Period.HoursPerDay,
			Token1Symbol: request.Period.Token1Symbol,
			Token1Amount: request.Period.Token1Amount,
			Token2Symbol: request.Period.Token2Symbol,
			Token2Amount: request.Period.Token2Amount,
			Target:       request.Period.Target,
		}
		err := model.CreatePaymentPeriod(tx, paymentPeriod)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func createApplicantsSkills(tx *gorm.DB, bountyID uint64, request *model.BountyRequest) error {
	for _, applicantsSkill := range request.ApplicantsSkills {
		tagID, err := model.GetAndUpdateTagID(tx, applicantsSkill)
		if err != nil {
			return err
		}
		tarTargetRel := &tag.TagTargetRel{
			TargetID: bountyID,
			TagID:    tagID,
			Target:   tag.Bounty,
		}
		err = model.CreateTagTargetRel(tx, tarTargetRel)
		if err != nil {
			return err
		}
	}
	return nil
}

func getContract(chainID uint64, txHash string, bountyID uint64) {
	var contractChan = make(chan *model.ContractInfoResponse, 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				s := string(debug.Stack())
				log.Error("recover: err=%v\n stack=%s", err, s)
			}
		}()

		// 重试100次
		retryCount := 100
		for retryCount > 0 {
			contractAddress, status := transaction.GetContractAddress(chainID, txHash)
			contractInfo := &model.ContractInfoResponse{
				ContractAddress: contractAddress,
				Status:          status,
			}
			// 网络失败,15s后重试
			if status == transaction.Failure {
				time.Sleep(15 * time.Second)
				retryCount--
				continue
			}

			select {
			case contractChan <- contractInfo:
				for contract := range contractChan {
					if contract.Status != transaction.Success {
						transaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, bountyID, transaction.Pending, contract.ContractAddress)
					} else {
						transaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, bountyID, int(contract.Status), contract.ContractAddress)
					}
				}
			case <-time.After(5 * time.Second):
				fmt.Println("get contract address time over!")
			}

			return
		}
	}()
	return
}

// unused
// func updateDepositStatusChain(chainID uint64, txHash string, bountyID uint64) {
// 	if len(txHash) == 0 {
// 		return
// 	}

// 	go func() {
// 		for {
// 			status := transaction.GetTxHashStatus(chainID, txHash)

// 			if status == transaction.Pending {
// 				time.Sleep(15 * time.Second)
// 				continue
// 			}

// 			model.UpdateDepositStatus(mysql.DB, txHash, status)
// 			break
// 		}
// 	}()
// }

func handlePayDetail(request model.PayDetail) (paymentMode, totalRewardToken int) {
	if len(request.Stages) > 0 {
		paymentMode = model.PaymentModeStage
		for _, stage := range request.Stages {
			totalRewardToken = stage.Token1Amount + stage.Token2Amount
		}
		return paymentMode, totalRewardToken
	} else {
		paymentMode = model.PaymentModePeriod
		totalRewardToken = request.Period.Token1Amount + request.Period.Token2Amount
		return paymentMode, totalRewardToken
	}
}

// QueryAllOnChainBounties query all bounties, display in bounty tab
func QueryAllOnChainBounties(request *model2.Pagination) (err error) {
	bounties, err := model.PageSelectOnChainBounties(mysql.DB, request)
	if err != nil {
		return err
	}

	if len(bounties) > 0 {
		if items, err := iter(bounties, 0); err != nil {
			return err
		} else {
			request.Rows = items
		}
	}

	return nil
}

type ItemType int

const (
	tabBounty = iota + 1
	startupBounty
	myPostedBounty
	myParticipatedBounty
)

func QueryBountiesByStartup(startupId uint64, pagination *model2.Pagination) (err error) {
	// todo check startup exist!
	// 按照发布顺序降序查询
	pagination.Sort = "created_at desc"
	bounties, err := model.PageSelectBountiesByStartupId(mysql.DB, startupId, pagination)
	if err != nil {
		return err
	}

	if len(bounties) > 0 {
		if items, err := iter(bounties, 0); err != nil {
			return err
		} else {
			pagination.Rows = items
		}
	}

	return nil
}

func QueryComerPostedBountyList(comerId uint64, pagination *model2.Pagination) (err error) {
	bounties, err := model.PageSelectPostedBounties(mysql.DB, comerId, pagination)
	if err != nil {
		return err
	}

	if len(bounties) > 0 {
		if items, err := iter(bounties, comerId); err != nil {
			return err
		} else {
			pagination.Rows = items
		}
	}

	return nil
}

func QueryComerParticipatedBountyList(comerId uint64, pagination *model2.Pagination) (err error) {
	bounties, err := model.PageSelectParticipatedBounties(mysql.DB, comerId, pagination)
	if err != nil {
		return err
	}
	if len(bounties) > 0 {
		if items, err := iter(bounties, comerId); err != nil {
			return err
		} else {
			pagination.Rows = items
		}
	}
	return nil
}

func iter(bounties []model.Bounty, crtComerId uint64) (items []*model.DetailItem, err error) {
	startupMap := make(map[uint64]startup.Startup)
	// 遍历
	for _, bounty := range bounties {
		item, err := packItem(bounty, &startupMap, tabBounty, crtComerId)
		if err != nil {
			log.Warnf("##### iteration occurred error: %v\n", err)
			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}

var bountyStatusMap = map[int8]string{
	1: "Ready to work",
	2: "Started working",
	3: "Completed",
	4: "Expired",
}

var bountyDepositStatusMap = map[int8]string{
	0: "Pending",
	1: "Success",
	2: "Failure",
}

func packItem(bounty model.Bounty, startupMap *map[uint64]startup.Startup, itemType ItemType, crtComerId uint64) (detailItem *model.DetailItem, err error) {
	detailItem = &model.DetailItem{}
	var st startup.Startup
	// 取出 logo
	if su, ok := (*startupMap)[bounty.StartupID]; ok {
		st = su
	} else {
		// 查询startup表，放入map
		if err := startup.GetStartup(mysql.DB, bounty.StartupID, &st); err != nil {
			return detailItem, err
		}
		(*startupMap)[bounty.StartupID] = st
	}
	logo := st.Logo
	detailItem.Logo = logo
	detailItem.Title = bounty.Title
	detailItem.StartupId = st.ID
	detailItem.ChainID = st.ChainID
	detailItem.BountyId = bounty.ID
	detailItem.CreatedTime = bounty.CreatedAt
	detailItem.ApplyCutoffDate = bounty.ApplyCutoffDate
	detailItem.DepositTokenSymbol = bounty.DepositTokenSymbol
	// paymentMode用以计算 rewards
	paymentMode := bounty.PaymentMode
	var rewards []model.Reward
	// stage , 查询paymentTerms并统计
	if paymentMode == 1 {
		var terms []model.BountyPaymentTerms
		if err := model.GetPaymentTermsByBountyId(mysql.DB, bounty.ID, &terms); err != nil {
			return nil, err
		}
		calcRewardWhenIsPaymentTerms(terms, &rewards)
		detailItem.PaymentType = "Stage"
	} else if paymentMode == 2 {
		var periods []model.BountyPaymentPeriod
		// period, 查询PaymentPeriod
		if err := model.GetPaymentPeriodsByBountyId(mysql.DB, bounty.ID, &periods); err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}
		calcRewardWhenIsPaymentPeriod(periods, &rewards)
		detailItem.PaymentType = "Period"
	}

	detailItem.Rewards = &rewards
	// 申请者deposit要求, 由bounty_id去tag_target_rel表查询
	requirementSkills, err := model.GetBountyTagNames(mysql.DB, bounty.ID)
	if err != nil {
		return nil, err
	}
	detailItem.ApplicationSkills = requirementSkills
	// 申请人数，统计bounty_applicant
	applicantCount, err := model.GetApplicantCountOfBounty(mysql.DB, bounty.ID)
	if err != nil {
		return nil, err
	}
	detailItem.ApplicantCount = int(applicantCount)
	var status = bountyStatusMap[bounty.Status]
	var onChainStatus string
	// bounty状态，bounty tab和startup bounty中是一致的；my posted和my participated中状态不一致
	if itemType == tabBounty || itemType == startupBounty {
		onChainStatus = bountyDepositStatusMap[1]
	} else if itemType == myPostedBounty {
		bountyDeposit, err := model.GetBountyDepositByBountyAndComer(mysql.DB, bounty.ID, crtComerId)
		if err != nil {
			return nil, err
		}
		onChainStatus = bountyDepositStatusMap[bountyDeposit.Status]
	} else if itemType == myParticipatedBounty {
		// bountyApplicant, err := model.GetApplicantByBountyAndComer(mysql.DB, bounty.ID, crtComerId)
		bountyDeposit, err := model.GetBountyDepositByBountyAndComer(mysql.DB, bounty.ID, crtComerId)
		if err != nil {
			return nil, err
		}
		log.Infof("#### my bounty deposit: %v\n", bountyDeposit)
		onChainStatus = bountyDepositStatusMap[bountyDeposit.Status]
		// switch bountyApplicant.Status {
		// case 1:
		//	// 已申请
		//	status = "Applied"
		// case 2:
		//	// 通过申请
		//	status = "Approved"
		// case 3:
		//	//
		//	status = "Submitted"
		// case 4:
		//	status = "Revoked"
		// case 5:
		//	status = "Rejected"
		// case 6:
		//	status = "Quited"
		// }
	}
	detailItem.DepositRequirements = bounty.ApplicantDeposit
	detailItem.Status = status
	detailItem.OnChainStatus = onChainStatus
	return detailItem, nil
}

func calcRewardWhenIsPaymentTerms(terms []model.BountyPaymentTerms, rewards *[]model.Reward) {
	if len(terms) > 0 {
		var token1Symbol string
		var token2Symbol string
		var Token1Amount = 0
		var Token2Amount = 0
		for _, term := range terms {
			if term.Token1Symbol != "" {
				token1Symbol = term.Token1Symbol
				Token1Amount = Token1Amount + term.Token1Amount
			}
			if term.Token2Symbol != "" {
				token2Symbol = term.Token2Symbol
				Token2Amount = Token2Amount + term.Token2Amount
			}
		}
		if token1Symbol != "" {
			*rewards = append(*rewards, model.Reward{
				TokenSymbol: token1Symbol,
				Amount:      Token1Amount,
			})
		}
		if token2Symbol != "" {
			*rewards = append(*rewards, model.Reward{
				TokenSymbol: token2Symbol,
				Amount:      Token2Amount,
			})
		}
	}
}

func calcRewardWhenIsPaymentPeriod(periods []model.BountyPaymentPeriod, rewards *[]model.Reward) {
	if len(periods) > 0 {
		byTokenSymbol := make(map[string]int)
		var token1Symbol string // 其实固定是 UVU !!
		var token2Symbol string
		for _, period := range periods {
			if period.Token1Symbol != "" {
				token1Symbol = period.Token1Symbol
				if v, ok := byTokenSymbol[token1Symbol]; ok {
					byTokenSymbol[token1Symbol] = v + period.Token1Amount
				} else {
					byTokenSymbol[token1Symbol] = period.Token1Amount
				}
			}
			if period.Token2Symbol != "" {
				token2Symbol = period.Token2Symbol
				if v, ok := byTokenSymbol[token2Symbol]; ok {
					byTokenSymbol[token2Symbol] = v + period.Token2Amount
				} else {
					byTokenSymbol[token2Symbol] = period.Token2Amount
				}
			}
		}
		if token1Symbol != "" {
			*rewards = append(*rewards, model.Reward{
				TokenSymbol: token1Symbol,
				Amount:      byTokenSymbol[token1Symbol],
			})
		}
		if token2Symbol != "" {
			*rewards = append(*rewards, model.Reward{
				TokenSymbol: token2Symbol,
				Amount:      byTokenSymbol[token2Symbol],
			})
		}
	}
}

func GetBountyDetailByID(bountyID uint64) (*model.DetailResponse, error) {
	detailResponse, err := model.GetDetailByBountyID(mysql.DB, bountyID)
	if err != nil {
		return nil, err
	}
	return detailResponse, nil
}

func GetPaymentByBountyID(bountyID uint64, comerID uint64) (*model.PaymentResponse, error) {
	response, err := model.GetPaymentByBountyID(mysql.DB, bountyID, comerID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetBountyState(bountyID uint64, comerID uint64) (*model.GetBountyStateResponse, error) {
	response, err := model.GetBountyState(mysql.DB, bountyID, comerID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func UpdateBountyStatusByID(bountyID uint64) (string, error) {
	response, err := model.UpdateBountyCloseStatusByID(mysql.DB, bountyID)
	if err != nil {
		return "", err
	}
	return response, nil
}

func AddDeposit(bountyID uint64, request *model.AddDepositRequest, comerID uint64) error {
	deposit := &model.BountyDeposit{
		ChainID:     request.ChainID,
		TxHash:      request.TxHash,
		Status:      transaction.Success,
		BountyID:    bountyID,
		ComerID:     comerID,
		Access:      model.AccessIn,
		TokenSymbol: request.TokenSymbol,
		TokenAmount: request.TokenAmount,
		Timestamp:   time.Now(),
	}

	transactionApplicant := &model4.Transaction{
		ChainID:    request.ChainID,
		TxHash:     request.TxHash, // test chain
		TimeStamp:  time.Now(),
		Status:     transaction.Success,
		SourceType: transaction.BountyDepositAccount,
		RetryTimes: 0,
		SourceID:   int64(bountyID),
	}
	err := mysql.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = model4.CreateTransaction(tx, transactionApplicant)
		if err != nil {
			return err
		}
		err = model.CreateDeposit(tx, deposit)
		if err != nil {
			return err
		}
		err = model.UpdateBountyDepositAmount(tx, bountyID, request.TokenAmount)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func PayReward(bountyID uint64, request *model.PaidRequest) error {
	err := model.UpdatePaidByBountyID(mysql.DB, bountyID, request)
	if err != nil {
		return err
	}
	return nil
}

func CreateActivities(bountyID uint64, req *model.ActivitiesRequest, comerID uint64) error {
	postUpdate := &model3.PostUpdate{
		SourceType: req.SourceType, // 1 bounty normal 2 send-paid-info
		SourceID:   bountyID,
		ComerID:    comerID,
		Content:    req.Content,
		TimeStamp:  time.Now(),
	}
	err := model3.CreatePostUpdate(mysql.DB, postUpdate)
	if err != nil {
		return err
	}
	return nil
}

func CreateApplicants(bountyID uint64, request *model.ApplicantsDepositRequest, comerID uint64) error {
	// async update
	// updateDepositStatusChain(request.ChainID, request.TxHash, bountyID)

	bountyApplicant := &model.BountyApplicantForBounty{
		BountyID:    bountyID,
		ComerID:     comerID,
		ApplyAt:     time.Now(),
		Description: request.Description,
		Status:      model.ApplicantStatusApplied,
	}
	deposit := &model.BountyDeposit{
		ChainID:     request.ChainID,
		TxHash:      request.TxHash,
		Status:      transaction.Success,
		BountyID:    bountyID,
		ComerID:     comerID,
		Access:      model.AccessIn,
		TokenSymbol: request.TokenSymbol,
		TokenAmount: request.TokenAmount,
		Timestamp:   time.Now(),
	}
	transaction := &model4.Transaction{
		ChainID:    request.ChainID,
		TxHash:     request.TxHash,
		TimeStamp:  time.Now(),
		Status:     transaction.Success,
		SourceType: transaction.BountyDepositAccount,
		RetryTimes: 1,
		SourceID:   int64(bountyID),
	}

	err := mysql.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = model.CreateApplicants(tx, bountyApplicant)
		if err != nil {
			log.Warn(err)
			return err
		}

		err = model.CreateDeposit(tx, deposit)
		if err != nil {
			log.Warn(err)
			return err
		}
		if len(request.TxHash) > 0 {

			if err := model4.CreateTransaction(tx, transaction); err != nil {
				log.Warn(err)
				return err
			}
		}
		return nil
	})
	return err
}

func GetActivitiesByBountyID(bountyID uint64) (*[]model.ActivitiesResponse, error) {
	response, err := model.GetActivitiesByBountyID(mysql.DB, bountyID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetAllApplicantsByBountyID(bountyID uint64) (*[]model.Applicant, error) {
	response, err := model.GetApplicants(mysql.DB, bountyID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetFounderByBountyID(bountyID uint64) (*model.FounderResponse, error) {
	response, err := model.GetFounderByBountyID(mysql.DB, bountyID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetApprovedApplicantByBountyID(bountyID uint64) (*model.ApprovedResponse, error) {
	response, err := model.GetApprovedApplicantByBountyID(mysql.DB, bountyID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetDepositRecords(bountyID uint64) (*[]model.DepositRecord, error) {
	response, err := model.GetDepositRecordsByBountyID(mysql.DB, bountyID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func UpdateApplicantApprovedStatus(request *model.ApplicantsApprovedRequst, bountyID, comerID uint64, applicantComerID uint64, depositLockStatus int) (err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) (err error) {

		if err := model.UpdateApplicantApprovedStatus(tx, bountyID, applicantComerID, model.ApplicantStatusApproved); err != nil {
			return err
		}

		if err := model.UpdateBountyStatus(tx, bountyID, model.BountyStatusWordStarted); err != nil {
			fmt.Println("err ", err)
			return err
		}

		if err := UpdateApplicantDepositLockStatus(bountyID, applicantComerID, depositLockStatus); err != nil {
			fmt.Println("err ", err)
			return err
		}

		if err := UpdateApplicantDepositLockStatus(bountyID, comerID, depositLockStatus); err != nil {
			fmt.Println("err ", err)
			return err
		}

		// release other deposit
		if len(request.TxHash) != 0 {
			applicantsRejectedDeposits, err := model.GetApplicantsRejectedDeposits(tx, bountyID, applicantComerID)
			if err != nil {
				return err
			}

			for _, rejectedApplicantInfo := range applicantsRejectedDeposits {
				if rejectedApplicantInfo.ComerID == comerID {
					continue
				}
				deposit := &model.BountyDeposit{
					ChainID:     request.ChainID,
					TxHash:      request.TxHash,
					Status:      transaction.Success,
					BountyID:    bountyID,
					ComerID:     rejectedApplicantInfo.ComerID,
					Access:      model.AccessOut,
					TokenSymbol: rejectedApplicantInfo.TokenSymbol,
					TokenAmount: rejectedApplicantInfo.TokenAmount,
					Timestamp:   time.Now(),
				}

				err = model.CreateDeposit(tx, deposit)
				if err != nil {
					return err
				}
			}

			transactionApplicant := &model4.Transaction{
				ChainID:    request.ChainID,
				TxHash:     request.TxHash,
				TimeStamp:  time.Now(),
				Status:     transaction.Success,
				SourceType: transaction.BountyDepositAccount,
				RetryTimes: 1,
				SourceID:   int64(bountyID),
			}
			err = model4.CreateTransaction(tx, transactionApplicant)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

func GetStartupByBountyID(bountyID uint64) (*model.StartupListResponse, error) {
	response, err := model.GetStartupByBountyID(mysql.DB, bountyID)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func UpdateApplicantDepositLockStatus(bountyID, comerID uint64, depositStatus int) error {
	err := model.UpdateApplicantDepositLockStatus(mysql.DB, bountyID, comerID, depositStatus)
	if err != nil {
		return err
	}
	return nil
}

func ReleaseFounderDeposit(request *model.ReleaseRequst, bountyID uint64, comerID uint64) error {
	fmt.Println("releaseFounderDeposit")

	// 更新申请者状态
	err := mysql.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err := model.UpdateApplicantDepositByBountyID(tx, bountyID); err != nil {
			return err
		}
		if err := model.UpdateFounderDepositByBountyID(tx, bountyID); err != nil {
			return err
		}

		// release other deposit
		if len(request.TxHash) != 0 {
			applicantsRejectedDeposits, _ := model.GetAllApplicantsReleaseDeposits(tx, bountyID)
			for _, rejectedApplicantInfo := range applicantsRejectedDeposits {
				deposit := &model.BountyDeposit{
					ChainID:     request.ChainID,
					TxHash:      request.TxHash,
					Status:      transaction.Success,
					BountyID:    bountyID,
					ComerID:     rejectedApplicantInfo.ComerID,
					Access:      model.AccessOut,
					TokenSymbol: rejectedApplicantInfo.TokenSymbol,
					TokenAmount: rejectedApplicantInfo.TokenAmount,
					Timestamp:   time.Now(),
				}

				transactionApplicant := &model4.Transaction{
					ChainID:    request.ChainID,
					TxHash:     request.TxHash,
					TimeStamp:  time.Now(),
					Status:     transaction.Success,
					SourceType: transaction.BountyDepositAccount,
					RetryTimes: 1,
					SourceID:   int64(bountyID),
				}
				err = model4.CreateTransaction(tx, transactionApplicant)
				if err != nil {
					return err
				}

				err = model.CreateDeposit(tx, deposit)
				if err != nil {
					return err
				}
				// 更新申请者状态为释放状态
				err = model.UpdateApplicantStatus(tx, bountyID, rejectedApplicantInfo.ComerID, model.ApplicantStatusRefunded)
				if err != nil {
					return err
				}
			}

		}

		return nil
	})
	return err
}

func UpdateApplicantUnApprovedStatus(bountyID, applicantComerID uint64) error {
	return model.UpdateApplicantStatus(mysql.DB, bountyID, applicantComerID, model.ApplicantStatusUnApproved)
}

func ReleaseComerDeposit(request *model.ReleaseMyDepositRequst, bountyID uint64, comerID uint64) error {
	// 更新申请者状态
	err := mysql.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err := model.UpdateApplicantStatus(tx, bountyID, comerID, model.ApplicantStatusWithdraw); err != nil {
			return err
		}

		// release other deposit
		// fmt.Println("request.TxHash", request.TxHash)
		if len(request.TxHash) != 0 {
			applicantsRejectedDeposit, err := model.GetApplicantsReleaseDeposits(tx, bountyID, comerID)

			if err == nil {
				deposit := &model.BountyDeposit{
					ChainID:     request.ChainID,
					TxHash:      request.TxHash,
					Status:      transaction.Success,
					BountyID:    bountyID,
					ComerID:     applicantsRejectedDeposit.ComerID,
					Access:      model.AccessOut,
					TokenSymbol: applicantsRejectedDeposit.TokenSymbol,
					TokenAmount: applicantsRejectedDeposit.TokenAmount,
					Timestamp:   time.Now(),
				}

				transactionApplicant := &model4.Transaction{
					ChainID:    request.ChainID,
					TxHash:     request.TxHash, // test chain
					TimeStamp:  time.Now(),
					Status:     transaction.Success,
					SourceType: transaction.BountyDepositAccount,
					RetryTimes: 1,
					SourceID:   int64(bountyID),
				}
				err = model4.CreateTransaction(tx, transactionApplicant)
				if err != nil {
					return err
				}

				err = model.CreateDeposit(tx, deposit)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	return err

}
