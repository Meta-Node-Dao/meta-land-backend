package transaction

import (
	"ceres/pkg/model"
	"ceres/pkg/model/crowdfunding"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qiniu/x/log"
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

// TxDetail 交易详情
type TxDetail struct {
	Status           uint64
	Block            string
	Timestamp        string
	From             string
	InteractedWithTo *common.Address
	TokensTransferred
	Value                uint64
	TransactionFee       uint64
	GasLimit             uint64
	GasUsedbyTransaction uint64
	BaseFeePerGas        uint64
	MaxFeePerGas         uint64
	MaxPriorityFeePerGas uint64
	BurnedFees           uint64
	TxnSavings           uint64
	GasPrice             uint64
	TxnType              uint8
	Nonce                uint64
	InputData            *hexutil.Bytes
}

// TokensTransferred ...
type TokensTransferred struct {
	To  *common.Address
	For uint64
}

// GetTXDetails ...
func GetTXDetails(txHash string) (string, error) {
	RPCClient, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	if err != nil {
		log.Panic(err)
	}

	blockHash := common.HexToHash(txHash)

	v, isPending, err := RPCClient.TransactionByHash(context.Background(), blockHash)
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return "", err
	}
	if isPending {
		log.Warn("is pending:", isPending)
	}

	receipt, err := RPCClient.TransactionReceipt(context.Background(), blockHash)
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return "", err
	}

	//fmt.Println("Status", receipt.Status) // Status

	var to *common.Address
	for _, v := range receipt.Logs {
		if v.Address.String() != "" {
			to = &(v.Address)
			break
		}
	}
	// fmt.Println("to", to)
	//
	// fmt.Println("Logs") // Logs
	//
	// fmt.Println("receipt.GasUsed", receipt.GasUsed) //Gas Used by Transaction
	//
	// fmt.Println("Block Number", receipt.BlockNumber.String()) // Block

	block, err := RPCClient.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return "", err
	}

	//for _, v := range block.Body().Transactions {
	//fmt.Println("Gas Price", v.GasPrice())      // Gas Price
	//fmt.Println("Gas limit", v.Gas())           // Gas limit
	//fmt.Println("Value", v.Value().Uint64())    // Value
	//fmt.Println("Nonce", v.Nonce())             // Nonce
	//fmt.Println("Interacted With (To)", v.To()) // Interacted With (To)
	//fmt.Println("Cost", v.Cost())
	//fmt.Println("Max Priority Fee Per Gas", v.GasTipCap()) // Max Priority Fee Per Gas
	//fmt.Println("Max Fee", v.GasFeeCap())                  // Max Fee Per Gas
	//fmt.Println("Txn Mode", v.Mode())                      // Txn Mode
	// From
	// Tokens Transferred
	// Base Fee Per Gas

	bts, err := v.MarshalJSON()
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return "", err
	}
	// fmt.Println("v", string(bts))

	type Input struct {
		Data *hexutil.Bytes `json:"input"`
	}
	var input Input

	if err := json.Unmarshal(bts, &input); err != nil {
		log.Errorf("unexpected error: %v", err)
		return "", err
	}

	fmt.Println("---------------------------------")
	//}

	tm := time.Unix(int64(block.Time()), 0).UTC()
	//fmt.Println("tm", tm)

	res := &TxDetail{
		Status:           receipt.Status,
		Block:            receipt.BlockNumber.String(),
		Timestamp:        tm.Format("Jan-02-2006 15:04:05 AM +UTC"),
		From:             "",
		InteractedWithTo: v.To(),
		TokensTransferred: TokensTransferred{
			To:  to,
			For: 0,
		},
		Value:                v.Value().Uint64(),
		TransactionFee:       0,
		GasLimit:             v.Gas(),
		GasUsedbyTransaction: receipt.GasUsed,
		BaseFeePerGas:        block.BaseFee().Uint64(),
		MaxFeePerGas:         v.GasFeeCap().Uint64(),
		MaxPriorityFeePerGas: v.GasTipCap().Uint64(),
		BurnedFees:           0,
		TxnSavings:           0,
		GasPrice:             0,
		TxnType:              v.Type(),
		Nonce:                v.Nonce(),
		InputData:            input.Data,
	}

	bts, err = json.MarshalIndent(res, "", "\t\r")
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return "", err
	}

	// fmt.Println("json body", string(bts))

	return string(bts), nil
}

func xxxxxxx(txHashString string) (contractAddress string, status int) {
	RPCClient, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	if err != nil {
		log.Panic(err)
	}

	txHash := common.HexToHash(txHashString)
	tx, isPending, err := RPCClient.TransactionByHash(context.Background(), txHash)
	log.Infof("#####[TRANSACTION] TransactionByHash: tx-> %v, isPending-> %v, err-> %v \n", tx, isPending, err)
	if err != nil {
		log.Warn(err)
		return "", Failure
	}
	if isPending == false {
		receipt, err := RPCClient.TransactionReceipt(context.Background(), tx.Hash())
		log.Infof("#####[TRANSACTION] TransactionReceipt-> %v,  err-> %v \n", receipt, err)
		if err != nil {
			log.Warn(err)
			return "", Failure
		}
		var to *common.Address
		for _, v := range receipt.Logs {
			if v.Address.String() != "" {
				to = &(v.Address)
				break
			}
		}
		if receipt.Status == ReceiptFailure {
			return "", Failure
		}

		contractAddress = to.String()

		return contractAddress, Success
	}
	return "", Pending
}
func TestGetContractAddress(t *testing.T) {
	// details, err := GetTXDetails("0x1edea9169af654b31735a61af4797bf7fcbd3defc6aabb6fea51d3b4f7bcc1cf")
	details, err := GetTXDetails("0x6cce3a47a23143ec62a69eab9db132f144787b262e581a59ceb19949b776f63e")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(details)
	// 0xae9b7e662439faeca8974ae9c083dd8b3e380c03
	// 0xae9B7e662439fAeca8974Ae9c083dD8b3e380c03

	address, status := xxxxxxx("0x6cce3a47a23143ec62a69eab9db132f144787b262e581a59ceb19949b776f63e")
	log.Println(address, status)
	type a struct {
		decimal.Decimal
	}
	var investor = crowdfunding.Investor{
		CrowdfundingId: 0,
		ComerId:        0,
	}
	investor.Invest(crowdfunding.Invest, decimal.NewFromInt(1), decimal.RequireFromString("2"))
	//fmt.Printf("investor: %v\n", investor)
	pagination := model.Pagination{}
	xxx(&pagination)
	//fmt.Printf("pagination: %v\n", pagination)

}

func xxx(p *model.Pagination) {
	p.TotalPages = 99
}
