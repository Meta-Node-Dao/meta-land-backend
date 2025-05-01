package crowdfunding

import (
	"ceres/pkg/initialization/utility"
	"ceres/pkg/model"
	"ceres/pkg/model/crowdfunding"
	"ceres/pkg/router"
	"encoding/json"
)

func GetCrowdfunding(c *router.Context) {
	var res model.PageData
	var listString = "[\n    {\n        \"buy_price\": 200,\n        \"buy_token_contract\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n        \"buy_token_symbol\": \"MLT\",\n        \"chain_id\": 1,\n        \"comer_id\": 1001,\n        \"crowdfunding_contract\": \"0x742d35Cc6634C0532925a3b844Bc454e4438f44e\",\n        \"dex_init_price\": 250,\n        \"dex_router\": \"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D\",\n        \"end_time\": \"2023-12-31T23:59:59Z\",\n        \"id\": 601,\n        \"investors\": 85,\n        \"max_buy_amount\": 5000,\n        \"max_sell_percent\": 20,\n        \"min_buy_amount\": 100,\n        \"pair_address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n        \"poster\": \"https://storage.metaland.xyz/posters/cf_601.jpg\",\n        \"raise_balance\": 420000,\n        \"raise_goal\": 500000,\n        \"sell_tax\": 3,\n        \"sell_token_contract\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n        \"sell_token_symbol\": \"USDC\",\n        \"start_time\": \"2023-07-01T00:00:00Z\",\n        \"startup\": {\n            \"banner\": \"https://storage.metaland.xyz/startups/banner_601.jpg\",\n            \"chain_id\": 1,\n            \"comer_id\": 1001,\n            \"contract_audit\": \"audit_v3.pdf\",\n            \"id\": 601,\n            \"is_connected\": true,\n            \"kyc\": \"kyc_601.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/startups/logo_601.png\",\n            \"mission\": \"开发跨链借贷协议\",\n            \"name\": \"CrossFi\",\n            \"on_chain\": true,\n            \"socials\": [\n                {\n                    \"social_tool\": {\n                        \"id\": 5,\n                        \"logo\": \"https://static.metaland.xyz/socials/github.png\",\n                        \"name\": \"GitHub\"\n                    },\n                    \"value\": \"crossfi-protocol\"\n                }\n            ],\n            \"tags\": [\n                {\n                    \"id\": 9,\n                    \"tag\": {\n                        \"id\": 9,\n                        \"name\": \"借贷协议\",\n                        \"type\": 3\n                    },\n                    \"type\": 1\n                }\n            ],\n            \"tx_hash\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n            \"type\": 2\n        },\n        \"startup_id\": 601,\n        \"status\": 1,\n        \"swap_percent\": 15,\n        \"team_wallet\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n        \"title\": \"CrossFi流动性众筹\",\n        \"tx_hash\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\"\n    },\n    {\n        \"buy_price\": 0.05,\n        \"buy_token_contract\": \"0xEA674DeDe5fE460663539C1bB0365bFfE9d444f8\",\n        \"buy_token_symbol\": \"ETH\",\n        \"chain_id\": 56,\n        \"comer_id\": 1002,\n        \"crowdfunding_contract\": \"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984\",\n        \"dex_init_price\": 0.055,\n        \"dex_router\": \"0x10ED43C718714eb63d5aA57B78B54704E256024E\",\n        \"end_time\": \"2024-03-31T23:59:59Z\",\n        \"id\": 602,\n        \"investors\": 120,\n        \"max_buy_amount\": 100,\n        \"max_sell_percent\": 15,\n        \"min_buy_amount\": 0.1,\n        \"pair_address\": \"0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c\",\n        \"poster\": \"https://storage.metaland.xyz/posters/cf_602.jpg\",\n        \"raise_balance\": 850,\n        \"raise_goal\": 1000,\n        \"sell_tax\": 2,\n        \"sell_token_contract\": \"0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56\",\n        \"sell_token_symbol\": \"BUSD\",\n        \"start_time\": \"2023-10-01T00:00:00Z\",\n        \"startup\": {\n            \"banner\": \"https://storage.metaland.xyz/startups/banner_602.jpg\",\n            \"chain_id\": 56,\n            \"comer_id\": 1002,\n            \"contract_audit\": \"audit_v4.pdf\",\n            \"id\": 602,\n            \"is_connected\": false,\n            \"kyc\": \"kyc_602.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/startups/logo_602.png\",\n            \"mission\": \"构建多链NFT交易平台\",\n            \"name\": \"NFT Nexus\",\n            \"on_chain\": true,\n            \"socials\": [\n                {\n                    \"social_tool\": {\n                        \"id\": 4,\n                        \"logo\": \"https://static.metaland.xyz/socials/discord.png\",\n                        \"name\": \"Discord\"\n                    },\n                    \"value\": \"discord.gg/nftnexus\"\n                }\n            ],\n            \"tags\": [\n                {\n                    \"id\": 12,\n                    \"tag\": {\n                        \"id\": 12,\n                        \"name\": \"NFT\",\n                        \"type\": 4\n                    },\n                    \"type\": 2\n                }\n            ],\n            \"tx_hash\": \"0xEA674DeDe5fE460663539C1bB0365bFfE9d444f8\",\n            \"type\": 4\n        },\n        \"startup_id\": 602,\n        \"status\": 2,\n        \"swap_percent\": 10,\n        \"team_wallet\": \"0xEA674DeDe5fE460663539C1bB0365bFfE9d444f8\",\n        \"title\": \"NFT交易平台众筹\",\n        \"tx_hash\": \"0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c\"\n    }\n]"
	var resList []crowdfunding.CrowdfundingBasicResponse
	err := json.Unmarshal([]byte(listString), &resList)
	if err != nil {
		c.HandleError(err)
		return
	}
	res.Total = len(resList)
	res.Page = 1
	res.Size = 15
	res.List = utility.ConvertToInterfaceSlice(resList)
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
	var resString = "{\n    \"buy_price\": 150,\n    \"buy_token_contract\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n    \"buy_token_decimals\": 18,\n    \"buy_token_name\": \"MetaLand Token\",\n    \"buy_token_supply\": 10000000,\n    \"buy_token_symbol\": \"MLT\",\n    \"chain_id\": 1,\n    \"comer_id\": 1001,\n    \"crowdfunding_contract\": \"0x742d35Cc6634C0532925a3b844Bc454e4438f44e\",\n    \"description\": \"去中心化元宇宙基础设施建设项目\",\n    \"detail\": \"项目包含土地NFT、虚拟商城和社交系统\",\n    \"dex_init_price\": 200,\n    \"dex_router\": \"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D\",\n    \"end_time\": \"2023-12-31T23:59:59Z\",\n    \"id\": 701,\n    \"investor\": {\n        \"buy_token_balance\": 5000,\n        \"buy_token_total\": 15000,\n        \"comer_id\": 2001,\n        \"crowdfunding_id\": 701,\n        \"id\": 301,\n        \"sell_token_balance\": 300,\n        \"sell_token_total\": 1000\n    },\n    \"investors\": 120,\n    \"max_buy_amount\": 10000,\n    \"max_sell_percent\": 30,\n    \"min_buy_amount\": 100,\n    \"pair_address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n    \"poster\": \"https://storage.metaland.xyz/posters/cf_701.jpg\",\n    \"raise_balance\": 850000,\n    \"raise_goal\": 1000000,\n    \"sell_tax\": 5,\n    \"sell_token_balance\": 25000,\n    \"sell_token_contract\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n    \"sell_token_decimals\": 6,\n    \"sell_token_deposit\": 50000,\n    \"sell_token_name\": \"Tether USD\",\n    \"sell_token_supply\": 10000000,\n    \"sell_token_symbol\": \"USDT\",\n    \"start_time\": \"2023-06-01T00:00:00Z\",\n    \"startup\": {\n        \"banner\": \"https://storage.metaland.xyz/startups/banner_701.jpg\",\n        \"chain_id\": 1,\n        \"comer_id\": 1001,\n        \"contract_audit\": \"audit_v3.pdf\",\n        \"id\": 701,\n        \"is_connected\": true,\n        \"kyc\": \"kyc_701.pdf\",\n        \"logo\": \"https://storage.metaland.xyz/startups/logo_701.png\",\n        \"mission\": \"构建开放的元宇宙生态\",\n        \"name\": \"MetaLand Core\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"social_tool\": {\n                    \"id\": 3,\n                    \"logo\": \"https://static.metaland.xyz/socials/discord.png\",\n                    \"name\": \"Discord\"\n                },\n                \"value\": \"discord.gg/metaland\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 7,\n                \"tag\": {\n                    \"id\": 7,\n                    \"name\": \"元宇宙\",\n                    \"type\": 3\n                }\n            }\n        ],\n        \"tx_hash\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n        \"type\": 3\n    },\n    \"startup_id\": 701,\n    \"status\": 1,\n    \"swap_percent\": 20,\n    \"swaps\": [\n        {\n            \"banner\": \"https://storage.metaland.xyz/swaps/swap_701_1.jpg\",\n            \"chain_id\": 1,\n            \"comer_id\": 1001,\n            \"contract_audit\": \"swap_audit_v1.pdf\",\n            \"id\": 1,\n            \"is_connected\": true,\n            \"kyc\": \"kyc_swap_701.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/swaps/logo_701_1.png\",\n            \"mission\": \"土地NFT交易市场\",\n            \"name\": \"LandSwap\",\n            \"on_chain\": true,\n            \"socials\": [\n                {\n                    \"social_tool\": {\n                        \"id\": 4,\n                        \"logo\": \"https://static.metaland.xyz/socials/twitter.png\",\n                        \"name\": \"Twitter\"\n                    },\n                    \"value\": \"@landswap\"\n                }\n            ],\n            \"tags\": [\n                {\n                    \"id\": 8,\n                    \"tag\": {\n                        \"id\": 8,\n                        \"name\": \"NFT\",\n                        \"type\": 4\n                    }\n                }\n            ],\n            \"tx_hash\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n            \"type\": 4\n        }\n    ],\n    \"team_wallet\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n    \"title\": \"MetaLand核心生态建设众筹\",\n    \"tx_hash\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n    \"youtube\": \"https://youtu.be/metaland_demo\"\n}"
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
	var res model.PageData
	var list []crowdfunding.CrowdfundingSwapResponse
	var resString = "[\n    {\n        \"banner\": \"https://storage.metaland.xyz/swap/banner_701.jpg\",\n        \"chain_id\": 1,\n        \"comer_id\": 1001,\n        \"contract_audit\": \"audit_v3.pdf\",\n        \"id\": 701,\n        \"is_connected\": true,\n        \"kyc\": \"kyc_701.pdf\",\n        \"logo\": \"https://storage.metaland.xyz/swap/logo_701.png\",\n        \"mission\": \"构建跨链DEX聚合器\",\n        \"name\": \"MetaSwap\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"social_tool\": {\n                    \"id\": 3,\n                    \"logo\": \"https://static.metaland.xyz/socials/twitter.png\",\n                    \"name\": \"Twitter\"\n                },\n                \"value\": \"@metaswap\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 7,\n                \"tag\": {\n                    \"id\": 7,\n                    \"name\": \"DeFi\",\n                    \"type\": 2\n                }\n            }\n        ],\n        \"tx_hash\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n        \"type\": 3\n    },\n    {\n        \"banner\": \"https://storage.metaland.xyz/swap/banner_702.jpg\",\n        \"chain_id\": 56,\n        \"comer_id\": 1002,\n        \"contract_audit\": \"audit_v4.pdf\",\n        \"id\": 702,\n        \"is_connected\": false,\n        \"kyc\": \"kyc_702.pdf\",\n        \"logo\": \"https://storage.metaland.xyz/swap/logo_702.png\",\n        \"mission\": \"NFT流动性解决方案\",\n        \"name\": \"ArtSwap\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"social_tool\": {\n                    \"id\": 4,\n                    \"logo\": \"https://static.metaland.xyz/socials/discord.png\",\n                    \"name\": \"Discord\"\n                },\n                \"value\": \"discord.gg/artswap\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 12,\n                \"tag\": {\n                    \"id\": 12,\n                    \"name\": \"NFT\",\n                    \"type\": 4\n                }\n            }\n        ],\n        \"tx_hash\": \"0x742d35Cc6634C0532925a3b844Bc454e4438f44e\",\n        \"type\": 4\n    },\n    {\n        \"banner\": \"https://storage.metaland.xyz/swap/banner_703.jpg\",\n        \"chain_id\": 137,\n        \"comer_id\": 1003,\n        \"contract_audit\": \"audit_v5.pdf\",\n        \"id\": 703,\n        \"is_connected\": true,\n        \"kyc\": \"kyc_703.pdf\",\n        \"logo\": \"https://storage.metaland.xyz/swap/logo_703.png\",\n        \"mission\": \"多链稳定币交换协议\",\n        \"name\": \"StableSwap\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"social_tool\": {\n                    \"id\": 5,\n                    \"logo\": \"https://static.metaland.xyz/socials/telegram.png\",\n                    \"name\": \"Telegram\"\n                },\n                \"value\": \"t.me/stableswap\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 9,\n                \"tag\": {\n                    \"id\": 9,\n                    \"name\": \"稳定币\",\n                    \"type\": 3\n                }\n            }\n        ],\n        \"tx_hash\": \"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984\",\n        \"type\": 2\n    }\n]"
	err := json.Unmarshal([]byte(resString), &list)
	if err != nil {
		c.HandleError(err)
		return
	}
	res.List = utility.ConvertToInterfaceSlice(list)
	res.Total = len(list)
	res.Page = 1
	res.Size = 24

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
