package account

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/utility/jwt"
	"fmt"
	"github.com/qiniu/x/log"
	"testing"
)

func TestGetComerProfile(t *testing.T) {
	type args struct {
		comerID  uint64
		response *model.ComerProfileResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetComerProfile(tt.args.comerID, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("GetComerProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXXX(t *testing.T) {
	uin, err := jwt.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21lcl91aW4iOiIxMjQ4NjgxMDg2OTM1MDQiLCJleHAiOjE2NTYwNTM5NDYsImlhdCI6MTY1NTc5NDc0Nn0.JRUY4URvntwg4yZaKdvKcKY2K74AC4K1rOkrUdnT3sY")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(uin)
}
