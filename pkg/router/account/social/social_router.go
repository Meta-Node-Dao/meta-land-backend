package social

import (
	"ceres/pkg/initialization/utility"
	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/router"
	"encoding/json"
)

func GetSocials(c *router.Context) {
	var listString = "[\n    {\n        \"id\": 4,\n        \"logo\": \"https://static.metaland.xyz/socials/discord.png\",\n        \"name\": \"Discord\"\n    },\n    {\n        \"id\": 5,\n        \"logo\": \"https://static.metaland.xyz/socials/github.png\",\n        \"name\": \"GitHub\"\n    }\n]"
	var socialTools []account.SocialToolResponse
	err := json.Unmarshal([]byte(listString), &socialTools)
	if err != nil {
		c.HandleError(err)
		return
	}
	var res model.PageData
	res.Total = len(socialTools)
	res.Page = 1
	res.Size = 15
	res.List = utility.ConvertToInterfaceSlice(socialTools)
	c.OK(res)
}

//func UpdateSocial(ctx *router.Context) {
//	var request account.SocialModifyRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	if err := accountService.UpdateSocial(comerId, request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
//
//func ClearSocial(ctx *router.Context) {
//	var request account.SocialRemoveRequest
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	comerId := ctx.Keys[middleware.ComerUinContextKey].(uint64)
//	if err := accountService.RemoveSocial(comerId, request.SocialType); err != nil {
//		ctx.HandleError(err)
//		return
//	}
//	ctx.OK(nil)
//}
