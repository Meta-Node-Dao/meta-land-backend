package pkg

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/utility"
	"github.com/gotomicro/ego/core/econf"
	"gorm.io/gorm"
	"os"
	"time"
)

import (
	p_mysql "ceres/pkg/initialization/mysql"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
)

func setupMysql() (err error) {

	testMysqlCfg := &config.MysqlConfig{}
	v := econf.New()
	// relative path
	file, err := os.Open("./config.dev.toml")
	if err != nil {
		return err
	}
	if err = v.LoadFromReader(file, toml.Unmarshal); err != nil {
		return err
	}

	if err = v.UnmarshalKey("ceres.mysql", testMysqlCfg); err != nil {
		return err
	}
	db, err := gorm.Open(mysql.Open(testMysqlCfg.Dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		return
	}
	if testMysqlCfg.Debug {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(testMysqlCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(testMysqlCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(testMysqlCfg.ConnMaxLifetime))
	p_mysql.DB = db
	config.Seq = &config.Sequence{Epoch: 1626023857}
	// for db insertion
	err = utility.Init()
	if err != nil {
		return err
	}
	return nil

}

// do not delete this
func DoTestWithMysql(t func()) {
	if err := setupMysql(); err != nil {
		log.Fatal("test terminated...", err)
		return
	}
	t()
}
