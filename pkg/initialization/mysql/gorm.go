package mysql

import (
	"ceres/pkg/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB Mysql instance with gorm
var DB *gorm.DB

// Init the mysql
func Init() (err error) {
	DB, err = gorm.Open(mysql.Open(config.Mysql.Dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Silent,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		return err
	}
	if config.Mysql.Debug {
		DB = DB.Debug()
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(config.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Mysql.ConnMaxLifetime))
	return err
}
