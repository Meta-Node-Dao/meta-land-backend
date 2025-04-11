package pkg

import (
	"ceres/pkg/service/transaction"
	"encoding/hex"
	"testing"
)

func TestSimpleEncoding(t *testing.T) {

	origin := "0xE2346ffF0ae08e172B2C6384F5Aa66C42c5527D9"

	res, err := hex.DecodeString(origin)
	t.Log(err)
	t.Log(res)
}

func TestGetContractAddress(t *testing.T) {
	doTest(func() {
		var chainID uint64
		txHash := "0xd1444b97e63335250ffb9117af8dc15c2b9f445ec47865467de4cbe89a5335f4"
		chainID = 43113
		contract, status := transaction.GetContractAddress(chainID, txHash)
		t.Log("contract:", contract)
		t.Log("status:", status)
	})
}
