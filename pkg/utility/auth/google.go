package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

var googleEndpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
}

var scopes = []string{"https://www.googleapis.com/auth/userinfo.profile",
	"https://www.googleapis.com/auth/userinfo.email"}

const googleOauthStateString = "random"

// Google REST client
// implemnetes the OauthClient interface
type Google struct {
	oauth2.Config
	OauthState   string
	CurrentState string
	Code         string
	client       *http.Client
}

type googleAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint64 `json:"expires_in"`
}

// GetAccessToken
// use the requestToken to get the access token which will be used to get the github user information
func (google *Google) getAccessToken() (accessToken string, err error) {
	token, err := google.Exchange(oauth2.NoContext, google.Code)
	if err != nil {
		err = fmt.Errorf("Code exchange failed with '%s'", err)
		return
	}
	return token.AccessToken, nil
}

type googleInspectBody struct {
	UserID string `json:"user_id"`
}

type googleInspectResponse struct {
	Data facebookInspectBody `json:"data"`
}

func (response *googleInspectResponse) GetUserID() string {
	return response.Data.UserID
}

// GoogleOauthAccount  Google Oauth account profile
type GoogleOauthAccount struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// GetUserID implement the OauthAccount interface
func (account *GoogleOauthAccount) GetUserID() string {
	return account.ID
}

// GetUserAvatar implement the OauthAccount interface
func (account *GoogleOauthAccount) GetUserAvatar() string {
	return account.Picture
}

// GetUserNick implement the OauthAccount interface
func (account *GoogleOauthAccount) GetUserNick() string {
	return account.Name
}

// GetUserProfile  Facebook Oauth get user profile logic
func (google *Google) GetUserProfile() (account OauthAccount, err error) {
	accessToken, err := google.getAccessToken()
	if err != nil {
		return
	}
	// firstly to intspect the access token to get the facebook user ID
	request, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", accessToken),
		nil,
	)

	response, err := google.client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	account = &GoogleOauthAccount{}
	err = json.Unmarshal(body, account)
	if err != nil {
		return nil, err
	}
	return
}
