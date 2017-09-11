package oauth

import (
	"errors"

	"github.com/wanliu/go-oauth2-server/models"
)

var (
	// ErrRoleNotFound ...
	ErrRoleNotFound = errors.New("Role not found")
)

// FindRoleByID looks up a role by ID and returns it
func (s *Service) FindRoleByID(id string) (*models.OauthRole, error) {
	role := new(models.OauthRole)
	if s.db.Where("id = ?", id).First(role).RecordNotFound() {
		return nil, ErrRoleNotFound
	}
	return role, nil
}

func (s *Service) CreateRole(id, name string) (*models.OauthRole, error) {
	var role = models.OauthRole{ID: id, Name: name}

	if err := s.db.FirstOrCreate(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
