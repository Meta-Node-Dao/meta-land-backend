package metrics

import "github.com/gotomicro/ego/server/egovernor"

// Vernor ego vernor component
var Vernor *egovernor.Component

// Init the vernor
func Init() (err error) {
	Vernor = egovernor.Load("server.governor").Build()

	return
}
