package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"errors"
	"strconv"
	"strings"

	"github.com/qiniu/x/log"
)

func CreateProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.CreateProfileRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.CreateComerProfile(comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

func GetProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response model.ComerProfileResponse
	if err := service.GetComerProfile(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

func UpdateProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.UpdateProfileRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	if err := service.UpdateComerProfile(comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

func UpdateCover(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.UpdateComerCoverRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(router.ErrBadRequest.WithMsg(err.Error()))
		return
	}

	if err := service.UpdateComerCover(comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

func GetModulesOfTargetComer(ctx *router.Context) {
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	if targetComerID == 0 {
		ctx.HandleError(router.ErrBadRequest.WithMsg("invalid comer id"))
		return
	}
	info, err := service.GetComerModuleInfo(targetComerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(info)
}

func UpdateComerSkill(ctx *router.Context) {
	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.UpdateSkillsRequest
	if er := ctx.ShouldBindJSON(&request); er != nil {
		ctx.HandleError(er)
		return
	}
	if err := service.UpdateComerSkill(comerId, request.Skills); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func UpdateComerBio(ctx *router.Context) {
	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.UpdateBioRequest
	if er := ctx.ShouldBindJSON(&request); er != nil {
		ctx.HandleError(er)
		return
	}
	if err := service.UpdateComerBio(comerId, request.Bio); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

//func UpdateComerLanguages(ctx *router.Context) {
//	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var request model.UpdateLanguageInfosRequest
//	if er := ctx.ShouldBindJSON(&request); er != nil {
//		ctx.HandleError(er)
//		return
//	}
//	if err := service.UpdateLanguages(comerId, request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}

func UpdateComerEducations(ctx *router.Context) {
	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.UpdateEducationsRequest
	if er := ctx.ShouldBindJSON(&request); er != nil {
		ctx.HandleError(er)
		return
	}

	if err := service.UpdateEducations(comerId, request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func UpdateBasic(ctx *router.Context) {
	var request model.UpdateBasicInfoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if err := service.CreateOrUpdateBasic(comerId, request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func ProfileComerConnectedInfo(ctx *router.Context) {
	comerId, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		ctx.HandleError(errors.New("invalid param comerID"))
	}
	info, err := model.ProfileComerConnectedInfo(mysql.DB, comerId)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(info)
}

func ProfileComerModuleDataCntInfo(ctx *router.Context) {
	comerId, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		ctx.HandleError(errors.New("invalid param comerID"))
	}
	dataTypeStr := ctx.Query("type")
	if strings.TrimSpace(dataTypeStr) == "" {
		ctx.HandleError(errors.New("invalid param dataType"))
	}
	dataType, err := strconv.Atoi(dataTypeStr)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	info, err := model.ProfileComerModuleDataInfo(mysql.DB, comerId, model.BusinessModuleDataType(dataType))
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(info)
}
