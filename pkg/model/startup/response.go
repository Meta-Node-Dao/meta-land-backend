package startup

import (
	"ceres/pkg/model/account"
	"ceres/pkg/model/tag"
)

type ListStartupsResponse struct {
	List  []Startup `json:"list"`
	Total int64     `json:"total"`
}

type StartupListResponse struct {
	List  []Startup `json:"list"`
	Total int64     `json:"total"`
	Page  int64     `json:"page"`
	Size  int64     `json:"size"`
}

type StartupBasicResponse struct {
	Banner         string                    `json:"banner"`
	ChainId        int                       `json:"chain_id"`
	ComerId        int                       `json:"comer_id"`
	ConnectedTotal int                       `json:"connected_total"`
	ContractAudit  string                    `json:"contract_audit"`
	Finance        StartupFinance            `json:"finance"`
	Id             int                       `json:"id"`
	IsConnected    bool                      `json:"is_connected"`
	Kyc            string                    `json:"kyc"`
	Logo           string                    `json:"logo"`
	Mission        string                    `json:"mission"`
	Name           string                    `json:"name"`
	OnChain        bool                      `json:"on_chain"`
	Tags           []tag.TagRelationResponse `json:"tags"`
	Team           []StartupTeam             `json:"team"`
	TxHash         string                    `json:"tx_hash"`
	Type           int                       `json:"type"`
}

type IsConnectedResponse struct {
	IsConnected bool `json:"is_connected"`
}

type StartupInfoResponse struct {
	Banner         string                    `json:"banner"`
	ChainId        int                       `json:"chain_id"`
	ComerId        int                       `json:"comer_id"`
	ConnectedTotal int                       `json:"connected_total"`
	ContractAudit  string                    `json:"contract_audit"`
	Finance        StartupFinance            `json:"finance"`
	Id             int                       `json:"id"`
	IsConnected    bool                      `json:"is_connected"`
	Kyc            string                    `json:"kyc"`
	Logo           string                    `json:"logo"`
	Mission        string                    `json:"mission"`
	Name           string                    `json:"name"`
	OnChain        bool                      `json:"on_chain"`
	Overview       string                    `json:"overview"`
	Socials        []SocialBookResponse      `json:"socials"`
	TabSequence    string                    `json:"tab_sequence"`
	Tags           []tag.TagRelationResponse `json:"tags"`
	Team           []StartupTeam             `json:"team"`
	TxHash         string                    `json:"tx_hash"`
	Type           int                       `json:"type"`
}

type SocialBookResponse struct {
	Id           int                `json:"id"`
	SocialTool   SocialToolResponse `json:"social_tool"`
	SocialToolId int                `json:"social_tool_id"`
	TargetId     int                `json:"target_id"`
	Type         int                `json:"type"`
	Value        string             `json:"value"`
}

type SocialToolResponse struct {
	Id   int    `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

type StartupFinance struct {
	ChainId          int                    `json:"chain_id"`
	ComerId          int                    `json:"comer_id"`
	ContractAddress  string                 `json:"contract_address"`
	Id               int                    `json:"id"`
	LaunchedAt       string                 `json:"launched_at"`
	Name             string                 `json:"name"`
	PresaleEndedAt   string                 `json:"presale_ended_at"`
	PresaleStartedAt string                 `json:"presale_started_at"`
	StartupId        int                    `json:"startup_id"`
	Supply           int                    `json:"supply"`
	Symbol           string                 `json:"symbol"`
	Wallets          []StartupFinanceWallet `json:"wallets"`
}

type StartupFinanceWallet struct {
	Address          string `json:"address"`
	Id               int    `json:"id"`
	Name             string `json:"name"`
	StartupFinanceId int    `json:"startup_finance_id"`
	StartupId        int    `json:"startup_id"`
}

type StartupTeam struct {
	Comer              account.ComerBasicResponse `json:"comer"`
	ComerId            int                        `json:"comer_id"`
	CreatedAt          string                     `json:"created_at"`
	Id                 int                        `json:"id"`
	Position           string                     `json:"position"`
	StartupId          int                        `json:"startup_id"`
	StartupTeamGroup   StartupTeamGroup           `json:"startup_team_group"`
	StartupTeamGroupId int                        `json:"startup_team_group_id"`
}

type StartupTeamGroup struct {
	ComerId   int    `json:"comer_id"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	StartupId int    `json:"startup_id"`
}

type StartupTeamGroupResponse struct {
	ComerId   int    `json:"comer_id"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	StartupId int    `json:"startup_id"`
}

type GetStartupResponse struct {
	Startup
}

type ExistStartupResponse struct {
	IsExist bool `json:"isExist"`
}

type FollowedByMeResponse struct {
	IsFollowed bool `json:"isFollowed"`
}

type ListComerStartupsResponse struct {
	List  []*ListComerStartup `json:"list"`
	Total int                 `json:"total"`
}

type ListComerStartup struct {
	StartupID uint64 `gorm:"column:id" json:"startupID"`
	ComerId   uint64 `gorm:"comer_id" json:"comerId"`
	OnChain   bool   `json:"onChain"`
	Name      string `gorm:"column:name" json:"name"`
	FollowedByMeResponse
}

type StartupModuleDataInfo struct {
	BountyCnt       int64 `json:"bountyCnt"`
	CrowdfundingCnt int64 `json:"crowdfundingCnt"`
	ProposalCnt     int64 `json:"proposalCnt"`
	OtherDappCnt    int64 `json:"otherDappCnt"`
}

type SimpleStartupInfo struct {
	StartupId   uint64 `gorm:"column:id" json:"startupId"`
	StartupName string `gorm:"column:name" json:"startupName"`
	StartupLogo string `gorm:"column:avatar" json:"startupLogo"`
	OnChain     bool   `gorm:"column:on_chain" json:"onChain"`
}
