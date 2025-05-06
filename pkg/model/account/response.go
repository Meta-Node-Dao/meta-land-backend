package account

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
)

// JwtAuthorizationResponse ACTIVE
type JwtAuthorizationResponse struct {
	Token string `json:"token"`
}

type ComerResponse struct {
	Activation     bool   `json:"activation"`
	Address        string `json:"address"`
	Avatar         string `json:"avatar"`
	Banner         string `json:"banner"`
	CustomDomain   string `json:"custom_domain"`
	Id             int    `json:"id"`
	InvitationCode string `json:"invitation_code"`
	IsConnected    bool   `json:"is_connected"`
	IsSeted        bool   `json:"is_seted"`
	Location       string `json:"location"`
	Name           string `json:"name"`
	TimeZone       string `json:"time_zone"`
}

type ComerInfoDetailResponse struct {
	Accounts       []ComerAccountResponse      `json:"accounts"`
	Activation     bool                        `json:"activation"`
	Address        string                      `json:"address"`
	Avatar         string                      `json:"avatar"`
	Banner         string                      `json:"banner"`
	ConnectedTotal ComerConnectedTotalResponse `json:"connected_total"`
	CustomDomain   string                      `json:"custom_domain"`
	Educations     []ComerEducation            `json:"educations"`
	Id             int                         `json:"id"`
	Info           ComerInfo                   `json:"info"`
	InvitationCode string                      `json:"invitation_code"`
	IsConnected    bool                        `json:"is_connected"`
	Languages      []ComerLanguageResponse     `json:"languages"`
	Location       string                      `json:"location"`
	Name           string                      `json:"name"`
	Skills         []tag.TagRelationResponse   `json:"skills"`
	Socials        []SocialBookResponse        `json:"socials"`
	TimeZone       string                      `json:"time_zone"`
}

type LanguageResponse struct {
	Code string `json:"code"`
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ComerInvitationCountResponse struct {
	ActivatedTotal int `json:"activated_total"`
	InactiveTotal  int `json:"inactive_total"`
}

type InvitationRecordsPageData struct {
	List  []interface{}                 `json:"list"`
	Page  int                           `json:"page"`
	Size  int                           `json:"size"`
	Total int                           `json:"total"`
	Data  ComerInvitationRecordResponse `json:"data"`
}

type StartupListResponse struct {
	List  []SimpleStartupInfo `json:"list"`
	Total int                 `json:"total"`
}

type SimpleStartupInfo struct {
	Avatar  string `json:"avatar"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
	OnChain bool   `json:"on_chain"`
}

type ComerAccountResponse struct {
	Avatar    string `json:"avatar"`
	ComerId   int    `json:"comer_id"`
	Id        int    `json:"id"`
	IsLinked  bool   `json:"is_linked"`
	IsPrimary bool   `json:"is_primary"`
	Nickname  string `json:"nickname"`
	Oin       string `json:"oin"`
	Type      int    `json:"type"`
}

type ComerConnectedTotalResponse struct {
	BeConnectComerTotal int `json:"be_connect_comer_total"`
	ConnectComerTotal   int `json:"connect_comer_total"`
	ConnectStartupTotal int `json:"connect_startup_total"`
}

type ComerEducation struct {
	ComerId     int    `json:"comer_id"`
	GraduatedAt string `json:"graduated_at"`
	Id          int    `json:"id"`
	Major       string `json:"major"`
	School      string `json:"school"`
}

type ComerInfo struct {
	Bio     string `json:"bio"`
	ComerId int    `json:"comer_id"`
	Id      int    `json:"id"`
}

type ComerLanguageResponse struct {
	ComerId    int      `json:"comer_id"`
	Id         int      `json:"id"`
	Language   Language `json:"language"`
	LanguageId int      `json:"language_id"`
	Level      int      `json:"level"`
}

type Language struct {
	Code string `json:"code"`
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SocialBookResponse struct {
	Id           int                `json:"id"`
	SocialTool   SocialToolResponse `json:"social_tool"`
	SocialToolId int                `json:"social_tool_id"`
	TargetId     int                `json:"target_id"`
	Type         int                `json:"type"`
	Value        string             `json:"value"`
}

type ThirdPartyVerifyResponse struct {
	Verify bool `json:"verify"`
}

type IsConnectedResponse struct {
	IsConnected bool `json:"is_connected"`
}

type ProjectCountResponse struct {
	BountyCount        int `json:"bounty_count"`
	CrowdfundingCount  int `json:"crowdfunding_count"`
	GovernanceCount    int `json:"governance_count"`
	OtherDappCount     int `json:"other_dapp_count"`
	SaleLaunchpadCount int `json:"sale_launchpad_count"`
	StartupCount       int `json:"startup_count"`
}

type ShareSetResponse struct {
	ShareCode string `json:"share_code"`
}

type SocialToolResponse struct {
	Id   int    `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

type ComerInvitationRecordResponse struct {
	Comer          ComerBasicResponse `json:"comer"`
	ComerId        int                `json:"comer_id"`
	CreatedAt      string             `json:"created_at"`
	Id             int                `json:"id"`
	InvitationCode string             `json:"invitation_code"`
	Invitee        ComerBasicResponse `json:"invitee"`
	InviteeId      int                `json:"invitee_id"`
}

type ComerBasicResponse struct {
	Activation     bool   `json:"activation"`
	Address        string `json:"address"`
	Avatar         string `json:"avatar"`
	Banner         string `json:"banner"`
	CustomDomain   string `json:"custom_Domain"`
	Id             int    `json:"id"`
	InvitationCode string `json:"invitation_code"`
	IsConnected    bool   `json:"is_connected"`
	Location       string `json:"location"`
	Name           string `json:"name"`
	TimeZone       string `json:"time_zone"`
}

type OauthLoginResponse struct {
	ComerID    uint64 `json:"comerID"`
	Nick       string `json:"nick"`
	Avatar     string `json:"avatar"`
	Address    string `json:"address"`
	Token      string `json:"token"`
	IsProfiled bool   `json:"isProfiled"`
	// OauthType wallet, oauthGithub, oauthGoogle
	// OauthType ComerAccountType `json:"oauthType"`
	// OauthLinked
	OauthLinked    bool   `json:"oauthLinked"`
	OauthAccountId uint64 `json:"oauthAccountId"`
}

// ComerLoginResponse comer login response
type ComerLoginResponse struct {
	ComerID       uint64                     `json:"comerID"`
	Nick          string                     `json:"nick"`
	Avatar        string                     `json:"avatar"`
	Address       string                     `json:"address"`
	Token         string                     `json:"token"`
	IsProfiled    bool                       `json:"isProfiled"`
	FirstLogin    bool                       `json:"firstLogin"`
	ComerAccounts []*OauthAccountBindingInfo `json:"comerAccounts"`
}

// WalletNonceResponse ACTIVE
type WalletNonceResponse struct {
	Nonce string `json:"nonce"`
}

// ComerProfileResponse return the profile of some comer
type ComerProfileResponse struct {
	*ComerProfile
	ComerAccounts []*OauthAccountBindingInfo `json:"comerAccounts"`
}
type OauthAccountBindingInfo struct {
	Linked      bool             `json:"linked"`
	AccountType ComerAccountType `json:"accountType"`
	AccountId   uint64           `json:"accountId"`
}

// ComerOuterAccountListResponse response of the comer outer account list
type ComerOuterAccountListResponse struct {
	List  []ComerAccount `json:"list"`
	Total uint64         `json:"total"`
}

type GetComerInfoResponse struct {
	Comer
	ComerProfile  ComerProfile    `json:"comerProfile"`
	Follows       []FollowComer   `json:"follows"`
	FollowsCount  int64           `json:"followsCount"`
	Followed      []FollowedComer `json:"followed"`
	FollowedCount int64           `json:"followedCount"`
}

type IsFollowedResponse struct {
	IsFollowed bool `json:"isFollowed"`
}

// LinkWalletResponse link wallet response
type LinkWalletResponse struct {
	IsProfiled bool   `json:"isProfiled"`
	Token      string `json:"token"`
}

// ComerFollowerInfo information of follower
type ComerFollowerInfo struct {
	ComerId      uint64 `json:"comerId"`
	ComerAvatar  string `json:"comerAvatar"`
	ComerName    string `json:"comerName"`
	FollowedByMe *bool  `json:"followedByMe,omitempty"`
}

type StartupFollowerInfo struct {
	StartupId    uint64 `json:"startupId"`
	StartupLogo  string `json:"startupLogo"`
	StartupName  string `json:"startupName"`
	FollowedByMe bool   `json:"followedByMe"`
}

type ModuleInfo struct {
	Module     model.BusinessModule `json:"module"`
	HasCreated bool                 `json:"hasCreated"`
}

type ComerConnectedInfo struct {
	StartupCnt  int64 `json:"startupCnt"`
	ComerCnt    int64 `json:"comerCnt"`
	FollowerCnt int64 `json:"followerCnt"`
}

type BusinessModuleDataType int

const (
	Posted BusinessModuleDataType = iota + 1
	Participated
)

type ComerModuleDataInfo struct {
	Type            BusinessModuleDataType `json:"type"`
	StartupCnt      int64                  `json:"startupCnt"`
	BountyCnt       int64                  `json:"bountyCnt"`
	CrowdfundingCnt int64                  `json:"crowdfundingCnt"`
	ProposalCnt     int64                  `json:"proposalCnt"`
}
