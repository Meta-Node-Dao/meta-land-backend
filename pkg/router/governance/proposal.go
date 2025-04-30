package governance

import (
	"ceres/pkg/initialization/utility"
	"ceres/pkg/model"
	"ceres/pkg/model/governance"
	"ceres/pkg/router"
	"encoding/json"
	"fmt"
)

func GetProposal(c *router.Context) {
	var res model.PageData
	var listString = "[\n    {\n        \"author_wallet_address\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n        \"block_number\": 15678943,\n        \"chain_id\": 1,\n        \"comer\": {\n            \"activation\": true,\n            \"address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n            \"avatar\": \"https://storage.metaland.xyz/avatars/1001.png\",\n            \"id\": 1001,\n            \"name\": \"社区治理员\"\n        },\n        \"comer_id\": 1001,\n        \"description\": \"提案升级DAO治理合约，增加多签功能\",\n        \"discussion_link\": \"https://forum.metaland.xyz/t/701\",\n        \"end_time\": \"2023-12-31T23:59:59Z\",\n        \"id\": 701,\n        \"ipfs_hash\": \"QmNq4oyFnxD5QeGozZ6M9vRqJ6XgS7JYJ8xzvJFqZkLQ9A\",\n        \"max_votes\": 1500,\n        \"max_votes_choice_item\": \"选项A\",\n        \"release_timestamp\": \"2023-06-01T00:00:00Z\",\n        \"start_time\": \"2023-06-15T00:00:00Z\",\n        \"startup\": {\n            \"banner\": \"https://storage.metaland.xyz/startups/banner_701.jpg\",\n            \"chain_id\": 1,\n            \"comer_id\": 1001,\n            \"contract_audit\": \"audit_v5.pdf\",\n            \"governance_setting\": {\n                \"allow_member\": true,\n                \"proposal_threshold\": 100,\n                \"strategies\": {\n                    \"chain_id\": 1,\n                    \"strategy_name\": \"代币加权投票\",\n                    \"token_contract_address\": \"0x742d35Cc6634C0532925a3b844Bc454e4438f44e\",\n                    \"token_min_balance\": 1000\n                },\n                \"vote_symbol\": \"GOV\"\n            },\n            \"id\": 701,\n            \"is_connected\": true,\n            \"kyc\": \"kyc_701.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/startups/logo_701.png\",\n            \"mission\": \"构建去中心化自治组织\",\n            \"name\": \"MetaDAO\",\n            \"on_chain\": true,\n            \"socials\": [\n                {\n                    \"social_tool\": {\n                        \"id\": 6,\n                        \"logo\": \"https://static.metaland.xyz/socials/telegram.png\",\n                        \"name\": \"Telegram\"\n                    },\n                    \"value\": \"t.me/metadao\"\n                }\n            ],\n            \"tags\": [\n                {\n                    \"id\": 15,\n                    \"tag\": {\n                        \"id\": 15,\n                        \"name\": \"DAO\",\n                        \"type\": 5\n                    }\n                }\n            ],\n            \"tx_hash\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n            \"type\": 5\n        },\n        \"startup_id\": 701,\n        \"status\": 1,\n        \"title\": \"DAO治理合约升级提案\",\n        \"vote_system\": \"quadratic\"\n    },\n    {\n        \"author_wallet_address\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n        \"block_number\": 15689234,\n        \"chain_id\": 137,\n        \"comer\": {\n            \"activation\": true,\n            \"address\": \"0xEA674DeDe5fE460663539C1bB0365bFfE9d444f8\",\n            \"avatar\": \"https://storage.metaland.xyz/avatars/1002.png\",\n            \"id\": 1002,\n            \"name\": \"生态建设者\"\n        },\n        \"comer_id\": 1002,\n        \"description\": \"调整流动性挖矿奖励分配比例\",\n        \"discussion_link\": \"https://forum.metaland.xyz/t/702\",\n        \"end_time\": \"2024-03-31T23:59:59Z\",\n        \"id\": 702,\n        \"ipfs_hash\": \"QmXyZ4oyFnxD5QeGozZ6M9vRqJ6XgS7JYJ8xzvJFqZkLQ9B\",\n        \"max_votes\": 2500,\n        \"max_votes_choice_item\": \"方案C\",\n        \"release_timestamp\": \"2023-09-01T00:00:00Z\",\n        \"start_time\": \"2023-09-15T00:00:00Z\",\n        \"startup\": {\n            \"banner\": \"https://storage.metaland.xyz/startups/banner_702.jpg\",\n            \"chain_id\": 137,\n            \"comer_id\": 1002,\n            \"contract_audit\": \"audit_v6.pdf\",\n            \"governance_setting\": {\n                \"allow_member\": false,\n                \"proposal_threshold\": 500,\n                \"strategies\": {\n                    \"chain_id\": 137,\n                    \"strategy_name\": \"LP代币投票\",\n                    \"token_contract_address\": \"0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270\",\n                    \"token_min_balance\": 500\n                },\n                \"vote_symbol\": \"POL\"\n            },\n            \"id\": 702,\n            \"is_connected\": false,\n            \"kyc\": \"kyc_702.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/startups/logo_702.png\",\n            \"mission\": \"优化DeFi协议流动性\",\n            \"name\": \"PolyFarm\",\n            \"on_chain\": true,\n            \"socials\": [\n                {\n                    \"social_tool\": {\n                        \"id\": 7,\n                        \"logo\": \"https://static.metaland.xyz/socials/discord.png\",\n                        \"name\": \"Discord\"\n                    },\n                    \"value\": \"discord.gg/polyfarm\"\n                }\n            ],\n            \"tags\": [\n                {\n                    \"id\": 18,\n                    \"tag\": {\n                        \"id\": 18,\n                        \"name\": \"DeFi\",\n                        \"type\": 2\n                    }\n                }\n            ],\n            \"tx_hash\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7\",\n            \"type\": 3\n        },\n        \"startup_id\": 702,\n        \"status\": 2,\n        \"title\": \"流动性挖矿参数调整提案\",\n        \"vote_system\": \"weighted\"\n    }\n]"
	var list []governance.GovernanceBasicResponse
	err := json.Unmarshal([]byte(listString), &list)
	if err != nil {
		c.HandleError(err)
		return
	}
	res.Total = len(list)
	res.Page = 1
	res.Size = 24
	res.List = utility.ConvertToInterfaceSlice(list)
	c.OK(res)
}

func CreateProposal(c *router.Context) {
	var res governance.CreateProposalResponse
	var resString = "{\n    \"author_comer_id\": 12345,\n    \"author_wallet_address\": \"0x1234567890abcdef1234567890abcdef12345678\",\n    \"block_number\": 123456,\n    \"chain_id\": 1,\n    \"choices\": [\n        {\n            \"id\": 1,\n            \"item_name\": \"Option A: Protocol Upgrade\",\n            \"proposal_id\": 1,\n            \"seq_num\": 1,\n            \"vote_total\": 1500\n        },\n        {\n            \"id\": 2,\n            \"item_name\": \"Option B: Parameter Adjustment\",\n            \"proposal_id\": 1,\n            \"seq_num\": 2,\n            \"vote_total\": 800\n        },\n        {\n            \"id\": 3,\n            \"item_name\": \"Option C: No Change\",\n            \"proposal_id\": 1,\n            \"seq_num\": 3,\n            \"vote_total\": 300\n        }\n    ],\n    \"description\": \"Proposal for DAO governance parameter optimization\",\n    \"discussion_link\": \"https://forum.example.com/proposals/1\",\n    \"end_time\": 1672502400,\n    \"ipfs_hash\": \"QmXzw5FwT9q4J7hVd8s6E3bZmK1S2mNpLqR\",\n    \"release_timestamp\": 1672416000,\n    \"start_time\": 1672339200,\n    \"startup_id\": 987,\n    \"title\": \"DAO Governance Parameter Optimization Proposal v1.0\",\n    \"vote_system\": \"snapshot\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func GetProposalInfo(c *router.Context) {
	var res governance.GovernanceResponse
	var resString = "{\n    \"author_wallet_address\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n    \"block_number\": 16789023,\n    \"chain_id\": 1,\n    \"choices\": [\n        {\n            \"id\": 1,\n            \"item_name\": \"赞成提案\",\n            \"proposal_id\": 701,\n            \"seq_num\": 1,\n            \"vote_total\": 850\n        },\n        {\n            \"id\": 2,\n            \"item_name\": \"反对提案\",\n            \"proposal_id\": 701,\n            \"seq_num\": 2,\n            \"vote_total\": 150\n        }\n    ],\n    \"comer\": {\n        \"activation\": true,\n        \"address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n        \"avatar\": \"https://storage.metaland.xyz/avatars/1001.png\",\n        \"id\": 1001,\n        \"name\": \"提案发起人\"\n    },\n    \"comer_id\": 1001,\n    \"description\": \"DAO治理系统升级提案，增加多重签名功能\",\n    \"discussion_link\": \"https://forum.metaland.xyz/t/701\",\n    \"end_time\": \"2023-12-31T23:59:59Z\",\n    \"id\": 701,\n    \"ipfs_hash\": \"QmXyZ4oyFnxD5QeGozZ6M9vRqJ6XgS7JYJ8xzvJFqZkLQ9B\",\n    \"release_timestamp\": \"2023-06-01T00:00:00Z\",\n    \"start_time\": \"2023-06-15T00:00:00Z\",\n    \"startup\": {\n        \"banner\": \"https://storage.metaland.xyz/startups/banner_701.jpg\",\n        \"chain_id\": 1,\n        \"comer_id\": 1001,\n        \"contract_audit\": \"audit_v5.pdf\",\n        \"governance_setting\": {\n            \"allow_member\": true,\n            \"proposal_threshold\": 1000,\n            \"strategies\": {\n                \"chain_id\": 1,\n                \"strategy_name\": \"代币持有量投票\",\n                \"token_contract_address\": \"0x742d35Cc6634C0532925a3b844Bc454e4438f44e\",\n                \"token_min_balance\": 500\n            },\n            \"vote_symbol\": \"GOV\"\n        },\n        \"id\": 701,\n        \"is_connected\": true,\n        \"kyc\": \"kyc_701.pdf\",\n        \"logo\": \"https://storage.metaland.xyz/startups/logo_701.png\",\n        \"mission\": \"构建去中心化治理生态\",\n        \"name\": \"MetaGovernance\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"social_tool\": {\n                    \"id\": 3,\n                    \"logo\": \"https://static.metaland.xyz/socials/discord.png\",\n                    \"name\": \"Discord\"\n                },\n                \"value\": \"discord.gg/metagov\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 15,\n                \"tag\": {\n                    \"id\": 15,\n                    \"name\": \"DAO\",\n                    \"type\": 5\n                }\n            }\n        ],\n        \"tx_hash\": \"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c\",\n        \"type\": 5\n    },\n    \"startup_id\": 701,\n    \"status\": 1,\n    \"title\": \"DAO治理系统升级提案\",\n    \"vote_system\": \"quadratic\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		fmt.Println("here")
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func DeleteProposal(c *router.Context) {
	c.OK("delete proposal successfully!")
}

func GetDataDict(c *router.Context) {
	c.OK("dict data")
}

//func CreateProposal(ctx *router.Context) {
//	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	if comerId == 0 {
//		ctx.HandleError(errors.New("invalid comerId"))
//		return
//	}
//	var request governance.CreateProposalRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if err := service.CreateProposal(comerId, &request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func GetProposal(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
//		info, err := service.GetProposal(proposalId)
//		if err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(info)
//	})
//}
//
//func DeleteProposal(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
//		comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//		if comerId == 0 {
//			ctx.HandleError(errors.New("invalid comer id"))
//			return
//		}
//		if err := service.DeleteProposal(comerId, proposalId); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(nil)
//	})
//}
//
//func PublicList(ctx *router.Context) {
//	var request governance.ProposalListRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if err := service.SelectProposalPublicList(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(request.Pagination)
//}
//
//func StartupProposalList(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "startupID", func(startupId uint64) {
//		var pagination model.Pagination
//		if err := ctx.ShouldBindJSON(&pagination); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		if err := service.GetStartupProposalList(startupId, &pagination); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(pagination)
//	})
//}
//
//func ComerPostProposalList(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "comerID", func(comerId uint64) {
//		var pagination model.Pagination
//		if err := ctx.ShouldBindJSON(&pagination); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		if err := service.GetComerPostProposalList(comerId, &pagination); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(pagination)
//	})
//}
//
//func ComerParticipateProposalList(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "comerID", func(comerId uint64) {
//		var pagination model.Pagination
//		if err := ctx.ShouldBindJSON(&pagination); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		if err := service.GetComerParticipateProposalList(comerId, &pagination); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(pagination)
//	})
//}
