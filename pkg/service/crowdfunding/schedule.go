package crowdfunding

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/crowdfunding"
	"context"
	"github.com/gotomicro/ego/task/ecron"
	"github.com/qiniu/x/log"
)

func LiveCrowdfundingStatusSchedule() ecron.Ecron {
	job := func(ctx context.Context) error {
		list, err := crowdfunding.SelectToBeStartedCrowdfundingListWithin1Min(mysql.DB)
		if err != nil {
			log.Errorf("#### update crowdfunding live status: %v\n", err)
		}
		if len(list) > 0 {
			for _, c := range list {
				log.Infof("#### update crowdfunding live status: %d\n", c.ID)
				err = crowdfunding.UpdateCrowdfundingStatus(mysql.DB, c.ID, crowdfunding.Live)
				if err != nil {
					log.Infof("#### update crowdfunding live status: %v\n", err)
				}
			}
		}
		return err
	}
	return ecron.Load("ceres.status.cron").Build(ecron.WithJob(job))
}

func EndedCrowdfundingStatusSchedule() ecron.Ecron {
	job := func(ctx context.Context) error {
		list, err := crowdfunding.SelectToBeEndedCrowdfundingList(mysql.DB)
		if err != nil {
			log.Errorf("#### update crowdfunding ended status: %v\n", err)
		}
		if len(list) > 0 {
			for _, c := range list {
				err = crowdfunding.UpdateCrowdfundingStatus(mysql.DB, c.ID, crowdfunding.Ended)
				if err != nil {
					log.Infof("#### update crowdfunding ended status: %v\n", err)
				}
			}
		}
		return err
	}
	return ecron.Load("ceres.status.cron").Build(ecron.WithJob(job))
}
