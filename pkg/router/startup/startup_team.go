package startup

import (
	model "ceres/pkg/model/startup_team"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/startup"
	"strconv"

	"github.com/qiniu/x/log"
)

// ListStartupTeamMembers get startup team member list
func ListStartupTeamMembers(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	crtComerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.ListStartupTeamMemberRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response model.ListStartupTeamMemberResponse
	if err := service.ListStartupTeamMembers(startupID, crtComerId, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// CreateStartupTeamMember add startup team member
func CreateStartupTeamMember(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var request model.CreateStartupTeamMemberRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.CreateStartupTeamMember(startupID, comerID, &request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// UpdateStartupTeamMember update startup team member's title
func UpdateStartupTeamMember(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var request model.UpdateStartupTeamMemberRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.UpdateStartupTeamMember(startupID, comerID, &request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// DeleteStartupTeamMember delete startup team member
func DeleteStartupTeamMember(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	if err := service.DeleteStartupTeamMember(startupID, comerID); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}
