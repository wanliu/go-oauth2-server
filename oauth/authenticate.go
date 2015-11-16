package oauth

import (
	"errors"
	"time"
)

// Authenticate checks the access token is valid
func (s *Service) Authenticate(token string) error {
	// Fetch the access token from the database
	accessToken := new(AccessToken)
	if s.db.Where(AccessToken{
		Token: token,
	}).First(accessToken).RecordNotFound() {
		return errors.New("Access token not found")
	}

	// Check the access token hasn't expired
	if time.Now().After(accessToken.ExpiresAt) {
		return errors.New("Access token expired")
	}

	return nil
}