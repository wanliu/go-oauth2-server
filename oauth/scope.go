package oauth

import (
	"errors"
	"sort"
	"strings"
	"time"

	"github.com/RichardKnop/uuid"
	"github.com/wanliu/go-oauth2-server/models"
)

var (
	// ErrInvalidScope ...
	ErrInvalidScope = errors.New("Invalid scope")
)

// GetScope takes a requested scope and, if it's empty, returns the default
// scope, if not empty, it validates the requested scope
func (s *Service) GetScope(requestedScope string) (string, error) {
	// Return the default scope if the requested scope is empty
	if requestedScope == "" {
		return s.GetDefaultScope(), nil
	}

	// If the requested scope exists in the database, return it
	if s.ScopeExists(requestedScope) {
		return requestedScope, nil
	}

	// Otherwise return error
	return "", ErrInvalidScope
}

// GetDefaultScope returns the default scope
func (s *Service) GetDefaultScope() string {
	// Fetch default scopes
	var scopes []string
	s.db.Model(new(models.OauthScope)).Where("is_default = ?", true).Pluck("scope", &scopes)

	// Sort the scopes alphabetically
	sort.Strings(scopes)

	// Return space delimited scope string
	return strings.Join(scopes, " ")
}

// ScopeExists checks if a scope exists
func (s *Service) ScopeExists(requestedScope string) bool {
	// Split the requested scope string
	scopes := strings.Split(requestedScope, " ")

	// Count how many of requested scopes exist in the database
	var count int
	s.db.Model(new(models.OauthScope)).Where("scope in (?)", scopes).Count(&count)

	// Return true only if all requested scopes found
	return count == len(scopes)
}

func (s *Service) CreateScope(name string, isDefault bool) (*models.OauthScope, error) {
	var scope = models.OauthScope{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Scope:     name,
		IsDefault: isDefault,
	}

	if err := s.db.FirstOrCreate(&scope).Error; err != nil {
		return nil, err
	}

	return &scope, nil
}
