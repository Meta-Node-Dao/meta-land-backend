package redis

import "github.com/gotomicro/ego-component/eredis"

// Client redis client
var Client *eredis.Component

// Init the redis client
func Init() (err error) {
	Client = eredis.Load("ceres.redis").Build()
	return
}
