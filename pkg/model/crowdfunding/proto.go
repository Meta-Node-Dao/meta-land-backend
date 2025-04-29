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

// Crowdfunding 众筹项目表结构
type Crowdfunding struct {
	model.Base
	ChainInfo
	SellInfo
	BuyInfo
	CrowdfundingContract string          `gorm:"column:crowdfunding_contract" json:"crowdfunding_contract"`     // 众筹合约地址
	StartupID            uint64          `gorm:"column:startup_id" json:"startup_id"`                           // 初创公司ID
	ComerID              uint64          `gorm:"column:comer_id" json:"comer_id"`                               // 创始人ID
	RaiseGoal            decimal.Decimal `gorm:"column:raise_goal;type:decimal(38,18)" json:"raise_goal"`       // 募资目标
	RaiseBalance         decimal.Decimal `gorm:"column:raise_balance;type:decimal(38,18)" json:"raise_balance"` // 已募资金额
	TeamWallet           string          `gorm:"column:team_wallet" json:"team_wallet"`                         // 团队钱包地址
	SwapPercent          decimal.Decimal `gorm:"column:swap_percent" json:"swap_percent"`                       // 兑换百分比
	StartTime            time.Time       `gorm:"column:start_time" json:"start_time"`                           // 开始时间
	EndTime              time.Time       `gorm:"column:end_time" json:"end_time"`                               // 结束时间
	Poster               string          `gorm:"column:poster" json:"poster"`                                   // 海报URL
	Youtube              string          `gorm:"column:youtube" json:"youtube"`                                 // YouTube链接
	Detail               string          `gorm:"column:detail" json:"detail"`                                   // 详情URL
	Description          string          `gorm:"column:description" json:"description"`                         // 项目描述
	Status               int8            `gorm:"column:status;default:0" json:"status"`                         // 状态: 0-待定 1-即将开始 2-进行中 3-已结束 4-已取消 5-失败
}

func (c Crowdfunding) Json() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (c Crowdfunding) TableName() string {
	return "crowdfunding"
}

type ChainInfo struct {
	ChainId uint64 `gorm:"column:chain_id;uniqueIndex:chain_tx_uindex" json:"chain_id"` // 链ID
	TxHash  string `gorm:"column:tx_hash;uniqueIndex:chain_tx_uindex" json:"tx_hash"`   // 交易哈希
}

type SellInfo struct {
	SellTokenContract string          `gorm:"column:sell_token_contract" json:"sell_token_contract"`                   // 出售代币合约地址
	SellTokenName     string          `gorm:"column:sell_token_name" json:"sell_token_name"`                           // 出售代币名称
	SellTokenSymbol   string          `gorm:"column:sell_token_symbol" json:"sell_token_symbol"`                       // 出售代币符号
	SellTokenDecimals int             `gorm:"column:sell_token_decimals" json:"sell_token_decimals"`                   // 出售代币精度
	SellTokenSupply   decimal.Decimal `gorm:"column:sell_token_supply;type:decimal(38,18)" json:"sell_token_supply"`   // 出售代币总量
	SellTokenDeposit  decimal.Decimal `gorm:"column:sell_token_deposit;type:decimal(38,18)" json:"sell_token_deposit"` // 出售代币质押量
	SellTokenBalance  decimal.Decimal `gorm:"column:sell_token_balance;type:decimal(38,18)" json:"sell_token_balance"` // 出售代币余额
	MaxSellPercent    decimal.Decimal `gorm:"column:max_sell_percent" json:"maxSellPercent"`
	SellTax           decimal.Decimal `gorm:"column:sell_tax" json:"sellTax"`
	//MaxSellPercent    float64         `gorm:"column:max_sell_percent" json:"max_sell_percent"`                         // 最大出售百分比
	//SellTax           float64         `gorm:"column:sell_tax" json:"sell_tax"`                                         // 出售税率
}

type BuyInfo struct {
	BuyTokenContract string          `gorm:"column:buy_token_contract" json:"buy_token_contract"`                 // 购买代币合约地址
	BuyTokenName     string          `gorm:"column:buy_token_name" json:"buy_token_name"`                         // 购买代币名称
	BuyTokenSymbol   string          `gorm:"column:buy_token_symbol" json:"buy_token_symbol"`                     // 购买代币符号
	BuyTokenDecimals int             `gorm:"column:buy_token_decimals" json:"buy_token_decimals"`                 // 购买代币精度
	BuyTokenSupply   decimal.Decimal `gorm:"column:buy_token_supply;type:decimal(38,18)" json:"buy_token_supply"` // 购买代币总量
	BuyPrice         decimal.Decimal `gorm:"column:buy_price;type:decimal(38,18)" json:"buy_price"`               // 购买价格
	MaxBuyAmount     decimal.Decimal `gorm:"column:max_buy_amount;type:decimal(38,18)" json:"max_buy_amount"`     // 最大购买量
}

// CrowdfundingIBORate 众筹项目IBO费率表结构
type CrowdfundingIBORate struct {
	model.RelationBase
	CrowdfundingID int64           `gorm:"column:crowdfunding_id" json:"crowdfunding_id"`                   // 关联的众筹项目ID
	EndTime        time.Time       `gorm:"column:end_time" json:"end_time"`                                 // 阶段结束时间
	MaxBuyAmount   decimal.Decimal `gorm:"column:max_buy_amount;type:decimal(38,18)" json:"max_buy_amount"` // 最大购买量
	MaxSellPercent float64         `gorm:"column:max_sell_percent" json:"max_sell_percent"`                 // 最大出售百分比
	BuyPrice       decimal.Decimal `gorm:"column:buy_price;type:decimal(38,18)" json:"buy_price"`           // IBO价格
	SwapPercent    float64         `gorm:"column:swap_percent" json:"swap_percent"`                         // 兑换百分比
}

// TableName 指定表名
func (CrowdfundingIBORate) TableName() string {
	return "crowdfunding_ibo_rate"
}

// CrowdfundingInvestor 众筹项目投资者表结构
type CrowdfundingInvestor struct {
	model.RelationBase
	CrowdfundingID   int64           `gorm:"column:crowdfunding_id;uniqueIndex:crowdfunding_comer_uindex" json:"crowdfunding_id"` // 众筹项目ID
	ComerID          int64           `gorm:"column:comer_id;uniqueIndex:crowdfunding_comer_uindex" json:"comer_id"`               // 投资者ID
	BuyTokenTotal    decimal.Decimal `gorm:"column:buy_token_total;type:decimal(38,18)" json:"buy_token_total"`                   // 购买代币总量
	BuyTokenBalance  decimal.Decimal `gorm:"column:buy_token_balance;type:decimal(38,18)" json:"buy_token_balance"`               // 购买代币余额
	SellTokenTotal   decimal.Decimal `gorm:"column:sell_token_total;type:decimal(38,18)" json:"sell_token_total"`                 // 出售代币总量
	SellTokenBalance decimal.Decimal `gorm:"column:sell_token_balance;type:decimal(38,18)" json:"sell_token_balance"`             // 出售代币余额
}

// TableName 指定表名
func (CrowdfundingInvestor) TableName() string {
	return "crowdfunding_investor"
}

// CrowdfundingSwap 众筹项目兑换记录表结构
type CrowdfundingSwap struct {
	model.RelationBase
	ChainInfo
	Timestamp       *time.Time      `gorm:"column:timestamp" json:"timestamp"`                                     // 交易时间戳
	Status          int8            `gorm:"column:status;default:0" json:"status"`                                 // 状态:0-待处理 1-成功 2-失败
	CrowdfundingID  uint64          `gorm:"column:crowdfunding_id" json:"crowdfunding_id"`                         // 众筹项目ID
	ComerID         uint64          `gorm:"column:comer_id" json:"comer_id"`                                       // 用户ID
	Access          int8            `gorm:"column:access" json:"access"`                                           // 操作类型:1-投资 2-赎回
	BuyTokenSymbol  string          `gorm:"column:buy_token_symbol" json:"buy_token_symbol"`                       // 买入代币符号
	BuyTokenAmount  decimal.Decimal `gorm:"column:buy_token_amount;type:decimal(38,18)" json:"buy_token_amount"`   // 买入代币数量
	SellTokenSymbol string          `gorm:"column:sell_token_symbol" json:"sell_token_symbol"`                     // 卖出代币符号
	SellTokenAmount decimal.Decimal `gorm:"column:sell_token_amount;type:decimal(38,18)" json:"sell_token_amount"` // 卖出代币数量
	Price           decimal.Decimal `gorm:"column:price;type:decimal(38,18)" json:"price"`                         // 兑换价格
}

// TableName 指定表名
func (CrowdfundingSwap) TableName() string {
	return "crowdfunding_swap"
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
