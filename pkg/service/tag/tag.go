package tag

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/tag"

	"github.com/qiniu/x/log"
)

// GetStartupTagList return the all startup tags in list
func GetStartupTagList(request model.ListRequest, response *model.ListResponse) (err error) {
	tagList := make([]model.Tag, 0)
	total, err := model.GetTagList(mysql.DB, request, &tagList)
	if err != nil {
		log.Warn(err)
		return
	}
	response.Total = total
	response.List = tagList
	return
}
