package crowdfunding

import (
	"ceres/pkg/model"
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type CrowdfundingStatus int

const (
	Pending CrowdfundingStatus = iota
	Upcoming
	Live
	Ended
	Cancelled
	OnChainFailure
)

type ChainInfo struct {
	ChainId uint64 `gorm:"column:chain_id;unique_index:chain_tx_uindex" json:"chainId"`
	TxHash  string `gorm:"tx_hash;unique_index:chain_tx_uindex" json:"txHash"`
}

type Crowdfunding struct {
	model.Base
	ChainInfo
	SellInfo
	BuyInfo
	// update by querying chain???
	CrowdfundingContract string          `gorm:"column:crowdfunding_contract" json:"crowdfundingContract,omitempty"`
	StartupId            uint64          `gorm:"column:startup_id" json:"startupId"`
	ComerId              uint64          `gorm:"column:comer_id" json:"comerId"`
	RaiseGoal            decimal.Decimal `gorm:"column:raise_goal" json:"raiseGoal"`
	// update by querying chain???
	RaiseBalance decimal.Decimal `gorm:"column:raise_balance" json:"raiseBalance"`

	TeamWallet  string             `gorm:"column:team_wallet" json:"teamWallet"`
	SwapPercent decimal.Decimal    `gorm:"column:swap_percent" json:"swapPercent"`
	StartTime   time.Time          `gorm:"column:start_time" json:"startTime"`
	EndTime     time.Time          `gorm:"column:end_time" json:"endTime"`
	Poster      string             `gorm:"column:poster" json:"poster"`
	Youtube     string             `gorm:"column:youtube" json:"youtube"`
	Detail      string             `gorm:"column:detail" json:"detail"`
	Description string             `gorm:"column:description" json:"description"`
	Status      CrowdfundingStatus `gorm:"column:status" json:"status"`
}

func (c Crowdfunding) Json() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type SellInfo struct {
	SellTokenContract string          `gorm:"column:sell_token_contract" json:"sellTokenContract"`
	SellTokenName     string          `gorm:"column:sell_token_name" json:"sellTokenName,omitempty"`
	SellTokenSymbol   string          `gorm:"column:sell_token_symbol" json:"sellTokenSymbol,omitempty"`
	SellTokenDecimals int             `gorm:"column:sell_token_decimals" json:"sellTokenDecimals,omitempty"`
	SellTokenSupply   decimal.Decimal `gorm:"column:sell_token_supply" json:"sellTokenSupply,omitempty"`
	SellTokenDeposit  decimal.Decimal `gorm:"column:sell_token_deposit" json:"sellTokenDeposit"`
	SellTokenBalance  decimal.Decimal `gorm:"column:sell_token_balance" json:"sellTokenBalance"`
	MaxSellPercent    decimal.Decimal `gorm:"column:max_sell_percent" json:"maxSellPercent"`
	SellTax           decimal.Decimal `gorm:"column:sell_tax" json:"sellTax"`
}

type BuyInfo struct {
	BuyTokenContract string          `gorm:"column:buy_token_contract" json:"buyTokenContract"`
	BuyTokenName     string          `gorm:"column:buy_token_name" json:"buyTokenName,omitempty"`
	BuyTokenSymbol   string          `gorm:"column:buy_token_symbol" json:"buyTokenSymbol,omitempty"`
	BuyTokenDecimals int             `gorm:"column:buy_token_decimals" json:"buyTokenDecimals,omitempty"`
	BuyTokenSupply   decimal.Decimal `gorm:"column:buy_token_supply" json:"buyTokenSupply,omitempty"`
	// IBO rate
	BuyPrice     decimal.Decimal `gorm:"column:buy_price" json:"buyPrice"`
	MaxBuyAmount decimal.Decimal `gorm:"column:max_buy_amount" json:"maxBuyAmount"`
}

func (c Crowdfunding) TableName() string {
	return "crowdfunding"
}

type CrowdfundingSwapStatus int
type SwapAccess int

func (receiver SwapAccess) String() string {
	switch receiver {
	case Invest:
		return "Invest"
	case Withdraw:
		return "Withdraw"
	default:
		panic("unsupported swapAccess")
	}
}

const (
	SwapPending CrowdfundingSwapStatus = iota
	SwapSuccess
	SwapFailure
)
const (
	Invest SwapAccess = iota + 1
	Withdraw
)

type CrowdfundingSwap struct {
	model.RelationBase
	ChainInfo
	Timestamp       time.Time              `gorm:"timestamp" json:"timestamp"`
	Status          CrowdfundingSwapStatus `gorm:"status" json:"status"`
	CrowdfundingId  uint64                 `gorm:"crowdfunding_id" json:"crowdfundingId"`
	ComerId         uint64                 `gorm:"comer_id" json:"comerId"`
	Access          SwapAccess             `gorm:"access" json:"access"`
	BuyTokenSymbol  string                 `gorm:"buy_token_symbol" json:"buyTokenSymbol"`
	BuyTokenAmount  decimal.Decimal        `gorm:"buy_token_amount" json:"buyTokenAmount"`
	SellTokenSymbol string                 `gorm:"sell_token_symbol" json:"sellTokenSymbol"`
	SellTokenAmount decimal.Decimal        `gorm:"sell_token_amount" json:"sellTokenAmount"`
	Price           decimal.Decimal        `gorm:"price" json:"price"`
}

func (c CrowdfundingSwap) TableName() string {
	return "crowdfunding_swap"
}

type IboRateHistory struct {
	model.RelationBase
	CrowdfundingId uint64    `gorm:"crowdfunding_id" json:"crowdfundingId"`
	EndTime        time.Time `gorm:"end_time" json:"endTime"`
	//BuyTokenSymbol  string          `gorm:"buy_token_symbol" json:"buyTokenSymbol"`
	MaxBuyAmount   decimal.Decimal `gorm:"max_buy_amount" json:"maxBuyAmount"`
	MaxSellPercent decimal.Decimal `gorm:"max_sell_percent" json:"maxSellPercent"`
	//SellTokenSymbol string          `gorm:"sell_token_symbol" json:"sellTokenSymbol"`
	BuyPrice    decimal.Decimal `gorm:"buy_price" json:"buyPrice"`
	SwapPercent decimal.Decimal `gorm:"swap_percent" json:"swapPercent"`
}

func (receiver IboRateHistory) TableName() string {
	return "crowdfunding_ibo_rate"
}

type Investor struct {
	model.RelationBase
	CrowdfundingId uint64 `gorm:"crowdfunding_id" json:"crowdfundingId"`
	ComerId        uint64 `gorm:"comer_id" json:"comerId"`
	// total bought token
	BuyTokenTotal decimal.Decimal `gorm:"buy_token_total" json:"buyTokenTotal"`
	// current balance of bought token
	BuyTokenBalance decimal.Decimal `gorm:"buy_token_balance" json:"buyTokenBalance"`
	// total sold token
	SellTokenTotal decimal.Decimal `gorm:"sell_token_total" json:"sellTokenTotal"`
	// current balance sold token
	SellTokenBalance decimal.Decimal `gorm:"sell_token_balance" json:"sellTokenBalance"`
}

func (i *Investor) Invest(access SwapAccess, buyTokenAmount, sellTokenAmount decimal.Decimal) {
	if access == Invest {
		i.BuyTokenTotal = i.BuyTokenTotal.Add(buyTokenAmount)
		i.SellTokenTotal = i.SellTokenTotal.Add(sellTokenAmount)
		i.BuyTokenBalance = i.BuyTokenBalance.Add(buyTokenAmount)
		i.SellTokenBalance = i.SellTokenBalance.Add(sellTokenAmount)
	} else {
		i.BuyTokenBalance = i.BuyTokenBalance.Sub(buyTokenAmount)
		i.SellTokenBalance = i.SellTokenBalance.Sub(sellTokenAmount)
	}
}

func (i *Investor) TableName() string {
	return "crowdfunding_investor"
}
