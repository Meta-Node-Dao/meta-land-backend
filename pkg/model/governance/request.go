package governance

import (
	"ceres/pkg/model"
	"github.com/shopspring/decimal"
)

type CreateOrUpdateGovernanceSettingRequest struct {
	SettingRequest
	Strategies []StrategyRequest `json:"strategies"`
	Admins     []AdminRequest    `json:"admins"`
}

type SettingRequest struct {
	VoteSymbol        string          `gorm:"column:vote_symbol" json:"voteSymbol"`
	AllowMember       bool            `gorm:"column:allow_member" json:"allowMember"`
	ProposalThreshold decimal.Decimal `gorm:"column:proposal_threshold" json:"proposalThreshold"`
	ProposalValidity  decimal.Decimal `gorm:"column:proposal_validity" json:"proposalValidity"`
}

type StrategyRequest struct {
	DictValue            string          `json:"dictValue" binding:"required"`
	StrategyName         string          `json:"strategyName" binding:"required"`
	ChainId              uint64          `json:"chainId" binding:"required"`
	VoteSymbol           string          `json:"voteSymbol"`
	TokenContractAddress string          `json:"tokenContractAddress" binding:"required"`
	VoteDecimals         int             `json:"voteDecimals"`
	TokenMinBalance      decimal.Decimal `json:"tokenMinBalance"`
}

type AdminRequest struct {
	WalletAddress string `json:"walletAddress"`
}

type VoteRequest struct {
	VoterWalletAddress string          `json:"voterWalletAddress"`
	ChoiceItemId       uint64          `json:"choiceItemId"`
	ChoiceItemName     string          `json:"choiceItemName"`
	Votes              decimal.Decimal `json:"votes"`
	IpfsHash           string          `json:"ipfsHash"`
}

type ProposalListRequest struct {
	*model.Pagination
	States []ProposalStatus `json:"states"`
}

type CreateProposalRequest struct {
	GovernanceProposalInfo
	Choices []ProposalChoice `json:"choices"`
}
