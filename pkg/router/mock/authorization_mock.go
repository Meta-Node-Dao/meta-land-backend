package mock

import (
	"ceres/pkg/model/account"
	"ceres/pkg/router"
)

func GithubOauth(ctx *router.Context) {
	var response account.JwtAuthorizationResponse
	response.Token = "test_token.GithubOauth"
	ctx.OK(response)
}

func GoogleOauth(ctx *router.Context) {
	var response account.JwtAuthorizationResponse
	response.Token = "test_token.GoogleOauth"
	ctx.OK(response)
}

func LoginByWalletAddress(ctx *router.Context) {
	var response account.JwtAuthorizationResponse
	response.Token = "test_token.LoginByWalletAddress"
	ctx.OK(response)
}

func GetNonceByAddress(ctx *router.Context) {
	var response account.WalletNonceResponse
	response.Nonce = "111111"
	ctx.OK(response)
}
