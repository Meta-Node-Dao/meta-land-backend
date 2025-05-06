package chain

import (
	"ceres/pkg/model"
)

// Chain 区块链网络表结构
type Chain struct {
	model.Base
	ChainID uint64 `gorm:"column:chain_id;uniqueIndex" json:"chain_id"` // 链ID（唯一索引）
	Name    string `gorm:"column:name;not null;default:''" json:"name"` // 链名称
	Logo    string `gorm:"column:logo;not null;default:''" json:"logo"` // 链Logo
	Status  int8   `gorm:"column:status;default:1" json:"status"`       // 状态：1-正常，2-禁用
}

// TableName identify the table name of this model.
func (Chain) TableName() string {
	return "chain"
}

// ChainContract 区块链合约表结构
type ChainContract struct {
	model.Base
	ChainID       uint64 `gorm:"column:chain_id;not null;default:0" json:"chain_id"`                // 链ID
	Address       string `gorm:"column:address;not null;default:''" json:"address"`                 // 合约地址
	Project       int8   `gorm:"column:project;not null;default:0" json:"project"`                  // 项目类型：1-Startup, 2-Bounty, 3-Crowdfunding, 4-Gover
	Type          int8   `gorm:"column:type;not null;default:0" json:"type"`                        // 合约类型：1-工厂合约, 2-子合约
	Version       string `gorm:"column:version;not null;default:''" json:"version"`                 // 合约版本
	ABI           string `gorm:"column:abi;type:text;not null" json:"abi"`                          // ABI JSON
	CreatedTxHash string `gorm:"column:created_tx_hash;not null;default:''" json:"created_tx_hash"` // 创建交易哈希
}

// TableName identify the table name of this model.
func (ChainContract) TableName() string {
	return "chain_contract"
}

// ChainEndpoint 区块链节点端点表结构
type ChainEndpoint struct {
	model.Base
	Protocol int8   `gorm:"column:protocol;not null;default:0" json:"protocol"` // 通信协议：1-rpc 2-wss
	ChainID  uint64 `gorm:"column:chain_id;not null;default:0" json:"chain_id"` // 链ID
	URL      string `gorm:"column:url;not null;default:''" json:"url"`          // 节点URL
	Status   int8   `gorm:"column:status;not null;default:1" json:"status"`     // 状态：1-正常 2-禁用
}

// TableName identify the table name of this model.
func (ChainEndpoint) TableName() string {
	return "chain_endpoint"
}
