/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/28 09:02
 */

package bounty

import (
	"ceres/pkg/model"
	"time"
)

type Bounty struct {
	model.Base
	ChainID            uint64    `gorm:"column:chain_id;unique_index:chain_tx_uindex" json:"chainID"`
	TxHash             string    `gorm:"column:tx_hash;unique_index:chain_tx_uindex" json:"txHash"`
	DepositContract    string    `gorm:"column:deposit_contract" json:"depositContract"`
	StartupID          uint64    `gorm:"column:startup_id" json:"startupID"`
	ComerID            uint64    `gorm:"column:comer_id" json:"comerID"`
	Title              string    `gorm:"column:title" json:"title"`
	ApplyCutoffDate    time.Time `gorm:"column:apply_cutoff_date" json:"expiresIn"`
	DiscussionLink     string    `gorm:"column:discussion_link" json:"discussionLink"`
	DepositTokenSymbol string    `gorm:"column:deposit_token_symbol" json:"depositTokenSymbol"`
	ApplicantDeposit   int       `gorm:"column:applicant_deposit" json:"applicationDeposit"`
	FounderDeposit     int       `gorm:"column:founder_deposit" json:"founderDeposit"`
	Description        string    `gorm:"column:description" json:"description"`
	PaymentMode        int       `gorm:"column:payment_mode" json:"paymentMode"`
	Status             int       `gorm:"column:status" json:"status"`
	TotalRewardToken   int       `gorm:"column:total_reward_token" json:"totalRewardToken"`
}

// TableName the Bounty table name for gorm
func (Bounty) TableName() string {
	return "bounty"
}

type BountyApplicant struct {
	model.RelationBase
	BountyID    uint64    `gorm:"column:bounty_id" json:"bountyID"`
	ComerID     uint64    `gorm:"column:comer_id" json:"comerID"`
	ApplyAt     time.Time `gorm:"column:apply_at" json:"applyAt"`
	RevokeAt    time.Time `gorm:"column:revoke_at" json:"revokeAt"`
	ApproveAt   time.Time `gorm:"column:approve_at" json:"approveAt"`
	QuitAt      time.Time `gorm:"column:quit_at" json:"quitAt"`
	SubmitAt    time.Time `gorm:"column:submit_at" json:"submitAt"`
	Status      int       `gorm:"column:status" json:"status"`
	Description string    `gorm:"column:description" json:"description"`
}

// TableName the BountyApplicant table name for gorm
func (BountyApplicant) TableName() string {
	return "bounty_applicant"
}

type BountyApplicantForBounty struct {
	model.RelationBase
	BountyID    uint64    `gorm:"column:bounty_id" json:"bountyID"`
	ComerID     uint64    `gorm:"column:comer_id" json:"comerID"`
	ApplyAt     time.Time `gorm:"column:apply_at"`
	Status      int       `gorm:"column:status" json:"status"`
	Description string    `gorm:"column:description" json:"description"`
}

func (BountyApplicantForBounty) TableName() string {
	return "bounty_applicant"
}

type BountyContact struct {
	model.RelationBase
	BountyID       uint64 `gorm:"column:bounty_id;unique_index:bounty_contact_uindex" json:"bountyID"`
	ContactType    uint8  `gorm:"column:contact_type;unique_index:bounty_contact_uindex" json:"contactType"`
	ContactAddress string `gorm:"column:contact_address;unique_index:bounty_contact_uindex" json:"contactAddress"`
}

// TableName the BountyContact table name for gorm
func (BountyContact) TableName() string {
	return "bounty_contact"
}

type BountyDeposit struct {
	model.RelationBase
	ChainID     uint64    `gorm:"column:chain_id;unique_index:chain_tx_uindex" json:"chainID"`
	TxHash      string    `gorm:"column:tx_hash;unique_index:chain_tx_uindex" json:"txHash"`
	Status      int       `gorm:"column:status" json:"status"`
	BountyID    uint64    `gorm:"column:bounty_id" json:"bountyID"`
	ComerID     uint64    `gorm:"column:comer_id" json:"comerID"`
	Access      int       `gorm:"column:access" json:"access"`
	TokenSymbol string    `gorm:"column:token_symbol" json:"tokenSymbol,omitempty"`
	TokenAmount int       `gorm:"column:token_amount" json:"tokenAmount,omitempty"`
	TimeStamp   time.Time `gorm:"column:timestamp"`
}

// TableName the BountyDeposit table name for gorm
func (BountyDeposit) TableName() string {
	return "bounty_deposit"
}

type BountyPaymentPeriod struct {
	model.RelationBase
	BountyID     uint64 `gorm:"column:bounty_id;unique_index:bounty_id_uindex" json:"bountyID"`
	PeriodType   uint8  `gorm:"column:period_type" json:"periodType"`
	PeriodAmount int64  `gorm:"column:period_amount" json:"periodAmount"`
	HoursPerDay  int    `gorm:"column:hours_per_day" json:"hoursPerDay"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
	Target       string `gorm:"column:target" json:"target"`
}

// TableName the BountyPaymentPeriod table name for gorm
func (BountyPaymentPeriod) TableName() string {
	return "bounty_payment_period"
}

type BountyPaymentTerms struct {
	model.RelationBase
	BountyID     uint64 `gorm:"column:bounty_id" json:"bountyID"`
	PaymentMode  int    `gorm:"column:payment_mode" json:"paymentMode"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
	Terms        string `gorm:"column:terms" json:"terms"`
	SeqNum       int    `gorm:"column:seq_num" json:"seqNum"`
	Status       int    `gorm:"column:status" json:"status"`
}

// TableName the BountyPaymentTerms table name for gorm
func (BountyPaymentTerms) TableName() string {
	return "bounty_payment_terms"
}

type Transaction struct {
	model.RelationBase
	ChainID    uint64    `gorm:"column:chain_id;unique_index:chain_tx_uindex" json:"chainID"`
	TxHash     string    `gorm:"column:tx_hash;unique_index:chain_tx_uindex" json:"txHash"`
	TimeStamp  time.Time `gorm:"column:timestamp"`
	Status     int       `gorm:"column:status" json:"status,omitempty"` // 0:Pending 1:Success 2:Failure
	SourceType int       `gorm:"column:source_type" json:"sourceType"`
	SourceID   int64     `gorm:"column:source_id" json:"sourceID"`
	RetryTimes int       `gorm:"column:retry_times" json:"retryTimes"`
}

// TableName the Transaction table name for gorm
func (Transaction) TableName() string {
	return "transaction"
}

type PostUpdate struct {
	model.RelationBase
	SourceType int       `gorm:"sourceType"`
	SourceID   uint64    `gorm:"sourceID"`
	ComerID    uint64    `gorm:"comerID"`
	Content    string    `gorm:"column:content"`
	TimeStamp  time.Time `gorm:"column:timestamp"` // post time
}

// TableName the PostUpdate table name for gorm
func (PostUpdate) TableName() string {
	return "post_update"
}
