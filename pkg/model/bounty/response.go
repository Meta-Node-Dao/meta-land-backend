/**
 * @Author Sun
 * @Description:
 * @File  response
 * @Version 1.0.0
 * @Date 2022/6/29 13:17
 */

package bounty

import (
	"ceres/pkg/model/account"
	"ceres/pkg/model/tag"
	"time"
)

type TagRelationResponse struct {
	ID       int         `json:"id"`
	Tag      TagResponse `json:"tag"`
	TagID    int         `json:"tag_id"`
	TargetID int         `json:"target_id"`
	Type     int         `json:"type"`
}

type TagResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

type StartupBasic struct {
	Banner        string `json:"banner"`
	ChainId       int    `json:"chain_id"`
	ComerId       int    `json:"comer_id"`
	ContractAudit string `json:"contract_audit"`
	Id            int    `json:"id"`
	IsConnected   bool   `json:"is_connected"`
	Kyc           string `json:"kyc"`
	Logo          string `json:"logo"`
	Mission       string `json:"mission"`
	Name          string `json:"name"`
	OnChain       bool   `json:"on_chain"`
	TxHash        string `json:"tx_hash"`
	Type          int    `json:"type"`
}

// BountyBasicResponse ACTIVE
type BountyBasicResponse struct {
	ApplicantCount              int                   `json:"applicant_count"`
	ApplicantDeposit            int                   `json:"applicants_deposit"`
	ApplicantMinDeposit         int                   `json:"applicant_min_deposit"`
	ApplyDeadline               string                `json:"apply_deadline"`
	ChainId                     int                   `json:"chain_id"`
	ComerId                     int                   `json:"comer_id"`
	ContractAddress             string                `json:"contract_address"`
	CreatedAt                   string                `json:"created_at"`
	DepositContractAddress      string                `json:"deposit_contract_address"`
	DepositContractTokenDecimal int                   `json:"deposit_contract_token_decimal"`
	DepositContractTokenSymbol  string                `json:"deposit_contract_token_symbol"`
	DiscussionLink              string                `json:"discussion_link"`
	ExpiredTime                 string                `json:"expired_time"`
	FounderDeposit              int                   `json:"founder_deposit"`
	Id                          int                   `json:"id"`
	IsLock                      int                   `json:"is_lock"`
	PaymentMode                 int                   `json:"payment_mode"`
	Reward                      BountyReward          `json:"reward"`
	Skills                      []TagRelationResponse `json:"skills"`
	Startup                     StartupBasic          `json:"startup"`
	StartupId                   int                   `json:"startup_id"`
	Status                      int                   `json:"status"`
	Title                       string                `json:"title"`
	TxHash                      string                `json:"tx_hash"`
}

type BountyInfoResponse struct {
	applicant_deposit              int
	applicant_min_deposit          int
	applicants                     []BountyApplicant
	apply_deadline                 string
	approved                       BountyApplicant
	chain_id                       int
	comer_id                       int
	contacts                       []BountyContact
	contract_address               string
	created_at                     string
	deposit_contract_address       string
	deposit_contract_token_decimal int
	deposit_contract_token_symbol  string
	deposit_records                []BountyDepositRecord
	description                    string
	discussion_link                string
	expired_time                   string
	founder                        BountyComer
	founder_deposit                int
	id                             int
	is_lock                        int
	my_deposit                     int
	my_role                        int
	my_status                      int
	payment_mode                   int
	period                         BountyPaymentPeriod
	post_updates                   []PostUpdate
	skills                         []TagRelationResponse
	startup                        StartupCardResponse
	startup_id                     int
	status                         int
	terms                          []BountyPaymentTerms
	title                          string
	tx_hash                        string
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

type BountyComer struct {
	Activation     bool                      `json:"activation"`
	Address        string                    `json:"address"`
	Avatar         string                    `json:"avatar"`
	Banner         string                    `json:"banner"`
	CustomDomain   string                    `json:"custom_domain"`
	Id             int                       `json:"id"`
	InvitationCode string                    `json:"invitation_code"`
	IsConnected    bool                      `json:"is_connected"`
	Location       string                    `json:"location"`
	Name           string                    `json:"name"`
	Skills         []tag.TagRelationResponse `json:"skills"`
	TimeZone       string                    `json:"time_zone"`
}

type BountyDepositRecord struct {
	Amount    int                        `json:"amount"`
	BountyId  int                        `json:"bounty_id"`
	Comer     account.ComerBasicResponse `json:"comer"`
	ComerId   int                        `json:"comer_id"`
	CreatedAt string                     `json:"created_at"`
	Id        int                        `json:"id"`
	Mode      int                        `json:"mode"`
	Status    int                        `json:"status"`
	TxHash    string                     `json:"tx_hash"`
}

type ContractInfoResponse struct {
	ContractAddress string
	Status          int
}

type DetailItem struct {
	BountyId            uint64    `json:"bountyId"`
	StartupId           uint64    `json:"startupId"`
	ChainID             uint64    `json:"chainID"`
	Logo                string    `json:"logo"`
	Title               string    `json:"title"`
	Status              string    `json:"status"`
	OnChainStatus       string    `json:"onChainStatus"`
	PaymentType         string    `json:"paymentType"`
	DepositTokenSymbol  string    `json:"depositTokenSymbol"`
	ApplyCutoffDate     time.Time `json:"applyCutoffDate"`
	CreatedTime         time.Time `json:"createdTime"`
	Rewards             *[]Reward `json:"rewards"`
	ApplicantCount      int       `json:"applicantCount"`
	ApplicationSkills   []string  `json:"applicationSkills"`
	DepositRequirements int       `json:"depositRequirements"`
}

type TabListResponse struct {
	BountyCount int `json:"bountyCount"`
	PageParam
	TotalPages int `json:"totalPages"`
	Records    []*DetailItem
}

type Reward struct {
	TokenSymbol string `json:"tokenSymbol"`
	Amount      int    `json:"amount"`
}

type DetailBounty struct {
	Title              string    `json:"title" gorm:"title"`
	Status             int       `json:"status" gorm:"status"`
	DiscussionLink     string    `json:"discussionLink" gorm:"discussion_link"`
	ApplyCutoffDate    string    `json:"expiresIn" gorm:"apply_cutoff_date"`
	ApplicantDeposit   int       `json:"applicantsDeposit" gorm:"applicant_deposit"`
	Description        string    `json:"description" gorm:"description"`
	DepositContract    string    `json:"depositContract" gorm:"deposit_contract"`
	DepositTokenSymbol string    `json:"depositTokenSymbol" gorm:"deposit_token_symbol"`
	ChainID            uint64    `json:"chainID" gorm:"chain_id"`
	CreatedAt          time.Time `json:"createdAt" gorm:"created_at"`
}

type DetailResponse struct {
	DetailBounty
	ApplicantSkills []string  `json:"applicantSkills"`
	Contacts        []Contact `json:"contact"`
}

type BountyPaymentInfo struct {
	PaymentMode int `json:"paymentMode" gorm:"payment_mode"`
}

type PaymentResponse struct {
	BountyPaymentInfo `json:"bountyPaymentInfo"`
	Rewards           BountyReward `json:"rewards"`
	StageTerms        []StageTerm  `json:"stageTerms,omitempty"`
	*PeriodTerms      `json:"periodTerms,omitempty"`
}

type BountyReward struct {
	BountyID     int    `json:"bounty_id"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1_symbol"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1_amount"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2_symbol"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2_amount"`
}

type StageTerm struct {
	SeqNum       int    `json:"seqNum" gorm:"seq_Num"`
	Status       int    `json:"status" gorm:"status"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
	Terms        string `json:"terms" gorm:"terms"`
}

type PeriodTerms struct {
	PeriodModes []PeriodMode `json:"periodModes"`
	Terms       string       `json:"terms"`
	HoursPerDay int          `json:"hoursPerDay" gorm:"hours_per_day"`
	PeriodType  int          `json:"periodType" gorm:"period_type"`
}

type PeriodMode struct {
	SeqNum       int    `json:"seqNum" gorm:"seq_Num"`
	Status       int    `json:"status" gorm:"status"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount"`
}

type PeriodInfo struct {
	Target      string `json:"target" gorm:"column:target"`
	HoursPerDay int    `json:"hoursPerDay" gorm:"column:hours_per_day"`
	PeriodType  int    `json:"periodType" gorm:"column:period_type"`
}
type ActivitiesResponse struct {
	ComerID    string    `json:"comerID" gorm:"comer_id"`
	Name       string    `json:"name" gorm:"name"`
	Avatar     string    `json:"avatar" gorm:"avatar"`
	SourceType int       `json:"sourceType" gorm:"source_type"`
	Content    string    `json:"content" gorm:"content"`
	Timestamp  time.Time `json:"timestamp" gorm:"timestamp"`
}

//type ActivityContent struct {
//	SourceType int       `json:"sourceType" gorm:"source_type"`
//	Content    string    `json:"content" gorm:"content"`
//	CreatedAt  time.Time `json:"createdAt" gorm:"created_at"`
//}

type ComerInfo struct {
	ComerID uint64 `json:"comerID" gorm:"comer_id"`
	Name    string `json:"comerName" gorm:"name"`
	Avatar  string `json:"comerImage" gorm:"avatar"`
}

type StartupListResponse struct {
	ID            uint64   `gorm:"column:id" json:"id"`
	Title         string   `gorm:"column:name" json:"title"`
	Mode          int      `gorm:"column:mode" json:"mode"`
	Logo          string   `gorm:"column:logo" json:"logo"`
	ChainID       uint64   `gorm:"column:chain_id" json:"chainID"`
	TxHash        string   `gorm:"column:tx_hash" json:"blockChainAddress"`
	ContractAudit string   `gorm:"column:contract_audit" json:"contractAudit"`
	Website       string   `gorm:"column:website" json:"website"`
	Discord       string   `gorm:"column:discord" json:"discord"`
	Twitter       string   `gorm:"column:twitter" json:"twitter"`
	Telegram      string   `gorm:"column:telegram" json:"telegram"`
	Docs          string   `gorm:"column:docs" json:"docs"`
	Mission       string   `gorm:"column:mission" json:"mission"`
	Tag           []string `gorm:"-" json:"tag"`
}

type BountyApplicantsResponse struct {
	Applicants []Applicant
}

type Applicant struct {
	ComerID     string    `json:"comerID"`
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	Applyat     time.Time `json:"applyAt"`
}

type FounderResponse struct {
	ComerID          uint64   `json:"comerID"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	ApplicantsSkills []string `json:"applicantsSkills"`
	TimeZone         string   `json:"timeZone"`
	Location         string   `gorm:"column:location" json:"location"`
	Email            string   `gorm:"column:email" json:"email"`
}

type ApprovedResponse struct {
	ComerID          uint64   `json:"comerID"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Address          string   `json:"address"`
	ApplicantsSkills []string `json:"applicantsSkills"`
}

type DepositRecordsResponse struct {
	DepositRecords []DepositRecord
}

type DepositRecord struct {
	ComerID     string    `json:"comerID"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"time" gorm:"create_at"`
	TokenAmount int       `json:"tokenAmount" gorm:"token_Amount"`
	Access      int       `json:"access"`
}

type BountyDetailStatus struct {
	Role    int  `json:"role"`
	Lock    bool `json:"lock"`
	Release bool `json:"release"`
}
