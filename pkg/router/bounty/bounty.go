package bounty

import (
	"ceres/pkg/model"
	"ceres/pkg/router"
)

const (
	DepositSuccessStatus = 2
	DepositLockStatus    = 3
	DepositUnLockStatus  = 4
)

var bountiesString = "{\"page\":1,\"size\":10,\"total\":4,\"list\":[{\"id\":262852933996544,\"startup_id\":260702765973504,\"comer_id\":260417511387136,\"chain_id\":57000,\"tx_hash\":\"0xc9c58e519842fbfe1923208a16c1ce7841a37778a080e5f51248847aa23ad45d\",\"contract_address\":\"0x2C676FbF47bFb8305c8e4D74e950f67a1DaACfca\",\"apply_deadline\":\"1690766459000\",\"title\":\"0707 test  bounty\",\"applicant_min_deposit\":\"0\",\"founder_deposit\":\"0.1\",\"applicant_deposit\":\"0.2\",\"deposit_contract_address\":\"0x0000000000000000000000000000000000000000\",\"deposit_contract_token_decimal\":18,\"deposit_contract_token_symbol\":\"TSYS\",\"discussion_link\":\"\",\"payment_mode\":1,\"expired_time\":\"0\",\"is_lock\":0,\"status\":1,\"created_at\":\"1688664080000\",\"applicant_count\":1,\"skills\":[{\"id\":262852933996547,\"type\":6,\"tag_id\":170810836987909,\"target_id\":262852933996544,\"tag\":{\"id\":170810836987909,\"name\":\"Developer\",\"type\":6}}],\"startup\":{\"id\":260702765973504,\"comer_id\":260417511387136,\"name\":\"Test #01\",\"logo\":\"\",\"type\":4,\"mission\":\"Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test\",\"kyc\":\"\",\"contract_audit\":\"\",\"chain_id\":57000,\"on_chain\":true,\"tx_hash\":\"0xd70e153f9e084df9d188ac6a1938c28480daa04f2a2e5d1f21626747160dc9e3\",\"is_connected\":false,\"banner\":\"\"},\"reward\":{\"bounty_id\":262852933996544,\"token1_symbol\":\"USDC\",\"token1_amount\":\"1\",\"token2_symbol\":\"\",\"token2_amount\":\"0\"}},{\"id\":261956929986560,\"startup_id\":260775855996929,\"comer_id\":258939681902592,\"chain_id\":57000,\"tx_hash\":\"0x35ad7b61dbca367cafa579840ccf452b8842baa6b09553373d1d4292f0b33623\",\"contract_address\":\"0x00d13EFa481f68d13e73941119C721C32C32Add0\",\"apply_deadline\":\"1689861612000\",\"title\":\"Has transferred liquidity from GoRollux to Pegasys v3 on Rollux , that is why TVL will increaseHas transferred liquidity from GoRollux to Pegasys v3 on Rollux , that is why TVL will increase\",\"applicant_min_deposit\":\"2\",\"founder_deposit\":\"0\",\"applicant_deposit\":\"0\",\"deposit_contract_address\":\"0x2Be160796F509CC4B1d76fc97494D56CF109C3f1\",\"deposit_contract_token_decimal\":6,\"deposit_contract_token_symbol\":\"USDC\",\"discussion_link\":\"\",\"payment_mode\":1,\"expired_time\":\"0\",\"is_lock\":0,\"status\":1,\"created_at\":\"1688450456000\",\"applicant_count\":0,\"skills\":[{\"id\":261956929986563,\"type\":6,\"tag_id\":170810836987910,\"target_id\":261956929986560,\"tag\":{\"id\":170810836987910,\"name\":\"UI/UE\",\"type\":6}},{\"id\":261956929986564,\"type\":6,\"tag_id\":170629596917763,\"target_id\":261956929986560,\"tag\":{\"id\":170629596917763,\"name\":\"Project Manager\",\"type\":6}},{\"id\":261956929986565,\"type\":6,\"tag_id\":170810836987911,\"target_id\":261956929986560,\"tag\":{\"id\":170810836987911,\"name\":\"Designer\",\"type\":6}}],\"startup\":{\"id\":260775855996929,\"comer_id\":258939681902592,\"name\":\"阿斯蒂芬按上级领导发生打飞机啊说法\",\"logo\":\"\",\"type\":0,\"mission\":\"\",\"kyc\":\"\",\"contract_audit\":\"\",\"chain_id\":57000,\"on_chain\":true,\"tx_hash\":\"0x9479684e50f953e050a341e922a0283aeb301ffbb897230dde0e56fcb8f56861\",\"is_connected\":false,\"banner\":\"\"},\"reward\":{\"bounty_id\":261956929986560,\"token1_symbol\":\"USDC\",\"token1_amount\":\"3\",\"token2_symbol\":\"\",\"token2_amount\":\"0\"}},{\"id\":261949279576064,\"startup_id\":261930283573249,\"comer_id\":167619177164800,\"chain_id\":57000,\"tx_hash\":\"0x2e6c207eedf05247c7bcb3995a7d4f985574a0140d4b1678eec7dd9dab979921\",\"contract_address\":\"0xC722A52D723FD0109220b0c7eCE6582B21862b66\",\"apply_deadline\":\"1690810035000\",\"title\":\"0704 test bounty\",\"applicant_min_deposit\":\"1\",\"founder_deposit\":\"0.1\",\"applicant_deposit\":\"0\",\"deposit_contract_address\":\"0x0000000000000000000000000000000000000000\",\"deposit_contract_token_decimal\":18,\"deposit_contract_token_symbol\":\"TSYS\",\"discussion_link\":\"\",\"payment_mode\":1,\"expired_time\":\"0\",\"is_lock\":0,\"status\":1,\"created_at\":\"1688448632000\",\"applicant_count\":0,\"skills\":[{\"id\":261949279576067,\"type\":6,\"tag_id\":170810836987909,\"target_id\":261949279576064,\"tag\":{\"id\":170810836987909,\"name\":\"Developer\",\"type\":6}}],\"startup\":{\"id\":261930283573249,\"comer_id\":167619177164800,\"name\":\"0704 project on testnet\",\"logo\":\"\",\"type\":1,\"mission\":\"const res = await this.contractStore.setStartupSuccessAfter({               startup_id: String(this.\",\"kyc\":\"\",\"contract_audit\":\"\",\"chain_id\":57000,\"on_chain\":true,\"tx_hash\":\"0x1815946ec3c37c0dac6c545fb36a96cad8bc3e4453e075023b0ab58e8e5de44b\",\"is_connected\":false,\"banner\":\"\"},\"reward\":{\"bounty_id\":261949279576064,\"token1_symbol\":\"USDC\",\"token1_amount\":\"1\",\"token2_symbol\":\"\",\"token2_amount\":\"0\"}},{\"id\":260778599071744,\"startup_id\":260702765973504,\"comer_id\":260417511387136,\"chain_id\":57000,\"tx_hash\":\"0x0610b1cc9cc5992911e1aab9e7c65b4132ed42646ea0ac3bf4e9e455d51db82b\",\"contract_address\":\"0x85bCa780E19BE2E904F3933f51068bdaaa2F69C6\",\"apply_deadline\":\"1689839833000\",\"title\":\"Test bountyTest bountyTest bounty\",\"applicant_min_deposit\":\"0\",\"founder_deposit\":\"0\",\"applicant_deposit\":\"0\",\"deposit_contract_address\":\"0x0000000000000000000000000000000000000000\",\"deposit_contract_token_decimal\":18,\"deposit_contract_token_symbol\":\"TSYS\",\"discussion_link\":\"\",\"payment_mode\":1,\"expired_time\":\"0\",\"is_lock\":0,\"status\":1,\"created_at\":\"1688169520000\",\"applicant_count\":3,\"skills\":[{\"id\":260778599071749,\"type\":6,\"tag_id\":170810836987910,\"target_id\":260778599071744,\"tag\":{\"id\":170810836987910,\"name\":\"UI/UE\",\"type\":6}},{\"id\":260778599071750,\"type\":6,\"tag_id\":170629596917763,\"target_id\":260778599071744,\"tag\":{\"id\":170629596917763,\"name\":\"Project Manager\",\"type\":6}},{\"id\":260778599071751,\"type\":6,\"tag_id\":170629596917765,\"target_id\":260778599071744,\"tag\":{\"id\":170629596917765,\"name\":\"Marketing\",\"type\":6}}],\"startup\":{\"id\":260702765973504,\"comer_id\":260417511387136,\"name\":\"Test #01\",\"logo\":\"\",\"type\":4,\"mission\":\"Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test #01Test\",\"kyc\":\"\",\"contract_audit\":\"\",\"chain_id\":57000,\"on_chain\":true,\"tx_hash\":\"0xd70e153f9e084df9d188ac6a1938c28480daa04f2a2e5d1f21626747160dc9e3\",\"is_connected\":false,\"banner\":\"\"},\"reward\":{\"bounty_id\":260778599071744,\"token1_symbol\":\"TSYS\",\"token1_amount\":\"39\",\"token2_symbol\":\"\",\"token2_amount\":\"0\"}}]}"

func GetBounties(ctx *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	ctx.OK(res)
}

func CreateBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "create bounty successful!"
	ctx.OK(message)
}

func GetBountyInfo(ctx *router.Context) {
	ctx.OK("")
}

func ApplyBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "apply bounty successful!"
	ctx.OK(message)
}

func CloseBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "apply bounty successful!"
	ctx.OK(message)
}

func PayBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "pay bounty successful!"
	ctx.OK(message)
}

func PostUpdateBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "post update bounty successful!"
	ctx.OK(message)
}

//func CreateBounty(ctx *router.Context) {
//	request := new(bounty.BountyRequest)
//	if err := ctx.ShouldBindJSON(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	if err := service.CreateComerBounty(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	response := "create bounty successful!"
//
//	ctx.OK(response)
//}
//
//
//func GetPublicBountyList(ctx *router.Context) {
//	var request model.Pagination
//	if err := model.ParsePagination(ctx, &request, 10); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	if err := service.QueryAllOnChainBounties(&request); err != nil {
//		ctx.HandleError(err)
//	} else {
//		ctx.OK(request)
//	}
//}
//
//
//func GetBountyListByStartup(ctx *router.Context) {
//	var pagination model.Pagination
//	if err := model.ParsePagination(ctx, &pagination, 3); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	startupId, err := strconv.ParseUint(ctx.Param("startupId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if startupId == 0 {
//		err := router.ErrBadRequest.WithMsg("Invalid startupId!")
//		ctx.HandleError(err)
//		return
//	}
//
//	if err := service.QueryBountiesByStartup(startupId, &pagination); err != nil {
//		ctx.HandleError(err)
//	} else {
//		ctx.OK(pagination)
//	}
//}
//
//
//func GetMyPostedBountyList(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var pagination model.Pagination
//	if err := model.ParsePagination(ctx, &pagination, 5); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	if err := service.QueryComerPostedBountyList(comerID, &pagination); err != nil {
//		ctx.HandleError(err)
//	} else {
//		ctx.OK(pagination)
//	}
//}
//
//
//func GetMyParticipatedBountyList(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var pagination model.Pagination
//	if err := model.ParsePagination(ctx, &pagination, 8); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	if err := service.QueryComerParticipatedBountyList(comerID, &pagination); err != nil {
//		ctx.HandleError(err)
//	} else {
//		ctx.OK(pagination)
//	}
//}
//
//func GetBountyDetailByID(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.GetBountyDetailByID(bountyID)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("get bounty detail fail")
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func GetPaymentByBountyID(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	response, err := service.GetPaymentByBountyID(bountyID, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func GetState(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	response, err := service.GetBountyState(bountyID, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func UpdateBountyCloseStatus(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.UpdateBountyStatusByID(bountyID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func AddDeposit(ctx *router.Context) {
//	request := new(bounty.AddDepositRequest)
//	if err := ctx.ShouldBindJSON(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	err = service.AddDeposit(bountyID, request, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("add deposit success")
//}
//
//// 支付奖金
//func PayReward(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	request := new(bounty.PaidRequest)
//	if err := ctx.ShouldBindJSON(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	err = service.PayReward(bountyID, request)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK("update paid success")
//}
//
//func CreateActivities(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	in := new(bounty.ActivitiesRequest)
//	if err := ctx.ShouldBindJSON(in); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	err = service.CreateActivities(bountyID, in, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("activities create success")
//}
//
//func CreateApplicants(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	request := new(bounty.ApplicantsDepositRequest)
//	if err := ctx.ShouldBindJSON(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	err = service.CreateApplicants(bountyID, request, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("create applicants success")
//}
//
//func GetActivitiesLists(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.GetActivitiesByBountyID(bountyID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func GetAllApplicantsByBountyID(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.GetAllApplicantsByBountyID(bountyID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func GetFounderByBountyID(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.GetFounderByBountyID(bountyID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func GetApprovedApplicantByBountyID(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.GetApprovedApplicantByBountyID(bountyID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func GetDepositRecords(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.GetDepositRecords(bountyID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func UpdateFounderApprovedApplicant(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	applicantComerID, err := strconv.ParseUint(ctx.Param("applicantComerID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	request := new(bounty.ApplicantsApprovedRequst)
//	if err := ctx.ShouldBindJSON(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	err = service.UpdateApplicantApprovedStatus(request, bountyID, comerID, applicantComerID, DepositLockStatus)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("approved success")
//}
//
//func UpdateFounderUnapprovedApplicant(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	applicantComerID, err := strconv.ParseUint(ctx.Param("applicantComerID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	err = service.UpdateApplicantUnApprovedStatus(bountyID, applicantComerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("unapproved success")
//}
//
//func GetStartupByBountyID(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//	response, err := service.GetStartupByBountyID(bountyID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//func UpdateApplicantsLockDeposit(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	err = service.UpdateApplicantDepositLockStatus(bountyID, comerID, DepositLockStatus)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("lock deposit success")
//}
//
//func UpdateApplicantsUnlockDeposit(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	err = service.UpdateApplicantDepositLockStatus(bountyID, comerID, DepositUnLockStatus)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("unlock deposit success")
//}
//
//func ReleaseDeposit(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	request := new(bounty.ReleaseRequst)
//	if err := ctx.ShouldBindJSON(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	err = service.ReleaseFounderDeposit(request, bountyID, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("release deposit success")
//}
//
//func ReleaseMyDeposit(ctx *router.Context) {
//	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, err := tool.GetComerIDByToken(ctx)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	request := new(bounty.ReleaseMyDepositRequst)
//	if err := ctx.ShouldBindJSON(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	err = service.ReleaseComerDeposit(request, bountyID, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("release my deposit success")
//}
//
//func GetComerPostedBountyList(ctx *router.Context) {
//	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var pagination model.Pagination
//	if err := model.ParsePagination(ctx, &pagination, 5); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	if err := service.QueryComerPostedBountyList(comerID, &pagination); err != nil {
//		ctx.HandleError(err)
//	} else {
//		ctx.OK(pagination)
//	}
//}
//
//func GetComerParticipatedBountyList(ctx *router.Context) {
//	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var pagination model.Pagination
//	if err := model.ParsePagination(ctx, &pagination, 8); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	if err := service.QueryComerParticipatedBountyList(comerID, &pagination); err != nil {
//		ctx.HandleError(err)
//	} else {
//		ctx.OK(pagination)
//	}
//}
