package mock

import (
	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/router"
	"encoding/json"
)

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
