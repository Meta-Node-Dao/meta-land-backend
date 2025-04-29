package governance

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/model/dict"
	"ceres/pkg/model/governance"
	"errors"
	"fmt"
	"github.com/qiniu/x/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func CreateProposal(comerId uint64, request *governance.CreateProposalRequest) error {
	if comerId != request.AuthorComerID {
		return errors.New("invalid comerId")
	}
	if err := request.Valid(); err != nil {
		return err
	}
	request.Status = int(governance.ProposalUpcoming)
	proposalModel := governance.GovernanceProposal{
		//GovernanceProposal: request.GovernanceProposal,
	}
	if len(request.Choices) == 0 {
		return errors.New("choices cannot be empty")
	}
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := governance.CreateProposal(mysql.DB, &proposalModel); err != nil {
			return err
		}
		var choices []*governance.GovernanceChoice
		for _, choice := range request.Choices {
			choices = append(choices, &governance.GovernanceChoice{
				ProposalChoice: governance.ProposalChoice{
					ProposalID: proposalModel.ID,
					ItemName:   choice.ItemName,
					SeqNum:     choice.SeqNum,
				},
			})
		}
		if err := governance.CreateProposalChoices(tx, choices); err != nil {
			return err
		}
		return nil
	})
}

func GetProposal(proposalId uint64) (detail governance.ProposalDetail, err error) {
	publicInfo, err := governance.GetProposalPublicInfo(mysql.DB, proposalId)
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
	choices, err := governance.GetProposalChoices(mysql.DB, proposalId)
	if err != nil {
		return
	}
	dictModel, err := dict.SelectByDictTypeAndLabel(mysql.DB, "voteSystem", publicInfo.VoteSystem)
	if err != nil {
		return
	}
	detail = governance.ProposalDetail{
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
	proposal, err := governance.GetProposalById(mysql.DB, proposalId)
	if err != nil {
		return err
	}
	if proposal.ID == 0 {
		return errors.New(fmt.Sprintf("proposal %d does not exist", proposalId))
	}
	admins, err := governance.GetGovernanceAdminsByStartupId(mysql.DB, proposal.StartupID)
	if err != nil {
		return err
	}
	var can bool
	if len(admins) == 0 {
		can = comerId == proposal.AuthorComerID
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
			if err := governance.DeleteProposal(mysql.DB, comerId, proposalId); err != nil {
				return err
			}
			if err := governance.DeleteProposalChoices(tx, proposalId); err != nil {
				return err
			}
			return nil
		})
	}
	return errors.New("cannot delete proposal")
}

func SelectProposalPublicList(request *governance.ProposalListRequest) error {
	proposals, err := governance.SelectProposalList(mysql.DB, request)
	request.Rows = packItem(proposals)
	return err
}

func GetStartupProposalList(startupId uint64, request *model.Pagination) error {
	proposals, err := governance.SelectProposalListByStartupId(mysql.DB, startupId, request)
	request.Rows = packItem(proposals)
	return err
}

func GetComerPostProposalList(comerId uint64, request *model.Pagination) error {
	proposals, err := governance.SelectProposalListByComerPosted(mysql.DB, comerId, request)
	request.Rows = packItem(proposals)
	return err
}

func GetComerParticipateProposalList(comerId uint64, request *model.Pagination) error {
	proposals, err := governance.SelectProposalListByComerParticipate(mysql.DB, comerId, request)
	request.Rows = packItem(proposals)
	return err
}

func packItem(proposals []governance.ProposalPublicInfo) []governance.ProposalItem {
	var list []governance.ProposalItem
	if len(proposals) > 0 {
		for _, proposal := range proposals {
			result, err := calculateProposalVoteResult(proposal.ProposalId)
			if err != nil {
				log.Warn(err)
				result = governance.ProposalVoteResult{}
			}
			list = append(list, governance.ProposalItem{
				ProposalPublicInfo: proposal,
				ProposalVoteResult: result,
			})
		}
	}
	return list
}

func calculateProposalVoteResult(proposalId uint64) (voteResult governance.ProposalVoteResult, err error) {
	result, err := GetProposalCrtVoteResult(proposalId)
	if err != nil {
		return
	}
	choices := result.ChoiceVoteInfos
	var (
		max *governance.ChoiceVoteInfo
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
	voteResult = governance.ProposalVoteResult{
		MaximumVotesChoice:   maxVotesChoice,
		MaximumVotesChoiceId: maxVotesChoiceId,
		Votes:                maxVotes,
		InvalidResult:        nil,
	}

	return
}
