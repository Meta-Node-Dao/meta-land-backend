package startup

import (
	model2 "ceres/pkg/model"
	model "ceres/pkg/model/startup"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	"ceres/pkg/service/startup"
	"strconv"
)

func CreateStartupGroup(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var request model.CreateOrUpdateStartupGroupRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	group, err := startup.CreateStartupGroup(startupID, comerID, request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(group)
}

func DeleteStartupGroup(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	groupID, err := strconv.ParseUint(ctx.Param("groupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid group ID")
		ctx.HandleError(err)
		return
	}
	err = startup.DeleteStartupGroup(groupID, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func UpdateStartupGroup(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	groupID, err := strconv.ParseUint(ctx.Param("groupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid group ID")
		ctx.HandleError(err)
		return
	}
	var request model.CreateOrUpdateStartupGroupRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := startup.UpdateStartupGroup(groupID, comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func GetStartupGroups(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid group ID")
		ctx.HandleError(err)
		return
	}

	groups, err := startup.GetStartupGroups(startupID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(groups)
}

func GetStartupGroupMembers(ctx *router.Context) {
	startupId, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	groupID, err := strconv.ParseUint(ctx.Param("groupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid group ID")
		ctx.HandleError(err)
		return
	}
	var pagination model2.Pagination
	err = model2.ParsePagination(ctx, &pagination, 10)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	err = startup.GetStartupGroupMembers(startupId, groupID, &pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

func ChangeComerGroupAndLocation(ctx *router.Context) {
	startupId, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	groupID, err := strconv.ParseUint(ctx.Param("groupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid group ID")
		ctx.HandleError(err)
		return
	}
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var request model.ModifyLocationRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	err = startup.ChangeComerGroupAndPosition(startupId, groupID, comerID, request)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}
