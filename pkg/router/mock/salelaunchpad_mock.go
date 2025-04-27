package mock

import (
	"ceres/pkg/model"
	"ceres/pkg/model/crowdfunding"
	"ceres/pkg/router"
	"encoding/json"
)

func GetSaleLaunchPad(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func UpdateSaleLaunchPad(c *router.Context) {
	var message model.MessageResponse
	message.Message = "update sale launchpad successful!"
	c.OK(message)
}

func CreateSaleLaunchPad(c *router.Context) {
	var message model.MessageResponse
	message.Message = "create sale launchpad successful!"
	c.OK(message)
}

func GetSaleLaunchPadSupplyDex(c *router.Context) {
	var res crowdfunding.SaleLaunchpadResponse
	var resString = "{\n    \"chain_id\": 1,\n    \"comer_id\": 12345,\n    \"contract_address\": \"0x1234567890abcdef\",\n    \"cycle\": 30,\n    \"cycle_release\": 10,\n    \"description\": \"Sample crowdfunding project description\",\n    \"detail\": \"Detailed information about the project\",\n    \"dex_init_price\": 100,\n    \"dex_pair_address\": \"0x0987654321fedcba\",\n    \"dex_router\": \"0xrouteraddress123\",\n    \"ended_at\": 1672502400,\n    \"first_release\": 20,\n    \"hard_cap\": 1000000,\n    \"id\": 1,\n    \"invest_token_balance\": 500000,\n    \"invest_token_contract\": \"0xinvesttoken123\",\n    \"invest_token_decimals\": 18,\n    \"invest_token_name\": \"Investment Token\",\n    \"invest_token_supply\": 10000000,\n    \"invest_token_symbol\": \"INV\",\n    \"investor\": {\n        \"buy_token_balance\": 2500,\n        \"buy_token_total\": 5000,\n        \"comer_id\": 12345,\n        \"crowdfunding_id\": 1,\n        \"id\": 1,\n        \"sell_token_balance\": 1000,\n        \"sell_token_total\": 2000\n    },\n    \"investors\": 100,\n    \"liquidity_rate\": 50,\n    \"max_invest_amount\": 5000,\n    \"min_invest_amount\": 100,\n    \"poster\": \"https://example.com/poster.jpg\",\n    \"presale_price\": 50,\n    \"presale_token_balance\": 200000,\n    \"presale_token_contract\": \"0xpresaletoken123\",\n    \"presale_token_decimals\": 18,\n    \"presale_token_deposit\": 100000,\n    \"presale_token_name\": \"Presale Token\",\n    \"presale_token_supply\": 5000000,\n    \"presale_token_symbol\": \"PRE\",\n    \"soft_cap\": 500000,\n    \"started_at\": 1672416000,\n    \"startup\": {\n        \"banner\": \"https://example.com/banner.jpg\",\n        \"chain_id\": 1,\n        \"comer_id\": 12345,\n        \"contract_audit\": \"audit-report.pdf\",\n        \"id\": 1,\n        \"is_connected\": true,\n        \"kyc\": \"kyc-document.pdf\",\n        \"logo\": \"https://example.com/logo.png\",\n        \"mission\": \"Company mission statement\",\n        \"name\": \"Startup Inc.\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"type\": \"twitter\",\n                \"url\": \"https://twitter.com/startup\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 1,\n                \"name\": \"Blockchain\"\n            }\n        ],\n        \"tx_hash\": \"0xstartuptxhash123\",\n        \"type\": 1\n    },\n    \"startup_id\": 1,\n    \"status\": 1,\n    \"swaps\": [\n        {\n            \"amount\": 1000,\n            \"chain_id\": 1,\n            \"comer\": {\n                \"avatar\": \"https://example.com/avatar.jpg\",\n                \"comer_id\": 12345,\n                \"name\": \"John Doe\",\n                \"profile_verified\": true,\n                \"uin\": \"UIN12345\"\n            },\n            \"comer_id\": 12345,\n            \"id\": 1,\n            \"sale_launchpad_id\": 1,\n            \"timestamp\": 1672416000,\n            \"token_symbol\": \"INV\",\n            \"tx_hash\": \"0xswaptxhash123\",\n            \"type\": 1\n        }\n    ],\n    \"team_wallet\": \"0xteamwallet123\",\n    \"title\": \"Sample Crowdfunding Project\",\n    \"tx_hash\": \"0xprojecttxhash123\",\n    \"youtube\": \"https://youtu.be/samplevideo\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func GetSaleLaunchPadInfo(c *router.Context) {
	var res crowdfunding.SaleLaunchpadResponse
	var resString = "{\n    \"chain_id\": 1,\n    \"comer_id\": 12345,\n    \"contract_address\": \"0x1234567890abcdef\",\n    \"cycle\": 30,\n    \"cycle_release\": 10,\n    \"description\": \"Sample crowdfunding project description\",\n    \"detail\": \"Detailed information about the project\",\n    \"dex_init_price\": 100,\n    \"dex_pair_address\": \"0x0987654321fedcba\",\n    \"dex_router\": \"0xrouteraddress123\",\n    \"ended_at\": 1672502400,\n    \"first_release\": 20,\n    \"hard_cap\": 1000000,\n    \"id\": 1,\n    \"invest_token_balance\": 500000,\n    \"invest_token_contract\": \"0xinvesttoken123\",\n    \"invest_token_decimals\": 18,\n    \"invest_token_name\": \"Investment Token\",\n    \"invest_token_supply\": 10000000,\n    \"invest_token_symbol\": \"INV\",\n    \"investor\": {\n        \"buy_token_balance\": 2500,\n        \"buy_token_total\": 5000,\n        \"comer_id\": 12345,\n        \"crowdfunding_id\": 1,\n        \"id\": 1,\n        \"sell_token_balance\": 1000,\n        \"sell_token_total\": 2000\n    },\n    \"investors\": 100,\n    \"liquidity_rate\": 50,\n    \"max_invest_amount\": 5000,\n    \"min_invest_amount\": 100,\n    \"poster\": \"https://example.com/poster.jpg\",\n    \"presale_price\": 50,\n    \"presale_token_balance\": 200000,\n    \"presale_token_contract\": \"0xpresaletoken123\",\n    \"presale_token_decimals\": 18,\n    \"presale_token_deposit\": 100000,\n    \"presale_token_name\": \"Presale Token\",\n    \"presale_token_supply\": 5000000,\n    \"presale_token_symbol\": \"PRE\",\n    \"soft_cap\": 500000,\n    \"started_at\": 1672416000,\n    \"startup\": {\n        \"banner\": \"https://example.com/banner.jpg\",\n        \"chain_id\": 1,\n        \"comer_id\": 12345,\n        \"contract_audit\": \"audit-report.pdf\",\n        \"id\": 1,\n        \"is_connected\": true,\n        \"kyc\": \"kyc-document.pdf\",\n        \"logo\": \"https://example.com/logo.png\",\n        \"mission\": \"Company mission statement\",\n        \"name\": \"Startup Inc.\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"type\": \"twitter\",\n                \"url\": \"https://twitter.com/startup\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 1,\n                \"name\": \"Blockchain\"\n            }\n        ],\n        \"tx_hash\": \"0xstartuptxhash123\",\n        \"type\": 1\n    },\n    \"startup_id\": 1,\n    \"status\": 1,\n    \"swaps\": [\n        {\n            \"amount\": 1000,\n            \"chain_id\": 1,\n            \"comer\": {\n                \"avatar\": \"https://example.com/avatar.jpg\",\n                \"comer_id\": 12345,\n                \"name\": \"John Doe\",\n                \"profile_verified\": true,\n                \"uin\": \"UIN12345\"\n            },\n            \"comer_id\": 12345,\n            \"id\": 1,\n            \"sale_launchpad_id\": 1,\n            \"timestamp\": 1672416000,\n            \"token_symbol\": \"INV\",\n            \"tx_hash\": \"0xswaptxhash123\",\n            \"type\": 1\n        }\n    ],\n    \"team_wallet\": \"0xteamwallet123\",\n    \"title\": \"Sample Crowdfunding Project\",\n    \"tx_hash\": \"0xprojecttxhash123\",\n    \"youtube\": \"https://youtu.be/samplevideo\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func GetSaleLaunchPadHistoryRecords(c *router.Context) {
	var res crowdfunding.SaleLaunchpadHistoryResponse
	var resString = "{\n    \"amount\": 500,\n    \"chain_id\": 1,\n    \"comer\": {\n        \"avatar\": \"https://example.com/avatars/123.jpg\",\n        \"comer_id\": 456,\n        \"name\": \"Alice Smith\",\n        \"profile_verified\": true,\n        \"uin\": \"UIN_ALICE_2023\"\n    },\n    \"comer_id\": 456,\n    \"id\": 789,\n    \"sale_launchpad_id\": 13579,\n    \"timestamp\": 1672531200,\n    \"token_symbol\": \"MLT\",\n    \"tx_hash\": \"0x4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a\",\n    \"type\": 2\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func GetSaleLaunchPadTransferLpSign(c *router.Context) {
	var res crowdfunding.SignResponse
	res.Data = "0x1234567890abcdef"
	res.Sign = "0x1234567890abcdef"
	c.OK(res)
}
