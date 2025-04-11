package upload

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/s3"
	"ceres/pkg/initialization/utility"
	"ceres/pkg/router"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/qiniu/x/log"
)

func Upload(ctx *router.Context) {
	file, fileInfo, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg("Invalidate file")
		ctx.HandleError(err)
		return
	}

	if fileInfo.Size > config.Aws.MaxSize {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg("The maximum size of upload file is 1MB")
		ctx.HandleError(err)
		return
	}

	fileName := ctx.Param("name")
	fileNames := strings.Split(fileInfo.Filename, ".")
	if len(fileNames) != 2 {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg("Invalidate file name")
		ctx.HandleError(err)
		return
	}
	fileName = fmt.Sprintf("%v.%v", utility.Sequence.Next(), fileNames[1])

	result, err := s3.S3Uploader.Upload(&s3manager.UploadInput{
		Bucket: &config.Aws.Bucket,
		Key:    &fileName,
		ACL:    aws.String("public-read"),
		Body:   file,
	})
	if err != nil {
		log.Warn(err)
		ctx.HandleError(err)
		return
	}

	ctx.OK(struct {
		Url string
	}{Url: result.Location})
}
