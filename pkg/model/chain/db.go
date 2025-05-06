package chain

import (
	"gorm.io/gorm"
)

// GetChainList get chain list
func GetChainList(db *gorm.DB, chains *[]ChainBasicResponse) (err error) {
	err = db.Preload("ChainContracts", func(tx *gorm.DB) *gorm.DB {
		return tx.Where(ChainContract{
			Type: 1,
		})
	}).Where(Chain{
		Status: 1,
	}).Order("created_at DESC").Find(chains).Error
	return
}

// GetChainCompleteList get chain complete list
func GetChainCompleteList(db *gorm.DB, chains *[]ChainBasicResponse) (err error) {
	err = db.
		Preload("ChainContracts", func(tx *gorm.DB) *gorm.DB {
			return tx.Where(ChainContract{
				Type: 1,
			})
		}).
		Preload("ChainEndpoints", func(tx *gorm.DB) *gorm.DB {
			return tx.Where(ChainEndpoint{
				Status: 1,
			})
		}).
		Where(Chain{
			Status: 1,
		}).Order("created_at DESC").Find(chains).Error
	return
}
