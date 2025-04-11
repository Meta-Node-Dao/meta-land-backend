package startup

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/startup"
	"ceres/pkg/router"
	"strconv"
)

func GetStartupModuleDataCount(ctx *router.Context) {
	startupId, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	info, err := startup.ProfileStartupModuleDataInfo(mysql.DB, startupId)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(info)
}
