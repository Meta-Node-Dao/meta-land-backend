package tag

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/chain"

	"github.com/qiniu/x/log"
)

// GetChainList return the all chain list
func GetChainList(response *model.ListResponse) (err error) {
	chainList := make([]model.Chain, 0)
	err = model.GetChainList(mysql.DB, &chainList)
	if err != nil {
		log.Warn(err)
		return
	}
	response.List = chainList
	return
}

// GetChainList return the all chain list
func GetChainCompleteList(response *model.ListResponse) (err error) {
	chainList := make([]model.Chain, 0)
	err = model.GetChainCompleteList(mysql.DB, &chainList)
	if err != nil {
		log.Warn(err)
		return
	}
	response.List = chainList
	return
}
