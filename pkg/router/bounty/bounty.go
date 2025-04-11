/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/25 22:54
 */

package bounty

import (
	"ceres/pkg/model"
	"ceres/pkg/model/bounty"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/bounty"
	"ceres/pkg/utility/tool"
	"strconv"
)

const (
	DepositSuccessStatus = 2
	DepositLockStatus    = 3
	DepositUnLockStatus  = 4
)

// CreateBounty create bounty
func CreateBounty(ctx *router.Context) {
	request := new(bounty.BountyRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}

	if err := service.CreateComerBounty(request); err != nil {
		ctx.HandleError(err)
		return
	}
	response := "create bounty successful!"

	ctx.OK(response)
}

// GetPublicBountyList bounty list displayed in bounty tab
func GetPublicBountyList(ctx *router.Context) {
	var request model.Pagination
	if err := model.ParsePagination(ctx, &request, 10); err != nil {
		ctx.HandleError(err)
		return
	}

	if err := service.QueryAllOnChainBounties(&request); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(request)
	}
}

// GetBountyListByStartup get bounty list belongs to startup
func GetBountyListByStartup(ctx *router.Context) {
	var pagination model.Pagination
	if err := model.ParsePagination(ctx, &pagination, 3); err != nil {
		ctx.HandleError(err)
		return
	}
	startupId, err := strconv.ParseUint(ctx.Param("startupId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	if startupId == 0 {
		err := router.ErrBadRequest.WithMsg("Invalid startupId!")
		ctx.HandleError(err)
		return
	}

	if err := service.QueryBountiesByStartup(startupId, &pagination); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(pagination)
	}
}

// GetMyPostedBountyList get bounty list posted by me
func GetMyPostedBountyList(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var pagination model.Pagination
	if err := model.ParsePagination(ctx, &pagination, 5); err != nil {
		ctx.HandleError(err)
		return
	}

	if err := service.QueryComerPostedBountyList(comerID, &pagination); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(pagination)
	}
}

// GetMyParticipatedBountyList get bounty list
func GetMyParticipatedBountyList(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var pagination model.Pagination
	if err := model.ParsePagination(ctx, &pagination, 8); err != nil {
		ctx.HandleError(err)
		return
	}

	if err := service.QueryComerParticipatedBountyList(comerID, &pagination); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(pagination)
	}
}

func GetBountyDetailByID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetBountyDetailByID(bountyID)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("get bounty detail fail")
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetPaymentByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	response, err := service.GetPaymentByBountyID(bountyID, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetState(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	response, err := service.GetBountyState(bountyID, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func UpdateBountyCloseStatus(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.UpdateBountyStatusByID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func AddDeposit(ctx *router.Context) {
	request := new(bounty.AddDepositRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	err = service.AddDeposit(bountyID, request, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("add deposit success")
}

// 支付奖金
func PayReward(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	request := new(bounty.PaidRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.PayReward(bountyID, request)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK("update paid success")
}

func CreateActivities(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	in := new(bounty.ActivitiesRequest)
	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.HandleError(err)
		return
	}
	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.CreateActivities(bountyID, in, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("activities create success")
}

func CreateApplicants(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	request := new(bounty.ApplicantsDepositRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	err = service.CreateApplicants(bountyID, request, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("create applicants success")
}

func GetActivitiesLists(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetActivitiesByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetAllApplicantsByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetAllApplicantsByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetFounderByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetFounderByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetApprovedApplicantByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetApprovedApplicantByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func GetDepositRecords(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetDepositRecords(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func UpdateFounderApprovedApplicant(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	applicantComerID, err := strconv.ParseUint(ctx.Param("applicantComerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	request := new(bounty.ApplicantsApprovedRequst)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.UpdateApplicantApprovedStatus(request, bountyID, comerID, applicantComerID, DepositLockStatus)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("approved success")
}

func UpdateFounderUnapprovedApplicant(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	applicantComerID, err := strconv.ParseUint(ctx.Param("applicantComerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}

	err = service.UpdateApplicantUnApprovedStatus(bountyID, applicantComerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("unapproved success")
}

func GetStartupByBountyID(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}
	response, err := service.GetStartupByBountyID(bountyID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(response)
}

func UpdateApplicantsLockDeposit(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.UpdateApplicantDepositLockStatus(bountyID, comerID, DepositLockStatus)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("lock deposit success")
}

func UpdateApplicantsUnlockDeposit(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	err = service.UpdateApplicantDepositLockStatus(bountyID, comerID, DepositUnLockStatus)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("unlock deposit success")
}

func ReleaseDeposit(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	request := new(bounty.ReleaseRequst)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}

	err = service.ReleaseFounderDeposit(request, bountyID, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("release deposit success")
}

func ReleaseMyDeposit(ctx *router.Context) {
	bountyID, err := strconv.ParseUint(ctx.Param("bountyID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid bounty ID")
		ctx.HandleError(err)
		return
	}

	comerID, err := tool.GetComerIDByToken(ctx)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	request := new(bounty.ReleaseMyDepositRequst)
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.HandleError(err)
		return
	}

	err = service.ReleaseComerDeposit(request, bountyID, comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK("release my deposit success")
}

// GetComerPostedBountyList get bounty list posted by me
func GetComerPostedBountyList(ctx *router.Context) {
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var pagination model.Pagination
	if err := model.ParsePagination(ctx, &pagination, 5); err != nil {
		ctx.HandleError(err)
		return
	}

	if err := service.QueryComerPostedBountyList(comerID, &pagination); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(pagination)
	}
}

// GetComerParticipatedBountyList get bounty list
func GetComerParticipatedBountyList(ctx *router.Context) {
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var pagination model.Pagination
	if err := model.ParsePagination(ctx, &pagination, 8); err != nil {
		ctx.HandleError(err)
		return
	}

	if err := service.QueryComerParticipatedBountyList(comerID, &pagination); err != nil {
		ctx.HandleError(err)
	} else {
		ctx.OK(pagination)
	}
}
