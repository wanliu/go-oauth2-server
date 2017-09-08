package oauth

import (
	"errors"
	"strings"
	"time"

	"github.com/RichardKnop/uuid"
	"github.com/jinzhu/gorm"
	"github.com/wanliu/go-oauth2-server/models"
	"github.com/wanliu/go-oauth2-server/util"
	"github.com/wanliu/go-oauth2-server/util/password"
)

var (
	// ErrClientNotFound ...
	ErrClientNotFound = errors.New("Client not found")
	// ErrInvalidClientSecret ...
	ErrInvalidClientSecret = errors.New("Invalid client secret")
	// ErrClientIDTaken ...
	ErrClientIDTaken = errors.New("Client ID taken")
)

// ClientExists returns true if client exists
func (s *Service) ClientExists(clientID string) bool {
	_, err := s.FindClientByClientID(clientID)
	return err == nil
}

// FindClientByClientID looks up a client by client ID
func (s *Service) FindClientByClientID(clientID string) (*models.OauthClient, error) {
	// Client IDs are case insensitive
	client := new(models.OauthClient)
	notFound := s.db.Where("key = LOWER(?)", clientID).
		First(client).RecordNotFound()

	// Not found
	if notFound {
		return nil, ErrClientNotFound
	}

	return client, nil
}

// CreateClient saves a new client to database
func (s *Service) CreateClient(clientID, secret, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(s.db, clientID, secret, redirectURI, "", "")
}

// CreateClient saves a new client to database
func (s *Service) CreateClientByUserID(userId, name, clientID, secret, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(s.db, clientID, secret, redirectURI, userId, name)
}

// CreateClientTx saves a new client to database using injected db object
func (s *Service) CreateClientTx(tx *gorm.DB, clientID, secret, redirectURI string) (*models.OauthClient, error) {
	return s.createClientCommon(tx, clientID, secret, redirectURI, "", "")
}

// AuthClient authenticates client
func (s *Service) AuthClient(clientID, secret string) (*models.OauthClient, error) {
	// Fetch the client
	client, err := s.FindClientByClientID(clientID)
	if err != nil {
		return nil, ErrClientNotFound
	}

	// Verify the secret
	if password.VerifyPassword(client.Secret, secret) != nil {
		return nil, ErrInvalidClientSecret
	}

	return client, nil
}

func (s *Service) createClientCommon(db *gorm.DB, clientID, pass, redirectURI, userId, name string) (*models.OauthClient, error) {
	// Check client ID
	if s.ClientExists(clientID) {
		return nil, ErrClientIDTaken
	}

	// Hash password
	secretHash, err := password.HashPassword(pass)
	if err != nil {
		return nil, err
	}

	client := &models.OauthClient{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		Name:        util.StringOrNull(name),
		UserID:      util.StringOrNull(userId),
		Key:         strings.ToLower(clientID),
		Secret:      string(secretHash),
		Password:    util.StringOrNull(pass),
		RedirectURI: util.StringOrNull(redirectURI),
	}
	if err := db.Create(client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

func (s *Service) ListClientByUserID(userId string, offset, count int) ([]models.OauthClient, error) {
	var clients []models.OauthClient
	if err := s.db.Find(&clients, "user_id = ?", userId).Offset(offset).Limit(count).Error; err != nil {
		return nil, err
	}

	return clients, nil
}
