/**
 * @Author: Sun
 * @Description:
 * @File:  eth_test
 * @Version: 1.0.0
 * @Date: 2022/7/6 23:27
 */

package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestInit(t *testing.T) {
	client, err := ethclient.Dial("wss://api.avax-test.network/ext/bc/C/ws")
	if err != nil {
		fmt.Println(err)
	}
	txHash := common.HexToHash("0x74c3486882ed8e7f0063165f5938b631f07cc7c0b780f31f6465b39956d14a32")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)

	fmt.Println(tx.Hash().Hex(), isPending, err)
}
