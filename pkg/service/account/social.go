package account

import (
	"ceres/pkg/initialization/mysql"
	accountModel "ceres/pkg/model/account"
)

func UpdateSocial(comerId uint64, request accountModel.SocialModifyRequest) error {
	return accountModel.UpdateComerSocial(mysql.DB, comerId, request)
}

func RemoveSocial(comerId uint64, socialType accountModel.SocialType) error {
	return UpdateSocial(comerId, accountModel.SocialModifyRequest{
		SocialType: socialType,
		SocialLink: "",
	})
}
