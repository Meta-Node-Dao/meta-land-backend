package governance

import (
	"ceres/pkg/model"
	"ceres/pkg/router"
)

func VoteProposal(c *router.Context) {
	c.OK("vote proposal successfully!")
}

func GetProposalInvestRecords(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

//func VoteProposal(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
//		comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//
//		var request governance.VoteRequest
//		if err := ctx.ShouldBindJSON(&request); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		err := service.VoteProposal(comerId, proposalId, request)
//		if err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(nil)
//	})
//}
//
//func ProposalVoteRecords(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
//		var pagination model.Pagination
//		if err := ctx.ShouldBindJSON(&pagination); err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		err := service.GetProposalVoteRecords(proposalId, &pagination)
//		if err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(pagination)
//	})
//}
//
//
//func ProposalVoteInfo(ctx *router.Context) {
//	uintPathVariableWithCb(ctx, "proposalID", func(proposalId uint64) {
//		result, err := service.GetProposalCrtVoteResult(proposalId)
//		if err != nil {
//			ctx.HandleError(err)
//			return
//		}
//		ctx.OK(result)
//	})
//}
