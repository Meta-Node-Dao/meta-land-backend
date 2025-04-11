package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gotomicro/ego/core/elog"
)

// Facebook REST client
// implemnetes the OauthClient interface
// https://developers.facebook.com/docs/facebook-login/advanced
type Facebook struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	RequestToken string
	client       *http.Client
}

type facebookAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint64 `json:"expires_in"`
}

// GetAccessToken
// use the requestToken to get the access token which will be used to get the github user information
func (facebook *Facebook) getAccessToken() (accessToken string, err error) {
	u := fmt.Sprintf(
		"https://graph.facebook.com/oauth/access_token?client_id=%s&redirect_uri=%s&client_secret=%s&code=%s",
		facebook.ClientID,
		facebook.RedirectURI,
		facebook.ClientSecret,
		facebook.RequestToken,
	)
	request, _ := http.NewRequest(
		"GET",
		u,
		nil,
	)
	response, err := facebook.client.Do(request)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	r := facebookAccessTokenResponse{}
	err = response.Body.Close()
	if err != nil {
		elog.Warnf("close the response body failed %v", err)
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}
	accessToken = r.AccessToken

	return
}

type facebookInspectBody struct {
	UserID string `json:"user_id"`
}

type facebookInspectResposne struct {
	Data facebookInspectBody `json:"data"`
}

func (response *facebookInspectResposne) GetUserID() string {
	return response.Data.UserID
}

// FacebookOauthAccount  Facebook Oauth account profile
type FacebookOauthAccount struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// GetUserID implement the OauthAccount interface
func (account *FacebookOauthAccount) GetUserID() string {
	return account.ID
}

// GetUserAvatar implement the OauthAccount interface
func (account *FacebookOauthAccount) GetUserAvatar() string {
	return account.Picture
}

// GetUserNick implement the OauthAccount interface
func (account *FacebookOauthAccount) GetUserNick() string {
	return account.Name
}

// GetUserProfile  Facebook Oauth get user profile logic
func (facebook *Facebook) GetUserProfile() (account OauthAccount, err error) {
	accessToken, err := facebook.getAccessToken()
	if err != nil {
		return
	}
	// firstly to intspect the access token to get the facebook user ID
	request, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("https://graph.facebook.com/debug_token?input_token={token-to-inspect}&access_token=%s", accessToken),
		nil,
	)
	response, err := facebook.client.Do(request)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	r := facebookInspectResposne{}
	err = response.Body.Close()
	if err != nil {
		elog.Warnf("close the body failed %v", err)
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}

	// get user avatar result have to try 3 times to confirm the avatar get it
	request, _ = http.NewRequest(
		"GET",
		fmt.Sprintf("https://graph.facebook.com/v10.0/%s?fields=id,name,picture&access_token=%s", r.GetUserID(), accessToken),
		nil,
	)
	for i := 0; i < 3; i++ {
		response, err = facebook.client.Do(request)
		if err != nil {
			return
		}
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		account = &FacebookOauthAccount{}
		err = json.Unmarshal(body, account)
		if err != nil {
			return
		}
	}
	err = errors.New("try to get the user profile too many times")
	return
}
