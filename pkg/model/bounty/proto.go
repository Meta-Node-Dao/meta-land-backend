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

// Bounty 赏金任务表结构
type Bounty struct {
	model.Base
	ChainID            uint64    `gorm:"column:chain_id;uniqueIndex:chain_tx_uindex" json:"chain_id"` // 链ID，复合唯一索引
	TxHash             string    `gorm:"column:tx_hash;uniqueIndex:chain_tx_uindex" json:"tx_hash"`   // 交易哈希，复合唯一索引
	DepositContract    string    `gorm:"column:deposit_contract" json:"deposit_contract"`             // 合约地址
	StartupID          uint64    `gorm:"column:startup_id" json:"startup_id"`                         // 初创公司ID
	ComerID            uint64    `gorm:"column:comer_id" json:"comer_id"`                             // 用户ID
	Title              string    `gorm:"column:title" json:"title"`                                   // 任务标题
	ApplyCutoffDate    time.Time `gorm:"column:apply_cutoff_date" json:"apply_cutoff_date"`           // 申请截止日期
	DiscussionLink     string    `gorm:"column:discussion_link" json:"discussion_link"`               // 讨论链接
	DepositTokenSymbol string    `gorm:"column:deposit_token_symbol" json:"deposit_token_symbol"`     // 质押代币符号
	ApplicantDeposit   int       `gorm:"column:applicant_deposit" json:"applicant_deposit"`           // 申请人质押金额
	FounderDeposit     int       `gorm:"column:founder_deposit" json:"founder_deposit"`               // 创始人质押金额
	Description        string    `gorm:"column:description" json:"description"`                       // 任务描述
	PaymentMode        int       `gorm:"column:payment_mode" json:"payment_mode"`                     // 支付方式
	Status             int8      `gorm:"column:status" json:"status"`                                 // 任务状态
	TotalRewardToken   int       `gorm:"column:total_reward_token" json:"total_reward_token"`         // 总奖励代币数
}

// TableName the Bounty table name for gorm
func (Bounty) TableName() string {
	return "bounty"
}

// BountyApplicant 赏金任务申请人表结构
type BountyApplicant struct {
	model.Base
	BountyID    uint64    `gorm:"column:bounty_id;index:idx_bounty" json:"bounty_id"` // 赏金任务ID
	ComerID     uint64    `gorm:"column:comer_id;index:idx_comer" json:"comer_id"`    // 申请人ID
	ApplyAt     time.Time `gorm:"column:apply_at" json:"apply_at"`                    // 申请时间
	RevokeAt    time.Time `gorm:"column:revoke_at" json:"revoke_at"`                  // 撤销时间
	ApproveAt   time.Time `gorm:"column:approve_at" json:"approve_at"`                // 批准时间
	QuitAt      time.Time `gorm:"column:quit_at" json:"quit_at"`                      // 退出时间
	SubmitAt    time.Time `gorm:"column:submit_at" json:"submit_at"`                  // 提交时间
	Status      int       `gorm:"column:status;index:idx_status" json:"status"`       // 申请状态
	Description string    `gorm:"column:description" json:"description"`              // 申请描述
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

// BountyContact 赏金任务联系方式表结构
type BountyContact struct {
	model.Base
	BountyID       uint64 `gorm:"column:bounty_id;uniqueIndex:bounty_contact_uindex" json:"bounty_id"`             // 赏金任务ID
	ContactType    uint8  `gorm:"column:contact_type;uniqueIndex:bounty_contact_uindex" json:"contact_type"`       // 联系方式类型
	ContactAddress string `gorm:"column:contact_address;uniqueIndex:bounty_contact_uindex" json:"contact_address"` // 联系地址
}

// TableName the BountyContact table name for gorm
func (BountyContact) TableName() string {
	return "bounty_contact"
}

// BountyDeposit 赏金任务质押记录表结构
type BountyDeposit struct {
	model.Base
	ChainID     uint64     `gorm:"column:chain_id;uniqueIndex:chain_tx_uindex" json:"chain_id"` // 链ID，复合唯一索引
	TxHash      string     `gorm:"column:tx_hash;uniqueIndex:chain_tx_uindex" json:"tx_hash"`   // 交易哈希，复合唯一索引
	Status      int8       `gorm:"column:status" json:"status"`                                 // 质押状态
	BountyID    uint64     `gorm:"column:bounty_id;index:idx_bounty" json:"bounty_id"`          // 关联的赏金任务ID
	ComerID     uint64     `gorm:"column:comer_id;index:idx_comer" json:"comer_id"`             // 用户ID
	Access      int        `gorm:"column:access" json:"access"`                                 // 访问权限
	TokenSymbol string     `gorm:"column:token_symbol" json:"token_symbol"`                     // 代币符号
	TokenAmount int        `gorm:"column:token_amount" json:"token_amount"`                     // 代币数量
	Timestamp   *time.Time `gorm:"column:timestamp" json:"timestamp"`                           // 时间戳(指针类型允许NULL)
}

// TableName the BountyDeposit table name for gorm
func (BountyDeposit) TableName() string {
	return "bounty_deposit"
}

// BountyPaymentPeriod 赏金任务支付周期表结构
type BountyPaymentPeriod struct {
	model.Base
	BountyID     uint64 `gorm:"column:bounty_id;uniqueIndex:bounty_id_uindex" json:"bounty_id"` // 赏金任务ID（唯一索引）
	PeriodType   uint8  `gorm:"column:period_type" json:"period_type"`                          // 周期类型
	PeriodAmount int64  `gorm:"column:period_amount" json:"period_amount"`                      // 周期数量
	HoursPerDay  int    `gorm:"column:hours_per_day" json:"hours_per_day"`                      // 每日小时数
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1_symbol"`                      // 代币1符号
	Token1Amount int    `gorm:"column:token1_amount" json:"token1_amount"`                      // 代币1数量
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2_symbol"`                      // 代币2符号
	Token2Amount int    `gorm:"column:token2_amount" json:"token2_amount"`                      // 代币2数量
	Target       string `gorm:"column:target" json:"target"`                                    // 目标描述
}

// TableName the BountyPaymentPeriod table name for gorm
func (BountyPaymentPeriod) TableName() string {
	return "bounty_payment_period"
}

// BountyPaymentTerms 赏金任务支付条款表结构
type BountyPaymentTerms struct {
	model.Base
	BountyID     uint64 `gorm:"column:bounty_id;index:idx_bounty" json:"bounty_id"` // 关联的赏金任务ID
	PaymentMode  int8   `gorm:"column:payment_mode" json:"payment_mode"`            // 支付方式
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1_symbol"`          // 第一种代币符号
	Token1Amount int    `gorm:"column:token1_amount" json:"token1_amount"`          // 第一种代币数量
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2_symbol"`          // 第二种代币符号
	Token2Amount int    `gorm:"column:token2_amount" json:"token2_amount"`          // 第二种代币数量
	Terms        string `gorm:"column:terms" json:"terms"`                          // 支付条款详情
	SeqNum       int    `gorm:"column:seq_num" json:"seq_num"`                      // 排序序号
	Status       int    `gorm:"column:status" json:"status"`                        // 状态
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
