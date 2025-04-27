package governance

import (
	"ceres/pkg/model/account"
	"ceres/pkg/model/tag"
	"github.com/shopspring/decimal"
	"time"
)

type GovernanceBasicResponse struct {
	AuthorWalletAddress string                        `json:"author_wallet_address"`
	BlockNumber         int                           `json:"block_number"`
	ChainId             int                           `json:"chain_id"`
	Comer               account.ComerBasicResponse    `json:"comer"`
	ComerId             int                           `json:"comer_id"`
	Description         string                        `json:"description"`
	DiscussionLink      string                        `json:"discussion_link"`
	EndTime             string                        `json:"end_time"`
	Id                  int                           `json:"id"`
	IpfsHash            string                        `json:"ipfs_hash"`
	MaxVotes            int                           `json:"max_votes"`
	MaxVotesChoiceItem  string                        `json:"max_votes_choice_item"`
	ReleaseTimestamp    string                        `json:"release_timestamp"`
	StartTime           string                        `json:"start_time"`
	Startup             GovernanceStartupCardResponse `json:"startup"`
	StartupId           int                           `json:"startup_id"`
	Status              int                           `json:"status"`
	Title               string                        `json:"title"`
	VoteSystem          string                        `json:"vote_system"`
}

type CreateProposalResponse struct {
	AuthorComerId       int                  `json:"author_comer_id"`
	AuthorWalletAddress string               `json:"author_wallet_address"`
	BlockNumber         int                  `json:"block_number"`
	ChainId             int                  `json:"chain_id"`
	Choices             []ProposalChoiceItem `json:"choices"`
	Description         string               `json:"description"`
	DiscussionLink      string               `json:"discussion_link"`
	EndTime             int                  `json:"end_time"`
	IpfsHash            string               `json:"ipfs_hash"`
	ReleaseTimestamp    int                  `json:"release_timestamp"`
	StartTime           int                  `json:"start_time"`
	StartupId           int                  `json:"startup_id"`
	Title               string               `json:"title"`
	VoteSystem          string               `json:"vote_system"`
}

type GovernanceResponse struct {
	AuthorWalletAddress string                        `json:"author_wallet_address"`
	BlockNumber         int                           `json:"block_number"`
	ChainId             int                           `json:"chain_id"`
	Choices             []ProposalChoiceItem          `json:"choices"`
	Comer               account.ComerBasicResponse    `json:"comer"`
	ComerId             int                           `json:"comer_id"`
	Description         string                        `json:"description"`
	DiscussionLink      string                        `json:"discussion_link"`
	EndTime             string                        `json:"end_time"`
	Id                  int                           `json:"id"`
	IpfsHash            string                        `json:"ipfs_hash"`
	ReleaseTimestamp    string                        `json:"release_timestamp"`
	StartTime           string                        `json:"start_time"`
	Startup             GovernanceStartupCardResponse `json:"startup"`
	StartupId           int                           `json:"startup_id"`
	Status              int                           `json:"status"`
	Title               string                        `json:"title"`
	VoteSystem          string                        `json:"vote_system"`
}

type GovernanceVoteResponse struct {
	ChoiceItemId       int                        `json:"choice_item_id"`
	ChoiceItemName     string                     `json:"choice_item_name"`
	Comer              account.ComerBasicResponse `json:"comer"`
	Id                 int                        `json:"id"`
	IpfsHash           string                     `json:"ipfs_hash"`
	ProposalId         int                        `json:"proposal_id"`
	VoterComerId       int                        `json:"voter_comer_id"`
	VoterWalletAddress string                     `json:"voter_wallet_address"`
	Votes              int                        `json:"votes"`
}

type ProposalChoiceItem struct {
	Id         int    `json:"id"`
	ItemName   string `json:"item_name"`
	ProposalId int    `json:"proposal_id"`
	SeqNum     int    `json:"seq_num"`
	VoteTotal  int    `json:"vote_total"`
}

type GovernanceStartupCardResponse struct {
	Banner            string                       `json:"banner"`
	ChainId           int                          `json:"chain_id"`
	ComerId           int                          `json:"comer_id"`
	ContractAudit     string                       `json:"contract_audit"`
	GovernanceSetting GovernanceSettingResponse    `json:"governance_setting"`
	Id                int                          `json:"id"`
	IsConnected       bool                         `json:"is_connected"`
	Kyc               string                       `json:"kyc"`
	Logo              string                       `json:"logo"`
	Mission           string                       `json:"mission"`
	Name              string                       `json:"name"`
	OnChain           bool                         `json:"on_chain"`
	Socials           []account.SocialBookResponse `json:"socials"`
	Tags              []tag.TagRelationResponse    `json:"tags"`
	TxHash            string                       `json:"tx_hash"`
	Type              int                          `json:"type"`
}

type GovernanceSettingResponse struct {
	AllowMember       bool                    `json:"allow_member"`
	ComerId           int                     `json:"comer_id"`
	Id                int                     `json:"id"`
	ProposalThreshold int                     `json:"proposal_threshold"`
	ProposalValidity  int                     `json:"proposal_validity"`
	StartupId         int                     `json:"startup_id"`
	Strategies        ModelGovernanceStrategy `json:"strategies"`
	VoteSymbol        string                  `json:"vote_symbol"`
}

type ModelGovernanceStrategy struct {
	ChainId              int    `json:"chain_id"`
	DictValue            string `json:"dict_value"`
	Id                   int    `json:"id"`
	SettingId            int    `json:"setting_id"`
	StrategyName         string `json:"strategy_name"`
	TokenContractAddress string `json:"token_contract_address"`
	TokenMinBalance      int    `json:"token_min_balance"`
	VoteDecimals         int    `json:"vote_decimals"`
	VoteSymbol           string `json:"vote_symbol"`
}

type GovernanceSettingDetailResponse struct {
	Admins            []GovernanceAdminResponse    `json:"admins"`
	AllowMember       bool                         `json:"allow_member"`
	ComerId           int                          `json:"comer_id"`
	Id                int                          `json:"id"`
	ProposalThreshold int                          `json:"proposal_threshold"`
	ProposalValidity  int                          `json:"proposal_validity"`
	StartupId         int                          `json:"startup_id"`
	Strategies        []GovernanceStrategyResponse `json:"strategies"`
	VoteSymbol        string                       `json:"vote_symbol"`
}

type GovernanceAdminResponse struct {
	Address   string `json:"address"`
	Id        int    `json:"id"`
	SettingId int    `json:"setting_id"`
}

type GovernanceStrategyResponse struct {
	ChainId              int    `json:"chain_id"`
	DictValue            string `json:"dict_value"`
	Id                   int    `json:"id"`
	SettingId            int    `json:"setting_id"`
	StrategyName         string `json:"strategy_name"`
	TokenContractAddress string `json:"token_contract_address"`
	TokenMinBalance      int    `json:"token_min_balance"`
	VoteDecimals         int    `json:"vote_decimals"`
	VoteSymbol           string `json:"vote_symbol"`
}

type GovernanceStrategies []GovernanceStrategy
type GovernanceAdmins []*GovernanceAdmin
type GovernanceSettingDetail struct {
	GovernanceSetting
	Strategies GovernanceStrategies `json:"strategies"`
	Admins     GovernanceAdmins     `json:"admins"`
}

type ProposalVoteResult struct {
	MaximumVotesChoice   *string          `json:"maximumVotesChoice"`
	MaximumVotesChoiceId *uint64          `json:"maximumVotesChoiceId"`
	Votes                *decimal.Decimal `json:"votes"`
	InvalidResult        *string          `json:"voteResult"`
}
type ProposalPublicInfo struct {
	ProposalId          uint64         `gorm:"column:id" json:"proposalId"`
	StartupId           uint64         `gorm:"column:startup_id" json:"startupId"`
	AllowMember         bool           `gorm:"column:allow_member" json:"allowMember"`
	VoteSystem          string         `gorm:"column:vote_system" json:"voteSystem"`
	VoteSymbol          string         `gorm:"column:vote_symbol" json:"voteSymbol"`
	BlockNumber         uint64         `gorm:"column:block_number" json:"blockNumber"`
	DiscussionLink      string         `gorm:"column:discussion_link" json:"discussionLink"`
	StartupLogo         string         `gorm:"column:startup_logo" json:"startupLogo"`
	StartupName         string         `gorm:"column:startup_name" json:"startupName"`
	AuthorComerId       uint64         `gorm:"column:author_comer_id" json:"authorComerId"`
	AuthorComerAvatar   string         `gorm:"column:author_comer_avatar" json:"authorComerAvatar"`
	AuthorComerName     string         `gorm:"column:author_comer_name" json:"authorComerName"`
	AuthorWalletAddress string         `gorm:"column:author_wallet_address" json:"authorWalletAddress"`
	Title               string         `gorm:"column:title" json:"title"`
	Description         string         `gorm:"column:description" json:"description"`
	Status              ProposalStatus `gorm:"column:status" json:"status"`
	StartTime           time.Time      `gorm:"column:start_time" json:"startTime"`
	EndTime             time.Time      `gorm:"column:end_time" json:"endTime"`
}
type ProposalItem struct {
	ProposalPublicInfo
	ProposalVoteResult
}

type ChoiceVoteInfo struct {
	ChoiceId uint64           `json:"choiceId"`
	ItemName string           `json:"itemName"`
	Votes    *decimal.Decimal `json:"votes"`
	Percent  *decimal.Decimal `json:"percent"`
}

type CurrentProposalVoteResult struct {
	ChoiceVoteInfos *[]ChoiceVoteInfo `json:"choiceVoteInfos"`
	TotalVotes      *decimal.Decimal  `json:"totalVotes"`
}

type ProposalDetail struct {
	ProposalPublicInfo
	VoteSystemId uint64               `json:"voteSystemId"`
	Strategies   GovernanceStrategies `json:"strategies"`
	Admins       GovernanceAdmins     `json:"admins"`
	CurrentProposalVoteResult
	Choices GovernanceChoices `json:"choices"`
}

type VoteDetail struct {
	VoteInfo
	VoterComerAvatar string `gorm:"column:voter_comer_avatar" json:"voterComerAvatar"`
	VoterComerName   string `gorm:"column:voter_comer_name" json:"voterComerName"`
}

type VoteRecords []VoteDetail
