package my

import (
	"ceres/pkg/config"
	"ceres/pkg/utility/jwt"
	"testing"
)

func Test_GenerateComunionAuthorization(t *testing.T) {
	config.JWT = &config.JWTConfig{
		Expired: 259200,
		Secret:  "Comunion-Ceres",
	}
	token := jwt.Sign(124257602580480)
	t.Logf("%s", token)
}
