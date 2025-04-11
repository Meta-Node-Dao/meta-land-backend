package account

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
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
	Address *string `gorm:"column:address" json:"address"`
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

// ComerAccount the account model of comer
type ComerAccount struct {
	model.Base
	ComerID   uint64           `gorm:"column:comer_id" json:"comerID"`
	OIN       string           `gorm:"column:oin" json:"oin"`
	IsPrimary bool             `gorm:"column:is_primary" json:"isPrimary"`
	Nick      string           `gorm:"column:nick" json:"nick"`
	Avatar    string           `gorm:"column:avatar" json:"avatar"`
	Type      ComerAccountType `gorm:"column:type" json:"type"`
	IsLinked  bool             `gorm:"column:is_linked" json:"isLinked"`
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

// ComerProfile model
type ComerProfile struct {
	model.Base
	ComerID    uint64         `gorm:"column:comer_id" json:"comerID"`
	Name       string         `gorm:"column:name" json:"name"`
	Avatar     string         `gorm:"column:avatar" json:"avatar"`
	Cover      string         `gorm:"column:cover" json:"cover"`
	Location   string         `gorm:"column:location" json:"location"`
	TimeZone   string         `gorm:"column:time_zone" json:"timeZone"`
	Website    string         `gorm:"column:website" json:"website"`
	Email      string         `gorm:"column:email" json:"email"`
	Twitter    string         `gorm:"column:twitter" json:"twitter"`
	Discord    string         `gorm:"column:discord" json:"discord"`
	Telegram   string         `gorm:"column:telegram" json:"telegram"`
	Medium     string         `gorm:"column:medium" json:"medium"`
	Facebook   string         `gorm:"column:facebook" json:"facebook"`
	Linktree   string         `gorm:"column:linktree" json:"linktree"`
	BIO        string         `gorm:"column:bio" json:"bio"`
	Skills     []tag.Tag      `gorm:"many2many:tag_target_rel;foreignKey:ComerID;joinForeignKey:TargetID;" json:"skills"`
	Languages  LanguageInfos  `gorm:"column:languages;type:json" json:"languages"`
	Educations EducationInfos `gorm:"column:educations;type:json" json:"educations"`
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
