package chain

import (
	"ceres/pkg/model"
)

// Chain  Comunion support Chains.
type Chain struct {
	model.Base
	ChainID        uint64          `gorm:"column:chain_id" json:"chain_id"` // 链ID
	Name           string          `gorm:"column:name" json:"name"`         // 链名称
	Logo           string          `gorm:"column:logo" json:"logo"`         // 链图标
	Status         int             `gorm:"column:status" json:"status"`     // 链状态
	ChainContracts []ChainContract `json:"chain_contracts" gorm:"foreignKey:ChainID;references:ChainID"`
	ChainEndpoints []ChainEndpoint `json:"chain_endpoints" gorm:"foreignKey:ChainID;references:ChainID"`
}

// TableName identify the table name of this model.
func (Chain) TableName() string {
	return "chain"
}

// ChainContract  Comunion contracts.
type ChainContract struct {
	model.Base
	ChainID       uint64 `gorm:"column:chain_id" json:"chain_id"`               // 链ID
	Project       int    `gorm:"column:project" json:"project"`                 // 合约项目：1 startup, 2 bounty, 3 crowdfunding, 4 gover
	Address       string `gorm:"column:address" json:"address"`                 // 合约地址
	Type          int    `gorm:"column:type" json:"type"`                       // 合约类型 1工厂合约, 2子合约
	Version       string `gorm:"column:version" json:"version"`                 // 合约版本
	Abi           string `gorm:"column:abi" json:"abi"`                         // 合约ABI JSON
	CreatedTxHash string `gorm:"column:created_tx_hash" json:"created_tx_hash"` // 合约创建交易HASH
}

// TableName identify the table name of this model.
func (ChainContract) TableName() string {
	return "chain_contract"
}

// ChainEndpoint  endpoint of all chain list.
type ChainEndpoint struct {
	model.Base
	ChainID  uint64 `gorm:"column:chain_id" json:"chain_id"` // 链ID
	Protocol int    `gorm:"column:protocol" json:"protocol"` // 节点协议 1 RPC, 2 WSS
	URL      string `gorm:"column:url" json:"url"`           // 节点URL地址
	Status   int    `gorm:"column:status" json:"status"`     // 节点状态
}

// TableName identify the table name of this model.
func (ChainEndpoint) TableName() string {
	return "chain_endpoint"
}
