package governance

import (
	"ceres/pkg/model"
	"ceres/pkg/model/governance"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/governance"
	"errors"
)

func CreateProposal(ctx *router.Context) {
	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if comerId == 0 {
		ctx.HandleError(errors.New("invalid comerId"))
		return
	}
	var request governance.CreateProposalRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := service.CreateProposal(comerId, &request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(nil)
}

func GetProposal(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
		info, err := service.GetProposal(proposalId)
		if err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(info)
	})
}

func DeleteProposal(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
		comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
		if comerId == 0 {
			ctx.HandleError(errors.New("invalid comer id"))
			return
		}
		if err := service.DeleteProposal(comerId, proposalId); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(nil)
	})
}

func PublicList(ctx *router.Context) {
	var request governance.ProposalListRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	if err := service.SelectProposalPublicList(&request); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(request.Pagination)
}

func StartupProposalList(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "startupID", func(startupId uint64) {
		var pagination model.Pagination
		if err := ctx.ShouldBindJSON(&pagination); err != nil {
			ctx.HandleError(err)
			return
		}
		if err := service.GetStartupProposalList(startupId, &pagination); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(pagination)
	})
}

func ComerPostProposalList(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "comerID", func(comerId uint64) {
		var pagination model.Pagination
		if err := ctx.ShouldBindJSON(&pagination); err != nil {
			ctx.HandleError(err)
			return
		}
		if err := service.GetComerPostProposalList(comerId, &pagination); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(pagination)
	})
}

func ComerParticipateProposalList(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "comerID", func(comerId uint64) {
		var pagination model.Pagination
		if err := ctx.ShouldBindJSON(&pagination); err != nil {
			ctx.HandleError(err)
			return
		}
		if err := service.GetComerParticipateProposalList(comerId, &pagination); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(pagination)
	})
}
