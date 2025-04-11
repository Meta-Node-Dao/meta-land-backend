/**
 * @Author: Sun
 * @Description:
 * @File:  proto
 * @Version: 1.0.0
 * @Date: 2022/7/3 11:11
 */

package transaction

import (
	"ceres/pkg/model"
	"time"
)

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
