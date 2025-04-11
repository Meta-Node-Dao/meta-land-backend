package governance

import (
	"ceres/pkg/model"
	governanceModel "ceres/pkg/model/governance"
	"ceres/pkg/service/dict"
	"ceres/pkg/service/governance"
	"ceres/pkg/service/startup"
	pkg "ceres/pkg/tests"
	"encoding/json"
	"github.com/shopspring/decimal"
	"testing"
)

func Test_GetComersFollowedThisStartup(t *testing.T) {
	pkg.DoTestWithMysql(func() {
		pg := model.Pagination{
			Limit: 99,
			Page:  1,
		}
		err := startup.GetComersFollowedThisStartup(125678234316800, 124637493276672, &pg)
		if err != nil {
			t.Fatal(err)
		}
		indent, _ := json.MarshalIndent(pg, "", "\t")
		t.Logf("%s\n", string(indent))

	})
}

func Test_JoinedAndFollowedStartups(t *testing.T) {
	pkg.DoTestWithMysql(func() {
		list, err := startup.GetComerJoinedOrFollowedStartups(129525702930432)
		if err != nil {
			t.Fatal(err)
		} else {
			indent, _ := json.MarshalIndent(list, "", "\t")
			t.Logf("%s", string(indent))
		}
	})
}

func Test_SelectDictDataByType(t *testing.T) {
	pkg.DoTestWithMysql(func() {
		// available types:1. voteSystem; 2. governanceStrategy
		var (
			typeVoteSymbol         = "voteSystem"
			typeGovernanceStrategy = "governanceStrategy"
		)
		voteSymbolDicts, err := dict.SelectDictDataByType(typeVoteSymbol)
		if err != nil {
			t.Fatal(err)
		}
		indent, _ := json.MarshalIndent(voteSymbolDicts, "", "\t")
		t.Logf("voteSymbol_dictDatas:\n%s\n", string(indent))
		gvSymbolDicts, err := dict.SelectDictDataByType(typeGovernanceStrategy)
		if err != nil {
			t.Fatal(err)
		}
		indent, _ = json.MarshalIndent(gvSymbolDicts, "", "\t")
		t.Logf("governanceStrategy_dictDatas:\n%s\n", string(indent))
	})
}

func Test_CreateStartupGovernanceSetting(t *testing.T) {
	pkg.DoTestWithMysql(func() {
		var (
			ci uint64 = 1
			si uint64 = 1
		)
		governance.CreateStartupGovernanceSetting(ci, si, governanceModel.CreateOrUpdateGovernanceSettingRequest{
			SettingRequest: governanceModel.SettingRequest{
				VoteSymbol:        "",
				AllowMember:       false,
				ProposalThreshold: decimal.NewFromFloat(0.8),
				ProposalValidity:  decimal.NewFromFloat(0.3),
			},
			Strategies: []governanceModel.StrategyRequest{
				{
					DictValue:            "",
					StrategyName:         "",
					ChainId:              0,
					TokenContractAddress: "",
					VoteDecimals:         0,
					TokenMinBalance:      decimal.Decimal{},
				},
			},
			Admins: nil,
		})
	})
}

func Test_GetStartupGovernanceSetting(t *testing.T) {

}
