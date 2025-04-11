package dict

import (
	"ceres/pkg/router"
	service "ceres/pkg/service/dict"
)

func GetDictListByType(ctx *router.Context) {
	dictType := ctx.Query("type")
	list, err := service.SelectDictDataByType(dictType)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(list)
}
