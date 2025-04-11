package startup

type ListStartupsResponse struct {
	List  []Startup `json:"list"`
	Total int64     `json:"total"`
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
