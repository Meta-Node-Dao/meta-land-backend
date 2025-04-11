/**
 * @Author: Sun
 * @Description:
 * @File:  transaction
 * @Version: 1.0.0
 * @Date: 2022/7/3 10:37
 */

package transaction

import (
	"ceres/pkg/initialization/eth"
	"ceres/pkg/model/bounty"
	model "ceres/pkg/model/transaction"
	"context"
	"time"

	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

const (
	Pending        = 0
	Success        = 1
	Failure        = 2
	ConfirmFailure = 3

	BountyDepositContractCreated = 1
	BountyDepositAccount         = 2
	CrowdfundingContractCreated  = 3
	CrowdfundingModified         = 4
	CrowdfundingCancelled        = 5
	CrowdfundingRemoved          = 6
	CrowdfundingBought           = 7
	CrowdfundingSold             = 8
	ReceiptSuccess               = 1
	ReceiptFailure               = 0
)

func CreateTransaction(db *gorm.DB, bountyID uint64, request *bounty.BountyRequest) error {
	transaction := &model.Transaction{
		ChainID:    request.ChainID,
		TxHash:     request.TxHash,
		TimeStamp:  time.Now(),
		Status:     Pending,
		SourceType: BountyDepositContractCreated,
		RetryTimes: 1,
		SourceID:   int64(bountyID),
	}
	if err := model.CreateTransaction(db, transaction); err != nil {
		return err
	}
	return nil
}

func UpdateBountyContractAndTransactoinStatus(tx *gorm.DB, bountyID uint64, status int, contractAddress string) {
	log.Infof("#####UpdateBountyContractAndTransactoinStatus: bountyId-> %d, status->%d, contractAddress->%s", bountyID, status, contractAddress)
	err := model.UpdateTransactionStatus(tx, bountyID, status)
	if err != nil {
		log.Warn(err)
	}
	err = bounty.UpdateBountyDepositContract(tx, bountyID, contractAddress)
	if err != nil {
		log.Warn(err)
	}
	err = bounty.UpdateBountyDepositStatus(tx, bountyID, status)
	if err != nil {
		log.Warn(err)
	}
}

func GetContractAddress(chainID uint64, txHashString string) (contractAddress string, status int) {
	txHash := common2.HexToHash(txHashString)
	client, err := eth.GetClient(chainID)
	if err != nil {
		log.Warn(err)
		return "", Failure
	}
	tx, isPending, err := client.RPCClient.TransactionByHash(context.Background(), txHash)
	log.Infof("#####[TRANSACTION] TransactionByHash:%s ->  isPending-> %v, err-> %v \n", txHash, isPending, err)
	if err != nil {
		log.Warn(err)
		return "", Failure
	}
	if !isPending {
		receipt, err := client.RPCClient.TransactionReceipt(context.Background(), tx.Hash())
		log.Infof("#####[TRANSACTION] TransactionReceipt: %s -> Mode: %v,  err-> %v \n", txHash, receipt.Type, err)
		if err != nil {
			log.Warn(err)
			return "", Failure
		}
		var to *common2.Address
		for _, v := range receipt.Logs {
			if v.Address.String() != "" {
				to = &(v.Address)
				break
			}
		}
		if receipt.Status == ReceiptFailure {
			return "", ConfirmFailure
		}

		if to != nil {
			contractAddress = to.String()
		} else {
			log.Warnf("#####[TRANSACTION] TransactionReceipt Has Empty Field To: %s \n", txHash)
		}

		return contractAddress, Success
	}
	return "", Pending
}

func GetTxHashStatus(chainID uint64, txHashString string) (status int) {
	txHash := common2.HexToHash(txHashString)
	client, err := eth.GetClient(chainID)
	if err != nil {
		log.Warn(err)
		return Failure
	}
	tx, isPending, err := client.RPCClient.TransactionByHash(context.Background(), txHash)
	log.Infof("#####[TRANSACTION] TransactionByHash: tx-> %v, isPending-> %v, err-> %v \n", tx, isPending, err)
	if err != nil {
		log.Warn(err)
		return Failure
	}
	if !isPending {
		receipt, err := client.RPCClient.TransactionReceipt(context.Background(), tx.Hash())
		log.Infof("#####[TRANSACTION] TransactionReceipt-> %v,  err-> %v \n", receipt, err)
		if err != nil {
			log.Warn(err)
			return Failure
		}
		if receipt.Status == ReceiptFailure {
			return Failure
		}

		return Success
	}
	return Pending
}
