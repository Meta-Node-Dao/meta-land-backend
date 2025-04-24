package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/mock"
	"github.com/gotomicro/ego/server/egin"
)

func InitMock() (err error) {
	Gin = egin.Load("server.http").Build()
	apiRoot := Gin.Group("/api")

	authorization := apiRoot.Group("/authorizations")
	{
		authorization.POST("github", router.Wrap(mock.GithubOauth))
		authorization.POST("google", router.Wrap(mock.GoogleOauth))
		authorization.POST("wallet", router.Wrap(mock.LoginByWalletAddress))
		authorization.GET(":wallet_address/nonce", router.Wrap(mock.GetNonceByAddress))
	}

	bounty := apiRoot.Group("/")
	{
		bounty.GET("/bounties", router.Wrap(mock.GetBounties))
		bounty.POST("/bounties", router.Wrap(mock.CreateBounty))
		bounty.GET("/bounties/:bounty_id", router.Wrap(mock.GetBountyInfo))
		bounty.POST("/bounties/:bounty_id/apply", router.Wrap(mock.ApplyBounty))
		bounty.PUT("/bounties/:bounty_id/close", router.Wrap(mock.CloseBounty))
		bounty.PUT("/bounties/:bounty_id/payments/:bounty_payment_terms_id", router.Wrap(mock.PayBounty))
		bounty.POST("/bounties/:bounty_id/post-update", router.Wrap(mock.PostUpdateBounty))
	}

	return
}
