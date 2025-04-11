package governance

import (
	"encoding/json"
	"github.com/qiniu/x/log"
	"testing"
)

func Test_GovernanceSettingDetailJson(t *testing.T) {
	indent, _ := json.MarshalIndent(
		GovernanceSettingDetail{
			Strategies: GovernanceStrategies{
				GovernanceStrategy{},
			},
			Admins: GovernanceAdmins{
				&GovernanceAdmin{},
			},
		}, "", "\t")
	// copy output and paste to import json for yapi
	log.Info(string(indent))
}
