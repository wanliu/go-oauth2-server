package oauth

import (
	"github.com/wanliu/go-oauth2-server/models"
)

// AccessTokenResponse ...
type AccessTokenResponse struct {
	UserID       string        `json:"user_id,omitempty"`
	User         *UserResponse `json:"user,omitempty"`
	AccessToken  string        `json:"access_token"`
	ExpiresIn    int           `json:"expires_in"`
	TokenType    string        `json:"token_type"`
	Scope        string        `json:"scope"`
	RefreshToken string        `json:"refresh_token,omitempty"`
}

// IntrospectResponse ...
type IntrospectResponse struct {
	Active       bool          `json:"active"`
	Scope        string        `json:"scope,omitempty"`
	ClientID     string        `json:"client_id,omitempty"`
	Username     string        `json:"username,omitempty"`
	User         *UserResponse `json:"user,omitempty"`
	TokenType    string        `json:"token_type,omitempty"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresAt    int           `json:"exp,omitempty"`
}

type UserResponse struct {
	ID       string `json:"id"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// NewAccessTokenResponse ...
func NewAccessTokenResponse(accessToken *models.OauthAccessToken, refreshToken *models.OauthRefreshToken, lifetime int, theTokenType string) (*AccessTokenResponse, error) {
	response := &AccessTokenResponse{
		AccessToken: accessToken.Token,
		ExpiresIn:   lifetime,
		TokenType:   theTokenType,
		Scope:       accessToken.Scope,
	}
	if accessToken.User != nil {
		response.UserID = accessToken.User.MetaUserID
		response.User = &UserResponse{
			ID:       accessToken.User.ID,
			NickName: accessToken.User.NickName.String,
			Avatar:   accessToken.User.Avatar(),
		}
	}
	if refreshToken != nil {
		response.RefreshToken = refreshToken.Token
	}
	return response, nil
}
