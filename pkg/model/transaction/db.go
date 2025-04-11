/**
 * @Author: Sun
 * @Description:
 * @File:  transaction
 * @Version: 1.0.0
 * @Date: 2022/7/3 10:57
 */

package transaction

import "gorm.io/gorm"

func CreateTransaction(db *gorm.DB, transaction *Transaction) error {
	return db.Create(&transaction).Error
}

// UpdateTransactionStatus update transaction status by corresponding source id
func UpdateTransactionStatus(db *gorm.DB, sourceID uint64, status int) error {
	return db.Model(&Transaction{}).Where("source_id = ?", sourceID).Update("status", status).Error
}

func UpdateTransactionStatusById(db *gorm.DB, transactionId uint64, status int) error {
	return db.Model(&Transaction{}).Where("id = ?", transactionId).
		Updates(map[string]interface{}{"status": status, "retry_times": gorm.Expr("retry_times + 1")}).Error
}

func UpdateTransactionStatusWithRetry(db *gorm.DB, sourceID uint64, sourceType, status int) error {
	return db.Model(&Transaction{}).Where("source_id = ? and source_type = ?", sourceID, sourceType).
		Updates(map[string]interface{}{"status": status, "retry_times": gorm.Expr("retry_times + 1")}).Error
}

func GetPendingTransactions(db *gorm.DB) (transactionResponse []*GetTransactions, err error) {
	err = db.Table("transaction").Select("id, chain_id, tx_hash, source_id, source_type, retry_times").
		Where("status = 0").Order("created_at asc").Find(&transactionResponse).Error
	if err != nil {
		return nil, err
	}
	return transactionResponse, nil
}
