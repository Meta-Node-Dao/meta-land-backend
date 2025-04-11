package pkg

import (
	"ceres/pkg/utility/auth"
	"fmt"
	"testing"
)

//
//import (
//	"ceres/pkg/utility/auth"
//	"fmt"
//	"github.com/coreos/etcd/pkg/stringutil"
//	"net/http"
//	"testing"
//)
//
//type MockOauthClient struct {
//	ClientID     string
//	ClientSecret string
//	client       *http.Client
//	Code         string
//}
//type MockUserProfile struct {
//	Login  string `json:"login"`
//	ID     uint64 `json:"id"`
//	Name   string `json:"name"`
//	Avatar string `json:"avatar_url"`
//}
//
//// GetUserID implement the OauthAccount interface
//func (account *MockUserProfile) GetUserID() string {
//	return account.Login
//}
//
//// GetUserAvatar implement the OauthAccount interface
//func (account *MockUserProfile) GetUserAvatar() string {
//	return account.Avatar
//}
//
//// GetUserNick implement the OauthAccount interface
//func (account *MockUserProfile) GetUserNick() string {
//	return account.Name
//}
//
//func (c MockOauthClient) GetUserProfile() (account auth.OauthAccount, err error) {
//	profile := MockUserProfile{
//		Login:  "MOCK_LOGIN" + stringutil.RandomStrings(8, 1)[0],
//		ID:     0,
//		Name:   "MOCK_NAME" + stringutil.RandomStrings(8, 1)[0],
//		Avatar: "MOCK_AVATAR" + stringutil.RandomStrings(8, 1)[0],
//	}
//	return &profile, nil
//}
//
func TestAuthGithub(t *testing.T) {
	github := auth.NewGithubClient("9f113bcc6db1cba82902", "e7fd1ebc44c80e301ea9f7531f8293602e045fa3")
	oauth, err := github.GetUserProfile()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("oauth: %v\n", oauth)
}
