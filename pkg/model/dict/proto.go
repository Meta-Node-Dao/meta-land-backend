package dict

import "ceres/pkg/model"

type DictStatus int

const (
	Enabled DictStatus = iota + 1
	Disabled
)

type DictDataModel struct {
	model.RelationBase
	DictData
}

type DictData struct {
	StartupId uint64     `gorm:"startup_id" json:"startupId"`
	DictType  string     `gorm:"dict_type" json:"dictType"`
	DictLabel string     `gorm:"dict_label" json:"dictLabel"`
	DictValue string     `gorm:"dict_value" json:"dictValue"`
	SeqNum    int        `gorm:"seq_num" json:"seqNum"`
	Status    DictStatus `gorm:"status" json:"status"`
	Remark    string     `gorm:"remark" json:"remark"`
}

func (receiver DictDataModel) TableName() string {
	return "dict_data"
}
