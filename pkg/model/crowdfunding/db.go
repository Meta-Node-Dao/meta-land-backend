package crowdfunding

import (
	"ceres/pkg/model"
	"ceres/pkg/model/startup"
	"fmt"
	"math"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateCrowdfunding(db *gorm.DB, c *Crowdfunding) error {
	return db.Create(c).Error
}

func CreateCrowdfundingSwap(db *gorm.DB, cs *CrowdfundingSwap) error {
	return db.Create(cs).Error
}
func GetCrowdfundingSwapById(db *gorm.DB, swapId uint64) (swap CrowdfundingSwap, err error) {
	err = db.Model(&CrowdfundingSwap{}).Where("id = ?", swapId).Find(&swap).Error
	return
}
func UpdateCrowdfundingSwapStatus(db *gorm.DB, id uint64, status CrowdfundingSwapStatus) error {
	return db.Model(&CrowdfundingSwap{}).Where("id = ?", id).Updates(map[string]interface{}{"status": status}).Error
}
func SelectOnGoingByStartupId(db *gorm.DB, startupId uint64) (crowdfundingList []Crowdfunding, err error) {
	if err = db.Model(&Crowdfunding{}).Where("startup_id = ? and is_deleted = 0 and status in (0, 1, 2)", startupId).Find(&crowdfundingList).Error; err != nil {
		return
	}
	return crowdfundingList, nil
}
func SelectStartupsWithNonCrowdfundingOnGoing(db *gorm.DB, comerId uint64) (startups []StartupCrowdfundingInfo, err error) {
	var sts []startup.Startup
	var onGoingByComer []uint64
	err = db.Model(&Crowdfunding{}).Select("startup_id").Where("comer_id=? and is_deleted=0 and status in (0, 1, 2)", comerId).Find(&onGoingByComer).Error
	if err != nil {
		return
	}
	err = db.Model(&startup.Startup{}).Where("comer_id = ? and is_deleted = 0 ", comerId).Find(&sts).Error

	if len(sts) > 0 {
		for _, st := range sts {
			can := true
			if len(onGoingByComer) > 0 {
				for _, u := range onGoingByComer {
					if u == st.ID {
						can = false
					}
				}
			}
			startups = append(startups, StartupCrowdfundingInfo{st.ID, st.Name, can, st.OnChain, st.TokenContractAddress})
		}
	}
	return
}

func UpdateCrowdfundingContractAddressAndStatus(db *gorm.DB, fundingID uint64, address string, status CrowdfundingStatus) error {
	return db.Model(&Crowdfunding{}).Where("id = ?", fundingID).Updates(map[string]interface{}{"crowdfunding_contract": address, "status": status}).Error
}

func SelectCrowdfundingList(db *gorm.DB, pagination *PublicCrowdfundingListPageRequest) (err error) {
	var list []Crowdfunding
	tx := db.Where("crowdfunding.is_deleted = 0")

	if strings.TrimSpace(pagination.Keyword) != "" {
		tx = tx.Joins("inner join startup on startup.id = crowdfunding.startup_id and startup.name like ? and startup.is_deleted = false", "%"+strings.TrimSpace(pagination.Keyword)+"%")
	}
	if pagination.Mode == 0 {
		tx = tx.Where("crowdfunding.status in ?", []CrowdfundingStatus{Upcoming, Live, Ended, Cancelled})
	} else {
		tx = tx.Where("crowdfunding.status = ?", pagination.Mode)
	}
	var totalRows int64
	if err := tx.Table("crowdfunding").Count(&totalRows).Error; err != nil {
		return err
	}
	if err := tx.Table("crowdfunding").Order(fmt.Sprintf("crowdfunding.%s", pagination.GetSort())).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Scan(&list).Error; err != nil {
		return err
	}
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	pagination.Rows = list
	return
}

func GetCrowdfundingById(db *gorm.DB, crowdfundingId uint64) (entity Crowdfunding, err error) {
	if err = db.Model(&Crowdfunding{}).Where("is_deleted=0 and id = ?", crowdfundingId).Find(&entity).Error; err != nil {
		return
	}
	return
}

func SelectCrowdfundingListByFounder(db *gorm.DB, comerId uint64, pagination *model.Pagination) (err error) {
	var list []Crowdfunding
	err = db.Scopes(model.Paginate(&Crowdfunding{}, pagination, db)).Where("is_deleted=0 and comer_id = ?", comerId).Find(&list).Error
	if err != nil {
		return
	}
	pagination.Rows = list
	return
}

func SelectCrowdfundingListByInvestor(db *gorm.DB, comerId uint64, pagination *model.Pagination) (err error) {
	var list []Crowdfunding
	var cnt int64
	cntSql := fmt.Sprintf("select count(c.id) from crowdfunding c left join crowdfunding_investor ci on c.id=ci.crowdfunding_id where c.is_deleted=0 and ci.comer_id=%d", comerId)
	err = db.Raw(cntSql).Scan(&cnt).Error
	if err != nil {
		return
	}
	pageSql := fmt.Sprintf("select c.* from crowdfunding c left join crowdfunding_investor ci on c.id=ci.crowdfunding_id where c.is_deleted=0 and ci.comer_id=%d order by c.created_at desc limit %d,%d", comerId, pagination.GetOffset(), pagination.GetLimit())
	err = db.Raw(pageSql).Scan(&list).Error
	if err != nil {
		return
	}
	pagination.Rows = list
	pagination.TotalRows = cnt
	totalPages := int(math.Ceil(float64(cnt) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return
}

func UpdateCrowdfundingStatus(db *gorm.DB, crowdfundingId uint64, status CrowdfundingStatus) error {
	return db.Model(&Crowdfunding{}).Where("is_deleted=0 and id = ?", crowdfundingId).Update("status", status).Error
}

func QueryModificationHistories(db *gorm.DB, crowdfundingId uint64, pagination *model.Pagination) (err error) {
	var list []IboRateHistory
	err = db.Scopes(model.Paginate(&IboRateHistory{}, pagination, db)).Where("crowdfunding_id = ?", crowdfundingId).Order("created_at desc").Find(&list).Error
	if err != nil {
		return
	}
	pagination.Rows = list
	return
}

func UpdateCrowdfunding(db *gorm.DB, crowdfundingId uint64, request ModifyRequest) error {
	return db.Model(&Crowdfunding{}).Where("is_deleted=0 and id = ?", crowdfundingId).Updates(Crowdfunding{
		SellInfo:    SellInfo{MaxSellPercent: request.MaxSellPercent},
		BuyInfo:     BuyInfo{BuyPrice: request.BuyPrice, MaxBuyAmount: request.MaxBuyAmount},
		SwapPercent: request.SwapPercent,
		EndTime:     request.EndTime,
	}).Error
}

func FirstOrCreateInvestor(tx *gorm.DB, crowdfundingId, comerId uint64) (investor Investor, err error) {
	err = tx.Debug().Where(&Investor{CrowdfundingId: crowdfundingId, ComerId: comerId}).FirstOrCreate(&investor).Error
	return
}

func UpdateCrowdfundingInvestor(db *gorm.DB, investor Investor) error {
	return db.Model(&Investor{}).Where("id = ?", investor.ID).Updates(investor).Error
}

func UpdateCrowdfundingRaiseBalance(tx *gorm.DB, swap CrowdfundingSwap) error {
	var p clause.Expr
	mp := map[string]interface{}{}
	if swap.Access == Invest {
		p = gorm.Expr("raise_balance + ?", swap.BuyTokenAmount)
		// funding, err := GetCrowdfundingById(tx, swap.CrowdfundingId)
		//if err != nil {
		//	return err
		//}
		//if funding.RaiseBalance.Add(swap.BuyTokenAmount).GreaterThanOrEqual(funding.RaiseGoal) {
		//	mp["status"] = Ended
		//}
		mp["raise_balance"] = p
		return tx.Model(&Crowdfunding{}).Where("is_deleted = false and id = ?", swap.CrowdfundingId).Updates(mp).Error
	}
	p = gorm.Expr("raise_balance - ?", swap.BuyTokenAmount)
	mp["raise_balance"] = p
	return tx.Model(&Crowdfunding{}).Where("is_deleted = false and id = ?", swap.CrowdfundingId).Updates(mp).Error
}

func QuerySwapListByCrowdfundingId(db *gorm.DB, crowdfundingId uint64, pagination *model.Pagination) (err error) {
	var swaps []CrowdfundingSwap
	err = db.Scopes(model.Paginate(&CrowdfundingSwap{}, pagination, db)).Where("crowdfunding_id = ? and status = ?", crowdfundingId, SwapSuccess).Find(&swaps).Error
	if err != nil {
		return err
	}
	pagination.Rows = swaps
	return
}

func SelectInvestorByCrowdfundingIdAndComerId(db *gorm.DB, crowdfundingId, comerId uint64) (investor Investor, err error) {
	err = db.Model(&Investor{}).Where("crowdfunding_id = ? and comer_id = ?", crowdfundingId, comerId).Find(&investor).Error
	return
}

func SelectCrowdfundingListByStartupId(db *gorm.DB, startupId uint64) (list []Crowdfunding, err error) {
	err = db.Model(&Crowdfunding{}).
		Where("startup_id = ? and is_deleted = false and status in ?", startupId, []CrowdfundingStatus{Upcoming, Live, Ended, Cancelled}).
		Find(&list).Error
	return
}

func CreateIboRateHistory(db *gorm.DB, history *IboRateHistory) error {
	return db.Model(&IboRateHistory{}).Create(&history).Error
}

func GetIboRateHistoryById(db *gorm.DB, historyId uint64) (history IboRateHistory, err error) {
	err = db.Model(&IboRateHistory{}).Where("id = ?", historyId).Find(&history).Error
	return
}

func CountCrowdfundingPostedByComer(db *gorm.DB, targetComerID uint64) (cnt int64, er error) {
	er = db.Model(&Crowdfunding{}).Where("is_deleted = false and comer_id = ?", targetComerID).Count(&cnt).Error
	return
}

func SelectToBeStartedCrowdfundingListWithin1Min(db *gorm.DB) (list []*Crowdfunding, err error) {
	now := time.Now()
	err = db.Model(&Crowdfunding{}).
		Where("is_deleted = false and status =  ? and (start_time between ? and ? or start_time < ? and end_time > ?)",
			Upcoming,
			now.Add(-30*time.Second),
			now.Add(30*time.Second),
			now,
			now).
		Find(&list).Error
	return
}

func SelectToBeEndedCrowdfundingList(db *gorm.DB) (list []*Crowdfunding, err error) {
	err = db.Model(&Crowdfunding{}).
		Where("is_deleted = false and status not in  ? and end_time <= ?",
			[]CrowdfundingStatus{Ended, Cancelled, OnChainFailure},
			time.Now().Add(time.Second*30)).
		Find(&list).Error
	return
}
