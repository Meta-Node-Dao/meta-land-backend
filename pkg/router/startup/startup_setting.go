package startup

import (
	model "ceres/pkg/model/startup"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/startup"
	"strconv"

	"github.com/qiniu/x/log"
)

// UpdateStartupBasicSetting update startup security and social setting
func UpdateStartupBasicSetting(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var request model.UpdateStartupBasicSettingRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.UpdateStartupBasicSetting(startupID, &request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// UpdateStartupFinanceSetting update startup finance setting
func UpdateStartupFinanceSetting(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var request model.UpdateStartupFinanceSettingRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.UpdateStartupFinanceSetting(startupID, comerID, &request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

func UpdateStartupTabSequence(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var request model.UpdateStartupTabSequenceRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(router.ErrBadRequest.WithMsg(err.Error()))
		return
	}

	if err := service.UpdateStartupTabSequence(startupID, comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

func UpdateStartupSocialAndTags(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var request model.UpdateStartupSocialsAndTagsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(router.ErrBadRequest.WithMsg(err.Error()))
		return
	}

	if err := service.UpdateSocialsAndTags(startupID, comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// RemoveStartupSocial to be deleted
func RemoveStartupSocial(ctx *router.Context) {
	//comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	//startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	//if err != nil {
	//	err = router.ErrBadRequest.WithMsg("Invalid startup ID")
	//	ctx.HandleError(err)
	//	return
	//}
	//var request account.SocialRemoveRequest
	//if err := ctx.ShouldBindJSON(&request); err != nil {
	//	ctx.HandleError(router.ErrBadRequest.WithMsg(err.Error()))
	//	return
	//}
	//
	//if err := service.UpdateSocialsAndTags(startupID, comerID, account.SocialModifyRequest{
	//	SocialType: request.SocialType,
	//	SocialLink: "",
	//}); err != nil {
	//	ctx.HandleError(err)
	//	return
	//}
	ctx.OK(nil)
}

func UpdateStartupBasicSetting1(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var request model.UpdateStartupBasicSettingRequestNew
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.UpdateStartupBasicSettingNew(startupID, comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}
