package pkg

//import (
//	"ceres/pkg/config"
//	"ceres/pkg/initialization/mysql"
//	"ceres/pkg/initialization/utility"
//	model2 "ceres/pkg/model"
//	account2 "ceres/pkg/model/account"
//	bountyModel "ceres/pkg/model/bounty"
//	crowdfundingModel "ceres/pkg/model/crowdfunding"
//	startupModel "ceres/pkg/model/startup"
//	"ceres/pkg/model/startup_team"
//	"ceres/pkg/service/account"
//	"ceres/pkg/service/crowdfunding"
//	"ceres/pkg/service/startup"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"gorm.io/gorm"
//	"log"
//	"testing"
//	"time"
//)
//
//func Test_CountStartupsPostedByComer(t *testing.T) {
//	if err := setup(); err != nil {
//		log.Fatal("test terminated...")
//		return
//	}
//	startupCnt, err := startupModel.CountStartupsPostedByComer(mysql.DB, 1)
//	if err != nil {
//		log.Println(err.Error())
//	} else {
//		log.Printf("startup count: %v\n", startupCnt)
//	}
//}
//
//func Test_CountBountiesPostedByComer(t *testing.T) {
//	if err := setup(); err != nil {
//		log.Fatal("test terminated...")
//		return
//	}
//	bountiesCnt, err := bountyModel.CountBountiesPostedByComer(mysql.DB, 1)
//	if err != nil {
//		log.Println(err.Error())
//	} else {
//		log.Printf("bounty count: %v\n", bountiesCnt)
//	}
//}
//
//func Test_CountCrowdfundingPostedByComer(t *testing.T) {
//	if err := setup(); err != nil {
//		log.Fatal("test terminated...")
//		return
//	}
//	crowdfundingCnt, err := crowdfundingModel.CountCrowdfundingPostedByComer(mysql.DB, 1)
//	if err != nil {
//		log.Println(err.Error())
//	} else {
//		log.Printf("crowdfunding count: %v\n", crowdfundingCnt)
//	}
//}
//
//func Test_GetComerModuleInfo(t *testing.T) {
//	if err := setup(); err != nil {
//		log.Fatal("test terminated...")
//		return
//	}
//	if info, err := account.GetComerModuleInfo(1); err != nil {
//		log.Fatal(err)
//	} else {
//		bytes, err := json.MarshalIndent(info, "", "\t")
//		if err != nil {
//			log.Fatalln(err)
//			return
//		}
//		log.Println(string(bytes))
//	}
//}
//
//func Test_GetConnectorsOfComer(t *testing.T) {
//	DoTestWithMysql(func() {
//		pagination := model2.Pagination{
//			Limit: 10,
//			Page:  1,
//		}
//
//		if err := account.GetConnectorsOfComer(1, 2, &pagination); err != nil {
//			log.Fatal(err)
//		} else {
//			bytes, err := json.MarshalIndent(pagination, "", "\t")
//			if err != nil {
//				log.Fatalln(err)
//				return
//			}
//			log.Println(string(bytes))
//		}
//	})
//}
//
//func Test_GetCrowdfundingList(t *testing.T) {
//	DoTestWithMysql(func() {
//		request := crowdfundingModel.PublicCrowdfundingListPageRequest{
//			Pagination: model2.Pagination{
//				Limit:   9,
//				Page:    1,
//				Keyword: "a",
//			},
//			Mode: 0,
//		}
//		err := crowdfunding.GetCrowdfundingList(&request)
//		if err != nil {
//			t.Fail()
//		}
//		log.Println(request)
//	})
//}
//
//func Test_UpdateLanguageInfos(t *testing.T) {
//	DoTestWithMysql(func() {
//		comerId := uint64(124274732118016)
//		err := account.UpdateLanguages(comerId, account2.UpdateLanguageInfosRequest{Languages: []account2.LanguageInfo{
//			{
//				Language: "Chinese",
//				Level:    "Beginner",
//			},
//			{
//				Language: "English",
//				Level:    "Advanced",
//			},
//		}})
//		if err != nil {
//			t.Fatal(err)
//		}
//		var a account2.ComerProfile
//		err = account2.GetComerProfile(mysql.DB, comerId, &a)
//		if err != nil {
//			t.Fatal(err)
//		}
//		fmt.Println(a)
//	})
//}
//
//func Test_ProfileComerConnectedInfo(t *testing.T) {
//	DoTestWithMysql(func() {
//		if info, err := account2.ProfileComerConnectedInfo(mysql.DB, 129525702930432); err != nil {
//			t.Fatal(err)
//		} else {
//			indent, err := json.MarshalIndent(info, "", "\t")
//			if err != nil {
//				t.Fatal(err)
//			}
//			fmt.Printf("--INFO--%s\v", string(indent))
//			pagination := model2.Pagination{Page: 1, Limit: 5}
//			err = account.GetComersFollowedByComer(129525702930432, 129525702930432, &pagination)
//			if err != nil {
//				t.Fatal(err)
//			}
//			indent, _ = json.MarshalIndent(pagination, "", "\t")
//			fmt.Printf("--current following comers--%s\n", string(indent))
//			err = account.GetConnectorsOfComer(129525702930432, 129525702930432, &pagination)
//			if err != nil {
//				t.Fatal(err)
//			}
//			indent, _ = json.MarshalIndent(pagination, "", "\t")
//			fmt.Printf("--fans--%s\n", string(indent))
//		}
//	})
//}
//
//func Test_ProfileComerModuleDataCntInfo(t *testing.T) {
//	DoTestWithMysql(func() {
//		if postedInfo, err := account2.ProfileComerModuleDataInfo(mysql.DB, 129525702930432, account2.Posted); err != nil {
//			t.Fatal(err)
//		} else {
//			fmt.Printf("Posted:::startup:%d, bounty:%d, crowdfunding:%d, proposal:%d\n", postedInfo.StartupCnt, postedInfo.BountyCnt, postedInfo.CrowdfundingCnt, postedInfo.ProposalCnt)
//		}
//		participatedInfo, err := account2.ProfileComerModuleDataInfo(mysql.DB, 129525702930432, account2.Participated)
//		if err != nil {
//			t.Fatal(err)
//		}
//		fmt.Printf("Participated:::startup:%d, bounty:%d, crowdfunding:%d, proposal:%d\n", participatedInfo.StartupCnt, participatedInfo.BountyCnt, participatedInfo.CrowdfundingCnt, participatedInfo.ProposalCnt)
//
//		request := startupModel.ListStartupRequest{
//			ListRequest: model2.ListRequest{Limit: 1000},
//			Keyword:     "",
//			Mode:        0,
//		}
//		var sts []startupModel.Startup
//		total, err := startupModel.ListParticipatedStartups(mysql.DB, 129525702930432, &request, &sts)
//		if err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(request, "", "\t")
//		fmt.Printf("--total: %d--startups: %s --", total, indent)
//		if participatedInfo.StartupCnt != total {
//			t.Fatal("..........")
//		}
//	})
//}
//
//func Test_GetFansOfComer(t *testing.T) {
//	DoTestWithMysql(func() {
//		pagination := model2.Pagination{
//			Limit: 10,
//			Page:  1,
//		}
//		if err := account.GetConnectorsOfComer(129525702930432, 129525702930432, &pagination); err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(pagination, "", "\t")
//		fmt.Printf("%v\n", string(indent))
//	})
//}
//
//func Test_GetFollowedByComer(t *testing.T) {
//	DoTestWithMysql(func() {
//		pagination := model2.Pagination{
//			Limit: 10,
//			Page:  1,
//		}
//		if err := account.GetComersFollowedByComer(143086227501056, 125678234316800, &pagination); err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(pagination, "", "\t")
//		fmt.Printf("%v\n", string(indent))
//	})
//}
//
//func Test_GetFollowedStartupsOfComer(t *testing.T) {
//	DoTestWithMysql(func() {
//		pagination := model2.Pagination{
//			Limit: 10,
//			Page:  1,
//		}
//		if err := account.GetComerFollowedStartups(129525702930432, 129525702930432, &pagination); err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(pagination, "", "\t")
//		fmt.Printf("%v\n", string(indent))
//	})
//}
//
//func Test_GetComperProfile(t *testing.T) {
//	DoTestWithMysql(func() {
//		var pro account2.ComerProfile
//		err := mysql.DB.Model(account2.ComerProfile{}).Where("id = ?", 111).Error
//		if err != nil {
//
//		}
//		// a non-existing record won't return  ErrRecordNotFound !!!!
//		err = account2.GetComerProfile(mysql.DB, 111, &pro)
//		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
//			t.Fatal("......not found.......")
//		}
//	})
//}
//
//func Test_Create(t *testing.T) {
//	var (
//		si uint64 = 148447613366272
//		ci uint64 = 147997673598976
//	)
//	config.Seq = &config.Sequence{Epoch: 1626023857}
//	err := utility.Init()
//	if err != nil {
//		t.Fatal(err)
//	}
//	createOrUpdateStartupGroupRequest := startupModel.CreateOrUpdateStartupGroupRequest{
//		Name: "I'm-a-group",
//	}
//	startupGroup, err := startup.CreateStartupGroup(si, ci, createOrUpdateStartupGroupRequest)
//	if err != nil {
//		t.Error(err)
//	}
//
//	fmt.Println("new created group id is ", startupGroup.ID)
//}
//
//func Test_GetStartupGroupMembers(t *testing.T) {
//	DoTestWithMysql(func() {
//		var (
//			si         uint64 = 148447613366272
//			gi         uint64 = 151040100085760
//			ci         uint64 = 129525702930432
//			pagination        = model2.Pagination{
//				Limit: 10,
//				Page:  1,
//			}
//		)
//
//		err := startup.GetStartupGroupMembers(si, gi, &pagination)
//		if err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(pagination, "", "\t")
//		t.Logf("group --- members: %s\n", string(indent))
//		err = startup.CreateStartupTeamMember(si, ci, &startup_team.CreateStartupTeamMemberRequest{
//			Position: "123",
//			GroupId:  0,
//		})
//		if err != nil {
//			t.Fatal(err)
//		}
//		err = startup.GetStartupGroupMembers(144893620203520, 0, &pagination)
//		if err != nil {
//			t.Fatal(err)
//		}
//		indent, _ = json.MarshalIndent(pagination, "", "\t")
//		log.Printf("startup --- members: %s\n", string(indent))
//	})
//}
//
//func Test_DeleteStartupGroup(t *testing.T) {
//	DoTestWithMysql(func() {
//		var (
//			si uint64 = 148447613366272
//			ci uint64 = 147997673598976
//			gi uint64
//		)
//
//		config.Seq = &config.Sequence{Epoch: 1626023857}
//		err := utility.Init()
//		if err != nil {
//			t.Fatal(err)
//		}
//		createOrUpdateStartupGroupRequest := startupModel.CreateOrUpdateStartupGroupRequest{
//			Name: fmt.Sprintf("I'm-a-group-%v", time.Now()),
//		}
//		startupGroup, err := startup.CreateStartupGroup(si, ci, createOrUpdateStartupGroupRequest)
//		if err != nil {
//			t.Error(err)
//		}
//
//		gi = startupGroup.ID
//		fmt.Println("new created group id is ", startupGroup.ID)
//		groups, err := startup.GetStartupGroups(si)
//		if err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(groups, "", "\t")
//		fmt.Printf("--groups--%s\n", string(indent))
//		err = startup.DeleteStartupGroup(gi, ci)
//		if err != nil {
//			t.Fatal(err)
//		}
//	})
//}
//
//func Test_ChangeComerGroupAndLocation(t *testing.T) {
//	DoTestWithMysql(func() {
//
//		var (
//			si uint64 = 148447613366272
//			ci uint64 = 147997673598976
//			gi uint64
//		)
//		config.Seq = &config.Sequence{Epoch: 1626023857}
//		err := utility.Init()
//		if err != nil {
//			t.Fatal(err)
//		}
//		createOrUpdateStartupGroupRequest := startupModel.CreateOrUpdateStartupGroupRequest{
//			Name: fmt.Sprintf("I'm-a-group-%v", time.Now()),
//		}
//		startupGroup, err := startup.CreateStartupGroup(si, ci, createOrUpdateStartupGroupRequest)
//		if err != nil {
//			t.Error(err)
//		}
//
//		gi = startupGroup.ID
//		fmt.Println("new created group id is ", startupGroup.ID)
//		err = startup.ChangeComerGroupAndPosition(si, gi, ci, startupModel.ModifyLocationRequest{})
//		if err != nil {
//			t.Fatal(err)
//		}
//	})
//}
//
//func Test_Social(t *testing.T) {
//	DoTestWithMysql(func() {
//		var (
//			si uint64 = 148447613366272
//			ci uint64 = 147997673598976
//		)
//		err := startup.UpdateSocialsAndTags(si, ci, startupModel.UpdateStartupSocialsAndTagsRequest{
//			HashTags: nil,
//			Socials: []account2.SocialModifyRequest{
//				{SocialType: account2.SocialEmail, SocialLink: "1111@qq.com"},
//				{SocialType: account2.SocialDiscord, SocialLink: "111111111111"},
//				{SocialType: account2.SocialFacebook, SocialLink: "111111111111"},
//				{SocialType: account2.SocialMedium, SocialLink: "111111111111"},
//			},
//			DeletedSocials: []account2.SocialType{account2.SocialMedium, account2.SocialTelegram},
//		})
//		if err != nil {
//			t.Fatal(err)
//		}
//	})
//}
//
//func Test_CreateStartupTeamMember(t *testing.T) {
//	DoTestWithMysql(func() {
//		stm := startup_team.StartupTeamMember{
//			ComerID:   151243842596864,
//			StartupID: 148447613366272,
//			Position:  "Haki",
//		}
//		err := startup.CreateStartupTeamMember(stm.StartupID, stm.ComerID, &startup_team.CreateStartupTeamMemberRequest{
//			ComerID:  151243842596864,
//			Position: "Haki",
//			GroupId:  151304416755712,
//		})
//		if err != nil {
//			t.Fatal(err)
//		}
//		// group id is 0
//		err = startup.CreateStartupTeamMember(stm.StartupID, stm.ComerID, &startup_team.CreateStartupTeamMemberRequest{
//			ComerID:  151243842596864,
//			Position: "Haki",
//			GroupId:  0,
//		})
//		if err != nil {
//			t.Fatal(err)
//		}
//	})
//}
//
//func Test_UpdateStartupTeamMember(t *testing.T) {
//	DoTestWithMysql(func() {
//		var (
//			ci uint64 = 151243842596864
//			gi uint64 = 151304416755712
//			si uint64 = 148447613366272
//		)
//		stm := startup_team.UpdateStartupTeamMemberRequest{
//			Position: "Haki",
//			GroupId:  gi,
//		}
//		err := startup.UpdateStartupTeamMember(si, ci, &stm)
//		if err != nil {
//			t.Fatal(err)
//		}
//		// group id is 0
//		stm.GroupId = 0
//		err = startup.UpdateStartupTeamMember(si, ci, &stm)
//		if err != nil {
//			t.Fatal(err)
//		}
//	})
//}
//
//func Test_GetComersFollowedThisStartup(t *testing.T) {
//	DoTestWithMysql(func() {
//		var (
//			ci   uint64 = 129525702930432
//			si   uint64 = 124896046952448
//			page        = model2.Pagination{Limit: 10, Page: 1}
//		)
//		err := startup.GetComersFollowedThisStartup(ci, si, &page)
//		if err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(page, "", "\t")
//		log.Printf("--page:\n%s", string(indent))
//	})
//}
//
//func Test_ListStartupTeamMembers(t *testing.T) {
//	DoTestWithMysql(func() {
//		var (
//			ci uint64 = 129525702930432
//			si uint64 = 148447613366272
//			// page        = model2.Pagination{Limit: 10, Page: 1}
//		)
//		response := startup_team.ListStartupTeamMemberResponse{}
//		request := startup_team.ListStartupTeamMemberRequest{ListRequest: model2.ListRequest{
//			Limit:  10,
//			Offset: 0,
//		}}
//		err := startup.ListStartupTeamMembers(si, ci, &request, &response)
//		if err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(response, "", "\t")
//		log.Printf("--page:\n%s", string(indent))
//	})
//}
//
//func Test_GetCrowdfundingListByStartup(t *testing.T) {
//	DoTestWithMysql(func() {
//		list, err := crowdfunding.GetCrowdfundingListByStartup(129597652021248)
//		if err != nil {
//			t.Fatal(err)
//		}
//		indent, _ := json.MarshalIndent(list, "", "\t")
//		t.Logf("\n%s\n", string(indent))
//	})
//}
//
//func Test_UpdateStartupFinanceSetting(t *testing.T) {
//	DoTestWithMysql(func() {
//		var (
//			si uint64 = 129597652021248
//			ci uint64 = 129525702930432
//		)
//		startup.UpdateStartupFinanceSetting(si, ci, &startupModel.UpdateStartupFinanceSettingRequest{
//			TokenContractAddress: "",
//			LaunchNetwork:        nil,
//			TokenName:            nil,
//			TokenSymbol:          nil,
//			TotalSupply:          nil,
//			PresaleStart:         nil,
//			PresaleEnd:           nil,
//			LaunchDate:           nil,
//			Wallets:              nil,
//		})
//	})
//}
