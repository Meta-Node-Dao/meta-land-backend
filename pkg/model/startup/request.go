package startup

import (
	"ceres/pkg/model"
	"ceres/pkg/model/account"
)

type StartupListRequest struct {
	Page               int      `form:"page"`
	Size               int      `form:"size"`
	AdminComerId       int      `form:"admin_comer_id"`
	ComerId            uint64   `form:"comer_id"`
	Connected          bool     `form:"connected"`
	Keyword            string   `form:"keyword"`
	Tags               []string `form:"tags"`
	Type               int      `form:"type"`
	StartupTeamComerId uint64   `form:"startup_team_comer_id"`
}

type ListStartupRequest struct {
	model.ListRequest
	Keyword string `form:"keyword"`
	Mode    Mode   `form:"mode"`
}

type CreateStartupRequest struct {
	Logo     string   `json:"logo"`
	Mode     Mode     `form:"mode" binding:"required"`
	Name     string   `json:"name" binding:"required"`
	Mission  string   `json:"mission" binding:"required"`
	Overview string   `json:"overview" binding:"required"`
	TxHash   string   `json:"txHash"`
	ChainID  uint64   `json:"chainId" binding:"required"`
	HashTags []string `json:"hashTags" binding:"required"`
}

type UpdateStartupBasicSettingRequest struct {
	KYC           *string  `json:"kyc"`
	ContractAudit *string  `json:"contractAudit"`
	HashTags      []string `json:"hashTags"`
	Website       *string  `json:"website"`
	Discord       *string  `json:"discord"`
	Twitter       *string  `json:"twitter"`
	Telegram      *string  `json:"telegram"`
	Docs          *string  `json:"docs"`
}
type UpdateStartupBasicSettingRequestNew struct {
	Name     string   `json:"name"`
	Logo     string   `json:"logo"`
	Cover    string   `json:"cover"`
	Mode     Mode     `json:"mode"`
	Mission  string   `json:"mission"`
	Overview string   `json:"overview"`
	TxHash   string   `json:"txHash"`
	ChainID  uint64   `json:"chainId"`
	HashTags []string `json:"hashTags"`
}

type UpdateStartupFinanceSettingRequest struct {
	TokenContractAddress *string `json:"tokenContractAddress" binding:"required"`
	LaunchNetwork        *int    `json:"launchNetwork" binding:"required"`
	TokenName            *string `json:"tokenName" binding:"required"`
	TokenSymbol          *string `json:"tokenSymbol" binding:"required"`
	TotalSupply          *int64  `json:"totalSupply" binding:"required"`
	PresaleStart         *string `json:"presaleStart" binding:"required"`
	PresaleEnd           *string `json:"presaleEnd" binding:"required"`
	LaunchDate           *string `json:"launchDate" binding:"required"`
	Wallets              []struct {
		WalletName    string `json:"walletName" binding:"min=1,max=50"`
		WalletAddress string `json:"walletAddress" binding:"len=42,startswith=0x"`
	} `json:"wallets" binding:"required"`
}

type UpdateStartupCoverRequest struct {
	Image string `json:"image"`
}

type UpdateStartupSecurityRequest struct {
	KYC           *string `json:"kyc"`
	ContractAudit *string `json:"contractAudit"`
}

type UpdateStartupTabSequenceRequest struct {
	Tabs []model.BusinessModule `json:"tabs"`
}

type CreateOrUpdateStartupGroupRequest struct {
	Name string `json:"name"`
}

type UpdateStartupSocialsAndTagsRequest struct {
	HashTags       []string                      `json:"hashTags"`
	Socials        []account.SocialModifyRequest `json:"socials"`
	DeletedSocials []account.SocialType          `json:"deletedSocials"`
}

type ModifyLocationRequest struct {
	Position string `json:"position"`
}
