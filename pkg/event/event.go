package event

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/eth"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/qiniu/x/log"
)

var (
	// key = "STARTUP:%v:BLOCK_NUMBER" // Store the current query block height of each chain
	// expire   = time.Second * 15          // If you can't get it within 15 second, get it again
	chainApi = map[uint64]string{
		80001: "https://api-testnet.polygonscan.com/api?module=logs&action=getLogs&fromBlock=%v&address=%v",
		97:    "https://api-testnet.bscscan.com/api?module=logs&action=getLogs&fromBlock=%v&address=%v",
	}
)

func StartListen() {
	var count = 0
	for {
		log.Info("event.StartListen No:", count)
		waitGroup := &sync.WaitGroup{}
		waitGroup.Add(1)

		listenEvent := func() {
			defer waitGroup.Done()
			SubAllEvent()
		}

		go listenEvent()
		if count != 0 {
			if err := eth.Init(); err != nil {
				log.Warn(err)
			}
		}

		log.Info("event.StartListen Wait start")
		waitGroup.Wait()
		log.Info("event.StartListen Wait over")

		eth.Close()

		log.Info("event.StartListen Sleep:", 5*time.Second)
		time.Sleep(5 * time.Second)

		count++
	}
}

func SubAllEvent() {
	log.Info("SubEvent enter")
	select {
	case <-eth.EthSubChanel:
		waitGroup := &sync.WaitGroup{}
		for chainID, client := range eth.Clients {
			if chainID > 0 && client.WSSClient != nil && client.ChainInfo.StartupContractAddress != "" && client.ChainInfo.Abi != "" {
				waitGroup.Add(1)
				listenEvent := func(client eth.Client) {
					defer waitGroup.Done()
					SubEvent(client)
				}
				go listenEvent(*client)
			} else {
				log.Warn(fmt.Sprintf("Error: ChainID: %v SubEvent error", chainID))
			}
		}
		waitGroup.Wait()
	}
}

func SubEvent(client eth.Client) {
	startupAbi, err := GetABI(client.ChainInfo.Abi)
	log.Info(fmt.Sprintf("%v: SubEvent GetABI Done", client.ChainInfo.ChainID))
	if err != nil {
		log.Error(err)
		log.Error(fmt.Sprintf("Error: ChainID: %v GetABI error", client.ChainInfo.ChainID))
		return
	}
	log.Info("listen for chainID:", client.ChainInfo.ChainID)
	log.Info("listen for contract:", client.ChainInfo.StartupContractAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(config.Eth.Epoch),
		Addresses: []common.Address{common.HexToAddress(client.ChainInfo.StartupContractAddress)},
	}
	logs := make(chan types.Log)
	sub, err := client.WSSClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Error(err)
		log.Error(fmt.Sprintf("Error: ChainID: %v client.WSSClient.SubscribeFilterLogs error", client.ChainInfo.ChainID))
		return
	}
	for {
		select {
		case err = <-sub.Err():
			log.Warn(err)
			return
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case StartupContract.EventHex:
				intr, err := startupAbi.Events[StartupContract.Event].Inputs.UnpackValues(vLog.Data)
				if err != nil {
					log.Info(err)
					continue
				}
				go HandleStartup(intr[2].(common.Address).String(), intr[1], client.ChainInfo.ChainID, vLog.TxHash.String())
			}
		}
	}
}

// In view of the query speed may ignore some unfinished transactions, so directly check the last 5000 blocks
func HandleAllClientStartup() error {
	waitGroup := &sync.WaitGroup{}
	for chainID, client := range eth.Clients {
		// cacheKey := fmt.Sprintf(key, chainID)

		// log.Info("Start handle startup:", cacheKey)
		var blockNumber *big.Int
		// var rn int64

		// Get the current highest block and store it
		recentBlockNumber, err := client.RPCClient.BlockNumber(context.Background())
		if err != nil {
			log.Error("Error: client.RPCClient.BlockNumber", err)
		} else {
			// Some chains only support checking 5000 blocks
			if recentBlockNumber > 5000 {
				blockNumber = big.NewInt(int64(recentBlockNumber) - 5000)
			}
		}

		// blockNumberStr, err := redis.Client.Get(context.Background(), cacheKey)
		// if err != nil {
		// 	if rn < 0 {
		// 		rn = 0
		// 	}
		// 	blockNumber = big.NewInt(rn)
		// 	log.Error("Error: redis.Client.Get BlockNumber error", cacheKey, err)
		// } else {
		// 	n := new(big.Int)
		// 	n, ok := n.SetString(blockNumberStr, 10)
		// 	if ok {
		// 		blockNumber = n
		// 	} else {
		// 		log.Error("Error: big.Int SetString error", err)
		// 	}
		// }

		if chainID > 0 && client.RPCClient != nil && client.ChainInfo.StartupContractAddress != "" && client.ChainInfo.Abi != "" {
			waitGroup.Add(1)
			listenEvent := func(client eth.Client, blockNumber *big.Int) {
				defer waitGroup.Done()
				HandleClientStartup(client, blockNumber)

				// err = redis.Client.Set(context.TODO(), cacheKey, recentBlockNumber, expire)
				// if err != nil {
				// 	log.Error(fmt.Sprintf("Error: redis.Client.Set: %v => %v", cacheKey, recentBlockNumber), err)
				// }
			}
			go listenEvent(*client, blockNumber)
		} else {
			log.Warn(fmt.Sprintf("Error: ChainID: %v SubEvent error", chainID))
		}
	}
	waitGroup.Wait()
	return nil
}

func HandleClientStartup(client eth.Client, blockNumber *big.Int) {
	startupAbi, err := GetABI(client.ChainInfo.Abi)
	// log.Info(fmt.Sprintf("%v: HandleClientStartup GetABI Done", client.ChainInfo.ChainID))
	if err != nil {
		log.Error(err)
		log.Error(fmt.Sprintf("Error: ChainID: %v GetABI error", client.ChainInfo.ChainID))
		return
	}
	// log.Info("FilterLogs for chainID:", client.ChainInfo.ChainID)
	// log.Info("FilterLogs for contract:", client.ChainInfo.StartupContractAddress)
	log.Info(fmt.Sprintf("FilterLogs: ChainID: %v / Contract: %v", client.ChainInfo.ChainID, client.ChainInfo.StartupContractAddress))

	// mumbai/bsc ：Query log needs to be processed separately
	var logs []types.Log
	if client.ChainInfo.ChainID == 80001 || client.ChainInfo.ChainID == 97 {
		resp, err := http.Get(fmt.Sprintf(chainApi[client.ChainInfo.ChainID], blockNumber, client.ChainInfo.StartupContractAddress))
		if err != nil {
			log.Error(err)
			return
		}
		body, _ := ioutil.ReadAll(resp.Body)
		err = resp.Body.Close()
		if err != nil {
			log.Error("resp.Body.Close error:", err)
		}
		result := make(map[string]interface{})
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Error(fmt.Sprintf("Error: ChainID: %v > Response json.Unmarshal error: %v", client.ChainInfo.ChainID, err))
			return
		}
		rl := result["result"].([]interface{})
		list := make([]interface{}, len(rl))
		for i, r := range rl {
			rMap := r.(map[string]interface{})
			// Unwanted fields affect serialization into types.Log fails
			delete(rMap, "logIndex")
			delete(rMap, "transactionIndex")
			list[i] = rMap
		}
		res, err := json.Marshal(list)
		if err != nil {
			log.Error(fmt.Sprintf("Error: ChainID: %v json.Marshal result error: %v", client.ChainInfo.ChainID, err))
		}
		err = json.Unmarshal(res, &logs)
		if err != nil {
			log.Error(fmt.Sprintf("Error: ChainID: %v  > types.Log json.Unmarshal error: %v", client.ChainInfo.ChainID, err))
			return
		}
	} else {
		query := ethereum.FilterQuery{
			FromBlock: blockNumber,
			Addresses: []common.Address{common.HexToAddress(client.ChainInfo.StartupContractAddress)},
		}
		logs, err = client.RPCClient.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error(err)
			log.Error(fmt.Sprintf("Error: ChainID: %v client.RPCClient.FilterLogs error", client.ChainInfo.ChainID))
			return
		}
	}
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case StartupContract.EventHex:
			intr, err := startupAbi.Events[StartupContract.Event].Inputs.UnpackValues(vLog.Data)
			if err != nil {
				log.Info(err)
				continue
			}
			go HandleStartup(intr[0].(common.Address).String(), intr[1], client.ChainInfo.ChainID, vLog.TxHash.String())
		}
	}
}

func GetABI(abiJSON string) (*abi.ABI, error) {
	wrapABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	return &wrapABI, nil
}
