package eth

import (
	model "ceres/pkg/model/chain"
	service "ceres/pkg/service/chain"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qiniu/x/log"
)

type ChainInfo struct {
	ChainID                uint64 `json:"chain_id"`
	RPCURL                 string `json:"rpc_url"`
	WSSURL                 string `json:"wss_url"`
	StartupContractAddress string `json:"startup_contract_address"`
	Abi                    string `json:"abi"`
}

type Client struct {
	RPCClient *ethclient.Client
	WSSClient *ethclient.Client
	ChainInfo *ChainInfo
}

var Clients map[uint64]*Client

var EthSubChanel = make(chan struct{})

func GetClient(chainID uint64) (*Client, error) {
	if client, ok := Clients[chainID]; ok {
		return client, nil
	}
	return nil, fmt.Errorf("error: GetEthClient error; chain id: %v", chainID)
}

// init all chain client
func Init() error {
	var resp model.ListResponse
	err := service.GetChainCompleteList(&resp)
	if err != nil {
		log.Warn(err)
		return err
	}
	chains := resp.List
	Clients = make(map[uint64]*Client)
	for _, chain := range chains {
		Clients[chain.ChainID] = &Client{}
		chainInfo := &ChainInfo{
			ChainID: chain.ChainID,
		}
		for _, endpoint := range chain.ChainEndpoints {
			if endpoint.Protocol == 1 {
				chainInfo.RPCURL = endpoint.URL
			} else {
				chainInfo.WSSURL = endpoint.URL
			}
		}
		for _, contract := range chain.ChainContracts {
			if contract.Project == 1 && contract.Type == 1 {
				chainInfo.StartupContractAddress = contract.Address
				chainInfo.Abi = contract.Abi
			}
		}
		Clients[chain.ChainID].ChainInfo = chainInfo
		setClient(Clients[chain.ChainID])
	}
	// EthSubChanel <- struct{}{}
	// log.Info("eth.Init EthSubChanel <- struct{}{}")
	return nil
}

// Init the eth client
func setClient(client *Client) {
	var err error

	// 介于有些节点找不到WSS服务 故不再使用 WSS
	// log.Info("eth.Init ethclient_wss.Dial:", client.ChainInfo.WSSURL)
	// client.WSSClient, err = ethclient.Dial(client.ChainInfo.WSSURL)
	// if err != nil {
	// 	client.WSSClient = nil
	// 	log.Warn(err)
	// }
	log.Info("eth.Init ethclient_rpc.Dial:", client.ChainInfo.RPCURL)
	client.RPCClient, err = ethclient.Dial(client.ChainInfo.RPCURL)
	if err != nil {
		client.RPCClient = nil
		log.Warn(err)
	}
	log.Info("eth.Init ethclient.Dial done")
}

func Close() {
	log.Info("eth.Close start")
	for _, client := range Clients {
		if client.WSSClient != nil {
			client.WSSClient.Close()
		}
	}
	EthSubChanel = make(chan struct{})
	log.Info("eth.Close end")
}
