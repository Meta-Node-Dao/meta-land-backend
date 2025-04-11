package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/utility/auth"
	"ceres/pkg/utility/jwt"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

// LoginWithOauth common oauth login logic in comunion
func LoginWithOauth(client auth.OauthClient, oauthType account.ComerAccountType, response *account.ComerLoginResponse) (err error) {
	oauth, err := client.GetUserProfile()
	if err != nil {
		log.Warn(err)
		return router.ErrBadRequest.WithMsg("Login authorization failed")
	}

	// try to find account
	var comerAccount account.ComerAccount
	if err = account.GetComerAccount(mysql.DB, oauthType, oauth.GetUserID(), &comerAccount); err != nil {
		log.Warn(err)
		return
	}

	//set default profile status
	var isProfiled bool
	var profile account.ComerProfile
	var comer account.Comer

	if comerAccount.ID == 0 {
		comer = account.Comer{}
		err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
			if er = account.CreateComer(tx, &comer); er != nil {
				log.Warn(er)
				return er
			}
			comerAccount = account.ComerAccount{
				ComerID:   comer.ID,
				OIN:       oauth.GetUserID(),
				IsPrimary: true,
				Nick:      oauth.GetUserNick(),
				Avatar:    oauth.GetUserAvatar(),
				Type:      oauthType,
				IsLinked:  true,
			}
			if er = account.CreateAccount(tx, &comerAccount); er != nil {
				log.Warn(er)
				return er
			}
			// also create profile ????
			return
		})
		if err != nil {
			return
		}
	} else {
		if err = account.GetComerByID(mysql.DB, comerAccount.ComerID, &comer); err != nil {
			log.Warn(err)
			return err
		}
		if err = account.GetComerProfile(mysql.DB, comerAccount.ComerID, &profile); err != nil {
			log.Warn(err)
			return err
		}
		if profile.ID != 0 {
			isProfiled = true
		}
	}

	// sign with jwt using the comer UIN
	token := jwt.Sign(comerAccount.ComerID)

	address := ""
	if comer.Address != nil {
		address = *comer.Address
	}

	*response = account.ComerLoginResponse{
		IsProfiled: isProfiled,
		Avatar:     comerAccount.Avatar,
		Nick:       comerAccount.Nick,
		Address:    address,
		Token:      token,
		ComerID:    comerAccount.ComerID,
	}
	return nil
}

// LinkOauthAccountToComer link a new Oauth account to the current comer
func LinkOauthAccountToComer(comerID uint64, client auth.OauthClient, oauthType account.ComerAccountType, response *account.ComerLoginResponse) (err error) {
	oauth, err := client.GetUserProfile()
	if err != nil {
		log.Warn(err)
		return
	}
	log.Debugf("LINK_OAUTH_ACCOUNT_TO_COMER--> comerId: %d\n", comerID)
	// try to find account
	var comerAccount account.ComerAccount
	if err = account.GetComerAccount(mysql.DB, oauthType, oauth.GetUserID(), &comerAccount); err != nil {
		log.Warn(err)
		return err
	}
	if comerAccount.ID == 0 {
		comerAccount = account.ComerAccount{
			ComerID:   comerID,
			OIN:       oauth.GetUserID(),
			IsPrimary: true,
			Nick:      oauth.GetUserNick(),
			Avatar:    oauth.GetUserAvatar(),
			Type:      oauthType,
			IsLinked:  true,
		}
		if err = account.CreateAccount(mysql.DB, &comerAccount); err != nil {
			return err
		}
	} else if comerAccount.ComerID != comerID {
		log.Debugf("Account %s(%s) has bind to another comer: %d", oauth.GetUserID(), oauth.GetUserNick(), comerAccount.ComerID)
		err = router.ErrBadRequest.WithMsg("Account has bind to another comerId")
		return err
	}

	var (
		comerProfile account.ComerProfile
		isProfiled   bool
	)
	if err := account.GetComerProfile(mysql.DB, comerID, &comerProfile); err != nil {
		isProfiled = true
	} else {
		isProfiled = false
	}

	// sign with jwt using the comer UIN
	token := jwt.Sign(comerAccount.ComerID)

	address := ""

	*response = account.ComerLoginResponse{
		IsProfiled: isProfiled,
		Avatar:     comerAccount.Avatar,
		Nick:       comerAccount.Nick,
		Address:    address,
		Token:      token,
		ComerID:    comerID,
	}
	return
}

func LinkOauthToComer(oauthAccountId, comerId uint64) error {
	return account.BindComerAccountToComerId(mysql.DB, oauthAccountId, comerId)
}
