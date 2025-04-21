package tag

import (
	"ceres/pkg/model/tag"
	"ceres/pkg/router"
	service "ceres/pkg/service/tag"
)

// GetTagList get tag list
func GetTagList(ctx *router.Context) {
	var request tag.ListRequest
	if err := ctx.BindQuery(&request); err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid data format")
		ctx.HandleError(err)
		return
	}

	if err := request.Validate(); err != nil {
		ctx.HandleError(err)
		return
	}

	var response tag.ListResponse
	//if err := service.GetStartupTagList(request, &response); err != nil {
	//	ctx.HandleError(err)
	//	return
	//}

	ctx.OK(response)
}

func GetsStartupTagList(ctx *router.Context) {
	var request tag.TagListRequest
	if err := ctx.BindQuery(&request); err != nil {
		//err = router.ErrBadRequest.WithMsg("Invalid data format")
		ctx.HandleError(err)
		return
	}

	if err := request.Validate(); err != nil {
		ctx.HandleError(err)
		return
	}

	var response tag.ListResponse
	if err := service.GetStartupTagList(request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}
