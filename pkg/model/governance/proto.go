package governance

import (
	"ceres/pkg/model"
	"errors"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type VoteSystem string

const (
	VoteSystemSingleChoiceVoting VoteSystem = "Single choice voting"
	VoteSystemBasicVoting        VoteSystem = "Basic voting"
)

type GovernanceSetting struct {
	model.RelationBase
	StartupId         uint64          `gorm:"startup_id" json:"startupId"`
	ComerId           uint64          `gorm:"comer_id" json:"comerId"`
	VoteSymbol        string          `gorm:"vote_symbol" json:"voteSymbol"`
	AllowMember       bool            `gorm:"allow_member" json:"allowMember"`
	ProposalThreshold decimal.Decimal `gorm:"proposal_threshold" json:"proposalThreshold"`
	ProposalValidity  decimal.Decimal `gorm:"proposal_validity" json:"proposalValidity"`
}

func (receiver GovernanceSetting) TableName() string {
	return "governance_setting"
}

type GovernanceStrategy struct {
	model.RelationBase
	SettingId            uint64          `gorm:"setting_id" json:"settingId"`
	DictValue            string          `gorm:"dict_value" json:"dictValue"`
	StrategyName         string          `gorm:"strategy_name" json:"strategyName"`
	ChainId              uint64          `gorm:"chain_id" json:"chainId"`
	TokenContractAddress string          `gorm:"token_contract_address" json:"tokenContractAddress"`
	VoteSymbol           string          `gorm:"vote_symbol" json:"voteSymbol"`
	VoteDecimals         int             `gorm:"vote_decimals" json:"voteDecimals"`
	TokenMinBalance      decimal.Decimal `gorm:"token_min_balance" json:"tokenMinBalance"`
}

func (receiver GovernanceStrategy) TableName() string {
	return "governance_strategy"
}

type GovernanceAdmin struct {
	model.Base
	SettingId     uint64 `gorm:"setting_id" json:"settingId"`
	WalletAddress string `gorm:"wallet_address" json:"walletAddress"`
}

func (receiver GovernanceAdmin) TableName() string {
	return "governance_admin"
}

type ProposalStatus int

const (
	ProposalPending ProposalStatus = iota
	ProposalUpcoming
	ProposalActive
	ProposalEnded
	ProposalInvalid
)

type GovernanceProposalModel struct {
	model.Base
	GovernanceProposalInfo
}

type GovernanceProposalInfo struct {
	StartupId           uint64         `gorm:"startup_id" json:"startupId"`
	AuthorComerId       uint64         `gorm:"author_comer_id" json:"authorComerId"`
	AuthorWalletAddress string         `gorm:"author_wallet_address" json:"authorWalletAddress"`
	ChainId             uint64         `gorm:"chain_id" json:"chainId"`
	BlockNumber         uint64         `gorm:"block_number" json:"blockNumber"`
	ReleaseTimestamp    time.Time      `gorm:"release_timestamp" json:"releaseTimestamp"`
	IpfsHash            string         `gorm:"ipfs_hash" json:"ipfsHash"`
	Title               string         `gorm:"title" json:"title"`
	Description         string         `gorm:"description" json:"description"`
	DiscussionLink      string         `gorm:"discussion_link" json:"discussionLink"`
	VoteSystem          string         `gorm:"vote_system" json:"voteSystem"`
	StartTime           time.Time      `gorm:"start_time" json:"startTime"`
	EndTime             time.Time      `gorm:"end_time" json:"endTime"`
	Status              ProposalStatus `gorm:"status" json:"status"`
}

func (request GovernanceProposalInfo) Valid() error {
	if request.AuthorComerId == 0 {
		return errors.New("invalid authorComerId")
	}
	if request.StartupId == 0 {
		return errors.New("invalid startupId")
	}
	if strings.TrimSpace(request.AuthorWalletAddress) == "" {
		return errors.New("authorWalletAddress can not be empty")
	}
	if request.ChainId == 0 {
		return errors.New("chainId can not be empty")
	}
	if request.BlockNumber == 0 {
		return errors.New("blockNumber can not be empty")
	}
	if strings.TrimSpace(request.IpfsHash) == "" {
		return errors.New("ipfsHash can not be empty")
	}
	if strings.TrimSpace(request.Title) == "" {
		return errors.New("title can not be empty")
	}
	return nil
}

func (receiver GovernanceProposalModel) TableName() string {
	return "governance_proposal"
}

type GovernanceChoices []*GovernanceChoice
type GovernanceChoice struct {
	model.Base
	ProposalChoice
}

type ProposalChoice struct {
	ProposalId uint64 `gorm:"proposal_id" json:"proposalId"`
	ItemName   string `gorm:"item_name" json:"itemName"`
	SeqNum     int    `gorm:"seq_num" json:"seqNum"`
}

func (receiver GovernanceChoice) TableName() string {
	return "governance_choice"
}

type GovernanceVote struct {
	model.RelationBase
	VoteInfo
}

type VoteInfo struct {
	ProposalId         uint64          `gorm:"proposal_id" json:"proposalId"`
	VoterComerId       uint64          `gorm:"voter_comer_id" json:"voterComerId"`
	VoterWalletAddress string          `gorm:"voter_wallet_address" json:"voterWalletAddress"`
	ChoiceItemId       uint64          `gorm:"choice_item_id" json:"choiceItemId"`
	ChoiceItemName     string          `gorm:"choice_item_name" json:"choiceItemName"`
	Votes              decimal.Decimal `gorm:"votes" json:"votes"`
	IpfsHash           string          `gorm:"ipfs_hash" json:"ipfsHash"`
}

func (receiver GovernanceVote) TableName() string {
	return "governance_vote"
}
