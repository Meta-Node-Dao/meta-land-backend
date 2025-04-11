package tag

import (
	"ceres/pkg/model"
)

type Category string

const (
	ComerSkill Category = "comerSkill"
	Startup    Category = "startup"
	Bounty     Category = "bounty"
)

// Tag  Comunion tag for startup bounty profile and other position need Tag.
type Tag struct {
	model.Base
	Name     string   `gorm:"column:name" json:"name"`
	Category Category `gorm:"column:category" json:"category"`
	IsIndex  bool     `gorm:"column:is_index" json:"isIndex"`
}

// TableName identify the table name of this model.
func (Tag) TableName() string {
	return "tag"
}

// TagTargetRel  Comunion tag for startup bounty profile and other position need TagTargetRel.
type TagTargetRel struct {
	model.RelationBase
	TargetID uint64   `column:"target_id"`
	Target   Category `column:"target"`
	TagID    uint64   `column:"tag_id"`
}

// TableName identify the table name of this model.
func (TagTargetRel) TableName() string {
	return "tag_target_rel"
}
