package chain

import (
	"ceres/pkg/model/chain"
	"ceres/pkg/router"
	service "ceres/pkg/service/chain"
)

// GetChainList get all chain list
func GetChainList(ctx *router.Context) {
	var response chain.ChainListResponse
	if err := service.GetChainList(&response); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}
