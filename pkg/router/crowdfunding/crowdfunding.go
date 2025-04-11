package crowdfunding

import (
	"ceres/pkg/model"
	crowdfundingModel "ceres/pkg/model/crowdfunding"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	"ceres/pkg/service/crowdfunding"
	"errors"
	"strconv"
)

func CreateCrowdfunding(ctx *router.Context) {
	var request crowdfundingModel.CreateCrowdfundingRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := request.ValidRequest(); err != nil {
		ctx.HandleError(err)
		return
	}
	request.ComerId = comerId(ctx)
	if err := crowdfunding.CreateCrowdfunding(request); err != nil {
		ctx.HandleError(err)
		return
	}

}

func SelectNonFundingStartups(ctx *router.Context) {
	comerId := comerId(ctx)
	startups, err := crowdfunding.SelectNonFundingStartups(comerId)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	response := map[string]interface{}{
		"list":  startups,
		"total": len(startups),
	}
	ctx.OK(response)
}

func comerId(ctx *router.Context) uint64 {
	return ctx.Keys[middleware.ComerUinContextKey].(uint64)
}

func GetCrowdfundingList(ctx *router.Context) {
	var pagination crowdfundingModel.PublicCrowdfundingListPageRequest
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}

	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	err := crowdfunding.GetCrowdfundingList(&pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

func GetCrowdfundingDetail(ctx *router.Context) {
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	detail, err := crowdfunding.GetCrowdfundingDetail(crowdfundingId)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(detail)
}

func GetMyPostedCrowdfundingList(ctx *router.Context) {
	var pagination model.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	comerId := comerId(ctx)
	err := crowdfunding.GetPostedCrowdfundingListByComer(comerId, &pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

func GetMyParticipatedCrowdfundingList(ctx *router.Context) {
	var pagination model.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	comerId := comerId(ctx)
	err := crowdfunding.GetParticipatedCrowdFundingListOfComer(comerId, &pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

func CancelCrowdfunding(ctx *router.Context) {
	var re crowdfundingModel.TransactionHashRequest
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	if err = ctx.ShouldBindJSON(&re); err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := comerId(ctx)
	err = crowdfunding.CancelCrowdfunding(comerId, crowdfundingId, re.TxHash)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func RemoveCrowdfunding(ctx *router.Context) {
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var re crowdfundingModel.TransactionHashRequest
	if err = ctx.ShouldBindJSON(&re); err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := comerId(ctx)
	err = crowdfunding.FinalizeCrowdFunding(comerId, crowdfundingId, re.TxHash)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func Invest(ctx *router.Context) {
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var re crowdfundingModel.InvestRequest
	if err = ctx.ShouldBindJSON(&re); err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := comerId(ctx)
	err = crowdfunding.Invest(comerId, crowdfundingId, re)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func ModifyCrowdfunding(ctx *router.Context) {
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var modifyRequest crowdfundingModel.ModifyRequest
	if err = ctx.ShouldBindJSON(&modifyRequest); err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := comerId(ctx)
	err = crowdfunding.ModifyCrowdfunding(comerId, crowdfundingId, modifyRequest)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func GetBuyPriceAndSwapModificationHistories(ctx *router.Context) {
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var pagination model.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	if pagination.Limit == 0 {
		pagination.Limit = 3
	}

	comerId := comerId(ctx)
	err = crowdfunding.GetBuyPriceAndSwapModificationHistories(comerId, crowdfundingId, &pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

func GetCrowdfundingSwapRecords(ctx *router.Context) {
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	var pagination model.Pagination
	if err := ctx.ShouldBindJSON(&pagination); err != nil {
		ctx.HandleError(err)
		return
	}
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	err = crowdfunding.GetCrowdfundingSwapRecords(crowdfundingId, &pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination)
}

// GetInvestProfile get investor by crowdfundingId and comerId
// not used now!
func GetInvestProfile(ctx *router.Context) {
	crowdfundingId, err := strconv.ParseUint(ctx.Param("crowdfundingId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	comerId := comerId(ctx)
	investor, err := crowdfunding.GetInvestorDetail(crowdfundingId, comerId)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(investor)
}

func GetCrowdfundingListOfStartup(ctx *router.Context) {
	startupId, err := strconv.ParseUint(ctx.Param("startupId"), 0, 64)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	x, err := crowdfunding.GetCrowdfundingListByStartup(startupId)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(x)
}

// GetComerPostedCrowdfundingList crowdfunding list posted by comer
func GetComerPostedCrowdfundingList(ctx *router.Context) {
	comerId, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		ctx.HandleError(errors.New("invalid comerId"))
	}
	pagination := model.Pagination{
		Limit: 100,
		Page:  1,
		Sort:  "created_at desc",
	}
	err = crowdfunding.GetPostedCrowdfundingListByComer(comerId, &pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination.Rows)
}

// GetComerParticipatedCrowdfundingList crowdfunding list participated by comer
func GetComerParticipatedCrowdfundingList(ctx *router.Context) {
	comerId, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		ctx.HandleError(errors.New("invalid comerId"))
	}
	pagination := model.Pagination{
		Limit: 100,
		Page:  1,
		Sort:  "created_at desc",
	}
	err = crowdfunding.GetParticipatedCrowdFundingListOfComer(comerId, &pagination)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(pagination.Rows)
}
