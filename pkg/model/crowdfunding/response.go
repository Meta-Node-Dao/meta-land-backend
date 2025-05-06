package crowdfunding

import (
	"ceres/pkg/model/account"
	"ceres/pkg/model/tag"
	"time"

	"github.com/shopspring/decimal"
)

type SaleLaunchpadBasicResponse struct {
	ChainId              int                 `json:"chain_id"`
	ComerId              int                 `json:"comer_id"`
	ContractAddress      string              `json:"contract_address"`
	Cycle                int                 `json:"cycle"`
	CycleRelease         int                 `json:"cycle_release"`
	DexInitPrice         int                 `json:"dex_init_price"`
	DexPairAddress       string              `json:"dex_pair_address"`
	DexRouter            string              `json:"dex_router"`
	EndedAt              int                 `json:"ended_at"`
	FirstRelease         int                 `json:"first_release"`
	HardCap              int                 `json:"hard_cap"`
	Id                   int                 `json:"id"`
	InvestTokenBalance   int                 `json:"invest_token_balance"`
	InvestTokenContract  string              `json:"invest_token_contract"`
	InvestTokenSymbol    string              `json:"invest_token_symbol"`
	Investors            int                 `json:"investors"`
	LiquidityRate        int                 `json:"liquidity_rate"`
	MaxInvestAmount      int                 `json:"max_invest_amount"`
	MinInvestAmount      int                 `json:"min_invest_amount"`
	Poster               string              `json:"poster"`
	PresalePrice         int                 `json:"presale_price"`
	PresaleTokenBalance  int                 `json:"presale_token_balance"`
	PresaleTokenContract string              `json:"presale_token_contract"`
	PresaleTokenSymbol   string              `json:"presale_token_symbol"`
	SoftCap              int                 `json:"soft_cap"`
	StartedAt            int                 `json:"started_at"`
	Startup              StartupCardResponse `json:"startup"`
	StartupId            int                 `json:"startup_id"`
	Status               int                 `json:"status"`
	TeamWallet           string              `json:"team_wallet"`
	Title                string              `json:"title"`
	TxHash               string              `json:"tx_hash"`
}

type SaleLaunchpadResponse struct {
	ChainId              int                          `json:"chain_id"`
	ComerId              int                          `json:"comer_id"`
	ContractAddress      string                       `json:"contract_address"`
	Cycle                int                          `json:"cycle"`
	CycleRelease         int                          `json:"cycle_release"`
	Description          string                       `json:"description"`
	Detail               string                       `json:"detail"`
	DexInitPrice         float64                      `json:"dex_init_price"`
	DexPairAddress       string                       `json:"dex_pair_address"`
	DexRouter            string                       `json:"dex_router"`
	EndedAt              int                          `json:"ended_at"`
	FirstRelease         int                          `json:"first_release"`
	HardCap              int                          `json:"hard_cap"`
	Id                   int                          `json:"id"`
	InvestTokenBalance   int                          `json:"invest_token_balance"`
	InvestTokenContract  string                       `json:"invest_token_contract"`
	InvestTokenDecimals  int                          `json:"invest_token_decimals"`
	InvestTokenName      string                       `json:"invest_token_name"`
	InvestTokenSupply    int                          `json:"invest_token_supply"`
	InvestTokenSymbol    string                       `json:"invest_token_symbol"`
	Investor             CrowdfundingInvestorResponse `json:"investor"`
	Investors            int                          `json:"investors"`
	LiquidityRate        int                          `json:"liquidity_rate"`
	MaxInvestAmount      int                          `json:"max_invest_amount"`
	MinInvestAmount      int                          `json:"min_invest_amount"`
	Poster               string                       `json:"poster"`
	PresalePrice         int                          `json:"presale_price"`
	PresaleTokenBalance  int                          `json:"presale_token_balance"`
	PresaleTokenContract string                       `json:"presale_token_contract"`
	PresaleTokenDecimals int                          `json:"presale_token_decimals"`
	PresaleTokenDeposit  int                          `json:"presale_token_deposit"`
	PresaleTokenName     string                       `json:"presale_token_name"`
	PresaleTokenSupply   int                          `json:"presale_token_supply"`
	PresaleTokenSymbol   string                       `json:"presale_token_symbol"`
	SoftCap              int                          `json:"soft_cap"`
	StartedAt            int                          `json:"started_at"`
	Startup              StartupCardResponse          `json:"startup"`
	StartupId            int                          `json:"startup_id"`
	Status               int                          `json:"status"`
	Swaps                []SaleLaunchpadHistoryResponse
	TeamWallet           string `json:"team_wallet"`
	Title                string `json:"title"`
	TxHash               string `json:"tx_hash"`
	Youtube              string `json:"youtube"`
}

type CrowdfundingInvestorResponse struct {
	BuyTokenBalance  int `json:"buy_token_balance"`
	BuyTokenTotal    int `json:"buy_token_total"`
	ComerId          int `json:"comer_id"`
	CrowdfundingId   int `json:"crowdfunding_id"`
	Id               int `json:"id"`
	SellTokenBalance int `json:"sell_token_balance"`
	SellTokenTotal   int `json:"sell_token_total"`
}

type SaleLaunchpadHistoryResponse struct {
	Amount          int                        `json:"amount"`
	ChainId         int                        `json:"chain_id"`
	Comer           account.ComerBasicResponse `json:"comer"`
	ComerId         int                        `json:"comer_id"`
	Id              int                        `json:"id"`
	SaleLaunchpadId int                        `json:"sale_launchpad_id"`
	Timestamp       int                        `json:"timestamp"`
	TokenSymbol     string                     `json:"token_symbol"`
	TxHash          string                     `json:"tx_hash"`
	Type            int                        `json:"type"`
}

type StartupCardResponse struct {
	Banner        string                       `json:"banner"`
	ChainId       int                          `json:"chain_id"`
	ComerId       int                          `json:"comer_id"`
	ContractAudit string                       `json:"contract_audit"`
	Id            int                          `json:"id"`
	IsConnected   bool                         `json:"is_connected"`
	Kyc           string                       `json:"kyc"`
	Logo          string                       `json:"logo"`
	Mission       string                       `json:"mission"`
	Name          string                       `json:"name"`
	OnChain       bool                         `json:"on_chain"`
	Socials       []account.SocialBookResponse `json:"socials"`
	Tags          []tag.TagRelationResponse    `json:"tags"`
	TxHash        string                       `json:"tx_hash"`
	Type          int                          `json:"type"`
}

type SignResponse struct {
	Data string `json:"data"`
	Sign string `json:"sign"`
}

type CrowdfundingBasicResponse struct {
	BuyPrice             float64             `json:"buy_price"`
	BuyTokenContract     string              `json:"buy_token_contract"`
	BuyTokenSymbol       string              `json:"buy_token_symbol"`
	ChainId              int                 `json:"chain_id"`
	ComerId              int                 `json:"comer_id"`
	CrowdfundingContract string              `json:"crowdfunding_contract"`
	DexInitPrice         float64             `json:"dex_init_price"`
	DexRouter            string              `json:"dex_router"`
	EndTime              string              `json:"end_time"`
	Id                   int                 `json:"id"`
	Investors            int                 `json:"investors"`
	MaxBuyAmount         int                 `json:"max_buy_amount"`
	MaxSellPercent       int                 `json:"max_sell_percent"`
	MinBuyAmount         float64             `json:"min_buy_amount"`
	PairAddress          string              `json:"pair_address"`
	Poster               string              `json:"poster"`
	RaiseBalance         int                 `json:"raise_balance"`
	RaiseGoal            int                 `json:"raise_goal"`
	SellTax              int                 `json:"sell_tax"`
	SellTokenContract    string              `json:"sell_token_contract"`
	SellTokenSymbol      string              `json:"sell_token_symbol"`
	StartTime            string              `json:"start_time"`
	Startup              StartupCardResponse `json:"startup"`
	StartupId            int                 `json:"startup_id"`
	Status               int                 `json:"status"`
	SwapPercent          int                 `json:"swap_percent"`
	TeamWallet           string              `json:"team_wallet"`
	Title                string              `json:"title"`
	TxHash               string              `json:"tx_hash"`
}

type CrowdfundingResponse struct {
	BuyPrice             int                        `json:"buy_price"`
	BuyTokenContract     string                     `json:"buy_token_contract"`
	BuyTokenDecimals     int                        `json:"buy_token_decimals"`
	BuyTokenName         string                     `json:"buy_token_name"`
	BuyTokenSupply       int                        `json:"buy_token_supply"`
	BuyTokenSymbol       string                     `json:"buy_token_symbol"`
	ChainId              int                        `json:"chain_id"`
	ComerId              int                        `json:"comer_id"`
	CrowdfundingContract string                     `json:"crowdfunding_contract"`
	Description          string                     `json:"description"`
	Detail               string                     `json:"detail"`
	DexInitPrice         int                        `json:"dex_init_price"`
	DexRouter            string                     `json:"dex_router"`
	EndTime              string                     `json:"end_time"`
	Id                   int                        `json:"id"`
	Investor             CrowdfundingInvestor       `json:"investor"`
	Investors            int                        `json:"investors"`
	MaxBuyAmount         int                        `json:"max_buy_amount"`
	MaxSellPercent       int                        `json:"max_sell_percent"`
	MinBuyAmount         int                        `json:"min_buy_amount"`
	PairAddress          string                     `json:"pair_address"`
	Poster               string                     `json:"poster"`
	RaiseBalance         int                        `json:"raise_balance"`
	RaiseGoal            int                        `json:"raise_goal"`
	SellTax              int                        `json:"sell_tax"`
	SellTokenBalance     int                        `json:"sell_token_balance"`
	SellTokenContract    string                     `json:"sell_token_contract"`
	SellTokenDecimals    int                        `json:"sell_token_decimals"`
	SellTokenDeposit     int                        `json:"sell_token_deposit"`
	SellTokenName        string                     `json:"sell_token_name"`
	SellTokenSupply      int                        `json:"sell_token_supply"`
	SellTokenSymbol      string                     `json:"sell_token_symbol"`
	StartTime            string                     `json:"start_time"`
	Startup              StartupCardResponse        `json:"startup"`
	StartupId            int                        `json:"startup_id"`
	Status               int                        `json:"status"`
	SwapPercent          int                        `json:"swap_percent"`
	Swaps                []CrowdfundingSwapResponse `json:"swaps"`
	TeamWallet           string                     `json:"team_wallet"`
	Title                string                     `json:"title"`
	TxHash               string                     `json:"tx_hash"`
	Youtube              string                     `json:"youtube"`
}

type CrowdfundingSwapResponse struct {
	Banner        string                       `json:"banner"`
	ChainId       int                          `json:"chain_id"`
	ComerId       int                          `json:"comer_id"`
	ContractAudit string                       `json:"contract_audit"`
	Id            int                          `json:"id"`
	IsConnected   bool                         `json:"is_connected"`
	Kyc           string                       `json:"kyc"`
	Logo          string                       `json:"logo"`
	Mission       string                       `json:"mission"`
	Name          string                       `json:"name"`
	OnChain       bool                         `json:"on_chain"`
	Socials       []account.SocialBookResponse `json:"socials"`
	Tags          []tag.TagRelationResponse    `json:"tags"`
	TxHash        string                       `json:"tx_hash"`
	Type          int                          `json:"type"`
}

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
