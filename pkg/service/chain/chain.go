package tag

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/chain"

	"github.com/qiniu/x/log"
)

// GetChainList return the all chain list
func GetChainList(response *chain.ChainListResponse) (err error) {
	chainList := make([]chain.ChainBasicResponse, 0)
	err = chain.GetChainList(mysql.DB, &chainList)
	if err != nil {
		log.Warn(err)
		return
	}
	response.List = chainList
	return
}

func GetChainCompleteList(response *chain.ChainListResponse) (err error) {
	chainList := make([]chain.ChainBasicResponse, 0)
	err = chain.GetChainCompleteList(mysql.DB, &chainList)
	if err != nil {
		log.Warn(err)
		return
	}
	response.List = chainList
	return
}
