package governance

import (
	"ceres/pkg/model/governance"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/governance"
	"errors"
)

func CreateGovernanceSetting(ctx *router.Context) {
	startupId, err := uintPathVariable(ctx, "startupID")
	if err != nil {
		ctx.HandleError(err)
		return
	}
	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if comerId == 0 {
		ctx.HandleError(errors.New("invalid comerId"))
		return
	}
	var request governance.CreateOrUpdateGovernanceSettingRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	if _, err := service.CreateStartupGovernanceSetting(comerId, startupId, request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

// GetGovernanceSetting get startup governance setting
func GetGovernanceSetting(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "startupID", func(startupId uint64) {
		detail, err := service.GetStartupGovernanceSetting(startupId)
		if err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(detail)
	})
}
