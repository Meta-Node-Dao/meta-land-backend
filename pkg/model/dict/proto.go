package dict

import (
	"ceres/pkg/model"
)

type DictStatus int

const (
	Enabled DictStatus = iota + 1
	Disabled
)

// DictData 字典数据表结构
type DictData struct {
	model.Base
	StartupID uint64 `gorm:"column:startup_id" json:"startup_id"`              // 关联的初创公司ID
	DictType  string `gorm:"column:dict_type;index:idx_type" json:"dict_type"` // 字典类型
	DictLabel string `gorm:"column:dict_label" json:"dict_label"`              // 字典标签
	DictValue string `gorm:"column:dict_value" json:"dict_value"`              // 字典键值
	SeqNum    int    `gorm:"column:seq_num" json:"seq_num"`                    // 显示顺序
	Status    int8   `gorm:"column:status" json:"status"`                      // 状态(1:启用 2:停用)
	Remark    string `gorm:"column:remark" json:"remark"`                      // 备注

}

// TableName 指定表名
func (DictData) TableName() string {
	return "dict_data"
}
