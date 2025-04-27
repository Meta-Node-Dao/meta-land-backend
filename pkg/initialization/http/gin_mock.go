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

	comer := apiRoot.Group("/comer")
	{
		comer.GET("/", router.Wrap(mock.GetComer))
		comer.PUT("/", router.Wrap(mock.UpdateComerInfo))
		comer.DELETE("/accounts/:comer_account_id", router.Wrap(mock.UnlinkOauthByComerAccountId))
		comer.PUT("/bio", router.Wrap(mock.UpdateComerInfoBio))
		comer.GET("detail", router.Wrap(mock.GetComerInfoDetail))
		comer.POST("educations", router.Wrap(mock.BindComerEducations))
		comer.PUT("educations/:comer_education_id", router.Wrap(mock.UpdateComerEducation))
		comer.DELETE("educations/:comer_education_id", router.Wrap(mock.UnbindComerEducations))
		comer.GET("/invitation-count", router.Wrap(mock.GetComerInvitationCount))
		comer.GET("/invitation-records", router.Wrap(mock.GetComerInvitationRecords))
		comer.POST("/languages", router.Wrap(mock.BindComerLanguages))
		comer.PUT("/languages/:comer_language_id", router.Wrap(mock.UpdateComerLanguages))
		comer.DELETE("/languages/:comer_language_id", router.Wrap(mock.UnbindComerLanguages))
		comer.GET("/related-startups", router.Wrap(mock.GetComerJoinedAndFollowedStartups))
		comer.POST("/skills", router.Wrap(mock.BindComerSkills))
		comer.POST("/socials", router.Wrap(mock.BindComerSocials))
		comer.PUT("/socials/:soical_book_id", router.Wrap(mock.UpdateComerSocials))
		comer.DELETE("/socials/:soical_book_id", router.Wrap(mock.UnbindComerSocials))
	}

	comers := apiRoot.Group("/comers")
	{
		comers.GET("/address/:address", router.Wrap(mock.GetComerByAddress))
		comers.PUT("/domains", router.Wrap(mock.SetUserCustomDomain))
		comers.GET("/domains/existence", router.Wrap(mock.GetUserCustomDomainExistence))
		comers.GET("/domains/:custom_domain", router.Wrap(mock.GetUserCustomDomain))
		comers.GET("/verify/profile", router.Wrap(mock.VerifyComerAddProfile))
		comers.GET("/:comer_id", router.Wrap(mock.GetComerByComerId))
		comers.GET("/:comer_id/be_connect/comers", router.Wrap(mock.GetComerBeConnectComersByComerId))
		comers.POST("/:comer_id/connect", router.Wrap(mock.ConnectComer))
		comers.DELETE("/:comer_id/connect", router.Wrap(mock.UnconnectComer))
		comers.GET("/:comer_id/connect/comers", router.Wrap(mock.GetComerConnectComersByComerId))
		comers.GET("/:comer_id/connect/startups", router.Wrap(mock.GetStartupConnectByComerId))
		comers.GET("/:comer_id/connected", router.Wrap(mock.ConnectedComer))
		comers.GET("/:comer_id/detail", router.Wrap(mock.GetComerInfoDetailByComerId))
		comers.GET("/:comer_id/participated/count", router.Wrap(mock.GetComerParticipatedCountByComerId))
		comers.GET("/:comer_id/posted/count", router.Wrap(mock.GetComerPostedCountByComerId))
	}

	crowdfunding := apiRoot.Group("/crowdfundings")
	{
		crowdfunding.GET("/", router.Wrap(mock.GetCrowdfunding))
		crowdfunding.PUT("/", router.Wrap(mock.UpdateCrowdfunding))
		crowdfunding.POST("/", router.Wrap(mock.CreateCrowdfunding))
		crowdfunding.GET("/:crowdfunding_id", router.Wrap(mock.GetCrowdfundingInfo))
		crowdfunding.GET("/:crowdfunding_id/sign", router.Wrap(mock.GetCrowdfundingTransferLpSign))
		crowdfunding.GET("/:crowdfunding_id/swap-records", router.Wrap(mock.GetCrowdfundingInvestRecords))
	}

	dataDict := apiRoot.Group("/dict")
	{
		dataDict.GET("/:type", router.Wrap(mock.GetDataDict))
	}

	governance := apiRoot.Group("/governance")
	{
		governance.GET("/setting/:startup_id", router.Wrap(mock.GetGovernanceSetting))
		governance.POST("/setting/:startup_id", router.Wrap(mock.CreateGovernanceSetting))
	}

	languages := apiRoot.Group("/languages")
	{
		languages.GET("/", router.Wrap(mock.GetLanguages))
	}

	proposal := apiRoot.Group("/proposals")
	{
		proposal.GET("/", router.Wrap(mock.GetProposal))
		proposal.POST("/", router.Wrap(mock.CreateProposal))
		proposal.GET("/:proposal_id", router.Wrap(mock.GetProposalInfo))
		proposal.DELETE("/:proposal_id", router.Wrap(mock.DeleteProposal))
		proposal.POST("/:proposal_id/vote", router.Wrap(mock.VoteProposal))
		proposal.GET("/:proposal_id/votes", router.Wrap(mock.GetProposalInvestRecords))
	}

	saleLaunchPad := apiRoot.Group("/sale_launchpads")
	{
		saleLaunchPad.GET("/", router.Wrap(mock.GetSaleLaunchPad))
		saleLaunchPad.PUT("/", router.Wrap(mock.UpdateSaleLaunchPad))
		saleLaunchPad.POST("/", router.Wrap(mock.CreateSaleLaunchPad))
		saleLaunchPad.GET("/supply_dex", router.Wrap(mock.GetSaleLaunchPadSupplyDex))
		saleLaunchPad.GET("/:sale_launchpad_id", router.Wrap(mock.GetSaleLaunchPadInfo))
		saleLaunchPad.GET("/:sale_launchpad_id/history", router.Wrap(mock.GetSaleLaunchPadHistoryRecords))
		saleLaunchPad.GET("/:sale_launchpad_id/sign", router.Wrap(mock.GetSaleLaunchPadTransferLpSign))

	}

	share := apiRoot.Group("/share")
	{
		share.PUT("/", router.Wrap(mock.SetShare))
		share.GET("/:share_code", router.Wrap(mock.GetSharePageHtml))
	}

	social := apiRoot.Group("/socials")
	{
		social.GET("/", router.Wrap(mock.GetSocials))
	}

	startup := apiRoot.Group("/startups")
	{
		startup.GET("/", router.Wrap(mock.GetStartups))
		startup.POST("/", router.Wrap(mock.CreateStartup))
		startup.GET("/existence", router.Wrap(mock.GetStartupIsExistence))
		startup.GET("/:startup_id", router.Wrap(mock.GetStartupInfo))
		startup.PUT("/:startup_id", router.Wrap(mock.UpdateStartup))
		startup.POST("/:startup_id/connect", router.Wrap(mock.ConnectStartup))
		startup.GET("/:startup_id/connect/comers", router.Wrap(mock.GetComerConnectStartupComersByStartupId))
		startup.GET("/:startup_id/connected", router.Wrap(mock.ConnectedStartup))
		startup.PUT("/:startup_id/finance", router.Wrap(mock.SetStartupFinance))
		startup.GET("/:startup_id/relation/count", router.Wrap(mock.GetStartupRelationCount))
		startup.PUT("/:startup_id/security", router.Wrap(mock.UpdateStartupSecurity))
		startup.POST("/:startup_id/socials", router.Wrap(mock.BindStartupSocials))
		startup.PUT("/:startup_id/tab_sequence", router.Wrap(mock.UpdateStartupTabSequence))
		startup.GET("/:startup_id/team/comers", router.Wrap(mock.GetStartupTeam))
		startup.POST("/:startup_id/team/comers", router.Wrap(mock.SaveComerToStartupTeam))
		startup.DELETE("/:startup_id/team/comers/:startup_team_comer_id", router.Wrap(mock.DeleteComerOfStartupTeam))
		startup.GET("/:startup_id/team/comers/:startup_team_comer_id/existence", router.Wrap(mock.StartupTeamComerExistence))
		startup.GET("/:startup_id/team/groups", router.Wrap(mock.GetStartupTeamGroups))
		startup.POST("/:startup_id/team/groups", router.Wrap(mock.SaveStartupTeamGroup))
		startup.DELETE("/:startup_id/unconnect", router.Wrap(mock.UnconnectStartup))
	}

	tags := apiRoot.Group("/tags")
	{
		tags.GET("/:type", router.Wrap(mock.GetTagsByTagType))
	}

	return
}
