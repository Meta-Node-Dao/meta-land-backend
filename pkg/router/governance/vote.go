package governance

import (
	"ceres/pkg/model"
	"ceres/pkg/model/governance"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/governance"
)

func VoteProposal(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
		comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)

		var request governance.VoteRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.HandleError(err)
			return
		}
		err := service.VoteProposal(comerId, proposalId, request)
		if err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(nil)
	})
}

func ProposalVoteRecords(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
		var pagination model.Pagination
		if err := ctx.ShouldBindJSON(&pagination); err != nil {
			ctx.HandleError(err)
			return
		}
		err := service.GetProposalVoteRecords(proposalId, &pagination)
		if err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(pagination)
	})
}

// todo not in 羽雀、yapi
func ProposalVoteInfo(ctx *router.Context) {
	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
		result, err := service.GetProposalCrtVoteResult(proposalId)
		if err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(result)
	})
}
