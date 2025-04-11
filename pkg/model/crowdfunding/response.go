package crowdfunding

import (
	"time"

	"github.com/shopspring/decimal"
)

// StartupCrowdfundingInfo 企业信息，带可否融资字段
type StartupCrowdfundingInfo struct {
	StartupId     uint64 `json:"startupId"`
	StartupName   string `json:"startupName"`
	CanRaise      bool   `json:"canRaise"`
	OnChain       bool   `json:"onChain"`
	TokenContract string `json:"tokenContract,omitempty"`
}

// PublicItem item
type PublicItem struct {
	CrowdfundingId       uint64             `json:"crowdfundingId"`
	StartupId            uint64             `json:"startupId"`
	ComerId              uint64             `json:"comerId"`
	StartupName          string             `json:"startupName"`
	StartupLogo          string             `json:"startupLogo"`
	StartTime            time.Time          `json:"startTime"`
	EndTime              time.Time          `json:"endTime"`
	RaiseBalance         decimal.Decimal    `json:"raiseBalance"`
	RaiseGoal            decimal.Decimal    `json:"raiseGoal"`
	RaisedPercent        decimal.Decimal    `json:"raisedPercent"`
	BuyPrice             decimal.Decimal    `json:"buyPrice"`
	SwapPercent          decimal.Decimal    `json:"swapPercent"`
	Poster               string             `json:"poster"`
	Status               CrowdfundingStatus `json:"status"`
	CrowdfundingContract string             `json:"crowdfundingContract"`
	ChainId              uint64             `json:"chainId"`
	KYC                  string             `json:"kyc"`
	ContractAudit        string             `json:"contractAudit"`
	SellTokenContract    string             `json:"sellTokenAddress"`
	SellTokenSymbol      string             `json:"sellTokenSymbol"`
	BuyTokenContract     string             `json:"buyTokenAddress"`
	BuyTokenSymbol       string             `json:"buyTokenSymbol"`
	BuyTokenAmount       *decimal.Decimal   `json:"buyTokenAmount"`
}

type PostedItem struct {
	CrowdfundingId       uint64             `json:"crowdfundingId"`
	CrowdfundingContract string             `json:"crowdfundingContract"`
	StartupId            uint64             `json:"startupId"`
	StartupName          string             `json:"startupName"`
	StartupLogo          string             `json:"startupLogo"`
	StartTime            time.Time          `json:"startTime"`
	EndTime              time.Time          `json:"endTime"`
	RaiseBalance         decimal.Decimal    `json:"raiseBalance"`
	RaisedPercent        decimal.Decimal    `json:"raisedPercent"`
	Status               CrowdfundingStatus `json:"status"`
}

type ParticipatedItem struct {
	CrowdfundingId       uint64             `json:"crowdfundingId"`
	CrowdfundingContract string             `json:"crowdfundingContract"`
	StartupId            uint64             `json:"startupId"`
	StartupName          string             `json:"startupName"`
	StartupLogo          string             `json:"startupLogo"`
	RaiseBalance         decimal.Decimal    `json:"raiseBalance"`
	RaisedPercent        decimal.Decimal    `json:"raisedPercent"`
	BuyTokenSymbol       string             `json:"buyTokenSymbol"`
	BuyTokenAmount       decimal.Decimal    `json:"buyTokenAmount"`
	Status               CrowdfundingStatus `json:"status"`
}

// Detail detail
type Detail struct {
	CrowdfundingId       uint64 `json:"crowdfundingId"`
	ChainId              uint64 `json:"chainId"`
	CrowdfundingContract string `json:"crowdfundingContract"`
	TeamWallet           string `json:"teamWallet"`

	SellTokenContract string          `json:"sellTokenContract"`
	SellTokenName     string          `json:"sellTokenName,omitempty"`
	SellTokenSymbol   string          `json:"sellTokenSymbol,omitempty"`
	SellTokenDecimals int             `json:"sellTokenDecimals,omitempty"`
	SellTokenSupply   decimal.Decimal `json:"sellTokenSupply,omitempty"`
	MaxSellPercent    decimal.Decimal `json:"maxSellPercent"`
	SellTax           decimal.Decimal `json:"sellTax"`

	BuyTokenContract string `json:"buyTokenContract"`

	MaxBuyAmount decimal.Decimal `json:"maxBuyAmount"`
	BuyPrice     decimal.Decimal `json:"buyPrice"`
	SwapPercent  decimal.Decimal `json:"swapPercent"`

	RaiseBalance  decimal.Decimal `json:"raiseBalance"`
	RaiseGoal     decimal.Decimal `json:"raiseGoal"`
	RaisedPercent decimal.Decimal `json:"raisedPercent"`

	StartupId uint64    `json:"startupId"`
	ComerId   uint64    `json:"comerId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`

	Poster      string             `json:"poster"`
	Description string             `json:"description"`
	Youtube     string             `json:"youtube"`
	Detail      string             `json:"detail"`
	Status      CrowdfundingStatus `json:"status"`
}

// IBOHistory modification history of IBO
type IBOHistory struct {
	BuyPrice    decimal.Decimal `json:"buyPrice"`
	SwapPercent decimal.Decimal `json:"swapPercent"`
	UpdatedOn   time.Time       `json:"updatedOn"`
}

type InvestmentRecord struct {
	ComerName      string          `json:"comerName"`
	ComerAvatar    string          `json:"comerAvatar"`
	ComerId        uint64          `json:"comerId"`
	CrowdfundingId uint64          `json:"crowdfundingId"`
	Access         SwapAccess      `json:"access"`
	InvestAmount   decimal.Decimal `json:"amount"`
	Time           time.Time       `json:"time"`
}
