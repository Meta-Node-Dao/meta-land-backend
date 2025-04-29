package governance

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model"
	"ceres/pkg/model/governance"
	"errors"
	"fmt"
	"strings"

	"github.com/qiniu/x/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func VoteProposal(comerId, proposalId uint64, request governance.VoteRequest) error {
	// todo check comer has rights to vote this proposal
	// todo check comer has voted the choiceId before, if has then overwrite that record
	if request.IpfsHash == "" || strings.TrimSpace(request.IpfsHash) == "" {
		return errors.New("invalid ipfsHash")
	}
	if request.ChoiceItemId == 0 {
		return errors.New("invalid choiceItemId")
	}
	if request.Votes == decimal.Zero || request.Votes.IsZero() || request.Votes.IsNegative() {
		return errors.New("illegal votes")
	}
	proposal, err := governance.GetProposalById(mysql.DB, proposalId)
	if err != nil {
		return err
	}
	if proposal.ID == 0 {
		return fmt.Errorf("proposal %d does not exist", proposalId)
	}
	if proposal.Status != int8(governance.ProposalActive) {
		return errors.New("invalid proposal status")
	}

	choice, err := governance.GetChoiceByProposalIdAndChoiceId(mysql.DB, proposalId, request.ChoiceItemId)
	if err != nil {
		return err
	}
	if choice.ID == 0 {
		return errors.New("invalid choice")
	}
	governanceVote := governance.GovernanceVote{
		VoteInfo: governance.VoteInfo{
			ProposalID:         proposalId,
			VoterComerID:       comerId,
			VoterWalletAddress: request.VoterWalletAddress,
			ChoiceItemID:       request.ChoiceItemId,
			ChoiceItemName:     request.ChoiceItemName,
			Votes:              request.Votes,
			IPFSHash:           request.IpfsHash,
		},
	}
	pVoteSystem := governance.VoteSystem(proposal.VoteSystem)

	if pVoteSystem == governance.VoteSystemSingleChoiceVoting || pVoteSystem == governance.VoteSystemBasicVoting {
		vote, err := governance.GetVoteRecordByProposalIdAndComerId(mysql.DB, proposalId, comerId)
		if err != nil {
			return err
		}
		if vote.ID == 0 {
			return governance.CreateProposalVote(mysql.DB, &governanceVote)
		}
		//if pVoteSystem == governance.VoteSystemBasicVoting {
		//	return errors.New("voteSystem is Basic Voting, you can't vote again")
		//}
		// update previous vote, does ipfs changed also?
		return mysql.DB.Transaction(func(tx *gorm.DB) error {
			if er := governance.DeleteVoteByProposalIdAndVoterComer(tx, proposalId, comerId); er != nil {
				return er
			}
			return governance.CreateProposalVote(tx, &governanceVote)
		})
	} else {
		log.Warnf("unsupported VoteSystem %v\n", pVoteSystem)
	}
	return nil
}

func GetProposalVoteRecords(proposalId uint64, pagination *model.Pagination) (err error) {
	_, err = governance.GetVoteRecordsByProposalId(mysql.DB, proposalId, pagination)
	return err
}

func GetProposalCrtVoteResult(proposalId uint64) (result governance.CurrentProposalVoteResult, err error) {
	pagination := model.Pagination{
		Limit: 999,
		Page:  1,
	}
	records, err := governance.GetVoteRecordsByProposalId(mysql.DB, proposalId, &pagination)
	if err != nil {
		return result, err
	}
	mp := map[uint64][]governance.VoteDetail{}
	totalVotes := decimal.Zero
	choices, err := governance.GetProposalChoices(mysql.DB, proposalId)
	if err != nil {
		return result, err
	}
	if len(choices) == 0 {
		return result, errors.New("invalid proposal without any choices")
	}
	if len(records) > 0 {
		for _, choice := range choices {
			for _, record := range records {
				if choice.ID == record.ChoiceItemID {
					totalVotes = totalVotes.Add(record.Votes)
					choiceItemId := record.ChoiceItemID
					if details, ok := mp[choiceItemId]; ok {
						details = append(details, record)
						mp[choiceItemId] = details
					} else {
						mp[choiceItemId] = []governance.VoteDetail{record}
					}
					continue
				} else {
					if _, ok := mp[choice.ID]; !ok {
						mp[choice.ID] = []governance.VoteDetail{
							{VoteInfo: governance.VoteInfo{
								ProposalID:     proposalId,
								ChoiceItemID:   choice.ID,
								ChoiceItemName: choice.ItemName,
								Votes:          decimal.Zero,
							}},
						}
					}
				}
			}
		}
	} else {
		for _, choice := range choices {
			mp[choice.ID] = []governance.VoteDetail{
				{VoteInfo: governance.VoteInfo{
					ProposalID:     proposalId,
					ChoiceItemID:   choice.ID,
					ChoiceItemName: choice.ItemName,
					Votes:          decimal.Zero,
				}},
			}
		}

	}
	var vmp []governance.ChoiceVoteInfo
	if len(mp) > 0 {
		for choiceId, details := range mp {
			cTotalVotes := decimal.Zero
			for _, detail := range details {
				cTotalVotes = cTotalVotes.Add(detail.Votes)
			}
			percent := decimal.Zero
			if totalVotes != decimal.Zero && !totalVotes.IsZero() {
				percent = cTotalVotes.Div(totalVotes)
			}
			vmp = append(vmp, governance.ChoiceVoteInfo{
				ChoiceId: choiceId,
				ItemName: details[0].ChoiceItemName,
				Votes:    &cTotalVotes,
				Percent:  &percent,
			})
		}
	}
	result = governance.CurrentProposalVoteResult{
		ChoiceVoteInfos: &vmp,
		TotalVotes:      &totalVotes,
	}
	return
}
