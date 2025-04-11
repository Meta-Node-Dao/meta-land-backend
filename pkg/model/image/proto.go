package image

import (
	"ceres/pkg/model"
)

type Category string

// Avatar const image category type
const (
	Avatar Category = "avatar"
)

type Image struct {
	model.Base
	Name     string   `gorm:"column:name" json:"name"`
	Category Category `gorm:"column:category" json:"category"`
}

// TableName identify the table name of this model.
func (Image) TableName() string {
	return "image"
}
