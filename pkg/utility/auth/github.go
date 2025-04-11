package auth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gotomicro/ego/core/elog"
)

// Github REST client
// implemnetes the OauthClient interface
type Github struct {
	ClientID     string
	ClientSecret string
	client       *http.Client
	Code         string
}

// NewGithubClient  create a new github client
func NewGithubClient(clientID, clientSecret string) (github *Github) {
	github = &Github{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
	return
}

type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GetAccessToken  use the requestToken to get the access token which will be used to get the github user information
func (github *Github) GetAccessToken(code string) (accessToken string, err error) {
	github.Code = code
	u, _ := url.Parse(
		fmt.Sprintf(
			"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
			github.ClientID,
			github.ClientSecret,
			code,
		),
	)
	request := &http.Request{
		Method: "GET",
		URL:    u,
		Header: map[string][]string{
			"Accept": {"application/json"},
		},
	}
	resp, err := github.client.Do(request)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	r := githubAccessTokenResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return
	}
	accessToken = r.AccessToken
	return
}

// GithubUserProfile  response of api.github.com/user
// this struct implement the ceres/session/OauthAccount interface
type GithubUserProfile struct {
	Login  string `json:"login"`
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
}

// GetUserID implement the OauthAccount interface
func (account *GithubUserProfile) GetUserID() string {
	return account.Login
}

// GetUserAvatar implement the OauthAccount interface
func (account *GithubUserProfile) GetUserAvatar() string {
	return account.Avatar
}

// GetUserNick implement the OauthAccount interface
func (account *GithubUserProfile) GetUserNick() string {
	return account.Name
}

// GetUserProfile  get user profile information from api.github.com/user
func (github *Github) GetUserProfile() (account OauthAccount, err error) {
	accessToken, err := github.GetAccessToken(github.Code)
	if err != nil {
		return
	}
	u, _ := url.Parse("https://api.github.com/user")
	request := &http.Request{
		Method: "GET",
		URL:    u,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {fmt.Sprintf("token %s", accessToken)},
		},
	}
	resp, err := github.client.Do(request)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = resp.Body.Close()
	if err != nil {
		elog.Warnf("close the response body false %v", err)
	}

	account = &GithubUserProfile{}
	err = json.Unmarshal(body, &account)
	return
}
