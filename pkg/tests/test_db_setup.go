package pkg

import (
	"ceres/pkg/initialization/config"
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/initialization/redis"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/gotomicro/ego/core/econf"
	"github.com/qiniu/x/log"
)

func setup() (err error) {
	file, err := os.Open("./config.dev.toml")
	if err != nil {
		return err
	}
	err = econf.LoadFromReader(file, toml.Unmarshal)
	if err != nil {
		return err
	}
	err = config.Init()
	if err != nil {
		return err
	}
	err = mysql.Init()
	if err != nil {
		return err
	}
	err = redis.Init()
	if err != nil {
		return err
	}

	// err = eth.Init()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func doTest(t func()) {
	if err := setup(); err != nil {
		log.Fatal("test terminated...", err)
		return
	}
	t()
}

func TestInitApp(t *testing.T) {
	doTest(func() {
		log.Info("init app ok...")
	})
}
