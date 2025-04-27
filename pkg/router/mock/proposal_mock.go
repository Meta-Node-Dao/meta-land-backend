package mock

import (
	"ceres/pkg/model"
	"ceres/pkg/model/governance"
	"ceres/pkg/router"
	"encoding/json"
)

func GetProposal(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
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
	var resString = "{\n    \"author_wallet_address\": \"0x1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b\",\n    \"block_number\": 987654,\n    \"chain_id\": 1,\n    \"choices\": [\n        {\n            \"id\": 101,\n            \"item_name\": \"增加流动性池奖励\",\n            \"proposal_id\": 2023,\n            \"seq_num\": 1,\n            \"vote_total\": 15000\n        },\n        {\n            \"id\": 102,\n            \"item_name\": \"调整治理参数\",\n            \"proposal_id\": 2023,\n            \"seq_num\": 2,\n            \"vote_total\": 8500\n        }\n    ],\n    \"comer\": {\n        \"avatar\": \"https://storage.metaland.xyz/avatars/001.png\",\n        \"comer_id\": 1001,\n        \"name\": \"区块链开发者\",\n        \"profile_verified\": true,\n        \"uin\": \"UIN_DEV_2023\"\n    },\n    \"comer_id\": 1001,\n    \"description\": \"关于协议升级的第二阶段治理提案，包含流动性激励和参数优化\",\n    \"discussion_link\": \"https://forum.metaland.xyz/t/proposal-2023-002\",\n    \"end_time\": \"2023-03-01T00:00:00Z\",\n    \"id\": 2023,\n    \"ipfs_hash\": \"QmTq4J7hVd8s6E3bZmK1S2mNpLqR\",\n    \"release_timestamp\": \"2023-02-15T00:00:00Z\",\n    \"start_time\": \"2023-02-20T00:00:00Z\",\n    \"startup\": {\n        \"banner\": \"https://storage.metaland.xyz/startups/banner_001.jpg\",\n        \"chain_id\": 1,\n        \"comer_id\": 1001,\n        \"contract_audit\": \"security_audit_v2.3.pdf\",\n        \"governance_setting\": {\n            \"allow_member\": true,\n            \"comer_id\": 1001,\n            \"id\": 1,\n            \"proposal_threshold\": 1000,\n            \"proposal_validity\": 30,\n            \"startup_id\": 1,\n            \"strategies\": {\n                \"chain_id\": 1,\n                \"dict_value\": \"token_holder\",\n                \"id\": 1,\n                \"setting_id\": 1,\n                \"strategy_name\": \"代币持有者投票\",\n                \"token_contract_address\": \"0x789abcdef0123456789abcdef0123456789abcd\",\n                \"token_min_balance\": 100,\n                \"vote_decimals\": 18,\n                \"vote_symbol\": \"GOV\"\n            },\n            \"vote_symbol\": \"GOV\"\n        },\n        \"id\": 1,\n        \"is_connected\": true,\n        \"kyc\": \"kyc_verified_v3.pdf\",\n        \"logo\": \"https://storage.metaland.xyz/startups/logo_001.png\",\n        \"mission\": \"构建去中心化治理基础设施\",\n        \"name\": \"MetaGovernanceLab\",\n        \"on_chain\": true,\n        \"socials\": [\n            {\n                \"type\": \"discord\",\n                \"url\": \"https://discord.gg/mglab\"\n            },\n            {\n                \"type\": \"medium\",\n                \"url\": \"https://medium.com/mglab\"\n            }\n        ],\n        \"tags\": [\n            {\n                \"id\": 1,\n                \"name\": \"DAO\"\n            },\n            {\n                \"id\": 2,\n                \"name\": \"Governance\"\n            }\n        ],\n        \"tx_hash\": \"0x5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1a2b3c4d5e\",\n        \"type\": 2\n    },\n    \"startup_id\": 1,\n    \"status\": 1,\n    \"title\": \"协议升级第二阶段治理提案\",\n    \"vote_system\": \"multisig\"\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func DeleteProposal(c *router.Context) {

	c.OK("delete proposal successfully!")
}

func VoteProposal(c *router.Context) {
	c.OK("vote proposal successfully!")
}

func GetProposalInvestRecords(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

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

func GetDataDict(c *router.Context) {
	c.OK("dict data")
}
