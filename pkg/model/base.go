package model

import (
	"ceres/pkg/initialization/utility"
	"ceres/pkg/router"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type PageData struct {
	List  []interface{} `json:"list"`
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Total int           `json:"total"`
}

type IsExistResponse struct {
	IsExist bool `json:"is_exist"`
}

// Base contains common columns for all tables.
type Base struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	IsDeleted bool      `gorm:"column:is_deleted;default:false" json:"is_deleted"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = utility.Sequence.Next()
	return
}

// RelationBase contains common columns for all tables.
type RelationBase struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (base *RelationBase) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = utility.Sequence.Next()
	return
}

// ListRequest list request
type ListRequest struct {
	Limit     int  `form:"limit" binding:"gt=0"`
	Offset    int  `form:"offset" binding:"gte=0"`
	IsDeleted bool `form:"isDeleted"`
}

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	Mode       int         `json:"mode,omitempty" query:"mode"`
	Keyword    string      `json:"keyword,omitempty"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" || strings.TrimSpace(p.Sort) == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
func ParsePagination(ctx *router.Context, pagination *Pagination, defaultLimit int) (err error) {
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	limitStr := ctx.Query("limit")
	var limit int
	if limitStr != "" && strings.TrimSpace(limitStr) != "" {
		limit, err = strconv.Atoi(limitStr)
	} else {
		limit = 0
	}

	if err != nil {
		return err
	}

	modeStr := ctx.Query("mode")
	var mode int
	if modeStr != "" && strings.TrimSpace(modeStr) != "" {
		mode, err = strconv.Atoi(modeStr)
	} else {
		mode = 0
	}
	if err != nil {
		return err
	}
	pagination.Mode = mode

	keyword := ctx.Query("keyword")
	pagination.Keyword = keyword

	if page == 0 {
		page = 1
	}
	pagination.Page = page
	_sort := ctx.Query("sort")
	switch _sort {
	case "Created:Recent":
		_sort = "created_at desc"
	case "Created:Oldest":
		_sort = "created_at asc"
	case "Value:Highest":
		_sort = "total_reward_token desc"
	case "Value:Lowest":
		_sort = "total_reward_token asc"
	case "Deposit:Highest":
		_sort = "founder_deposit desc"
	case "Deposit:Lowest":
		_sort = "founder_deposit asc"
	default:
		_sort = "created_at desc"
	}

	pagination.Sort = _sort
	pagination.Limit = limit
	if pagination.Limit == 0 {
		pagination.Limit = defaultLimit
	}
	return nil
}

type BusinessModule int

const (
	ModuleStartup BusinessModule = iota + 1
	ModuleBounty
	ModuleCrowdfunding
	ModuleGovernance
	ModuleOtherDapp
)

var DefaultModules = []BusinessModule{ModuleBounty, ModuleCrowdfunding, ModuleGovernance, ModuleOtherDapp}

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var err error
	// 指定时区
	d.Time, err = time.ParseInLocation(`"2006-01-02"`, string(b), time.Local)
	if err != nil {
		return err
	}
	return nil
}
