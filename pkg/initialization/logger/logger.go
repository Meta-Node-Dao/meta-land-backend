package logger

import (
	"github.com/gotomicro/ego/core/elog"
	"github.com/qiniu/x/log"

	C "ceres/pkg/config"
)

// Logger elog.Logger instancd
var Logger *elog.Component

// Init init the logger
func Init() error {
	Logger = elog.Load("ceres.logger").Build()
	elog.DefaultLogger = Logger

	// init qiniu debug log
	switch C.Log.Level {
	case "debug":
		log.SetOutputLevel(log.Ldebug)
	}
	return nil
}
