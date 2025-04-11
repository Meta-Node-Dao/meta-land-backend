package pkg

import (
	"ceres/pkg/utility/net"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	ip := net.GetDomianIP()
	t.Log(ip)
}
