package dict

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/dict"
	"errors"
	"strings"
)

func SelectDictDataByType(dictType string) ([]dict.DictData, error) {
	if strings.TrimSpace(dictType) == "" {
		return nil, errors.New("dict type can not be empty")
	}
	return dict.SelectByDictType(mysql.DB, dictType)
}
