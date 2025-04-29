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

// GovernanceSetting 治理设置表结构
type GovernanceSetting struct {
	model.Base
	StartupId         uint64          `gorm:"column:startup_id;index:idx_startup" json:"startup_id"` // 关联的初创公司ID
	ComerId           uint64          `gorm:"column:comer_id" json:"comer_id"`                       // 创建者用户ID
	VoteSymbol        string          `gorm:"column:vote_symbol" json:"vote_symbol"`                 // 投票代币符号
	AllowMember       bool            `gorm:"column:allow_member" json:"allow_member"`               // 是否允许成员投票：0-否 1-是
	ProposalThreshold decimal.Decimal `gorm:"column:proposal_threshold" json:"proposal_threshold"`   // 提案阈值
	ProposalValidity  decimal.Decimal `gorm:"column:proposal_validity" json:"proposal_validity"`     // 提案有效期(天)

}

func (receiver GovernanceSetting) TableName() string {
	return "governance_setting"
}

// GovernanceStrategy 治理策略表结构
type GovernanceStrategy struct {
	model.Base
	SettingId            uint64          `gorm:"column:setting_id;index:idx_setting" json:"setting_id"`       // 关联的设置ID
	DictValue            string          `gorm:"column:dict_value" json:"dict_value"`                         // 字典值
	StrategyName         string          `gorm:"column:strategy_name" json:"strategy_name"`                   // 策略名称
	ChainId              uint64          `gorm:"column:chain_id" json:"chain_id"`                             // 链ID
	TokenContractAddress string          `gorm:"column:token_contract_address" json:"token_contract_address"` // 代币合约地址
	VoteSymbol           string          `gorm:"column:vote_symbol" json:"vote_symbol"`                       // 投票代币符号
	VoteDecimals         int             `gorm:"column:vote_decimals" json:"vote_decimals"`                   // 投票代币精度
	TokenMinBalance      decimal.Decimal `gorm:"column:token_min_balance" json:"token_min_balance"`           // 代币最小余额

}

func (receiver GovernanceStrategy) TableName() string {
	return "governance_strategy"
}

// GovernanceAdmin 治理管理员表结构
type GovernanceAdmin struct {
	model.Base
	SettingId     uint64 `gorm:"column:setting_id" json:"setting_id"`         // 关联的设置ID
	WalletAddress string `gorm:"column:wallet_address" json:"wallet_address"` // 钱包地址

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

func (request GovernanceProposal) Valid() error {
	if request.AuthorComerID == 0 {
		return errors.New("invalid authorComerId")
	}
	if request.StartupID == 0 {
		return errors.New("invalid startupId")
	}
	if strings.TrimSpace(request.AuthorWalletAddress) == "" {
		return errors.New("authorWalletAddress can not be empty")
	}
	if request.ChainID == 0 {
		return errors.New("chainId can not be empty")
	}
	if request.BlockNumber == 0 {
		return errors.New("blockNumber can not be empty")
	}
	if strings.TrimSpace(request.IPFSHash) == "" {
		return errors.New("ipfsHash can not be empty")
	}
	if strings.TrimSpace(request.Title) == "" {
		return errors.New("title can not be empty")
	}
	return nil
}

// GovernanceProposal 治理提案表结构
type GovernanceProposal struct {
	model.Base
	StartupID           uint64    `gorm:"column:startup_id;index:idx_startup" json:"startup_id"`     // 关联的初创公司ID
	AuthorComerID       uint64    `gorm:"column:author_comer_id" json:"author_comer_id"`             // 作者用户ID
	AuthorWalletAddress string    `gorm:"column:author_wallet_address" json:"author_wallet_address"` // 作者钱包地址
	ChainID             uint64    `gorm:"column:chain_id" json:"chain_id"`                           // 链ID
	BlockNumber         uint64    `gorm:"column:block_number" json:"block_number"`                   // 区块高度
	ReleaseTimestamp    time.Time `gorm:"column:release_timestamp" json:"release_timestamp"`         // 发布时间戳
	IPFSHash            string    `gorm:"column:ipfs_hash" json:"ipfs_hash"`                         // IPFS哈希
	Title               string    `gorm:"column:title" json:"title"`                                 // 提案标题
	Description         string    `gorm:"column:description;type:text" json:"description"`           // 提案描述
	DiscussionLink      string    `gorm:"column:discussion_link" json:"discussion_link"`             // 讨论链接
	VoteSystem          string    `gorm:"column:vote_system" json:"vote_system"`                     // 投票系统
	StartTime           time.Time `gorm:"column:start_time" json:"start_time"`                       // 开始时间
	EndTime             time.Time `gorm:"column:end_time" json:"end_time"`                           // 结束时间
	Status              int8      `gorm:"column:status" json:"status"`                               // 状态:0-待定 1-即将开始 2-活跃 3-已结束
}

// TableName 指定表名
func (GovernanceProposal) TableName() string {
	return "governance_proposal"
}

type GovernanceChoices []*GovernanceChoice

type ProposalChoice struct {
	ProposalID uint64 `gorm:"column:proposal_id;index:idx_proposal" json:"proposal_id"` // 关联的提案ID
	ItemName   string `gorm:"column:item_name" json:"item_name"`                        // 选项名称
	SeqNum     int    `gorm:"column:seq_num" json:"seq_num"`                            // 排序序号
}

// GovernanceChoice 治理投票选项表结构
type GovernanceChoice struct {
	model.Base
	ProposalChoice
}

func (receiver GovernanceChoice) TableName() string {
	return "governance_choice"
}

type VoteInfo struct {
	ProposalID         uint64          `gorm:"column:proposal_id;index:idx_proposal" json:"proposal_id"`    // 关联的提案ID
	VoterComerID       uint64          `gorm:"column:voter_comer_id;index:idx_voter" json:"voter_comer_id"` // 投票者用户ID
	VoterWalletAddress string          `gorm:"column:voter_wallet_address" json:"voter_wallet_address"`     // 投票者钱包地址
	ChoiceItemID       uint64          `gorm:"column:choice_item_id" json:"choice_item_id"`                 // 选择的选项ID
	ChoiceItemName     string          `gorm:"column:choice_item_name" json:"choice_item_name"`             // 选项名称
	Votes              decimal.Decimal `gorm:"column:votes" json:"votes"`                                   // 投票权重/数量
	IPFSHash           string          `gorm:"column:ipfs_hash" json:"ipfs_hash"`                           // IPFS哈希(存证)
}

// GovernanceVote 治理投票记录表结构
type GovernanceVote struct {
	model.Base
	VoteInfo
}

func (receiver GovernanceVote) TableName() string {
	return "governance_vote"
}
