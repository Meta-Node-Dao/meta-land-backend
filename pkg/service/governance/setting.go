package governance

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	"ceres/pkg/model/governance"
	startupModel "ceres/pkg/model/startup"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strings"
)

func CreateStartupGovernanceSetting(comerId, startupId uint64, request governance.CreateOrUpdateGovernanceSettingRequest) (setting governance.GovernanceSetting, err error) {
	startup, err := startupModel.GetStartupById(mysql.DB, startupId)
	if err != nil {
		return setting, err
	}
	if startup.ComerID != comerId {
		return setting, errors.New("comer is not founder of startup")
	}
	if request.VoteSymbol == "" || strings.TrimSpace(request.VoteSymbol) == "" {
		return setting, errors.New("vote symbol can not be empty")
	}
	if len(request.Strategies) == 0 {
		return setting, errors.New("governance strategies can not be empty")
	}
	mayBeExistedSetting, err := governance.GetGovernanceSetting(mysql.DB, startupId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return setting, err
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		if mayBeExistedSetting.ID == 0 {
			mayBeExistedSetting = governance.GovernanceSetting{
				StartupId:         startupId,
				ComerId:           comerId,
				VoteSymbol:        request.VoteSymbol,
				AllowMember:       request.AllowMember,
				ProposalThreshold: request.ProposalThreshold,
				ProposalValidity:  request.ProposalValidity,
			}
			er = governance.CreateGovernanceSetting(tx, &mayBeExistedSetting)
			if er != nil {
				return
			}
			return createStrategiesOrAdmins(tx, request, mayBeExistedSetting)
		}
		// updating
		er = governance.UpdateGovernanceSetting(tx, mayBeExistedSetting.ID, &governance.GovernanceSetting{
			VoteSymbol:        request.VoteSymbol,
			AllowMember:       request.AllowMember,
			ProposalThreshold: request.ProposalThreshold,
			ProposalValidity:  request.ProposalValidity,
		})
		if er != nil {
			return
		}
		// select again
		mayBeExistedSetting, _ = governance.GetGovernanceSetting(tx, startupId)
		er = governance.DeleteAdminsBySettingId(tx, mayBeExistedSetting.ID)
		if er != nil {
			return
		}
		er = governance.DeleteStrategiesBySettingId(tx, mayBeExistedSetting.ID)
		if er != nil {
			return
		}
		if err := createStrategiesOrAdmins(tx, request, mayBeExistedSetting); err != nil {
			return err
		}
		return nil
	})
	return mayBeExistedSetting, err
}

func createStrategiesOrAdmins(tx *gorm.DB, request governance.CreateOrUpdateGovernanceSettingRequest, mayBeExistedSetting governance.GovernanceSetting) (er error) {
	var strategies []*governance.GovernanceStrategy
	if len(request.Strategies) > 0 {
		for _, strategy := range request.Strategies {
			strategies = append(strategies, &governance.GovernanceStrategy{
				SettingId:            mayBeExistedSetting.ID,
				DictValue:            strategy.DictValue,
				StrategyName:         strategy.StrategyName,
				ChainId:              strategy.ChainId,
				TokenContractAddress: strategy.TokenContractAddress,
				VoteSymbol:           strategy.VoteSymbol,
				VoteDecimals:         strategy.VoteDecimals,
				TokenMinBalance:      strategy.TokenMinBalance,
			})
		}
	}
	er = governance.CreateGovernanceStrategies(tx, strategies)
	if er != nil {
		return
	}
	var admins []*governance.GovernanceAdmin
	if len(request.Admins) > 0 {
		for _, admin := range request.Admins {
			admins = append(admins, &governance.GovernanceAdmin{
				SettingId:     mayBeExistedSetting.ID,
				WalletAddress: admin.WalletAddress,
			})
		}
		return governance.CreateGovernanceAdmins(tx, admins)
	}
	return
}

func GetStartupGovernanceSetting(startupId uint64) (detail governance.GovernanceSettingDetail, err error) {
	setting, err := governance.GetGovernanceSetting(mysql.DB, startupId)
	if err != nil {
		return detail, err
	}
	st, err := startupModel.GetStartupById(mysql.DB, startupId)
	if err != nil {
		return detail, err
	}
	var startupFounder account.Comer
	if err := account.GetComerByID(mysql.DB, st.ComerID, &startupFounder); err != nil {
		return detail, err
	}
	// if setting.ID == 0 , create default governance-setting for startup
	if setting.ID == 0 {
		setting, err = CreateStartupGovernanceSetting(st.ComerID, startupId, governance.CreateOrUpdateGovernanceSettingRequest{
			SettingRequest: governance.SettingRequest{
				VoteSymbol:        "vote",
				AllowMember:       true,
				ProposalThreshold: decimal.Zero,
				ProposalValidity:  decimal.Zero,
			},
			Strategies: []governance.StrategyRequest{
				// ChainId of ethereum
				{DictValue: "ticket", StrategyName: "ticket", VoteSymbol: "", ChainId: config.Eth.EthereumChainID},
			},
			Admins: []governance.AdminRequest{
				// default admin is the funder of startup
				{WalletAddress: *startupFounder.Address},
			},
		})
		if err != nil {
			return governance.GovernanceSettingDetail{}, err
		}
	}
	strategies, err := governance.GetGovernanceStrategies(mysql.DB, setting.ID)
	if err != nil {
		return detail, err
	}
	admins, err := governance.GetGovernanceAdmins(mysql.DB, setting.ID)
	if err != nil {
		return detail, err
	}
	if len(admins) == 0 {
		governanceAdmins := []*governance.GovernanceAdmin{{SettingId: setting.ID, WalletAddress: *startupFounder.Address}}
		if err := governance.CreateGovernanceAdmins(mysql.DB, governanceAdmins); err != nil {
			return detail, err
		}
		admins = append(admins, governanceAdmins...)
	}

	detail = governance.GovernanceSettingDetail{
		GovernanceSetting: setting,
		Strategies:        strategies,
		Admins:            admins,
	}

	return detail, nil
}
