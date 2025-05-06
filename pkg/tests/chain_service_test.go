package pkg

import (
	"testing"
)

func TestGetChainList(t *testing.T) {
	doTest(func() {
		//var response model.ListResponse
		//err := service.GetChainList(&response)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//data, err := json.Marshal(response)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//t.Log(string(data))

		// for _, item := range response.List {
		// 	t.Log("=========================================================>")
		// 	t.Log("=====>chain:", item.Name, item.ChainID)
		// 	t.Log("=========>contract:")
		// 	for _, contract := range item.ChainContracts {
		// 		t.Log(contract.Address, contract.Project)
		// 	}
		// 	t.Log("=========>endpoint:")
		// 	for _, endpoint := range item.ChainEndpoints {
		// 		t.Log(endpoint.Protocol, endpoint.URL)
		// 	}
		// }
	})
}
