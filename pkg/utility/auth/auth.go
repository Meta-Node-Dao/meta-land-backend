package auth

import (
	"ceres/pkg/config"
	"crypto/tls"
	"net/http"

	"golang.org/x/oauth2"
)

// Comunion Oauth interface
// Comunion Ceres only do the final legged in all Oauth2 processing
// The Frontend will handle the other two legged using the standard Oauth2 API

// OauthAccount  Oauth account interface to get the Oauth user unique ID nick name and the avatar
type OauthAccount interface {

	// GetUserID
	// get the user unique ID for every userID
	GetUserID() string

	// GetUserNick
	// get user nick name from Oauth Account
	GetUserNick() string

	// GetUserAvatar
	// get user avatar from Oauth Account
	GetUserAvatar() string
}

// OauthClient  Abstraction of Oauth Login logic
type OauthClient interface {
	// GetUserProfile return oauth account profile from third website
	GetUserProfile() (account OauthAccount, err error)
}

// FIXME：should replace with ceres http library in the future
var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

// NewGithubOauthClient  build a new Github client with the request token from login
func NewGithubOauthClient(code string) (client OauthClient) {
	return &Github{
		ClientID:     config.Github.ClientID,
		ClientSecret: config.Github.ClientSecret,
		client:       httpClient,
		Code:         code,
	}
}

// NewFacebookClient build a new Facebook client with the request token from login
func NewFacebookClient(requestToken string) (client OauthClient) {
	return &Facebook{
		ClientID:     config.Facebook.ClientID,
		ClientSecret: config.Facebook.ClientSecret,
		RedirectURI:  config.Facebook.CallbackURL,
		client:       httpClient,
		RequestToken: requestToken,
	}
}

// NewGoogleClient build a new Google client with the request token from login
func NewGoogleClient(code string) (client *Google) {
	return &Google{
		Config: oauth2.Config{
			ClientID:     config.Google.ClientID,
			ClientSecret: config.Google.ClientSecret,
			RedirectURL:  config.Google.CallbackURL,
			Endpoint:     googleEndpoint,
			Scopes:       scopes,
		},
		client: httpClient,
		Code:   code,
	}
}

// NewTwitterClient  build a new Twitter client with the request token from login
func NewTwitterClient() (client OauthClient) {
	return
}

// NewLinkedinClient  build a new LinkedIn client with the request token from login
func NewLinkedinClient() (client OauthClient) {
	return
}
