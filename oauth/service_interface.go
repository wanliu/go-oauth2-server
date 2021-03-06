package oauth

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/wanliu/go-oauth2-server/config"
	"github.com/wanliu/go-oauth2-server/models"
	"github.com/wanliu/go-oauth2-server/session"
	"github.com/wanliu/go-oauth2-server/util/routes"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	// Exported methods
	GetConfig() *config.Config
	RestrictToRoles(allowedRoles ...string)
	IsRoleAllowed(role string) bool
	FindRoleByID(id string) (*models.OauthRole, error)
	CreateRole(id, name string) (*models.OauthRole, error)
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	ClientExists(clientID string) bool
	FindClientByClientID(clientID string) (*models.OauthClient, error)
	ListClientByUserID(userId string, offset, count int) ([]models.OauthClient, error)
	GetUserByToken(string) (*models.OauthUser, error)
	CreateClient(clientID, secret, redirectURI string) (*models.OauthClient, error)
	CreateClientByUserID(userId, name, clientID, secret, redirectURI string) (*models.OauthClient, error)
	CreateClientTx(tx *gorm.DB, clientID, secret, redirectURI string) (*models.OauthClient, error)
	AuthClient(clientID, secret string) (*models.OauthClient, error)
	UserExists(username string) bool
	FindUserByUsername(username string) (*models.OauthUser, error)
	FindUserByID(username string) (*models.OauthUser, error)
	CreateUser(roleID, username, password string) (*models.OauthUser, error)
	CreateUserTx(tx *gorm.DB, roleID, username, password string) (*models.OauthUser, error)
	SetPassword(user *models.OauthUser, password string) error
	SetPasswordTx(tx *gorm.DB, user *models.OauthUser, password string) error
	UpdateUsername(user *models.OauthUser, username string) error
	UpdateUser(user *models.OauthUser, params Params) error
	UpdateUsernameTx(db *gorm.DB, user *models.OauthUser, username string) error
	AuthUser(username, thePassword string) (*models.OauthUser, error)
	GetScope(requestedScope string) (string, error)
	GetDefaultScope() string
	ScopeExists(requestedScope string) bool
	CreateScope(scope string, isdefault bool) (*models.OauthScope, error)
	Login(client *models.OauthClient, user *models.OauthUser, scope string) (*models.OauthAccessToken, *models.OauthRefreshToken, error)
	GrantAuthorizationCode(client *models.OauthClient, user *models.OauthUser, expiresIn int, redirectURI, scope string) (*models.OauthAuthorizationCode, error)
	GrantAccessToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthAccessToken, error)
	GetOrCreateRefreshToken(client *models.OauthClient, user *models.OauthUser, expiresIn int, scope string) (*models.OauthRefreshToken, error)
	GetValidRefreshToken(token string, client *models.OauthClient) (*models.OauthRefreshToken, error)
	Authenticate(token string) (*models.OauthAccessToken, error)
	NewIntrospectResponseFromAccessToken(accessToken *models.OauthAccessToken) (*IntrospectResponse, error)
	NewIntrospectResponseFromRefreshToken(refreshToken *models.OauthRefreshToken) (*IntrospectResponse, error)
	ClearUserTokens(userSession *session.UserSession)
	Close()
}
