package governance

import (
	"ceres/pkg/model/governance"
	"ceres/pkg/router"
	"encoding/json"
)

func GetGovernanceSetting(c *router.Context) {
	var res governance.GovernanceSettingDetailResponse
	var resString = "{\n    \"admins\": [\n        {\n            \"address\": \"0x9c0d1e2f3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b\",\n            \"id\": 3001,\n            \"setting_id\": 5001\n        },\n        {\n            \"address\": \"0x1e2f3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d\",\n            \"id\": 3002,\n            \"setting_id\": 5001\n        }\n    ],\n    \"allow_member\": false,\n    \"comer_id\": 2005,\n    \"id\": 5001,\n    \"proposal_threshold\": 10000,\n    \"proposal_validity\": 604800,\n    \"startup_id\": 7001,\n    \"strategies\": [\n        {\n            \"chain_id\": 56,\n            \"dict_value\": \"bsc_holder\",\n            \"id\": 55,\n            \"setting_id\": 5001,\n            \"strategy_name\": \"BSC持币者治理\",\n            \"token_contract_address\": \"0x5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1a2b3c4d\",\n            \"token_min_balance\": 500,\n            \"vote_decimals\": 18,\n            \"vote_symbol\": \"BSCGOV\"\n        },\n        {\n            \"chain_id\": 42161,\n            \"dict_value\": \"nft_holder\",\n            \"id\": 56,\n            \"setting_id\": 5001,\n            \"strategy_name\": \"NFT持有者投票\",\n            \"token_contract_address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n            \"token_min_balance\": 1,\n            \"vote_decimals\": 0,\n            \"vote_symbol\": \"NFT-VOTE\"\n        }\n    ],\n    \"vote_symbol\": \"META-GOV\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func CreateGovernanceSetting(c *router.Context) {
	c.OK("create governance setting successfully!")
}

//func CreateGovernanceSetting(ctx *router.Context) {
//	startupId, err := uintPathVariable(ctx, "startupID")
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	if comerId == 0 {
//		ctx.HandleError(errors.New("invalid comerId"))
//		return
//	}
//	var request governance.CreateOrUpdateGovernanceSettingRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if _, err := service.CreateStartupGovernanceSetting(comerId, startupId, request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//// GetGovernanceSetting get startup governance setting
//func GetGovernanceSetting(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "startupID", func(startupId uint64) {
//		detail, err := service.GetStartupGovernanceSetting(startupId)
//		if err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(detail)
//	})
//}
