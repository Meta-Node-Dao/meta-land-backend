package image

import (
	"ceres/pkg/model/image"
	"ceres/pkg/router"
	service "ceres/pkg/service/image"
	"fmt"
)

// GetImageList get image list
func GetImageList(ctx *router.Context) {
	var request image.ListRequest
	if err := ctx.BindQuery(&request); err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid data format")
		ctx.HandleError(err)
		return
	}

	fmt.Println(request)

	if err := request.Validate(); err != nil {
		ctx.HandleError(err)
		return
	}

	var response image.ListResponse
	if err := service.GetImageList(request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}
