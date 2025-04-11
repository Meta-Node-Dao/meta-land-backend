package governance

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/model/dict"
	governanceModel "ceres/pkg/model/governance"
	"errors"
	"fmt"
	"github.com/qiniu/x/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func CreateProposal(comerId uint64, request *governanceModel.CreateProposalRequest) error {
	if comerId != request.AuthorComerId {
		return errors.New("invalid comerId")
	}
	if err := request.Valid(); err != nil {
		return err
	}
	request.Status = governanceModel.ProposalUpcoming
	proposalModel := governanceModel.GovernanceProposalModel{
		GovernanceProposalInfo: request.GovernanceProposalInfo,
	}
	if len(request.Choices) == 0 {
		return errors.New("choices cannot be empty")
	}
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := governanceModel.CreateProposal(mysql.DB, &proposalModel); err != nil {
			return err
		}
		var choices []*governanceModel.GovernanceChoice
		for _, choice := range request.Choices {
			choices = append(choices, &governanceModel.GovernanceChoice{
				ProposalChoice: governanceModel.ProposalChoice{
					ProposalId: proposalModel.ID,
					ItemName:   choice.ItemName,
					SeqNum:     choice.SeqNum,
				},
			})
		}
		if err := governanceModel.CreateProposalChoices(tx, choices); err != nil {
			return err
		}
		return nil
	})
}

func GetProposal(proposalId uint64) (detail governanceModel.ProposalDetail, err error) {
	publicInfo, err := governanceModel.GetProposalPublicInfo(mysql.DB, proposalId)
	if err != nil {
		return
	}
	setting, err := GetStartupGovernanceSetting(publicInfo.StartupId)
	if err != nil {
		return detail, err
	}

	strategies := setting.Strategies
	admins := setting.Admins
	voteResult, err := GetProposalCrtVoteResult(proposalId)
	if err != nil {
		return
	}
	choices, err := governanceModel.GetProposalChoices(mysql.DB, proposalId)
	if err != nil {
		return
	}
	dictModel, err := dict.SelectByDictTypeAndLabel(mysql.DB, "voteSystem", publicInfo.VoteSystem)
	if err != nil {
		return
	}
	detail = governanceModel.ProposalDetail{
		ProposalPublicInfo:        publicInfo,
		VoteSystemId:              dictModel.ID,
		Strategies:                strategies,
		Admins:                    admins,
		CurrentProposalVoteResult: voteResult,
		Choices:                   choices,
	}
	return
}

func DeleteProposal(comerId, proposalId uint64) error {
	// todo - who can delete a proposal? the founder?
	var cm account.Comer
	if err := account.GetComerByID(mysql.DB, comerId, &cm); err != nil {
		return err
	}
	if cm.ID == 0 {
		return errors.New(fmt.Sprintf("invalid comer %d", comerId))
	}
	if cm.Address == nil {
		return errors.New(fmt.Sprintf("invalid comer %d without walletAddress", comerId))
	}
	proposal, err := governanceModel.GetProposalById(mysql.DB, proposalId)
	if err != nil {
		return err
	}
	if proposal.ID == 0 {
		return errors.New(fmt.Sprintf("proposal %d does not exist", proposalId))
	}
	admins, err := governanceModel.GetGovernanceAdminsByStartupId(mysql.DB, proposal.StartupId)
	if err != nil {
		return err
	}
	var can bool
	if len(admins) == 0 {
		can = comerId == proposal.AuthorComerId
	} else {
		for _, admin := range admins {
			can = admin.WalletAddress == *cm.Address
			if can {
				break
			}
		}
	}
	if can {
		return mysql.DB.Transaction(func(tx *gorm.DB) error {
			if err := governanceModel.DeleteProposal(mysql.DB, comerId, proposalId); err != nil {
				return err
			}
			if err := governanceModel.DeleteProposalChoices(tx, proposalId); err != nil {
				return err
			}
			return nil
		})
	}
	return errors.New("cannot delete proposal")
}

func SelectProposalPublicList(request *governanceModel.ProposalListRequest) error {
	proposals, err := governanceModel.SelectProposalList(mysql.DB, request)
	request.Rows = packItem(proposals)
	return err
}

func GetStartupProposalList(startupId uint64, request *model.Pagination) error {
	proposals, err := governanceModel.SelectProposalListByStartupId(mysql.DB, startupId, request)
	request.Rows = packItem(proposals)
	return err
}

func GetComerPostProposalList(comerId uint64, request *model.Pagination) error {
	proposals, err := governanceModel.SelectProposalListByComerPosted(mysql.DB, comerId, request)
	request.Rows = packItem(proposals)
	return err
}

func GetComerParticipateProposalList(comerId uint64, request *model.Pagination) error {
	proposals, err := governanceModel.SelectProposalListByComerParticipate(mysql.DB, comerId, request)
	request.Rows = packItem(proposals)
	return err
}

func packItem(proposals []governanceModel.ProposalPublicInfo) []governanceModel.ProposalItem {
	var list []governanceModel.ProposalItem
	if len(proposals) > 0 {
		for _, proposal := range proposals {
			result, err := calculateProposalVoteResult(proposal.ProposalId)
			if err != nil {
				log.Warn(err)
				result = governanceModel.ProposalVoteResult{}
			}
			list = append(list, governanceModel.ProposalItem{
				ProposalPublicInfo: proposal,
				ProposalVoteResult: result,
			})
		}
	}
	return list
}

func calculateProposalVoteResult(proposalId uint64) (voteResult governanceModel.ProposalVoteResult, err error) {
	result, err := GetProposalCrtVoteResult(proposalId)
	if err != nil {
		return
	}
	choices := result.ChoiceVoteInfos
	var (
		max *governanceModel.ChoiceVoteInfo
		//maxChoiceVotes = decimal.Zero
	)
	if choices != nil && len(*choices) > 0 {
		for _, info := range *choices {
			tmp := info
			if max == nil {
				max = &tmp
				continue
			}
			if (*max.Votes).LessThanOrEqual(*tmp.Votes) {
				max = &tmp
			}
		}
	}
	var (
		maxVotesChoice   *string
		maxVotesChoiceId *uint64
		maxVotes         = &(decimal.Zero)
	)
	if max != nil {
		maxVotesChoice = &max.ItemName
		maxVotesChoiceId = &max.ChoiceId
		maxVotes = max.Votes
	}
	voteResult = governanceModel.ProposalVoteResult{
		MaximumVotesChoice:   maxVotesChoice,
		MaximumVotesChoiceId: maxVotesChoiceId,
		Votes:                maxVotes,
		InvalidResult:        nil,
	}

	return
}
