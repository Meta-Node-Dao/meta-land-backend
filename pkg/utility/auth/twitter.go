package auth

/// TODO: need help to complete the twitter oauth login logic

// Twitter
type Twitter struct {
	ConsumerKey    string
	ConsumerSecret string
}

// TwitterOauthAccount  Twitter Oauth user profile account
type TwitterOauthAccount struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url_https"`
}

// GetUserID implements the OauthAccount interface
func (account *TwitterOauthAccount) GetUserID() string {
	return account.ID
}

// GetUserNick implements the OauthAccount interface
func (account *TwitterOauthAccount) GetUserNick() string {
	return account.Name
}

// GetUserAvatar implements the OauthAccount interface
func (account *TwitterOauthAccount) GetUserAvatar() string {
	return account.ProfileImageURL
}
