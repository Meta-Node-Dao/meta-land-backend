package governance

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/governance"
	"context"
	"github.com/gotomicro/ego/task/ecron"
	"github.com/qiniu/x/log"
)

func ActiveProposalStatusSchedule() ecron.Ecron {
	job := func(ctx context.Context) error {
		list, err := governance.SelectToBeStartedProposalListWithin1Min(mysql.DB)
		if err != nil {
			log.Errorf("#### update proposal active status: %v\n", err)
		}
		if len(list) > 0 {
			for _, c := range list {
				log.Infof("#### update proposal active status: %d\n", c.ID)
				err = governance.UpdateProposalStatus(mysql.DB, c.ID, governance.ProposalActive)
				if err != nil {
					log.Infof("#### update proposal active status: %v\n", err)
				}
			}
		}
		return err
	}
	return ecron.Load("ceres.crowdfunding.cron").Build(ecron.WithJob(job))
}

func EndProposalStatusSchedule() ecron.Ecron {
	job := func(ctx context.Context) error {
		list, err := governance.SelectToEndedProposalListWithin1Min(mysql.DB)
		if err != nil {
			log.Errorf("#### update proposal ended status: %v\n", err)
		}

		if len(list) > 0 {
			for _, c := range list {
				setting, err := GetStartupGovernanceSetting(c.StartupId)
				if err != nil {
					return err
				}
				detail, err := GetProposal(c.ID)
				if err != nil {
					return err
				}
				status := governance.ProposalEnded
				if detail.TotalVotes.LessThan(setting.ProposalValidity) {
					status = governance.ProposalInvalid
				}
				err = governance.UpdateProposalStatus(mysql.DB, c.ID, status)
				if err != nil {
					log.Infof("#### update proposal %v  status: %v\n", status, err)
				}
			}
		}
		return err
	}
	return ecron.Load("ceres.crowdfunding.cron").Build(ecron.WithJob(job))
}
