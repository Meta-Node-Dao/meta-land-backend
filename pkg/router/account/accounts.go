package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"ceres/pkg/service/startup"
	"ceres/pkg/utility/jwt"
	"errors"
	"github.com/qiniu/x/log"
	"strconv"
)

// UserInfo list all accounts of the Comer
func UserInfo(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response account.ComerLoginResponse
	if err := service.UserInfo(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// ListAccounts list all accounts of the Comer
func ListAccounts(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response model.ComerOuterAccountListResponse
	if err := service.GetComerAccounts(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// UnlinkAccount unlink accounts for the Comer
func UnlinkAccount(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	accountID, err := strconv.ParseUint(ctx.Param("accountID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid account ID")
		ctx.HandleError(err)
		return
	}
	err = service.UnlinkComerAccount(comerID, accountID)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// LinkWithWallet link current account with wallet
func LinkWithWallet(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var ethLoginRequest model.EthLoginRequest
	if err := ctx.BindJSON(&ethLoginRequest); err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid data format")
		ctx.HandleError(err)
		return
	}

	err, finalComerId := service.LinkEthAccountToComer(comerID, ethLoginRequest.Address, ethLoginRequest.Signature)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	var (
		profile model.ComerProfile
		res     model.LinkWalletResponse
	)
	if err := account.GetComerProfile(mysql.DB, finalComerId, &profile); err != nil {
		ctx.HandleError(err)
		return
	}
	token := jwt.Sign(finalComerId)
	log.Infof("regenerate token %s after wallet link ......incoming comerId: %d, finalComerId: %d\n", token, comerID, finalComerId)
	if profile.ID != 0 {
		res = model.LinkWalletResponse{IsProfiled: true, Token: token}
	} else {
		res = model.LinkWalletResponse{IsProfiled: false, Token: token}
	}

	ctx.OK(res)
}

// GetComerInfo get comer
func GetComerInfo(ctx *router.Context) {
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var response model.GetComerInfoResponse
	if err := service.GetComerInfo(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// GetComerInfoByAddress get comer by address
func GetComerInfoByAddress(ctx *router.Context) {
	address := ctx.Param("address")
	if address == "" {
		err := router.ErrBadRequest.WithMsg("Comer's address required")
		ctx.HandleError(err)
		return
	}
	var response model.GetComerInfoResponse
	if err := service.GetComerInfoByAddress(address, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	if response.Comer.ID == 0 {
		ctx.OK(nil)
	} else {
		ctx.OK(response)
	}
}

// FollowComer follow Comer
func FollowComer(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	if comerID == targetComerID {
		err = router.ErrBadRequest.WithMsg("Can not follow myself")
		ctx.HandleError(err)
		return
	}
	if err = service.FollowComer(comerID, targetComerID); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// UnfollowComer unfollow Comer
func UnfollowComer(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	if err = service.UnfollowComer(comerID, targetComerID); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// ComerFollowedByMe get comer is followed by me
func ComerFollowedByMe(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	targetComerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	isFollowed, err := service.FollowedByComer(comerID, targetComerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(model.IsFollowedResponse{IsFollowed: isFollowed})
}

func JoinedAndFollowedStartups(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if comerID == 0 {
		ctx.HandleError(errors.New("invalid comer"))
		return
	}
	list, err := startup.GetComerJoinedOrFollowedStartups(comerID)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(list)
}
