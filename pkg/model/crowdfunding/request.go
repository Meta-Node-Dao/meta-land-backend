package crowdfunding

import (
	"ceres/pkg/model"
	"ceres/pkg/router"
	"strings"
	"time"

	"github.com/qiniu/x/log"
	"github.com/shopspring/decimal"
)

type InvestRequest struct {
	TransactionHashRequest
	Access          SwapAccess      `json:"access"`
	BuyTokenSymbol  string          `json:"buyTokenSymbol"`
	BuyTokenAmount  decimal.Decimal `json:"buyTokenAmount"`
	SellTokenSymbol string          `json:"sellTokenSymbol"`
	SellTokenAmount decimal.Decimal `json:"sellTokenAmount"`
	Price           decimal.Decimal `json:"price"`
}

type TransactionHashRequest struct {
	TxHash string `json:"txHash" binding:"required"`
}

type ModifyRequest struct {
	TransactionHashRequest
	SwapPercent  decimal.Decimal `json:"swapPercent,omitempty"`
	BuyPrice     decimal.Decimal `json:"buyPrice,omitempty"`
	MaxBuyAmount decimal.Decimal `json:"maxBuyAmount,omitempty"`
	//BuyTokenSymbol  string          `json:"buyTokenSymbol,omitempty"`
	//SellTokenSymbol string          `json:"sellTokenSymbol,omitempty"`
	MaxSellPercent decimal.Decimal `json:"maxSellPercent,omitempty"`
	EndTime        time.Time       `json:"endTime,omitempty"`
}

type PublicCrowdfundingListPageRequest struct {
	model.Pagination
	Mode CrowdfundingStatus `json:"mode,omitempty"`
}

type CreateCrowdfundingRequest struct {
	ChainInfo
	SellInfo
	BuyInfo
	// update by contract callback
	// CrowdfundingContract string          `gorm:"crowdfunding_contract" json:"crowdfundingContract, omitempty"`
	StartupId  uint64          `gorm:"startup_id" json:"startupId"`
	ComerId    uint64          `gorm:"comer_id" json:"comerId"`
	TeamWallet string          `gorm:"team_wallet" json:"teamWallet"`
	RaiseGoal  decimal.Decimal `gorm:"raise_goal" json:"raiseGoal"`
	// form
	SwapPercent decimal.Decimal `gorm:"swap_percent" json:"swapPercent"`
	StartTime   time.Time       `gorm:"start_time" json:"startTime"`
	EndTime     time.Time       `gorm:"end_time" json:"endTime"`
	Poster      string          `gorm:"poster" json:"poster"`
	Youtube     string          `gorm:"youtube" json:"youtube"`
	Detail      string          `gorm:"detail" json:"detail"`
	Description string          `gorm:"description" json:"description"`
}

type StringField struct {
	field, value string
}
type MoneyField struct {
	field string
	value decimal.Decimal
}
type IntField struct {
	field string
	value uint64
}

func (r CreateCrowdfundingRequest) ValidRequest() error {
	var stringFields = []StringField{{"Team wallet address", r.TeamWallet}, {"Token Contract", r.BuyTokenContract}, {"Poster", r.Poster}, {"Description", r.Description}, {"TxHash", r.TxHash}}
	if err := anyEmptyStringError(stringFields); err != nil {
		return err
	}
	if err := anyEmptyIntError([]IntField{{"StartupId", r.StartupId}, {"ChainId", r.ChainId}}); err != nil {
		return err
	}
	if err := anyZeroDecimalError([]MoneyField{{"Raise Goal", r.RaiseGoal}, {"IBO Rate", r.BuyPrice}, {"Maximum Buy Amount", r.MaxBuyAmount}, {"Maximum Sell", r.MaxSellPercent}}); err != nil {
		return err
	}

	if !strings.HasPrefix(r.BuyTokenContract, "0x") || len(r.BuyTokenContract) > 64 {
		return router.ErrBadRequest.WithMsgf("Invalid token contract: %s", r.BuyTokenContract)
	}

	if !strings.HasPrefix(r.TeamWallet, "0x") || len(r.TeamWallet) > 64 {
		return router.ErrBadRequest.WithMsgf("Invalid team wallet address: %s", r.TeamWallet)
	}

	if r.RaiseGoal.IsNegative() {
		return router.ErrBadRequest.WithMsg("Raise goal must be positive number")
	}

	if r.SwapPercent.IsNegative() {
		return router.ErrBadRequest.WithMsg("Swap must be positive number")
	}
	if r.BuyPrice.IsNegative() {
		return router.ErrBadRequest.WithMsg("IBO Rate must be positive  number")
	}
	if r.MaxBuyAmount.IsNegative() {
		return router.ErrBadRequest.WithMsg("Maximum Buy must be positive number")
	}
	if r.SellTax.IsNegative() {
		return router.ErrBadRequest.WithMsg("Sell Tax must be positive number")
	}

	if r.MaxSellPercent.IsNegative() {
		return router.ErrBadRequest.WithMsg("Maximum Sell must be positive number")
	}

	if !r.StartTime.Before(r.EndTime) {
		return router.ErrBadRequest.WithMsg("Start time needs to be before End time")
	}

	return nil
}

func anyEmptyIntError(targets []IntField) error {
	for _, target := range targets {
		if target.value == 0 {
			return router.ErrBadRequest.WithMsgf("Invalid parameter: %s", target.field)
		}
	}
	return nil
}

func anyEmptyStringError(targets []StringField) error {
	for _, target := range targets {
		if target.value == "" || strings.TrimSpace(target.value) == "" {
			return router.ErrBadRequest.WithMsgf("%s cannot be blank", target.field)
		}
	}
	return nil
}

func anyZeroDecimalError(targets []MoneyField) error {
	for _, target := range targets {
		log.Infof("#### %s: %v\n", target.field, target.value)
		if target.value.Cmp(decimal.Zero) == 0 {
			return router.ErrBadRequest.WithMsgf("%s cannot be empty", target.field)
		}
	}
	return nil
}
