package minio

import (
	"ceres/pkg/config"

	"github.com/gotomicro/ego/core/elog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Client minio client
var Client *minio.Client

// Init the minio
func Init() (err error) {

	Client, err = minio.New(
		config.Minio.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(config.Minio.AccessKey, config.Minio.SecretKey, ""),
			Secure: true,
		},
	)

	if err != nil {
		elog.Panicf("Configure the client failed %v", err)
	}

	return
}
