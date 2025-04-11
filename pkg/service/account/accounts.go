package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
	"fmt"

	"github.com/qiniu/x/log"
)

// UserInfo  unlink the comer account
func UserInfo(comerID uint64, UserInfo *model.ComerLoginResponse) error {
	var comer model.Comer
	if err := model.GetComerByID(mysql.DB, comerID, &comer); err != nil {
		log.Warn(err)
		return err
	}
	if comer.Address != nil {
		UserInfo.Address = *comer.Address
	} else {
		UserInfo.Address = ""
	}

	var isProfiled bool
	var profile model.ComerProfile
	if err := model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		log.Warn(err)
		return err
	}

	if profile.ID != 0 {
		isProfiled = true
	}

	UserInfo.Nick = profile.Name
	UserInfo.Avatar = profile.Avatar
	UserInfo.IsProfiled = isProfiled
	UserInfo.ComerID = comer.ID

	var accounts []model.ComerAccount
	if err := model.GetComerAccountsByComerId(mysql.DB, comerID, &accounts); err != nil {
		log.Warn(err)
		return err
	}
	log.Infof("comer accounts for %d : %v \n", comerID, accounts)
	var accountBindingInfos = []*model.OauthAccountBindingInfo{
		{Linked: false, AccountType: 1},
		{Linked: false, AccountType: 2},
	}
	if len(accounts) > 0 {
		mp := make(map[model.ComerAccountType]uint64)
		for _, account := range accounts {
			mp[account.Type] = account.ID
		}
		for _, info := range accountBindingInfos {
			if v, ok := mp[info.AccountType]; ok {
				info.AccountId = v
				info.Linked = true
			}
		}
	}
	log.Infof("comer accounts bidingInfos for %d : %v \n", comerID, accountBindingInfos)
	UserInfo.ComerAccounts = accountBindingInfos
	return nil
}

// GetComerAccounts get current comer accounts
func GetComerAccounts(comerID uint64, response *model.ComerOuterAccountListResponse) (err error) {
	accountList := make([]model.ComerAccount, 0)
	if err = model.ListAccount(mysql.DB, comerID, &accountList); err != nil {
		log.Warn(err)
		return
	}
	*response = model.ComerOuterAccountListResponse{
		List:  accountList,
		Total: uint64(len(accountList)),
	}
	return
}

// UnlinkComerAccount  unlink the comer account
func UnlinkComerAccount(comerID, accountID uint64) error {
	var comer model.Comer
	if err := model.GetComerByID(mysql.DB, comerID, &comer); err != nil {
		return err
	}
	if comer.HasAddress() {
		return model.DeleteAccount(mysql.DB, comerID, accountID)
	}
	return fmt.Errorf("can not unlink comer account")

}

// GetComerInfo get comer info
func GetComerInfo(comerID uint64, response *model.GetComerInfoResponse) (err error) {
	if err = model.GetComerProfile(mysql.DB, comerID, &response.ComerProfile); err != nil {
		log.Warn(err)
		return err
	}
	if err = model.GetComerByID(mysql.DB, comerID, &response.Comer); err != nil {
		log.Warn(err)
		return err
	}
	if response.FollowsCount, err = model.ListFollowComer(mysql.DB, comerID, &response.Follows); err != nil {
		log.Warn(err)
		return err
	}
	if response.FollowedCount, err = model.ListFollowedComer(mysql.DB, comerID, &response.Followed); err != nil {
		log.Warn(err)
		return err
	}
	return
}

// GetComerInfoByAddress get comer info by address
func GetComerInfoByAddress(address string, response *model.GetComerInfoResponse) (err error) {
	if err = model.GetComerByAddress(mysql.DB, address, &response.Comer); err != nil {
		log.Warn(err)
		return err
	}
	if response.Comer.ID != 0 {
		if err = model.GetComerProfile(mysql.DB, response.Comer.ID, &response.ComerProfile); err != nil {
			log.Warn(err)
			return err
		}
	}
	return
}
