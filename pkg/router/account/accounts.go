package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"ceres/pkg/service/startup"
	"ceres/pkg/utility/jwt"
	"encoding/json"
	"errors"
	"github.com/qiniu/x/log"
	"strconv"
)

func UserInfo(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response account.ComerLoginResponse
	if err := service.UserInfo(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

func ListAccounts(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response account.ComerOuterAccountListResponse
	if err := service.GetComerAccounts(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

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

func LinkWithWallet(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var ethLoginRequest account.EthLoginRequest
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
		profile account.ComerProfile
		res     account.LinkWalletResponse
	)
	if err := account.GetComerProfile(mysql.DB, finalComerId, &profile); err != nil {
		ctx.HandleError(err)
		return
	}
	token := jwt.Sign(finalComerId)
	log.Infof("regenerate token %s after wallet link ......incoming comerId: %d, finalComerId: %d\n", token, comerID, finalComerId)
	if profile.ID != 0 {
		res = account.LinkWalletResponse{IsProfiled: true, Token: token}
	} else {
		res = account.LinkWalletResponse{IsProfiled: false, Token: token}
	}

	ctx.OK(res)
}

func GetComerInfo(ctx *router.Context) {
	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
		ctx.HandleError(err)
		return
	}
	var response account.GetComerInfoResponse
	if err := service.GetComerInfo(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

func GetComerInfoByAddress(ctx *router.Context) {
	address := ctx.Param("address")
	if address == "" {
		err := router.ErrBadRequest.WithMsg("Comer's address required")
		ctx.HandleError(err)
		return
	}
	var response account.GetComerInfoResponse
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

	ctx.OK(account.IsFollowedResponse{IsFollowed: isFollowed})
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

var comerResString = "{\n  \"activation\": true,\n  \"address\": \"0x71C7656EC7ab88b098defB751B7401B5f6d8976F\",\n  \"avatar\": \"https://example.com/avatars/user123.jpg\",\n  \"banner\": \"https://example.com/banners/user123-banner.jpg\",\n  \"custom_domain\": \"user123.myplatform.com\",\n  \"id\": 42,\n  \"invitation_code\": \"INVITE-7XK9P2\",\n  \"is_connected\": true,\n  \"is_seted\": false,\n  \"location\": \"San Francisco, CA\",\n  \"name\": \"John Doe\",\n  \"time_zone\": \"America/Los_Angeles\"\n}"
var comerInfoDetailString = "{\n  \"id\": 494269249085440,\n  \"address\": \"0xD1b53493D2612168B3a80740698d6e9e6196FA4B\",\n  \"name\": \"WEconomy_277a15711ede\",\n  \"avatar\": \"https://comunion-avatars.s3.ap-northeast-1.amazonaws.com/users/female3.svg\",\n  \"banner\": \"\",\n  \"location\": \"\",\n  \"time_zone\": \"\",\n  \"invitation_code\": \"0xe74d978f\",\n  \"custom_domain\": \"494269249085440\",\n  \"activation\": true,\n  \"is_connected\": false,\n  \"info\": {\n    \"id\": 495703373897728,\n    \"comer_id\": 494269249085440,\n    \"bio\": \"WEB xiaobai !!!\"\n  },\n  \"educations\": null,\n  \"accounts\": null,\n  \"languages\": null,\n  \"socials\": null,\n  \"skills\": null,\n  \"connected_total\": {\n    \"connect_startup_total\": 1,\n    \"connect_comer_total\": 1,\n    \"be_connect_comer_total\": 0\n  }\n}"

func GetComer(ctx *router.Context) {
	var res account.ComerResponse
	err := json.Unmarshal([]byte(comerResString), &res)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(res)
}

func UpdateComerInfo(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "update comer info successful!"
	ctx.OK(message)
}

func UnlinkOauthByComerAccountId(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "unlink oauth by comer account id successful!"
	ctx.OK(message)
}

func UpdateComerInfoBio(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "update comer info bio successful!"
	ctx.OK(message)
}

func GetComerInfoDetail(ctx *router.Context) {
	var comerInfoDetail account.ComerInfoDetailResponse
	err := json.Unmarshal([]byte(comerInfoDetailString), &comerInfoDetail)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(comerInfoDetail)
}

func BindComerEducations(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "bind comer educations successful!"
	ctx.OK(message)
}

func UpdateComerEducation(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "update comer education successful!"
	ctx.OK(message)
}

func UnbindComerEducations(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "unbind comer education successful!"
	ctx.OK(message)
}

func GetComerInvitationCount(ctx *router.Context) {
	var res account.ComerInvitationCountResponse
	res.ActivatedTotal = 100
	res.InactiveTotal = 200
	ctx.OK(res)
}

func GetComerInvitationRecords(ctx *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	ctx.OK(res)
}

func BindComerLanguages(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "bind comer languages successful!"
	ctx.OK(message)
}

func UpdateComerLanguages(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "update comer languages successful!"
	ctx.OK(message)
}

func UnbindComerLanguages(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "unbind comer languages successful!"
	ctx.OK(message)
}

func GetComerJoinedAndFollowedStartups(ctx *router.Context) {
	var res account.StartupListResponse
	res.Total = 0
	ctx.OK(res)
}

func BindComerSkills(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "bind comer skills successful!"
	ctx.OK(message)
}

func BindComerSocials(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "bind comer socials successful!"
	ctx.OK(message)
}

func UpdateComerSocials(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "update comer socials successful!"
	ctx.OK(message)
}

func UnbindComerSocials(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "unbind comer socials successful!"
	ctx.OK(message)
}

func GetSocials(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func SetShare(c *router.Context) {
	var res account.ShareSetResponse
	res.ShareCode = "123456"
	c.OK(res)
}

func GetSharePageHtml(c *router.Context) {

	c.HTML(200, "share.html", nil)
}

func GetLanguages(c *router.Context) {
	var res account.LanguageResponse
	res.Code = "en"
	res.Id = 1
	res.Name = "English"
	c.OK(res)
}

func GetComerByAddress(c *router.Context) {
	var res account.ComerBasicResponse
	c.OK(res)
}

func SetUserCustomDomain(c *router.Context) {
	var message model.MessageResponse
	message.Message = "set user custom domain successful!"
	c.OK(message)
}

func GetUserCustomDomainExistence(c *router.Context) {
	var res model.IsExistResponse
	res.IsExist = false
	c.OK(res)
}

func GetUserCustomDomain(c *router.Context) {
	var res account.ComerInfoDetailResponse
	c.OK(res)
}

func VerifyComerAddProfile(c *router.Context) {
	var res account.ThirdPartyVerifyResponse
	c.OK(res)
}

func GetComerByComerId(c *router.Context) {
	var res account.ComerInfoDetailResponse
	c.OK(res)
}

func GetComerBeConnectComersByComerId(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func ConnectComer(c *router.Context) {
	var message model.MessageResponse
	message.Message = "connect comer successful!"
	c.OK(message)
}

func UnconnectComer(c *router.Context) {
	var message model.MessageResponse
	message.Message = "unconnect comer successful!"
	c.OK(message)
}

func GetComerConnectComersByComerId(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func GetStartupConnectByComerId(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func ConnectedComer(c *router.Context) {
	var res account.IsConnectedResponse
	res.IsConnected = true
	c.OK(res)
}

func GetComerInfoDetailByComerId(c *router.Context) {
	var res account.ComerInfoDetailResponse
	c.OK(res)
}

func GetComerParticipatedCountByComerId(c *router.Context) {
	var res account.ProjectCountResponse
	var resString = "{\n    \"bounty_count\": 12,\n    \"crowdfunding_count\": 7,\n    \"governance_count\": 5,\n    \"other_dapp_count\": 3,\n    \"sale_launchpad_count\": 3,\n    \"startup_count\": 9\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}

func GetComerPostedCountByComerId(c *router.Context) {
	var res account.ProjectCountResponse
	var resString = "{\n    \"bounty_count\": 12,\n    \"crowdfunding_count\": 7,\n    \"governance_count\": 5,\n    \"other_dapp_count\": 3,\n    \"sale_launchpad_count\": 3,\n    \"startup_count\": 9\n}"
	err := json.Unmarshal([]byte(resString), &res)
	if err != nil {
		c.HandleError(err)
		return
	}
	c.OK(res)
}
