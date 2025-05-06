package bounty

import (
	"ceres/pkg/initialization/utility"
	"ceres/pkg/model"
	"ceres/pkg/model/bounty"
	"ceres/pkg/router"
	"encoding/json"
)

const (
	DepositSuccessStatus = 2
	DepositLockStatus    = 3
	DepositUnLockStatus  = 4
)

var bountiesString = "[\n    {\n        \"applicant_count\": 8,\n        \"applicants_deposit\": 3200,\n        \"applicant_min_deposit\": 400,\n        \"apply_deadline\": \"2023-11-30T23:59:59Z\",\n        \"chain_id\": 1,\n        \"comer_id\": 2001,\n        \"contract_address\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n        \"created_at\": \"2023-03-01T08:00:00Z\",\n        \"deposit_contract_address\": \"0x742d35Cc6634C0532925a3b844Bc454e4438f44e\",\n        \"deposit_contract_token_decimal\": 6,\n        \"deposit_contract_token_symbol\": \"USDC\",\n        \"discussion_link\": \"https://forum.metaland.xyz/bounties/301\",\n        \"expired_time\": \"2024-03-01T00:00:00Z\",\n        \"founder_deposit\": 5000,\n        \"id\": 301,\n        \"is_lock\": 0,\n        \"payment_mode\": 1,\n        \"reward\": {\n            \"token1_symbol\": \"MLT\",\n            \"token1_amount\": 15000,\n            \"token2_symbol\": \"ETH\", \n            \"token2_amount\": 2\n        },\n        \"skills\": [\n            {\n                \"id\": 8,\n                \"tag\": {\n                    \"id\": 8,\n                    \"name\": \"智能合约审计\",\n                    \"type\": 2\n                },\n                \"tag_id\": 8,\n                \"target_id\": 301,\n                \"type\": 3\n            }\n        ],\n        \"startup\": {\n            \"banner\": \"https://storage.metaland.xyz/startups/banner_301.jpg\",\n            \"chain_id\": 1,\n            \"comer_id\": 2001,\n            \"contract_audit\": \"audit_v3.pdf\",\n            \"id\": 301,\n            \"is_connected\": true,\n            \"kyc\": \"kyc_301.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/startups/logo_301.png\",\n            \"mission\": \"构建去中心化金融基础设施\",\n            \"name\": \"DeFi Lab\",\n            \"on_chain\": true,\n            \"tx_hash\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n            \"type\": 1\n        },\n        \"startup_id\": 301,\n        \"status\": 2,\n        \"title\": \"开发跨链资产桥接协议\",\n        \"tx_hash\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\"\n    },\n    {\n        \"applicant_count\": 15,\n        \"applicants_deposit\": 7500,\n        \"applicant_min_deposit\": 500,\n        \"apply_deadline\": \"2024-01-15T23:59:59Z\",\n        \"chain_id\": 137,\n        \"comer_id\": 2002,\n        \"contract_address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n        \"created_at\": \"2023-09-01T10:00:00Z\",\n        \"deposit_contract_address\": \"0xda0C8C8d94a7bD15D9166B9d3FfDd2e9850Fb97B\",\n        \"deposit_contract_token_decimal\": 18,\n        \"deposit_contract_token_symbol\": \"MATIC\",\n        \"discussion_link\": \"https://forum.metaland.xyz/bounties/302\",\n        \"expired_time\": \"2024-09-01T00:00:00Z\",\n        \"founder_deposit\": 10000,\n        \"id\": 302,\n        \"is_lock\": 1,\n        \"payment_mode\": 3,\n        \"reward\": {\n            \"token1_symbol\": \"MLT\",\n            \"token1_amount\": 25000,\n            \"token2_symbol\": \"USDT\",\n            \"token2_amount\": 500\n        },\n        \"skills\": [\n            {\n                \"id\": 12,\n                \"tag\": {\n                    \"id\": 12,\n                    \"name\": \"ZK-Rollup\",\n                    \"type\": 4\n                },\n                \"tag_id\": 12,\n                \"target_id\": 302,\n                \"type\": 2\n            }\n        ],\n        \"startup\": {\n            \"banner\": \"https://storage.metaland.xyz/startups/banner_302.jpg\",\n            \"chain_id\": 137,\n            \"comer_id\": 2002,\n            \"contract_audit\": \"audit_v4.pdf\",\n            \"id\": 302,\n            \"is_connected\": false,\n            \"kyc\": \"kyc_302.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/startups/logo_302.png\",\n            \"mission\": \"开发Layer2扩容解决方案\",\n            \"name\": \"ZK Labs\",\n            \"on_chain\": true,\n            \"tx_hash\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n            \"type\": 3\n        },\n        \"startup_id\": 302,\n        \"status\": 1,\n        \"title\": \"零知识证明电路优化\",\n        \"tx_hash\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\"\n    }\n]"

func GetBounties(ctx *router.Context) {
	var res model.PageData
	var bounties []bounty.BountyBasicResponse
	err := json.Unmarshal([]byte(bountiesString), &bounties)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	res.Total = len(bounties)
	res.Page = 1
	res.Size = 15
	res.List = utility.ConvertToInterfaceSlice(bounties)
	ctx.OK(res)
}

func CreateBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "create bounty successful!"
	ctx.OK(message)
}

func GetBountyInfo(ctx *router.Context) {
	var res bounty.BountyInfoResponse
	var resString = "{\n    \"applicant_deposit\": 5000,\n    \"applicant_min_deposit\": 100,\n    \"applicants\": [\n        {\n            \"comerID\": \"user_1001\",\n            \"image\": \"https://storage.metaland.xyz/avatars/1001.png\",\n            \"name\": \"开发者小明\",\n            \"address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n            \"description\": \"拥有3年Solidity开发经验\",\n            \"status\": 1,\n            \"applyAt\": \"2023-08-01T10:00:00Z\"\n        }\n    ],\n    \"apply_deadline\": \"2023-12-31T23:59:59Z\",\n    \"approved\": {\n        \"comerID\": \"user_1002\",\n        \"image\": \"https://storage.metaland.xyz/avatars/1002.png\",\n        \"name\": \"审计员安娜\",\n        \"address\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n        \"description\": \"智能合约安全专家\",\n        \"status\": 2\n    },\n    \"chain_id\": 1,\n    \"comer_id\": 1001,\n    \"contacts\": [\n        {\n            \"type\": 1,\n            \"value\": \"contact@project.io\"\n        }\n    ],\n    \"contract_address\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n    \"created_at\": \"2023-06-01T00:00:00Z\",\n    \"deposit_contract_address\": \"0x742d35Cc6634C0532925a3b844Bc454e4438f44e\",\n    \"deposit_contract_token_decimal\": 18,\n    \"deposit_contract_token_symbol\": \"MLT\",\n    \"deposit_records\": [\n        {\n            \"amount\": 1000,\n            \"bounty_id\": 501,\n            \"comer\": {\n                \"activation\": true,\n                \"address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n                \"avatar\": \"https://storage.metaland.xyz/avatars/1001.png\",\n                \"id\": 1001,\n                \"name\": \"开发者小明\"\n            },\n            \"created_at\": \"2023-08-01T10:00:00Z\",\n            \"status\": 1\n        }\n    ],\n    \"description\": \"开发去中心化投票系统，要求支持多链部署\",\n    \"discussion_link\": \"https://forum.metaland.xyz/t/501\",\n    \"expired_time\": \"2024-06-01T00:00:00Z\",\n    \"founder\": {\n        \"activation\": true,\n        \"address\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n        \"avatar\": \"https://storage.metaland.xyz/avatars/founder_501.png\",\n        \"name\": \"项目发起人\",\n        \"skills\": [\n            {\n                \"id\": 5,\n                \"tag\": {\n                    \"id\": 5,\n                    \"name\": \"区块链开发\",\n                    \"type\": 1\n                }\n            }\n        ]\n    },\n    \"founder_deposit\": 10000,\n    \"id\": 501,\n    \"is_lock\": 0,\n    \"my_deposit\": 500,\n    \"my_role\": 2,\n    \"my_status\": 1,\n    \"payment_mode\": 2,\n    \"period\": {\n        \"hours_per_day\": 8,\n        \"period_type\": 3\n    },\n    \"skills\": [\n        {\n            \"id\": 5,\n            \"tag\": {\n                \"id\": 5,\n                \"name\": \"Solidity\",\n                \"type\": 1\n            },\n            \"type\": 1\n        }\n    ],\n    \"startup\": {\n        \"banner\": \"https://storage.metaland.xyz/startups/banner_501.jpg\",\n        \"logo\": \"https://storage.metaland.xyz/startups/logo_501.png\",\n        \"name\": \"去中心化治理实验室\",\n        \"socials\": [\n            {\n                \"social_tool\": {\n                    \"id\": 3,\n                    \"logo\": \"https://static.metaland.xyz/socials/twitter.png\",\n                    \"name\": \"Twitter\"\n                },\n                \"value\": \"@governance_lab\"\n            }\n        ]\n    },\n    \"startup_id\": 501,\n    \"status\": 1,\n    \"terms\": [\n        {\n            \"token1_amount\": 5000,\n            \"token1_symbol\": \"MLT\",\n            \"terms\": \"第一阶段交付基础合约\"\n        }\n    ],\n    \"title\": \"去中心化投票系统开发\",\n    \"tx_hash\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(res)
}

func ApplyBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "apply bounty successful!"
	ctx.OK(message)
}

func CloseBounty(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "close bounty successful!"
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
