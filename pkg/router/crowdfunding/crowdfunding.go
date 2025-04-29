package crowdfunding

import (
	"ceres/pkg/model"
	"ceres/pkg/model/crowdfunding"
	"ceres/pkg/router"
	"encoding/json"
)

func GetCrowdfunding(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func UpdateCrowdfunding(c *router.Context) {
	var message model.MessageResponse
	message.Message = "update crowd funding successful!"
	c.OK(message)
}

func CreateCrowdfunding(c *router.Context) {
	var message model.MessageResponse
	message.Message = "create crowd funding successful!"
	c.OK(message)
}

func GetCrowdfundingInfo(c *router.Context) {
	var res crowdfunding.CrowdfundingResponse
	var resString = "{\n    \"buy_price\": 150,\n    \"buy_token_contract\": \"0x7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6\",\n    \"buy_token_decimals\": 18,\n    \"buy_token_name\": \"MetaLand Token\",\n    \"buy_token_supply\": 10000000,\n    \"buy_token_symbol\": \"MLT\",\n    \"chain_id\": 1,\n    \"comer_id\": 10001,\n    \"crowdfunding_contract\": \"0x8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7\",\n    \"description\": \"Decentralized Metaverse Infrastructure Funding Round\",\n    \"detail\": \"Funding for development of cross-chain metaverse SDK\",\n    \"dex_init_price\": 180,\n    \"dex_router\": \"0x9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8\",\n    \"end_time\": \"2024-03-31T23:59:59Z\",\n    \"id\": 2024,\n    \"investor\": {\n        \"buy_token_balance\": 5000,\n        \"buy_token_total\": 15000,\n        \"comer_id\": 20001,\n        \"crowdfunding_id\": 2024,\n        \"id\": 3001,\n        \"sell_token_balance\": 800,\n        \"sell_token_total\": 2400\n    },\n    \"investors\": 356,\n    \"max_buy_amount\": 1000,\n    \"max_sell_percent\": 20,\n    \"min_buy_amount\": 100,\n    \"pair_address\": \"0xa1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0\",\n    \"poster\": \"https://storage.metaland.xyz/posters/crowdfunding/2024.jpg\",\n    \"raise_balance\": 534000,\n    \"raise_goal\": 2000000,\n    \"sell_tax\": 5,\n    \"sell_token_balance\": 250000,\n    \"sell_token_contract\": \"0x1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1\",\n    \"sell_token_decimals\": 6,\n    \"sell_token_deposit\": 500000,\n    \"sell_token_name\": \"MetaLand Equity\",\n    \"sell_token_supply\": 1000000,\n    \"sell_token_symbol\": \"MLE\",\n    \"start_time\": \"2024-03-01T00:00:00Z\",\n    \"startup\": {\n        \"banner\": \"https://storage.metaland.xyz/startups/banner_ml.jpg\",\n        \"chain_id\": 1,\n        \"comer_id\": 10001,\n        \"contract_audit\": \"audit_report_v3.pdf\",\n        \"id\": 501,\n        \"is_connected\": true,\n        \"kyc\": \"kyc_verified_v4.pdf\",\n        \"logo\": \"https://storage.metaland.xyz/startups/logo_ml.png\",\n        \"mission\": \"Building decentralized metaverse infrastructure\",\n        \"name\": \"MetaLab\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"type\": \"twitter\",\n                \"url\": \"https://twitter.com/metalab\"\n            },\n            {\n                \"type\": \"telegram\",\n                \"url\": \"https://t.me/metalab_chat\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 5,\n                \"name\": \"Metaverse\"\n            },\n            {\n                \"id\": 8,\n                \"name\": \"Infrastructure\"\n            }\n        ],\n        \"tx_hash\": \"0x2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2\",\n        \"type\": 3\n    },\n    \"startup_id\": 501,\n    \"status\": 2,\n    \"swap_percent\": 75,\n    \"swaps\": [\n        {\n            \"banner\": \"https://storage.metaland.xyz/swaps/swap_001.jpg\",\n            \"chain_id\": 1,\n            \"comer_id\": 20002,\n            \"contract_audit\": \"swap_audit_v2.pdf\",\n            \"id\": 7001,\n            \"is_connected\": true,\n            \"kyc\": \"kyc_swap_001.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/swaps/logo_swap1.png\",\n            \"mission\": \"Liquidity provider for metaverse assets\",\n            \"name\": \"MetaSwap\",\n            \"on_chain\": true,\n            \"socials\": [\n                {\n                    \"type\": \"discord\",\n                    \"url\": \"https://discord.gg/metaswap\"\n                }\n            ],\n            \"tags\": [\n                {\n                    \"id\": 12,\n                    \"name\": \"DEX\"\n                }\n            ],\n            \"tx_hash\": \"0x3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3\",\n            \"type\": 4\n        }\n    ],\n    \"team_wallet\": \"0x4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4\",\n    \"title\": \"MetaLab Infrastructure Funding Round 2024\",\n    \"tx_hash\": \"0x5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5\",\n    \"youtube\": \"https://youtu.be/metaverse-funding-2024\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func GetCrowdfundingTransferLpSign(c *router.Context) {
	var res crowdfunding.SignResponse
	res.Sign = "0x1234567890abcdef1234567890abcdef1234567890abcdef"
	c.OK(res)
}

func GetCrowdfundingInvestRecords(c *router.Context) {
	var res crowdfunding.CrowdfundingSwapResponse
	var resString = "{\n    \"banner\": \"https://storage.metaland.xyz/swaps/bsc_swap_banner.jpg\",\n    \"chain_id\": 56,\n    \"comer_id\": 3005,\n    \"contract_audit\": \"bsc_audit_v2.1.pdf\",\n    \"id\": 5501,\n    \"is_connected\": true,\n    \"kyc\": \"kyc_bsc_swap.pdf\",\n    \"logo\": \"https://storage.metaland.xyz/swaps/bsc_swap_logo.png\",\n    \"mission\": \"提供跨链流动性聚合服务\",\n    \"name\": \"BSC跨链兑换协议\",\n    \"on_chain\": true,\n    \"socials\": [\n        {\n            \"type\": \"twitter\",\n            \"url\": \"https://twitter.com/bsc_swap\"\n        },\n        {\n            \"type\": \"telegram\",\n            \"url\": \"https://t.me/bscswap_chat\"\n        }\n    ],\n    \"tags\": [\n        {\n            \"id\": 15,\n            \"name\": \"BSC生态\"\n        },\n        {\n            \"id\": 22,\n            \"name\": \"跨链协议\"\n        }\n    ],\n    \"tx_hash\": \"0x6d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0\",\n    \"type\": 5\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

//func CreateCrowdfunding(ctx *router.Context) {
//	var request crowdfundingModel.CreateCrowdfundingRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if err := request.ValidRequest(); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	request.ComerId = comerId(ctx)
//	if err := crowdfunding.CreateCrowdfunding(request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//}
//
//func SelectNonFundingStartups(ctx *router.Context) {
//	comerId := comerId(ctx)
//	startups, err := crowdfunding.SelectNonFundingStartups(comerId)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	response := map[string]interface{}{
//		"list":  startups,
//		"total": len(startups),
//	}
//	ctx.OK(response)
//}
//
//func comerId(ctx *router.Context) uint64 {
//	return ctx.Keys[middleware.ComerUinContextKey].(uint64)
//}
//
//func GetCrowdfundingList(ctx *router.Context) {
//	var pagination crowdfundingModel.PublicCrowdfundingListPageRequest
//	if err := ctx.ShouldBindQuery(&pagination); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	if pagination.Limit == 0 {
//		pagination.Limit = 10
//	}
//	err := crowdfunding.GetCrowdfundingList(&pagination)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(pagination)
//}
//
//func GetCrowdfundingDetail(ctx *router.Context) {
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	detail, err := crowdfunding.GetCrowdfundingDetail(crowdfundingId)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(detail)
//}
//
//func GetMyPostedCrowdfundingList(ctx *router.Context) {
//	var pagination model.Pagination
//	if err := ctx.ShouldBindJSON(&pagination); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if pagination.Limit == 0 {
//		pagination.Limit = 10
//	}
//	comerId := comerId(ctx)
//	err := crowdfunding.GetPostedCrowdfundingListByComer(comerId, &pagination)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(pagination)
//}
//
//func GetMyParticipatedCrowdfundingList(ctx *router.Context) {
//	var pagination model.Pagination
//	if err := ctx.ShouldBindJSON(&pagination); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if pagination.Limit == 0 {
//		pagination.Limit = 10
//	}
//	comerId := comerId(ctx)
//	err := crowdfunding.GetParticipatedCrowdFundingListOfComer(comerId, &pagination)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(pagination)
//}
//
//func CancelCrowdfunding(ctx *router.Context) {
//	var re crowdfundingModel.TransactionHashRequest
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if err = ctx.ShouldBindJSON(&re); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId := comerId(ctx)
//	err = crowdfunding.CancelCrowdfunding(comerId, crowdfundingId, re.TxHash)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func RemoveCrowdfunding(ctx *router.Context) {
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var re crowdfundingModel.TransactionHashRequest
//	if err = ctx.ShouldBindJSON(&re); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId := comerId(ctx)
//	err = crowdfunding.FinalizeCrowdFunding(comerId, crowdfundingId, re.TxHash)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func Invest(ctx *router.Context) {
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var re crowdfundingModel.InvestRequest
//	if err = ctx.ShouldBindJSON(&re); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId := comerId(ctx)
//	err = crowdfunding.Invest(comerId, crowdfundingId, re)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func ModifyCrowdfunding(ctx *router.Context) {
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var modifyRequest crowdfundingModel.ModifyRequest
//	if err = ctx.ShouldBindJSON(&modifyRequest); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId := comerId(ctx)
//	err = crowdfunding.ModifyCrowdfunding(comerId, crowdfundingId, modifyRequest)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func GetBuyPriceAndSwapModificationHistories(ctx *router.Context) {
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var pagination model.Pagination
//	if err := ctx.ShouldBindJSON(&pagination); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if pagination.Limit == 0 {
//		pagination.Limit = 3
//	}
//
//	comerId := comerId(ctx)
//	err = crowdfunding.GetBuyPriceAndSwapModificationHistories(comerId, crowdfundingId, &pagination)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(pagination)
//}
//
//func GetCrowdfundingSwapRecords(ctx *router.Context) {
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var pagination model.Pagination
//	if err := ctx.ShouldBindJSON(&pagination); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if pagination.Limit == 0 {
//		pagination.Limit = 10
//	}
//	err = crowdfunding.GetCrowdfundingSwapRecords(crowdfundingId, &pagination)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(pagination)
//}
//
//func GetInvestProfile(ctx *router.Context) {
//	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId := comerId(ctx)
//	investor, err := crowdfunding.GetInvestorDetail(crowdfundingId, comerId)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(investor)
//}
//
//func GetCrowdfundingListOfStartup(ctx *router.Context) {
//	startupId, err := strconv.ParseUint(ctx.Param("startupId"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	x, err := crowdfunding.GetCrowdfundingListByStartup(startupId)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(x)
//}
//
//func GetComerPostedCrowdfundingList(ctx *router.Context) {
//	comerId, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		ctx.HandleError(errors.New("invalid comerId"))
//	}
//	pagination := model.Pagination{
//		Limit: 100,
//		Page:  1,
//		Sort:  "created_at desc",
//	}
//	err = crowdfunding.GetPostedCrowdfundingListByComer(comerId, &pagination)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(pagination.Rows)
//}
//
//func GetComerParticipatedCrowdfundingList(ctx *router.Context) {
//	comerId, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		ctx.HandleError(errors.New("invalid comerId"))
//	}
//	pagination := model.Pagination{
//		Limit: 100,
//		Page:  1,
//		Sort:  "created_at desc",
//	}
//	err = crowdfunding.GetParticipatedCrowdFundingListOfComer(comerId, &pagination)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(pagination.Rows)
//}
