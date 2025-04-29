package account

import (
	"database/sql/driver"
	"encoding/json"
)

// EthLoginRequest the standard result of the web3.js signature
type EthLoginRequest struct {
	Address   string `json:"address" binding:"len=42,startswith=0x"`
	Signature string `json:"signature" binding:"required"`
}

// CreateProfileRequest create a new profile then will let the entity to backend
type CreateProfileRequest struct {
	Name     string   `json:"name" binding:"min=1,max=24"`
	Avatar   string   `json:"avatar"`
	Location string   `json:"location"`
	TimeZone string   `json:"timeZone"`
	SKills   []string `json:"skills" binding:"min=1"`
	ComerSocials
	BIO string `json:"bio" binding:"min=100"`
}

type ComerSocials struct {
	Email    string `json:"email"`
	Website  string `json:"website"`
	Twitter  string `json:"twitter"`
	Discord  string `json:"discord"`
	Telegram string `json:"telegram"`
	Medium   string `json:"medium"`
	Facebook string `json:"facebook"`
	Linktree string `json:"linktree"`
}

// UpdateProfileRequest  update the comer profile
type UpdateProfileRequest struct {
	Name     string   `json:"name" binding:"min=1,max=24"`
	Avatar   string   `json:"avatar"`
	Location string   `json:"location"`
	TimeZone string   `json:"timeZone"`
	SKills   []string `json:"skills" binding:"min=1"`
	ComerSocials
	BIO string `json:"bio" binding:"min=100"`
}

// LinkOauth2WalletRequest link oauth with given wallet
type LinkOauth2WalletRequest struct {
	OauthCode string           `json:"oauthCode" binding:"required"`
	OauthType ComerAccountType `json:"oauthType" binding:"required"`
}

// RegisterWithOauthRequest register with oauth
type RegisterWithOauthRequest struct {
	OauthAccountId uint64               `json:"oauthAccountId" biding:"required"`
	Profile        CreateProfileRequest `json:"profile" biding:"required"`
}

type UpdateBasicInfoRequest struct {
	Name     string `json:"name"`
	Cover    string `json:"cover"`
	Avatar   string `json:"avatar"`
	TimeZone string `json:"timeZone"`
	Location string `json:"location"`
}

type SocialModifyRequest struct {
	SocialType SocialType `json:"socialType"`
	SocialLink string     `json:"socialLink"`
}

type SocialRemoveRequest struct {
	SocialType SocialType `json:"socialType"`
}

type UpdateComerCoverRequest struct {
	Image string `json:"image"`
}

type UpdateSkillsRequest struct {
	Skills []string `json:"skills"`
}

type UpdateBioRequest struct {
	Bio string `json:"bio"`
}

type LanguageInfos []LanguageInfo

// UpdateLanguageInfosRequest create/update/delete
type UpdateLanguageInfosRequest struct {
	Languages LanguageInfos `json:"languages"`
}

func (c LanguageInfos) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *LanguageInfos) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type EducationInfos []EducationInfo
type UpdateEducationsRequest struct {
	Educations EducationInfos `json:"educations"`
}

func (c EducationInfos) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *EducationInfos) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
