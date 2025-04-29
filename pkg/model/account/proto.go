package account

import (
	"ceres/pkg/model"
	"fmt"
	"strings"
)

type ComerAccountType int

const (
	GithubOauth   ComerAccountType = 1
	GoogleOauth   ComerAccountType = 2
	TwitterOauth  ComerAccountType = 3
	FacebookOauth ComerAccountType = 4
	LikedinOauth  ComerAccountType = 5
)

type SocialType int

const (
	SocialEmail SocialType = iota + 1
	SocialWebsite
	SocialTwitter
	SocialDiscord
	SocialTelegram
	SocialMedium
	SocialFacebook
	SocialLinktree
)

func (receiver SocialType) String() string {
	switch receiver {
	case SocialDiscord:
		return "discord"
	case SocialEmail:
		return "email"
	case SocialFacebook:
		return "facebook"
	case SocialWebsite:
		return "website"
	case SocialTwitter:
		return "twitter"
	case SocialTelegram:
		return "telegram"
	case SocialMedium:
		return "medium"
	case SocialLinktree:
		return "linktree"
	default:
		panic("unsupported socialType")
	}
}

// Comer the comer model of comunion inner account
type Comer struct {
	model.Base
	Address *string `gorm:"column:address;uniqueIndex" json:"address"`
}

func (c Comer) HasAddress() bool {
	add := c.Address
	if add != nil && strings.TrimSpace(*add) != "" {
		return true
	}
	return false
}

func (c Comer) AddressStr() string {
	if c.HasAddress() {
		return *c.Address
	}
	return ""
}

// TableName Comer table name for gorm
func (Comer) TableName() string {
	return "comer"
}

type ComerAccount struct {
	model.Base
	ComerID   uint64           `gorm:"column:comer_id;index" json:"comer_id"`
	Oin       string           `gorm:"column:oin;uniqueIndex;not null" json:"oin"`
	IsPrimary bool             `gorm:"column:is_primary;not null" json:"is_primary"`
	Nick      string           `gorm:"column:nick;not null" json:"nick"`
	Avatar    string           `gorm:"column:avatar;not null" json:"avatar"`
	Type      ComerAccountType `gorm:"column:type;not null" json:"type"`
	IsLinked  bool             `gorm:"column:is_linked;not null" json:"is_linked"`
}

// TableName the ComerAccount table name for gorm
func (ComerAccount) TableName() string {
	return "comer_account"
}

type ComerAccounts []ComerAccount

func (a *ComerAccounts) HasSameOauthType(accounts *ComerAccounts) (has bool) {
	if a != nil && len(*a) > 0 && accounts != nil && len(*accounts) > 0 {
		for _, byAddress := range *a {
			for _, comerAccount := range *accounts {
				if byAddress.Type == comerAccount.Type {
					has = true
					break
				}
			}
		}
	}
	return
}

func (a ComerAccounts) AccountIds() []uint64 {
	var ids []uint64
	for _, comerAccount := range a {
		ids = append(ids, comerAccount.ID)
	}
	return ids
}

// ComerProfile 用户资料表结构
type ComerProfile struct {
	model.Base
	ComerID    uint64 `gorm:"column:comer_id;uniqueIndex" json:"comer_id"`
	Name       string `gorm:"column:name;not null" json:"name"`           // 用户名
	Avatar     string `gorm:"column:avatar;not null" json:"avatar"`       // 头像URL
	Cover      string `gorm:"column:cover" json:"cover"`                  // 封面图URL
	Location   string `gorm:"column:location;default:''" json:"location"` // 所在城市
	TimeZone   string `gorm:"column:time_zone" json:"time_zone"`          // 时区(如UTC-09:30)
	Website    string `gorm:"column:website;default:''" json:"website"`   // 个人网站
	Email      string `gorm:"column:email" json:"email"`                  // 电子邮箱
	Twitter    string `gorm:"column:twitter" json:"twitter"`              // Twitter账号
	Discord    string `gorm:"column:discord" json:"discord"`              // Discord账号
	Telegram   string `gorm:"column:telegram" json:"telegram"`            // Telegram账号
	Medium     string `gorm:"column:medium" json:"medium"`                // Medium账号
	Facebook   string `gorm:"column:facebook" json:"facebook"`            // Facebook账号
	Linktree   string `gorm:"column:linktree" json:"linktree"`            // Linktree链接
	Bio        string `gorm:"column:bio;type:text" json:"bio"`            // 个人简介
	Languages  string `gorm:"column:languages" json:"languages"`          // 使用语言
	Educations string `gorm:"column:educations" json:"educations"`        // 教育经历
}

// TableName the Profile table name for gorm
func (ComerProfile) TableName() string {
	return "comer_profile"
}

type FollowRelation struct {
	model.RelationBase
	ComerID       uint64 `gorm:"comer_id" json:"comerID"`
	TargetComerID uint64 `gorm:"target_comer_id" json:"targetComerID"`
}

// TableName Followed table name for gorm
func (FollowRelation) TableName() string {
	return "comer_follow_rel"
}

type FollowComer struct {
	TargetComerID uint64       `gorm:"target_comer_id" json:"comerID"`
	Comer         Comer        `gorm:"foreignkey:ID;references:TargetComerID" json:"comer"`
	ComerProfile  ComerProfile `gorm:"foreignkey:ComerID;references:TargetComerID" json:"comerProfile"`
}

// TableName FollowComer table name for gorm
func (FollowComer) TableName() string {
	return "comer_follow_rel"
}

type FollowedComer struct {
	ComerID      uint64       `gorm:"comer_id" json:"comerID"`
	Comer        Comer        `gorm:"foreignkey:ID;references:ComerID" json:"comer"`
	ComerProfile ComerProfile `gorm:"foreignkey:ComerID;references:ComerID" json:"comerProfile"`
}

// TableName FollowComer table name for gorm
func (FollowedComer) TableName() string {
	return "comer_follow_rel"
}

type LanguageLevel string

const (
	Beginner     LanguageLevel = "Beginner"
	Elementary   LanguageLevel = "Elementary"
	Intermediate LanguageLevel = "Intermediate"
	Advanced     LanguageLevel = "Advanced"
)

func (receiver LanguageLevel) Check() error {
	switch receiver {
	case Advanced:
		return nil
	case Intermediate:
		return nil
	case Elementary:
		return nil
	case Beginner:
		return nil
	default:
		return fmt.Errorf("unknown language level %v\n", receiver)
	}
}

type LanguageInfo struct {
	Id       string        `json:"id"`
	Language string        `json:"language"`
	Level    LanguageLevel `json:"level"`
}

type EducationInfo struct {
	Id          string `json:"id"`
	School      string `json:"school"`
	Major       string `json:"major"`
	GraduatedAt uint64 `json:"graduatedAt"`
}
