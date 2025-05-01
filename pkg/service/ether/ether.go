/**
* @Author: Sun
* @Description:
* @File:  ether
* @Version: 1.0.0
* @Date: 2022/7/3 10:39
 */

package ether

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/crowdfunding"
	modelTransaction "ceres/pkg/model/transaction"
	crowdfunding2 "ceres/pkg/service/crowdfunding"
	serviceTransaction "ceres/pkg/service/transaction"
	"context"
	"fmt"
	"github.com/gotomicro/ego/task/ecron"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
	"runtime/debug"
	"strings"
	"time"
)

const RetryThreshold = 3

func Recover() {
	if err := recover(); err != nil {
		s := string(debug.Stack())
		log.Error("recover: err=%v\n stack=%s", err, s)
	}
}

// GetAllContractAddresses  todo need refactor...
func GetAllContractAddresses() ecron.Ecron {
	job := func(ctx context.Context) error {
		transactions, err := modelTransaction.GetPendingTransactions(mysql.DB)
		log.Info("####GET ALL TRANSACTION_BY_STATUS:", transactions)
		if err != nil {
			return err
		}

		for _, tran := range transactions {
			contractAddress, status := serviceTransaction.GetContractAddress(tran.ChainID, tran.TxHash)
			time.Sleep(5 * time.Second)
			log.Info("##### handle transaction ", tran, "the contractAddress is:", contractAddress, "the status is :", status)
			sourceID := tran.SourceID
			switch tran.SourceType {
			case serviceTransaction.BountyDepositContractCreated:
				go serviceTransaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, sourceID, status, contractAddress)
			case serviceTransaction.CrowdfundingContractCreated:
				go func() {
					defer Recover()

					err := func(onChainStatus int, address string) error {
						if onChainStatus == 1 && strings.TrimSpace(address) != "" {

							entity, err := crowdfunding.GetCrowdfundingById(mysql.DB, sourceID)
							if err != nil {
								return err
							}
							var startTime = entity.StartTime
							var st = crowdfunding.Upcoming
							if startTime.After(time.Now().Add(-10*time.Second)) && startTime.Before(time.Now().Add(10*time.Second)) {
								st = crowdfunding.Live
							}
							if err := crowdfunding.UpdateCrowdfundingContractAddressAndStatus(mysql.DB, sourceID, address, st); err != nil {
								return err
							}

							if err := modelTransaction.UpdateTransactionStatusById(mysql.DB, tran.TransactionId, status); err != nil {
								return err
							}
							return nil
						}
						return postOnChainFailure(mysql.DB, tran, onChainStatus)
					}(status, contractAddress)
					if err != nil {
						log.Errorf("#### handle transaction CrowdfundingContractCreated error: %v\n", err)
					}
				}()
			case serviceTransaction.CrowdfundingModified:
				go func() {
					defer Recover()

					err := func(onChainStatus int, address string) error {
						if onChainStatus == 1 {

							log.Infof("#### Crowdfunding modified successfully, then update crowdfunding with history, %d\n", sourceID)

							history, err := crowdfunding.GetIboRateHistoryById(mysql.DB, sourceID)
							if err != nil {
								return err
							}
							if err := crowdfunding.UpdateCrowdfunding(mysql.DB, history.CrowdfundingId, crowdfunding.ModifyRequest{
								TransactionHashRequest: crowdfunding.TransactionHashRequest{},
								SwapPercent:            history.SwapPercent,
								BuyPrice:               history.BuyPrice,
								MaxBuyAmount:           history.MaxBuyAmount,
								MaxSellPercent:         history.MaxSellPercent,
								EndTime:                history.EndTime,
							}); err != nil {
								return err
							}

							if err := modelTransaction.UpdateTransactionStatusById(mysql.DB, tran.TransactionId, status); err != nil {
								return err
							}
							return nil
						}
						return postOnChainFailure(mysql.DB, tran, onChainStatus)
					}(status, contractAddress)
					if err != nil {
						log.Errorf("#### handle transaction CrowdfundingContractCreated error: %v\n", err)
					}
				}()
			case serviceTransaction.CrowdfundingRemoved:
				go func() {
					defer Recover()

					err := func(onChainStatus int, address string) error {
						if onChainStatus == 1 {

							if err := crowdfunding.UpdateCrowdfundingStatus(mysql.DB, sourceID, crowdfunding.Ended); err != nil {
								return err
							}

							if err := modelTransaction.UpdateTransactionStatusById(mysql.DB, tran.TransactionId, status); err != nil {
								return err
							}
							return nil
						}
						return postOnChainFailure(mysql.DB, tran, onChainStatus)
					}(status, contractAddress)
					if err != nil {
						log.Errorf("#### handle transaction CrowdfundingContractCreated error: %v\n", err)
					}
				}()
			case serviceTransaction.CrowdfundingCancelled:
				go func() {
					defer Recover()

					err := func(onChainStatus int, address string) error {
						if onChainStatus == 1 {

							if err := crowdfunding.UpdateCrowdfundingStatus(mysql.DB, sourceID, crowdfunding.Cancelled); err != nil {
								return err
							}

							if err := modelTransaction.UpdateTransactionStatusById(mysql.DB, tran.TransactionId, status); err != nil {
								return err
							}
							return nil
						}
						return postOnChainFailure(mysql.DB, tran, onChainStatus)
					}(status, contractAddress)
					if err != nil {
						log.Errorf("#### handle transaction CrowdfundingContractCreated error: %v\n", err)
					}
				}()
			case serviceTransaction.CrowdfundingBought, serviceTransaction.CrowdfundingSold:
				go func() {
					defer Recover()

					err := func(onChainStatus int, address string) error {
						swap, err := crowdfunding.GetCrowdfundingSwapById(mysql.DB, sourceID)
						if err != nil {

							log.Errorf("##### GetCrowdfundingSwapById: %s, %d, %v\n", tran.TxHash, sourceID, err)
							return err
						}
						if err := crowdfunding2.HandleOnChainStateForInvestment(onChainStatus, swap, *tran); err != nil {
							return err
						}
						return postOnChainFailure(mysql.DB, tran, onChainStatus)
					}(status, contractAddress)
					if err != nil {
						log.Errorf("#### handle transaction CrowdfundingContractCreated error: %v\n", err)
					}
				}()
			default:
				panic(fmt.Sprintf("unsupported source type：%d", tran.SourceType))
			}
		}
		return nil
	}
	cron := ecron.Load("ceres.cron").Build(ecron.WithJob(job))
	return cron
}

func postOnChainFailure(tx *gorm.DB, tran *modelTransaction.GetTransactions, status int) error {
	log.Infof("#### post on chain failure %d, %s[SOURCE_TYPE:%d]...\n", tran.TransactionId, tran.TxHash, tran.SourceType)

	if status == 2 && tran.RetryTimes < RetryThreshold {
		log.Infof("#### will retry %d, %s[SOURCE_TYPE:%d]...\n", tran.TransactionId, tran.TxHash, tran.SourceType)
		return modelTransaction.UpdateTransactionStatusById(tx, tran.TransactionId, 0)
	}
	if status == 2 && tran.RetryTimes >= RetryThreshold-1 {
		if err := modelTransaction.UpdateTransactionStatusById(tx, tran.TransactionId, status); err != nil {
			log.Warnf("#### on-chain-failure: %d, %s[SOURCE_TYPE:%d]...\n", tran.TransactionId, tran.TxHash, tran.SourceType)
			return err
		}
		switch tran.SourceType {
		case serviceTransaction.CrowdfundingContractCreated:
			return crowdfunding.UpdateCrowdfundingStatus(tx, tran.SourceID, crowdfunding.OnChainFailure)
		case serviceTransaction.CrowdfundingModified:
			log.Warnf("#### modify crowdfunding failure..%d.%s..%d", tran.TransactionId, tran.TxHash, tran.SourceType)
			return nil
		case serviceTransaction.CrowdfundingRemoved:
			log.Warnf("#### modify crowdfunding failure..%d.%s..%d", tran.TransactionId, tran.TxHash, tran.SourceType)
			return nil
		case serviceTransaction.CrowdfundingCancelled:
			log.Warnf("#### modify crowdfunding failure..%d.%s..%d", tran.TransactionId, tran.TxHash, tran.SourceType)
			return nil
		case serviceTransaction.CrowdfundingBought:
			log.Warnf("#### modify crowdfunding failure..%d.%s..%d", tran.TransactionId, tran.TxHash, tran.SourceType)
			return crowdfunding.UpdateCrowdfundingSwapStatus(tx, tran.SourceID, crowdfunding.SwapFailure)
		case serviceTransaction.CrowdfundingSold:
			log.Warnf("#### modify crowdfunding failure..%d.%s..%d", tran.TransactionId, tran.TxHash, tran.SourceType)
			return crowdfunding.UpdateCrowdfundingSwapStatus(tx, tran.SourceID, crowdfunding.SwapFailure)
		default:
			log.Warnf("#### unknown failure..%d.%s..%d", tran.TransactionId, tran.TxHash, tran.SourceType)
			return nil
		}
	}

	return nil
}
