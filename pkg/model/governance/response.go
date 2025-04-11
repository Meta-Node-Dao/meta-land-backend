package governance

import (
	"github.com/shopspring/decimal"
	"time"
)

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
