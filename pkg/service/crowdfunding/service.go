package crowdfunding

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model"
	"ceres/pkg/model/account"
	"ceres/pkg/model/crowdfunding"
	"ceres/pkg/model/startup"
	modelTransaction "ceres/pkg/model/transaction"
	"ceres/pkg/router"
	serviceTransaction "ceres/pkg/service/transaction"
	"fmt"
	"time"

	"github.com/qiniu/x/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func CreateCrowdfunding(request crowdfunding.CreateCrowdfundingRequest) error {
	var st startup.Startup
	if err := startup.GetStartup(mysql.DB, request.StartupId, &st); err != nil {
		return err
	}
	if st.ComerID != request.ComerId {
		return router.ErrBadRequest.WithMsg("Current Comer is not founder of startup")
	}
	crowdfundingList, er := crowdfunding.SelectOnGoingByStartupId(mysql.DB, request.StartupId)
	if er != nil {
		return er
	}
	// double check
	if len(crowdfundingList) > 0 {
		return router.ErrBadRequest.WithMsg("Startup has not ended crowdfunding")
	}
	funding := crowdfunding.Crowdfunding{
		ChainInfo:    request.ChainInfo,
		SellInfo:     request.SellInfo,
		BuyInfo:      request.BuyInfo,
		StartupId:    request.StartupId,
		ComerId:      request.ComerId,
		RaiseGoal:    request.RaiseGoal,
		RaiseBalance: decimal.Zero,
		TeamWallet:   request.TeamWallet,
		SwapPercent:  request.SwapPercent,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
		Poster:       request.Poster,
		Youtube:      request.Youtube,
		Detail:       request.Detail,
		Description:  request.Description,
		Status:       0,
	}
	funding.SellTokenBalance = funding.SellTokenDeposit
	return mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		log.Infof("#### TO_BE_CREATED_CROWDFUNDING:: %s", funding.Json())
		if er = crowdfunding.CreateCrowdfunding(tx, &funding); er != nil {
			return
		}
		return createTransactionAndQueryContract(tx, funding.ChainId, funding.ID, funding.TxHash, serviceTransaction.CrowdfundingContractCreated, func(address string, onChainStatus int) error {
			if onChainStatus == 1 && address != "" {
				var startTime = request.StartTime
				var st = crowdfunding.Upcoming
				if startTime.After(time.Now().Add(-10*time.Second)) && startTime.Before(time.Now().Add(10*time.Second)) {
					st = crowdfunding.Live
				}
				if err := crowdfunding.UpdateCrowdfundingContractAddressAndStatus(tx, funding.ID, address, st); err != nil {
					return err
				}
				if err := modelTransaction.UpdateTransactionStatusWithRetry(tx, funding.ID, serviceTransaction.CrowdfundingContractCreated, 1); err != nil {
					return err
				}
				return nil
			}
			return nil
		})
	})
}
func SelectNonFundingStartups(comerId uint64) (startups []crowdfunding.StartupCrowdfundingInfo, err error) {
	return crowdfunding.SelectStartupsWithNonCrowdfundingOnGoing(mysql.DB, comerId)
}

func GetCrowdfundingList(pagination *crowdfunding.PublicCrowdfundingListPageRequest) (err error) {
	err = crowdfunding.SelectCrowdfundingList(mysql.DB, pagination)
	if err != nil {
		return err
	}
	var items []crowdfunding.PublicItem
	if pagination.Rows != nil {
		for _, c := range pagination.Rows.([]crowdfunding.Crowdfunding) {
			fmt.Println("==================", c.BuyTokenSymbol, c.SellTokenSymbol)
			var st startup.Startup
			if err = startup.GetStartup(mysql.DB, c.StartupId, &st); err != nil {
				return
			}
			items = append(items, crowdfunding.PublicItem{
				CrowdfundingId:       c.ID,
				CrowdfundingContract: c.CrowdfundingContract,
				StartupId:            c.StartupId,
				ComerId:              c.ComerId,
				StartupName:          st.Name,
				StartTime:            c.StartTime,
				EndTime:              c.EndTime,
				RaiseBalance:         c.RaiseBalance,
				RaiseGoal:            c.RaiseGoal,
				RaisedPercent:        c.RaiseBalance.Div(c.RaiseGoal),
				BuyPrice:             c.BuyPrice,
				SwapPercent:          c.SwapPercent,
				Poster:               c.Poster,
				Status:               c.Status,
				ChainId:              c.ChainId,
				KYC:                  st.KYC,
				ContractAudit:        st.ContractAudit,
				BuyTokenContract:     c.BuyTokenContract,
				BuyTokenSymbol:       c.BuyTokenSymbol,
				SellTokenContract:    c.SellTokenContract,
				SellTokenSymbol:      c.SellTokenSymbol,
			})
		}
		pagination.Rows = items
	}

	return nil
}

func GetCrowdfundingDetail(crowdfundingId uint64) (detail crowdfunding.Detail, err error) {
	entity, err := crowdfunding.GetCrowdfundingById(mysql.DB, crowdfundingId)
	if err != nil {
		return crowdfunding.Detail{}, err
	}
	detail = crowdfunding.Detail{
		CrowdfundingId:       entity.ID,
		ChainId:              entity.ChainId,
		CrowdfundingContract: entity.CrowdfundingContract,
		TeamWallet:           entity.TeamWallet,
		SellTokenContract:    entity.SellTokenContract,
		SellTokenName:        entity.SellTokenName,
		SellTokenSymbol:      entity.SellTokenSymbol,
		SellTokenDecimals:    entity.SellTokenDecimals,
		SellTokenSupply:      entity.SellTokenSupply,
		MaxSellPercent:       entity.MaxSellPercent,
		SellTax:              entity.SellTax,
		MaxBuyAmount:         entity.MaxBuyAmount,
		BuyTokenContract:     entity.BuyTokenContract,
		BuyPrice:             entity.BuyPrice,
		SwapPercent:          entity.SwapPercent,
		RaiseBalance:         entity.RaiseBalance,
		RaiseGoal:            entity.RaiseGoal,
		RaisedPercent:        entity.RaiseBalance.Div(entity.RaiseGoal),
		StartupId:            entity.StartupId,
		ComerId:              entity.ComerId,
		StartTime:            entity.StartTime,
		EndTime:              entity.EndTime,
		Poster:               entity.Poster,
		Youtube:              entity.Youtube,
		Detail:               entity.Detail,
		Description:          entity.Description,
		Status:               entity.Status,
	}
	return
}

func GetPostedCrowdfundingListByComer(comerId uint64, pagination *model.Pagination) error {
	err := crowdfunding.SelectCrowdfundingListByFounder(mysql.DB, comerId, pagination)
	if err != nil {
		return err
	}
	var items []crowdfunding.PublicItem
	if pagination.Rows != nil {
		for _, c := range pagination.Rows.([]crowdfunding.Crowdfunding) {
			var st startup.Startup
			if err = startup.GetStartup(mysql.DB, c.StartupId, &st); err != nil {
				return err
			}
			items = append(items, crowdfunding.PublicItem{
				CrowdfundingId:       c.ID,
				CrowdfundingContract: c.CrowdfundingContract,
				StartupId:            c.StartupId,
				ComerId:              c.ComerId,
				StartupName:          st.Name,
				StartupLogo:          st.Logo,
				StartTime:            c.StartTime,
				EndTime:              c.EndTime,
				RaiseBalance:         c.RaiseBalance,
				RaiseGoal:            c.RaiseGoal,
				RaisedPercent:        c.RaiseBalance.Div(c.RaiseGoal),
				BuyPrice:             c.BuyPrice,
				SwapPercent:          c.SwapPercent,
				Poster:               c.Poster,
				Status:               c.Status,
				ChainId:              c.ChainId,
				KYC:                  st.KYC,
				ContractAudit:        st.ContractAudit,
				BuyTokenContract:     c.BuyTokenContract,
				BuyTokenSymbol:       c.BuyTokenSymbol,
				SellTokenContract:    c.SellTokenContract,
				SellTokenSymbol:      c.SellTokenSymbol,
			})
		}
		pagination.Rows = items
	}
	return nil
}

func GetParticipatedCrowdFundingListOfComer(comerId uint64, pagination *model.Pagination) error {

	err := crowdfunding.SelectCrowdfundingListByInvestor(mysql.DB, comerId, pagination)
	if err != nil {
		return err
	}
	var items []crowdfunding.PublicItem
	if pagination.Rows != nil {
		for _, c := range pagination.Rows.([]crowdfunding.Crowdfunding) {
			var st startup.Startup
			if err = startup.GetStartup(mysql.DB, c.StartupId, &st); err != nil {
				return err
			}
			investor, err := crowdfunding.SelectInvestorByCrowdfundingIdAndComerId(mysql.DB, c.ID, comerId)
			if err != nil {
				return err
			}
			items = append(items, crowdfunding.PublicItem{
				CrowdfundingId:       c.ID,
				CrowdfundingContract: c.CrowdfundingContract,
				StartupId:            c.StartupId,
				ComerId:              c.ComerId,
				StartupName:          st.Name,
				StartupLogo:          st.Logo,
				StartTime:            c.StartTime,
				EndTime:              c.EndTime,
				RaiseBalance:         c.RaiseBalance,
				RaiseGoal:            c.RaiseGoal,
				RaisedPercent:        c.RaiseBalance.Div(c.RaiseGoal),
				BuyPrice:             c.BuyPrice,
				SwapPercent:          c.SwapPercent,
				Poster:               c.Poster,
				Status:               c.Status,
				ChainId:              c.ChainId,
				KYC:                  st.KYC,
				ContractAudit:        st.ContractAudit,
				BuyTokenContract:     c.BuyTokenContract,
				SellTokenContract:    c.SellTokenContract,
				BuyTokenSymbol:       c.BuyTokenSymbol,
				BuyTokenAmount:       &investor.BuyTokenTotal,
			})
		}
		pagination.Rows = items
	}
	return nil
}

func CancelCrowdfunding(comerId, crowdfundingId uint64, txHash string) (err error) {
	entity, err := crowdfunding.GetCrowdfundingById(mysql.DB, crowdfundingId)
	if err != nil {
		return err
	}
	if entity.ComerId != comerId {
		return router.ErrBadRequest.WithMsg("current comer is not funder of the crowdfunding")
	}
	if entity.Status != crowdfunding.Upcoming && entity.Status != crowdfunding.Pending {
		return router.ErrBadRequest.WithMsg("current crowdfunding can not be cancelled")
	}
	return createTransactionAndQueryContract(mysql.DB, entity.ChainId, entity.ID, txHash, serviceTransaction.CrowdfundingCancelled, func(address string, onChainStatus int) error {
		// set crowdfunding status to `Cancelled`(todo: But what will happen if on-chain-failed???)
		if onChainStatus == 1 {
			return crowdfunding.UpdateCrowdfundingStatus(mysql.DB, crowdfundingId, crowdfunding.Cancelled)
		}
		return nil
	})

}

func FinalizeCrowdFunding(comerId, crowdfundingId uint64, txHash string) (err error) {
	entity, err := crowdfunding.GetCrowdfundingById(mysql.DB, crowdfundingId)
	if err != nil {
		return err
	}
	if entity.ComerId != comerId {
		return router.ErrBadRequest.WithMsg("current comer is not funder of the crowdfunding")
	}
	if entity.Status == crowdfunding.Ended || entity.RaiseBalance.Equal(entity.RaiseGoal) {
		return createTransactionAndQueryContract(mysql.DB, entity.ChainId, entity.ID, txHash, serviceTransaction.CrowdfundingRemoved, func(address string, onChainStatus int) error {
			// set crowdfunding status to `Ended`
			if onChainStatus == 1 {
				return crowdfunding.UpdateCrowdfundingStatus(mysql.DB, crowdfundingId, crowdfunding.Ended)
			}
			return nil
		})
	}
	return router.ErrBadRequest.WithMsg("current crowdfunding can not be removed")
}

func createTransactionAndQueryContract(tx *gorm.DB, chainId, crowdfundingId uint64, txHash string, sourceType int, callback func(address string, onChainStatus int) error) error {
	if err := modelTransaction.CreateTransaction(tx, &modelTransaction.Transaction{
		ChainID:    chainId,
		TxHash:     txHash,
		TimeStamp:  time.Now(),
		Status:     int(crowdfunding.Pending),
		SourceType: sourceType,
		RetryTimes: 0,
		SourceID:   int64(crowdfundingId),
	}); err != nil {
		return err
	}
	// query contract
	address, status := serviceTransaction.GetContractAddress(chainId, txHash)
	log.Infof("#### TX_HASH:%s --> CONTRACT_ADDRESS_AND_ON_CHAIN_STATUS_OF_CREATED_CROWDFUNDING:: %s, %d\n", txHash, address, status)
	if status != 2 {
		if err := modelTransaction.UpdateTransactionStatusWithRetry(tx, crowdfundingId, sourceType, status); err != nil {
			return err
		}
	}
	return callback(address, status)
}

func Invest(comerId, crowdfundingId uint64, request crowdfunding.InvestRequest) error {
	entity, err := crowdfunding.GetCrowdfundingById(mysql.DB, crowdfundingId)
	if err != nil {
		return err
	}
	if entity.Status != crowdfunding.Live {
		log.Errorf("Invalid crowdfunding status: %v\n", entity.Status)
		var m string
		if request.Access == crowdfunding.Invest {
			m = "buy"
		} else {
			m = "sell"
		}
		return router.ErrBadRequest.WithMsgf("can not %s, invalid crowdfunding status", m)
	}
	var swap = &crowdfunding.CrowdfundingSwap{
		ChainInfo:       crowdfunding.ChainInfo{ChainId: entity.ChainId, TxHash: request.TxHash},
		Status:          crowdfunding.SwapPending,
		CrowdfundingId:  entity.ID,
		ComerId:         comerId,
		Access:          request.Access,
		Timestamp:       time.Now(),
		BuyTokenSymbol:  request.BuyTokenSymbol,
		BuyTokenAmount:  request.BuyTokenAmount,
		SellTokenSymbol: request.SellTokenSymbol,
		SellTokenAmount: request.SellTokenAmount,
		Price:           request.Price,
	}
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		if err := crowdfunding.CreateCrowdfundingSwap(tx, swap); err != nil {
			return err
		}
		sourceType := serviceTransaction.CrowdfundingBought
		if request.Access == crowdfunding.Withdraw {
			sourceType = serviceTransaction.CrowdfundingSold
		}
		return createTransactionAndQueryContract(tx, swap.ChainId, swap.ID, request.TxHash, sourceType, func(address string, onChainStatus int) error {
			swapStatus := crowdfunding.SwapPending
			if onChainStatus == 1 {
				swapStatus = crowdfunding.SwapSuccess
			}
			// update crowdfunding swap onChainStatus
			if err := crowdfunding.UpdateCrowdfundingSwapStatus(tx, swap.ID, swapStatus); err != nil {
				return err
			}
			// insert or update crowdfundingInvestor while onChainsStatus is 1
			if onChainStatus == 1 {
				if err := modelTransaction.UpdateTransactionStatusWithRetry(tx, swap.ID, sourceType, 1); err != nil {
					return err
				}
				investor, err := crowdfunding.FirstOrCreateInvestor(tx, crowdfundingId, comerId)
				if err != nil {
					return err
				}
				investor.Invest(request.Access, request.BuyTokenAmount, request.SellTokenAmount)
				if err := crowdfunding.UpdateCrowdfundingInvestor(tx, investor); err != nil {
					return err
				}
				if err := crowdfunding.UpdateCrowdfundingRaiseBalance(tx, *swap); err != nil {
					return err
				}
			}
			return nil
		})
	})
}

func HandleOnChainStateForInvestment(onChainStatus int, swap crowdfunding.CrowdfundingSwap, tran modelTransaction.GetTransactions) error {
	sourceID := swap.ID
	txHash := swap.TxHash
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		// insert or update crowdfundingInvestor while onChainsStatus is 1
		if onChainStatus == 1 {
			log.Infof("##### Process on-chain-successful-investment: %s, %d, investAccess %s\n", txHash, sourceID, swap.Access.String())
			// update crowdfunding swap onChainStatus:
			if err := crowdfunding.UpdateCrowdfundingSwapStatus(tx, sourceID, crowdfunding.SwapSuccess); err != nil {
				log.Errorf("##### UpdateCrowdfundingSwapStatus: %s, %d, %v\n", txHash, sourceID, err)
				return err
			}
			investor, err := crowdfunding.FirstOrCreateInvestor(tx, swap.CrowdfundingId, swap.ComerId)
			if err != nil {
				log.Errorf("##### FirstOrCreateInvestor: %s, %d, %v\n", txHash, sourceID, err)
				return err
			}

			(&investor).Invest(swap.Access, swap.BuyTokenAmount, swap.SellTokenAmount)
			log.Infof("##### Update investor --- %v \n", investor)
			if err := crowdfunding.UpdateCrowdfundingInvestor(tx, investor); err != nil {
				log.Errorf("##### UpdateCrowdfundingInvestor: %s, %d, %v\n", txHash, sourceID, err)
				return err
			}

			if err := crowdfunding.UpdateCrowdfundingRaiseBalance(tx, swap); err != nil {
				log.Errorf("##### UpdateCrowdfundingRaiseBalance: %s, %d, %v\n", txHash, sourceID, err)
				return err
			}

			if err := modelTransaction.UpdateTransactionStatusById(tx, tran.TransactionId, 1); err != nil {
				return err
			}
			return nil
		} else if onChainStatus == 2 && tran.RetryTimes >= 3 {
			return crowdfunding.UpdateCrowdfundingSwapStatus(mysql.DB, sourceID, crowdfunding.SwapFailure)
		} else {
			log.Infof("##### [SwapId:%d, TxHash: %s]Retry .....", swap.ID, swap.TxHash)
		}
		return nil
	})
}

func ModifyCrowdfunding(comerId, crowdfundingId uint64, request crowdfunding.ModifyRequest) error {
	entity, err := crowdfunding.GetCrowdfundingById(mysql.DB, crowdfundingId)
	if err != nil {
		return err
	}
	if entity.ComerId != comerId {
		return router.ErrBadRequest.WithMsg("current comer is not funder of the crowdfunding")
	}
	if entity.Status != crowdfunding.Live && entity.Status != crowdfunding.Upcoming {
		return router.ErrBadRequest.WithMsg("the crowdfunding can not be modified")
	}
	return mysql.DB.Transaction(func(tx *gorm.DB) error {
		history := crowdfunding.IboRateHistory{
			CrowdfundingId: crowdfundingId,
			EndTime:        request.EndTime,
			//BuyTokenSymbol:  request.BuyTokenSymbol,
			MaxBuyAmount:   request.MaxBuyAmount,
			MaxSellPercent: request.MaxSellPercent,
			//SellTokenSymbol: request.SellTokenSymbol,
			BuyPrice:    request.BuyPrice,
			SwapPercent: request.SwapPercent,
		}
		if err := crowdfunding.CreateIboRateHistory(tx, &history); err != nil {
			return err
		}
		sourceType := serviceTransaction.CrowdfundingModified
		return createTransactionAndQueryContract(tx, entity.ChainId, history.ID, request.TxHash, sourceType, func(address string, onChainStatus int) error {
			if onChainStatus == 1 {
				log.Infof("#### Crowdfunding modified successfully, then update crowdfunding entity, %d\n", crowdfundingId)
				if err := crowdfunding.UpdateCrowdfunding(tx, crowdfundingId, request); err != nil {
					return err
				}
			}
			// do nothing
			return nil
		})
	})
}

func GetBuyPriceAndSwapModificationHistories(comerId, crowdfundingId uint64, pagination *model.Pagination) (err error) {
	entity, err := crowdfunding.GetCrowdfundingById(mysql.DB, crowdfundingId)
	if err != nil {
		return
	}
	if entity.ComerId != comerId {
		err = router.ErrBadRequest.WithMsg("current comer is not funder of the crowdfunding")
		return
	}

	return crowdfunding.QueryModificationHistories(mysql.DB, crowdfundingId, pagination)
}

func GetCrowdfundingSwapRecords(crowdfundingId uint64, pagination *model.Pagination) (err error) {
	var records []crowdfunding.InvestmentRecord
	err = crowdfunding.QuerySwapListByCrowdfundingId(mysql.DB, crowdfundingId, pagination)
	if err != nil {
		return
	}

	for _, swap := range pagination.Rows.([]crowdfunding.CrowdfundingSwap) {
		var profile account.ComerProfile
		comerId := swap.ComerId
		if err = account.GetComerProfile(mysql.DB, comerId, &profile); err != nil {
			return
		}
		records = append(records, crowdfunding.InvestmentRecord{
			ComerName:      profile.Name,
			ComerAvatar:    profile.Avatar,
			ComerId:        comerId,
			CrowdfundingId: swap.CrowdfundingId,
			Access:         swap.Access,
			InvestAmount:   swap.BuyTokenAmount,
			Time:           swap.Timestamp,
		})
	}
	pagination.Rows = records
	return
}

func GetInvestorDetail(crowdfundingId, comerId uint64) (crowdfunding.Investor, error) {
	return crowdfunding.SelectInvestorByCrowdfundingIdAndComerId(mysql.DB, crowdfundingId, comerId)
}

func GetCrowdfundingListByStartup(startupId uint64) (list []crowdfunding.PublicItem, err error) {
	crowdfundingList, err := crowdfunding.SelectCrowdfundingListByStartupId(mysql.DB, startupId)
	if err != nil {
		return nil, err
	}
	if len(crowdfundingList) > 0 {
		var st startup.Startup
		if err = startup.GetStartup(mysql.DB, startupId, &st); err != nil {
			return
		}
		for _, c := range crowdfundingList {
			list = append(list, crowdfunding.PublicItem{
				CrowdfundingId:       c.ID,
				CrowdfundingContract: c.CrowdfundingContract,
				StartupId:            c.StartupId,
				ComerId:              c.ComerId,
				StartupName:          st.Name,
				StartupLogo:          st.Logo,
				StartTime:            c.StartTime,
				EndTime:              c.EndTime,
				RaiseBalance:         c.RaiseBalance,
				RaiseGoal:            c.RaiseGoal,
				RaisedPercent:        c.RaiseBalance.Div(c.RaiseGoal),
				BuyPrice:             c.BuyPrice,
				SwapPercent:          c.SwapPercent,
				Poster:               c.Poster,
				Status:               c.Status,
				ChainId:              c.ChainId,
				KYC:                  st.KYC,
				ContractAudit:        st.ContractAudit,
				BuyTokenContract:     c.BuyTokenContract,
				BuyTokenSymbol:       c.BuyTokenSymbol,
				SellTokenContract:    c.SellTokenContract,
				SellTokenSymbol:      c.SellTokenSymbol,
			})
		}
	}
	return
}
