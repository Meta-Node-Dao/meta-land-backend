package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/account"
	"ceres/pkg/router/account/social"
	"ceres/pkg/router/bounty"
	"ceres/pkg/router/chain"
	"ceres/pkg/router/crowdfunding"
	"ceres/pkg/router/dict"
	"ceres/pkg/router/governance"
	"ceres/pkg/router/image"
	"ceres/pkg/router/middleware"
	"ceres/pkg/router/startup"
	"ceres/pkg/router/tag"
	"ceres/pkg/router/upload"

	"github.com/gotomicro/ego/server/egin"
)

// Gin instance
var Gin *egin.Component

// Init the Gin instance and the routers
func Init() (err error) {
	Gin = egin.Load("server.http").Build()

	apiRoot := Gin.Group("/api")

	startupsGroup := apiRoot.Group("/")
	{
		startupsGroup.GET("/startups", router.Wrap(startup.ListStartups))
	}

	// oauth login router
	oauthLogin := apiRoot.Group("/account/oauth")
	{
		oauthLogin.Use(middleware.GuestAuthorizationMiddleware())
		oauthLogin.GET("/github/login/callback", router.Wrap(account.LoginWithGithubCallback))
		oauthLogin.GET("/google/login/callback", router.Wrap(account.LoginWithGoogleCallback))

		oauthLogin.Use(middleware.ComerAuthorizationMiddleware())
		oauthLogin.POST("/register", router.Wrap(account.RegisterWithOauth))
	}

	// web3 login router
	web3Login := apiRoot.Group("/authorizations")
	{
		web3Login.Use(middleware.GuestAuthorizationMiddleware())
		web3Login.GET(":address/nonce", router.Wrap(account.GetBlockchainLoginNonce))
		web3Login.POST("/wallet", router.Wrap(account.LoginWithWallet))
	}

	// accounts operation router
	accountPriv := apiRoot.Group("/")
	{
		accountPriv.Use(middleware.ComerAuthorizationMiddleware())
		// basic operations
		accountPriv.GET("/list", router.Wrap(account.ListAccounts))
		accountPriv.GET("/user/info", router.Wrap(account.UserInfo))
		accountPriv.POST("/eth/wallet/link", router.Wrap(account.LinkWithWallet))
		accountPriv.DELETE("/:accountID/unlink", router.Wrap(account.UnlinkAccount))
		// profile operations
		//accountPriv.GET("/profile", router.Wrap(account.GetProfile))
		accountPriv.POST("/profile", router.Wrap(account.CreateProfile))
		accountPriv.PUT("/profile", router.Wrap(account.UpdateProfile))
		//////////////
		accountPriv.PUT("/profile/basic", router.Wrap(account.UpdateBasic))
		accountPriv.POST("/profile/social", router.Wrap(social.UpdateSocial))
		accountPriv.DELETE("/profile/social", router.Wrap(social.ClearSocial))
		accountPriv.POST("/profile/skills", router.Wrap(account.UpdateComerSkill))
		accountPriv.POST("/profile/cover", router.Wrap(account.UpdateCover))
		accountPriv.POST("/profile/bio", router.Wrap(account.UpdateComerBio))
		accountPriv.POST("/profile/languages", router.Wrap(account.UpdateComerLanguages))
		accountPriv.POST("/profile/educations", router.Wrap(account.UpdateComerEducations))

		accountPriv.GET("/profile/modules/:comerID", router.Wrap(account.GetModulesOfTargetComer))
		////////////
		accountPriv.GET("/related-startups", router.Wrap(account.JoinedAndFollowedStartups))

		// comer operations
		accountPriv.POST("/comer/:comerID/follow", router.Wrap(account.FollowComer))
		accountPriv.DELETE("/comer/:comerID/unfollow", router.Wrap(account.UnfollowComer))
		accountPriv.POST("/comer/:comerID/fans", router.Wrap(account.GetConnectorsOfComer))
		accountPriv.POST("/comer/:comerID/followed-startups", router.Wrap(account.GetComerFollowedStartups))
		accountPriv.POST("/comer/:comerID/following", router.Wrap(account.GetComersFollowedByComer))
		accountPriv.GET("/comer/:comerID/followedByMe", router.Wrap(account.ComerFollowedByMe))
		accountPriv.GET("/comer", router.Wrap(account.GetProfile))

	}

	// accounts operation router
	accountsPub := apiRoot.Group("")
	{
		accountsPub.Use(middleware.GuestAuthorizationMiddleware())
		accountsPub.GET("/comer/:comerID", router.Wrap(account.GetComerInfo))
		//accountsPub.GET("/comer", router.Wrap(account.GetComerInfo))
		accountsPub.GET("/comer/address/:address", router.Wrap(account.GetComerInfoByAddress))
		accountsPub.GET("/comer/:comerID/posted-crowdfundings", router.Wrap(crowdfunding.GetComerPostedCrowdfundingList))
		accountsPub.GET("/comer/:comerID/participated-crowdfundings", router.Wrap(crowdfunding.GetComerParticipatedCrowdfundingList))
		accountsPub.GET("/comer/:comerID/connected-count", router.Wrap(account.ProfileComerConnectedInfo))
		accountsPub.GET("/comer/:comerID/data-count", router.Wrap(account.ProfileComerModuleDataCntInfo))
	}

	coresPriv := apiRoot.Group("/")
	{
		coresPriv.Use(middleware.ComerAuthorizationMiddleware())
		coresPriv.GET("/startups/me", router.Wrap(startup.ListStartupsMe))
		coresPriv.GET("/startups/existence", router.Wrap(startup.Existence))
		coresPriv.POST("/startups", router.Wrap(startup.CreateStartup))
		coresPriv.GET("/startups/comer/:comerID/posted", router.Wrap(startup.ListStartupsPostedByComer))
		// coresPriv.GET("/startups", router.Wrap(crowdfunding.SelectNonFundingStartups))
		//coresPriv.GET("/startups/crowdfundable", router.Wrap(crowdfunding.SelectNonFundingStartups))
		coresPriv.POST("/startups/:startupID/follow", router.Wrap(startup.FollowStartup))
		coresPriv.DELETE("/startups/:startupID/unfollow", router.Wrap(startup.UnfollowStartup))
		coresPriv.GET("/startups/follow", router.Wrap(startup.ListFollowStartups))
		coresPriv.GET("/startups/participate", router.Wrap(startup.ListParticipateStartups))
		coresPriv.GET("/startups/comer/:comerID/participate", router.Wrap(startup.ListParticipateStartupsOfComer))
		coresPriv.GET("/startups/:startupID/teamMembers", router.Wrap(startup.ListStartupTeamMembers))
		coresPriv.POST("/startups/:startupID/teamMembers/:comerID", router.Wrap(startup.CreateStartupTeamMember))
		coresPriv.PUT("/startups/:startupID/teamMembers/:comerID", router.Wrap(startup.UpdateStartupTeamMember))
		coresPriv.DELETE("/startups/:startupID/teamMembers/:comerID", router.Wrap(startup.DeleteStartupTeamMember))
		coresPriv.PUT("/startups/:startupID/basicSetting", router.Wrap(startup.UpdateStartupBasicSetting))
		coresPriv.PUT("/startups/:startupID/basicSetting1", router.Wrap(startup.UpdateStartupBasicSetting1))
		coresPriv.PUT("/startups/:startupID/financeSetting", router.Wrap(startup.UpdateStartupFinanceSetting))
		coresPriv.GET("/startups/:startupID/followedByMe", router.Wrap(startup.StartupFollowedByMe))
		coresPriv.GET("/startups/:startupID/fans", router.Wrap(startup.ComersFollowedThisStartup))
		///////
		coresPriv.POST("/startups/:startupID/cover", router.Wrap(startup.UpdateStartupCover))
		coresPriv.POST("/startups/:startupID/security", router.Wrap(startup.UpdateStartupSecurity))
		coresPriv.POST("/startups/:startupID/sequence", router.Wrap(startup.UpdateStartupTabSequence))
		coresPriv.POST("/startups/:startupID/groups", router.Wrap(startup.CreateStartupGroup))
		coresPriv.POST("/startups/:startupID/social", router.Wrap(startup.UpdateStartupSocialAndTags))
		coresPriv.DELETE("/startups/:startupID/social", router.Wrap(startup.RemoveStartupSocial))
		coresPriv.DELETE("/startups/group/:groupID", router.Wrap(startup.DeleteStartupGroup))
		coresPriv.GET("/startups/:startupID/group/:groupID/members", router.Wrap(startup.GetStartupGroupMembers))
		coresPriv.PUT("/startups/group/:groupID", router.Wrap(startup.UpdateStartupGroup))
		coresPriv.GET("/startups/:startupID/groups", router.Wrap(startup.GetStartupGroups))
		coresPriv.POST("/startups/:startupID/group/:groupID/member/:comerID", router.Wrap(startup.ChangeComerGroupAndLocation))
		coresPriv.GET("/startups/:startupID/data-count", router.Wrap(startup.GetStartupModuleDataCount))
		/////////
		// bounty
		coresPriv.GET("/bounties", router.Wrap(bounty.GetPublicBountyList))
		coresPriv.GET("/bounties/startup/:startupId", router.Wrap(bounty.GetBountyListByStartup))
		coresPriv.GET("/bounties/me/participated", router.Wrap(bounty.GetMyParticipatedBountyList))
		coresPriv.GET("/bounties/me/posted", router.Wrap(bounty.GetMyPostedBountyList))
		coresPriv.GET("/bounties/comer/:comerID/participated", router.Wrap(bounty.GetComerParticipatedBountyList))
		coresPriv.GET("/bounties/comer/:comerID/posted", router.Wrap(bounty.GetComerPostedBountyList))
		// 1. crowdfunding
		coresPriv.POST("/crowdfunding", router.Wrap(crowdfunding.CreateCrowdfunding))
		// 2. public
		coresPriv.GET("/crowdfundings", router.Wrap(crowdfunding.GetCrowdfundingList))
		// 3. detail
		coresPriv.GET("/crowdfundings/:crowdfundingId", router.Wrap(crowdfunding.GetCrowdfundingDetail))
		// 4. cancel
		coresPriv.POST("/crowdfundings/:crowdfundingId/cancel", router.Wrap(crowdfunding.CancelCrowdfunding))
		// 5. finalize/remove
		coresPriv.POST("/crowdfundings/:crowdfundingId/remove", router.Wrap(crowdfunding.RemoveCrowdfunding))
		// 6. posted
		coresPriv.POST("/crowdfundings/posted", router.Wrap(crowdfunding.GetMyPostedCrowdfundingList))
		// 7. participated
		coresPriv.POST("/crowdfundings/participated", router.Wrap(crowdfunding.GetMyParticipatedCrowdfundingList))
		// 8. buy or sell
		coresPriv.POST("/crowdfundings/:crowdfundingId/invest", router.Wrap(crowdfunding.Invest))
		// 9. history
		coresPriv.PUT("/crowdfundings/:crowdfundingId/modify", router.Wrap(crowdfunding.ModifyCrowdfunding))
		// 10. history-list
		coresPriv.POST("/crowdfundings/:crowdfundingId/histories", router.Wrap(crowdfunding.GetBuyPriceAndSwapModificationHistories))
		// 11. invest records
		coresPriv.POST("/crowdfundings/:crowdfundingId/investments", router.Wrap(crowdfunding.GetCrowdfundingSwapRecords))
		// 12. investor profile
		coresPriv.GET("/crowdfundings/:crowdfundingId/investor", router.Wrap(crowdfunding.GetInvestProfile))
		// 13. crowdfunding list of startup (pagination) todo !
		coresPriv.POST("/crowdfundings/startup/:startupId", router.Wrap(crowdfunding.GetCrowdfundingListOfStartup))
		// -- governance start -- //
		coresPriv.POST("/governance-setting/:startupID", router.Wrap(governance.CreateGovernanceSetting))
		coresPriv.GET("/startups/:startupID/governance-setting", router.Wrap(governance.GetGovernanceSetting))
		coresPriv.POST("/proposals", router.Wrap(governance.CreateProposal))
		coresPriv.GET("/proposals/:proposalID", router.Wrap(governance.GetProposal))
		coresPriv.DELETE("/proposals/:proposalID", router.Wrap(governance.DeleteProposal))
		coresPriv.POST("/proposals/public-list", router.Wrap(governance.PublicList))
		coresPriv.POST("/proposals/startup/:startupID", router.Wrap(governance.StartupProposalList))
		coresPriv.POST("/proposals/comer/:comerID/participate", router.Wrap(governance.ComerParticipateProposalList))
		coresPriv.POST("/proposals/comer/:comerID/post", router.Wrap(governance.ComerPostProposalList))
		coresPriv.POST("/proposals/:proposalID/vote", router.Wrap(governance.VoteProposal))
		coresPriv.POST("/proposals/:proposalID/vote-records", router.Wrap(governance.ProposalVoteRecords))
		coresPriv.GET("/proposals/:proposalID/vote-info", router.Wrap(governance.ProposalVoteInfo))
		// -- governance end -- //

	}

	coresPub := apiRoot.Group("/cores")
	{
		coresPub.Use(middleware.GuestAuthorizationMiddleware())
		coresPub.GET("/startups", router.Wrap(startup.ListStartups))
		coresPub.GET("/startups/:startupID", router.Wrap(startup.GetStartup))
		coresPub.GET("/startups/name/:name/isExist", router.Wrap(startup.StartupNameIsExist))
		coresPub.GET("/startups/tokenContract/:tokenContract/isExist", router.Wrap(startup.StartupTokenContractIsExist))
		coresPub.GET("/startups/member/:comerID", router.Wrap(startup.ListBeMemberStartups))
		coresPub.GET("/startups/comer/:comerID", router.Wrap(startup.ListStartupsCreatedByComer))

		//coresPub.GET("/startups/:startupId/setting", router.Wrap(startup.GetStartupSetting))

	}

	// misc operation router
	misc := apiRoot.Group("/misc")
	{
		misc.Use(middleware.ComerAuthorizationMiddleware())
		misc.POST("/upload", router.Wrap(upload.Upload))
	}

	// meta information
	meta := apiRoot.Group("/")
	{
		meta.Use(middleware.GuestAuthorizationMiddleware())
		meta.GET("/tags", router.Wrap(tag.GetTagList))
		meta.GET("/tags/startup", router.Wrap(tag.GetsStartupTagList))

		meta.GET("/images", router.Wrap(image.GetImageList))
		meta.GET("/dicts", router.Wrap(dict.GetDictListByType))
	}

	// chain information
	chainRouter := apiRoot.Group("/chain")
	{
		chainRouter.Use(middleware.GuestAuthorizationMiddleware())
		chainRouter.GET("/list", router.Wrap(chain.GetChainList))
	}

	bounties := apiRoot.Group("/bounty")
	{
		// bounties.Use(middleware.GuestAuthorizationMiddleware())
		bounties.Use(middleware.ComerAuthorizationMiddleware())
		bounties.POST("/create", router.Wrap(bounty.CreateBounty))
		// detail: Temporary reservation
		bounties.POST("/detail", router.Wrap(bounty.CreateBounty))
		bounties.GET("/:bountyID/detail", router.Wrap(bounty.GetBountyDetailByID))
		bounties.GET("/:bountyID/payment", router.Wrap(bounty.GetPaymentByBountyID))
		bounties.POST("/:bountyID/paid", router.Wrap(bounty.PayReward))
		bounties.GET("/:bountyID/state", router.Wrap(bounty.GetState))
		bounties.GET("/:bountyID/activities", router.Wrap(bounty.GetActivitiesLists))
		bounties.POST("/:bountyID/postUpdate", router.Wrap(bounty.CreateActivities))
		bounties.GET("/:bountyID/founder", router.Wrap(bounty.GetFounderByBountyID))
		bounties.POST("/:bountyID/applicants/apply", router.Wrap(bounty.CreateApplicants))
		bounties.GET("/:bountyID/applicants", router.Wrap(bounty.GetAllApplicantsByBountyID))
		bounties.POST("/:bountyID/approve/:applicantComerID", router.Wrap(bounty.UpdateFounderApprovedApplicant))
		bounties.POST("/:bountyID/unapprove/:applicantComerID", router.Wrap(bounty.UpdateFounderUnapprovedApplicant))
		bounties.GET("/:bountyID/approved", router.Wrap(bounty.GetApprovedApplicantByBountyID))
		bounties.POST("/:bountyID/addDeposit", router.Wrap(bounty.AddDeposit))
		bounties.POST("/:bountyID/release", router.Wrap(bounty.ReleaseDeposit))
		bounties.POST("/:bountyID/releaseMyDeposit", router.Wrap(bounty.ReleaseMyDeposit))
		bounties.POST("/:bountyID/applicant/lock", router.Wrap(bounty.UpdateApplicantsLockDeposit))
		bounties.POST("/:bountyID/applicant/unlock", router.Wrap(bounty.UpdateApplicantsUnlockDeposit))
		bounties.POST("/:bountyID/close", router.Wrap(bounty.UpdateBountyCloseStatus))
		bounties.GET("/:bountyID/startup", router.Wrap(bounty.GetStartupByBountyID))
		bounties.GET("/:bountyID/deposits", router.Wrap(bounty.GetDepositRecords))
	}

	//testR := Gin.Group("/test")
	//{
	//	testR.Use(middleware.GuestAuthorizationMiddleware())
	//	testR.GET("/token/:comerId", router.Wrap(my.Token))
	//}
	return
}
