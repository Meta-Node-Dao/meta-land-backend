package startup

import (
	"ceres/pkg/event"
	"context"

	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/core/etrace"
	"github.com/gotomicro/ego/task/ecron"
)

func UpdateStartupOnChain() ecron.Ecron {
	job := func(ctx context.Context) error {
		elog.Info("#### UpdateStartupOnChain\n", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		return event.HandleAllClientStartup()
	}
	return ecron.Load("ceres.startup.cron").Build(ecron.WithJob(job))
}
