package pkg

import (
	"ceres/pkg/utility/net"
	"ceres/pkg/utility/sequence"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestSeq(t *testing.T) {
	machineIP := net.GetDomianIP()
	machineSignature := strings.Replace(machineIP, ".", "", 4)
	machineID, err := strconv.ParseInt(machineSignature, 10, 64)
	machineID %= 32
	if err != nil {
		return
	}
	flake := sequence.NewSnowflake(1525705533000, uint64(machineID))

	for i := 0; i <= 4097; i++ {
		fmt.Printf("%d\n", flake.Next())
		fmt.Println(flake)
	}
}
