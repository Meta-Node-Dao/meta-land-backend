package social

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	accountService "ceres/pkg/service/account"
)

func UpdateSocial(ctx *router.Context) {
	var request model.SocialModifyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if err := accountService.UpdateSocial(comerId, request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func ClearSocial(ctx *router.Context) {
	var request model.SocialRemoveRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if err := accountService.RemoveSocial(comerId, request.SocialType); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}
