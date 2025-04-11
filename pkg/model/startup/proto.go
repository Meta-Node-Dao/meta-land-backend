package startup

import (
	"ceres/pkg/model"
	"ceres/pkg/model/startup_team"
	"ceres/pkg/model/tag"
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/datatypes"
)

type Mode uint8

const (
	ModeESG Mode = 1
	ModeNGO Mode = 2
	ModeDAO Mode = 3
	ModeCOM Mode = 4
)

type Startup struct {
	model.Base
	ComerID              uint64    `gorm:"comer_id" json:"comerID"`
	Name                 string    `gorm:"name" json:"name"`
	Mode                 Mode      `gorm:"mode" json:"mode"`
	Logo                 string    `gorm:"logo" json:"logo"`
	Cover                string    `gorm:"Cover" json:"cover"`
	Mission              string    `gorm:"mission" json:"mission"`
	TokenContractAddress string    `gorm:"token_contract_address" json:"tokenContractAddress"`
	Overview             string    `gorm:"overview" json:"overview"`
	ChainID              uint64    `gorm:"chain_id" json:"chainID"`
	TxHash               string    `gorm:"tx_hash" json:"blockChainAddress"`
	OnChain              bool      `gorm:"on_chain" json:"onChain"`
	KYC                  string    `gorm:"kyc" json:"kyc"`
	ContractAudit        string    `gorm:"contract_audit" json:"contractAudit"`
	HashTags             []tag.Tag `gorm:"many2many:tag_target_rel;foreignKey:ID;joinForeignKey:TargetID;" json:"hashTags"`
	Website              string    `gorm:"website" json:"website"`
	Discord              string    `gorm:"discord" json:"discord"`
	Twitter              string    `gorm:"twitter" json:"twitter"`
	Telegram             string    `gorm:"telegram" json:"telegram"`
	Docs                 string    `gorm:"docs" json:"docs"`

	Email    string `gorm:"email" json:"email"`
	Facebook string `gorm:"facebook" json:"facebook"`
	Medium   string `gorm:"medium" json:"medium"`
	Linktree string `gorm:"linktree" json:"linktree"`

	LaunchNetwork int                              `gorm:"launch_network" json:"launchNetwork"`
	TokenName     string                           `gorm:"token_name" json:"tokenName"`
	TokenSymbol   string                           `gorm:"token_symbol" json:"tokenSymbol"`
	TotalSupply   int64                            `gorm:"total_supply" json:"totalSupply"`
	PresaleStart  NullTime                         `gorm:"presale_start" json:"presaleStart"`
	PresaleEnd    NullTime                         `gorm:"presale_end" json:"presaleEnd"`
	LaunchDate    NullTime                         `gorm:"launch_date" json:"launchDate"`
	TabSequence   datatypes.JSON                   `gorm:"tab_sequence" json:"tabSequence"`
	Wallets       []Wallet                         `json:"wallets"`
	Members       []startup_team.StartupTeamMember `json:"members"`
	MemberCount   int                              `gorm:"-" json:"memberCount"`
	Follows       []FollowRelation                 `json:"follows"`
	FollowCount   int                              `gorm:"-" json:"followCount"`
}

// TableName Startup table name for gorm
func (Startup) TableName() string {
	return "startup"
}

type Wallet struct {
	model.Base
	ComerID       uint64 `gorm:"comer_id" json:"comerID"`
	StartupID     uint64 `gorm:"startup_id" json:"startupID"`
	WalletName    string `gorm:"wallet_name" json:"walletName"`
	WalletAddress string `gorm:"wallet_address" json:"walletAddress"`
}

// TableName wallet table name for gorm
func (Wallet) TableName() string {
	return "startup_wallet"
}

type FollowRelation struct {
	model.RelationBase
	ComerID   uint64 `gorm:"comer_id" json:"comerID"`
	StartupID uint64 `gorm:"startup_id" json:"startupID"`
}

// TableName Followed table name for gorm
func (FollowRelation) TableName() string {
	return "startup_follow_rel"
}

// BasicSetting Startup security and social setting
type BasicSetting struct {
	KYC           string `gorm:"kyc" json:"kyc"`
	ContractAudit string `gorm:"contract_audit" json:"contractAudit"`
	Website       string `gorm:"website" json:"website"`
	Discord       string `gorm:"discord" json:"discord"`
	Twitter       string `gorm:"twitter" json:"twitter"`
	Telegram      string `gorm:"telegram" json:"telegram"`
	Docs          string `gorm:"docs" json:"docs"`
}

// FinanceSetting Startup finance setting
type FinanceSetting struct {
	TokenContractAddress string    `gorm:"token_contract_address" json:"tokenContractAddress"`
	LaunchNetwork        int       `gorm:"launch_network" json:"launchNetwork"`
	TokenName            string    `gorm:"token_name" json:"tokenName"`
	TokenSymbol          string    `gorm:"token_symbol" json:"tokenSymbol"`
	TotalSupply          int64     `gorm:"total_supply" json:"totalSupply"`
	PresaleStart         time.Time `gorm:"presale_start" json:"presaleStart"`
	PresaleEnd           time.Time `gorm:"presale_end" json:"presaleEnd"`
	LaunchDate           time.Time `gorm:"launch_date" json:"launchDate"`
	//Wallets              []WalletAddress  `json:"wallets"`
}

type NullTime struct {
	sql.NullTime
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal("")
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	var s *time.Time
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Time = *s
	} else {
		v.Valid = false
	}
	return nil
}

type StartupGroup struct {
	model.RelationBase
	StartupId uint64 `gorm:"startup_id" json:"startupId"`
	ComerId   uint64 `gorm:"comer_id" json:"comerId"`
	Name      string `gorm:"name" json:"name"`
}

func (s StartupGroup) TableName() string {
	return "startup_group"
}

type StartupGroupMemberRel struct {
	model.RelationBase
	StartupId uint64 `gorm:"startup_id" json:"startupId"`
	ComerId   uint64 `gorm:"comer_id" json:"comerId"`
	GroupId   uint64 `gorm:"group_id" json:"groupId"`
}

func (s StartupGroupMemberRel) TableName() string {
	return "startup_group_member_rel"
}
