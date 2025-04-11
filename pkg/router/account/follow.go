package account

import (
	"ceres/pkg/model"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	"ceres/pkg/service/account"
	"strconv"
)

func GetConnectorsOfComer(ctx *router.Context) {
	currentComerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var pagination model.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := account.GetConnectorsOfComer(currentComerId, targetComerID, &pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

// GetComersFollowedByComer get comers followed-by-comer
func GetComersFollowedByComer(ctx *router.Context) {
	currentComerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var pagination model.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := account.GetComersFollowedByComer(currentComerId, targetComerID, &pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

func GetComerFollowedStartups(ctx *router.Context) {
	currentComerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var pagination model.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := account.GetComerFollowedStartups(currentComerId, targetComerID, &pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}
