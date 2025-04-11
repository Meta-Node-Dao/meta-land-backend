package governance

import (
	"ceres/pkg/router"
	"strconv"
)

func uintPathVariable(ctx *router.Context, key string) (uint64, error) {
	return strconv.ParseUint(ctx.Param(key), 0, 64)
}

func uintPathVariableWithCb(ctx *router.Context, key string, fun func(pathVariable uint64)) {
	variable, err := strconv.ParseUint(ctx.Param(key), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	fun(variable)
}
