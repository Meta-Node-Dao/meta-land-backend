package pkg

import (
	"os"
	"testing"
)

func TestEnvDispatch(t *testing.T) {
	db := os.Getenv("TEST_CERES_WITH_DB")
	if db == "" {
		t.Log("database missed")
	} else {
		t.Log("do mock api test")
	}
}
