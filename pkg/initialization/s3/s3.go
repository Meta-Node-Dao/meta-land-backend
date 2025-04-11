package s3

import (
	"ceres/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var S3Uploader *s3manager.Uploader

// Init the mysql
func Init() (err error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.Aws.AccessKey, config.Aws.AccessSecret, ""),
		Endpoint:         aws.String(config.Aws.EndPoint),
		Region:           aws.String(config.Aws.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})
	S3Uploader = s3manager.NewUploader(sess)
	return
}
