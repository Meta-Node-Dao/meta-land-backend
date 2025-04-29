package startup

import (
	"ceres/pkg/model"
	"ceres/pkg/model/startup"
	"ceres/pkg/router"
	"encoding/json"
)

var startupPageDataString = "{\n    \"list\": [\n        {\n            \"banner\": \"https://storage.metaland.xyz/startups/banner_001.jpg\",\n            \"chain_id\": 1,\n            \"comer_id\": 1001,\n            \"connected_total\": 245,\n            \"contract_audit\": \"audit_report_v3.pdf\",\n            \"finance\": {\n                \"chain_id\": 1,\n                \"comer_id\": 1001,\n                \"contract_address\": \"0x1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b\",\n                \"id\": 1,\n                \"launched_at\": \"2023-06-01T00:00:00Z\",\n                \"name\": \"MetaLab Token\",\n                \"presale_ended_at\": \"2023-05-31T23:59:59Z\",\n                \"presale_started_at\": \"2023-05-01T00:00:00Z\",\n                \"startup_id\": 501,\n                \"supply\": 10000000,\n                \"symbol\": \"MLT\",\n                \"wallets\": [\n                    {\n                        \"address\": \"0x9c0d1e2f3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b\",\n                        \"id\": 1,\n                        \"name\": \"Team Wallet\",\n                        \"startup_finance_id\": 1,\n                        \"startup_id\": 501\n                    }\n                ]\n            },\n            \"id\": 501,\n            \"is_connected\": true,\n            \"kyc\": \"kyc_verified_v4.pdf\",\n            \"logo\": \"https://storage.metaland.xyz/startups/logo_ml.png\",\n            \"mission\": \"Building decentralized metaverse infrastructure\",\n            \"name\": \"MetaLab\",\n            \"on_chain\": true,\n            \"tags\": [\n                {\n                    \"id\": 5,\n                    \"name\": \"Metaverse\"\n                }\n            ],\n            \"team\": [\n                {\n                    \"comer\": {\n                        \"activation\": true,\n                        \"address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n                        \"avatar\": \"https://storage.metaland.xyz/avatars/1001.png\",\n                        \"banner\": \"https://storage.metaland.xyz/banners/1001.jpg\",\n                        \"custom_Domain\": \"founder.metalab.xyz\",\n                        \"id\": 1001,\n                        \"invitation_code\": \"MLAB2023\",\n                        \"is_connected\": true,\n                        \"location\": \"Singapore\",\n                        \"name\": \"James Wang\",\n                        \"time_zone\": \"UTC+8\"\n                    },\n                    \"comer_id\": 1001,\n                    \"created_at\": \"2023-01-15T09:30:00Z\",\n                    \"id\": 1,\n                    \"position\": \"CEO\",\n                    \"startup_id\": 501,\n                    \"startup_team_group\": {\n                        \"comer_id\": 1001,\n                        \"id\": 1,\n                        \"name\": \"Core Team\",\n                        \"startup_id\": 501\n                    },\n                    \"startup_team_group_id\": 1\n                }\n            ],\n            \"tx_hash\": \"0x5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1a2b3c4d5e\",\n            \"type\": 3\n        }\n    ],\n    \"page\": 1,\n    \"size\": 10,\n    \"total\": 15\n}"

func GetStartups(ctx *router.Context) {
	var res startup.StartupPageData
	err := json.Unmarshal([]byte(startupPageDataString), &res)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(res)
}

func CreateStartup(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "create startup successful!"
	ctx.OK(message)
}

func GetStartupIsExistence(ctx *router.Context) {
	var res model.IsExistResponse
	res.IsExist = false
	ctx.OK(res)
}

func GetStartupInfo(ctx *router.Context) {
	startupInfoString := "{\n    \"banner\": \"https://storage.metaland.xyz/startups/banner_ml.jpg\",\n    \"chain_id\": 1,\n    \"comer_id\": 1001,\n    \"connected_total\": 356,\n    \"contract_audit\": \"security_audit_v2.pdf\",\n    \"finance\": {\n        \"chain_id\": 1,\n        \"comer_id\": 1001,\n        \"contract_address\": \"0x3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f\",\n        \"id\": 501,\n        \"launched_at\": \"2023-09-01T00:00:00Z\",\n        \"name\": \"MetaLab Governance Token\",\n        \"presale_ended_at\": \"2023-08-31T23:59:59Z\",\n        \"presale_started_at\": \"2023-08-01T00:00:00Z\",\n        \"startup_id\": 1001,\n        \"supply\": 1000000000,\n        \"symbol\": \"MGT\",\n        \"wallets\": [\n            {\n                \"address\": \"0x5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1a2b3c4d\",\n                \"id\": 1,\n                \"name\": \"Liquidity Pool\",\n                \"startup_finance_id\": 501,\n                \"startup_id\": 1001\n            }\n        ]\n    },\n    \"id\": 1001,\n    \"is_connected\": true,\n    \"kyc\": \"kyc_metaverse_v3.pdf\",\n    \"logo\": \"https://storage.metaland.xyz/startups/logo_mgt.png\",\n    \"mission\": \"构建去中心化元宇宙基础设施协议\",\n    \"name\": \"MetaLab Core\",\n    \"on_chain\": true,\n    \"overview\": \"领先的元宇宙底层协议开发商，专注于跨链互操作解决方案\",\n    \"socials\": [\n        {\n            \"id\": 1,\n            \"social_tool\": {\n                \"id\": 3,\n                \"logo\": \"https://static.metaland.xyz/socials/twitter.png\",\n                \"name\": \"Twitter\"\n            },\n            \"social_tool_id\": 3,\n            \"target_id\": 1001,\n            \"type\": 1,\n            \"value\": \"https://twitter.com/MetaLabCore\"\n        }\n    ],\n    \"tab_sequence\": \"overview,team,finance\",\n    \"tags\": [\n        {\n            \"id\": 5,\n            \"tag\": {\n                \"id\": 5,\n                \"name\": \"跨链协议\",\n                \"type\": 1\n            },\n            \"tag_id\": 5,\n            \"target_id\": 1001,\n            \"type\": 2\n        }\n    ],\n    \"team\": [\n        {\n            \"comer\": {\n                \"activation\": true,\n                \"address\": \"0x9c0d1e2f3a4b5c6d7e8f9a0b1a2b3c4d5e6f7a8b\",\n                \"avatar\": \"https://storage.metaland.xyz/avatars/ceo.png\",\n                \"banner\": \"https://storage.metaland.xyz/banners/ceo_bg.jpg\",\n                \"custom_domain\": \"ceo.metalab.xyz\",\n                \"id\": 2001,\n                \"invitation_code\": \"META2023\",\n                \"is_connected\": true,\n                \"location\": \"Singapore\",\n                \"name\": \"张伟\",\n                \"time_zone\": \"UTC+8\"\n            },\n            \"comer_id\": 2001,\n            \"created_at\": \"2023-01-01T08:00:00Z\",\n            \"id\": 1,\n            \"position\": \"首席技术官\",\n            \"startup_id\": 1001,\n            \"startup_team_group\": {\n                \"comer_id\": 2001,\n                \"id\": 10,\n                \"name\": \"技术委员会\",\n                \"startup_id\": 1001\n            },\n            \"startup_team_group_id\": 10\n        }\n    ],\n    \"tx_hash\": \"0x8f9a0b1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b\",\n    \"type\": 1\n}"
	var res startup.StartupInfoResponse
	err := json.Unmarshal([]byte(startupInfoString), &res)
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.OK(res)
}

func UpdateStartup(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "update startup successful!"
	ctx.OK(message)
}

func ConnectStartup(ctx *router.Context) {
	var message model.MessageResponse
	message.Message = "connect startup successful!"
	ctx.OK(message)
}

func GetComerConnectStartupComersByStartupId(ctx *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	ctx.OK(res)
}

func ConnectedStartup(c *router.Context) {
	var res startup.IsConnectedResponse
	res.IsConnected = false
	c.OK(res)
}

func SetStartupFinance(c *router.Context) {
	var message model.MessageResponse
	message.Message = "set startup finance successful!"
	c.OK(message)
}

func GetStartupRelationCount(c *router.Context) {
	var res model.IsExistResponse
	res.IsExist = false
	c.OK(res)
}

func UpdateStartupSecurity(c *router.Context) {
	var message model.MessageResponse
	message.Message = "update startup security successful!"
	c.OK(message)
}

func BindStartupSocials(c *router.Context) {
	var message model.MessageResponse
	message.Message = "bind startup socials successful!"
	c.OK(message)
}

func GetStartupTeam(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func SaveComerToStartupTeam(c *router.Context) {
	var message model.MessageResponse
	message.Message = "save comer to startup team successful!"
	c.OK(message)
}

func DeleteComerOfStartupTeam(c *router.Context) {
	var message model.MessageResponse
	message.Message = "delete comer of startup team successful!"
	c.OK(message)
}

func StartupTeamComerExistence(c *router.Context) {
	var res model.IsExistResponse
	res.IsExist = false
	c.OK(res)
}

func GetStartupTeamGroups(c *router.Context) {
	var res model.PageData
	res.Total = 0
	res.Page = 0
	res.Size = 0
	c.OK(res)
}

func SaveStartupTeamGroup(c *router.Context) {
	var message model.MessageResponse
	message.Message = "save startup team group successful!"
	c.OK(message)
}

func UnconnectStartup(c *router.Context) {
	var message model.MessageResponse
	message.Message = "unconnect startup successful!"
	c.OK(message)
}

//func ListStartups(ctx *router.Context) {
//	var request model.StartupListRequest
//	if err := ctx.ShouldBindQuery(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//
//	var response model.StartupListResponse
//	if err := service.StartupLists(&request, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	for i, startup := range response.List {
//		response.List[i].MemberCount = len(startup.Members)
//		response.List[i].FollowCount = len(startup.Follows)
//	}
//	ctx.OK(response)
//}
//
//func CreateStartup(ctx *router.Context) {
//	var request model.CreateStartupRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	if err := service.CreateStartup(comerID, &request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK("create startup successful")
//}
//
//func Existence(ctx *router.Context) {
//
//	response := make(map[string]bool)
//	response["is_exist"] = false
//	ctx.OK(response)
//	return
//}
//
//
//func ListStartupsMe(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var request model.ListStartupRequest
//	if err := ctx.ShouldBindQuery(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//
//	var response model.ListStartupsResponse
//	if err := service.ListStartups(comerID, &request, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//
//func ListStartupsPostedByComer(ctx *router.Context) {
//	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var request model.ListStartupRequest
//	if err := ctx.ShouldBindQuery(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//
//	var response model.ListStartupsResponse
//	if err := service.ListStartups(comerID, &request, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(response)
//}
//
//
//func GetStartup(ctx *router.Context) {
//	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
//		ctx.HandleError(err)
//		return
//	}
//	var response model.GetStartupResponse
//	if err := service.GetStartup(startupID, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(response)
//}
//
//
//func StartupNameIsExist(ctx *router.Context) {
//	name := ctx.Param("name")
//	if name == "" {
//		err := router.ErrBadRequest.WithMsg("Startup's name has been used")
//		ctx.HandleError(err)
//		return
//	}
//	isExist, err := service.StartupNameIsExist(name)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(model.ExistStartupResponse{IsExist: isExist})
//}
//
//
//func StartupTokenContractIsExist(ctx *router.Context) {
//	name := ctx.Param("tokenContract")
//	if name == "" {
//		err := router.ErrBadRequest.WithMsg("Startup's token contract has been used")
//		ctx.HandleError(err)
//		return
//	}
//
//	isExist, err := service.StartupTokenContractIsExist(name)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(model.ExistStartupResponse{IsExist: isExist})
//}
//
//
//func FollowStartup(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
//		ctx.HandleError(err)
//		return
//	}
//	if err = service.FollowStartup(comerID, startupID); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(nil)
//}
//
//
//func UnfollowStartup(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
//		ctx.HandleError(err)
//		return
//	}
//	if err = service.UnfollowStartup(comerID, startupID); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(nil)
//}
//
//
//func ListFollowStartups(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var request model.ListStartupRequest
//	if err := ctx.ShouldBindQuery(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//
//	var response model.ListStartupsResponse
//	if err := service.ListFollowStartups(comerID, &request, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(response)
//}
//
//
//func ListParticipateStartups(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var request model.ListStartupRequest
//	if err := ctx.ShouldBindQuery(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//
//	var response model.ListStartupsResponse
//	if err := service.ListParticipateStartups(comerID, &request, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(response)
//}
//
//
//func ListParticipateStartupsOfComer(ctx *router.Context) {
//	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var request model.ListStartupRequest
//	if err := ctx.ShouldBindQuery(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//
//	var response model.ListStartupsResponse
//	if err := service.ListParticipateStartups(comerID, &request, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(response)
//}
//
//
//func ListBeMemberStartups(ctx *router.Context) {
//	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
//		ctx.HandleError(err)
//		return
//	}
//	var request model.ListStartupRequest
//	if err := ctx.ShouldBindQuery(&request); err != nil {
//		log.Warn(err)
//		err = router.ErrBadRequest.WithMsg(err.Error())
//		ctx.HandleError(err)
//		return
//	}
//	var response model.ListStartupsResponse
//	if err := service.ListBeMemberStartups(comerID, &request, &response); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(response)
//}
//
//
//func StartupFollowedByMe(ctx *router.Context) {
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
//		ctx.HandleError(err)
//		return
//	}
//	isFollowed, err := service.StartupFollowedByComer(startupID, comerID)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	ctx.OK(model.FollowedByMeResponse{IsFollowed: isFollowed})
//}
//
//func ListStartupsCreatedByComer(ctx *router.Context) {
//	comerID, err := strconv.ParseUint(ctx.Param("comerID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid comer ID")
//		ctx.HandleError(err)
//		return
//	}
//	currentComerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	responseTmp, err := service.GetStartupsByComerID(comerID, currentComerId)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//
//	response := &model.ListComerStartupsResponse{
//		Total: len(responseTmp),
//		List:  responseTmp,
//	}
//	ctx.OK(response)
//}
//
//func UpdateStartupCover(ctx *router.Context) {
//	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
//		ctx.HandleError(err)
//		return
//	}
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var request model.UpdateStartupCoverRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if err := service.UpdateStartupCover(startupID, comerID, request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func UpdateStartupSecurity(ctx *router.Context) {
//	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
//	if err != nil {
//		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
//		ctx.HandleError(err)
//		return
//	}
//	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	var request model.UpdateStartupSecurityRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if err := service.UpdateStartupSecurity(startupID, comerID, request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func ComersFollowedThisStartup(ctx *router.Context) {
//	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
//	comerId, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	if err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	var page model2.Pagination
//	if err := model2.ParsePagination(ctx, &page, 10); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	if err = service.GetComersFollowedThisStartup(comerId, startupID, &page); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(page)
//}
