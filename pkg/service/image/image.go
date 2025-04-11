package image

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/image"

	"github.com/qiniu/x/log"
)

func GetImageList(request model.ListRequest, response *model.ListResponse) (err error) {
	tagList := make([]model.Image, 0)
	total, err := model.GetImageList(mysql.DB, request, &tagList)
	if err != nil {
		log.Warn(err)
		return
	}
	response.Total = total
	response.List = tagList
	return
}
