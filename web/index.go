package web

import (
	// "fmt"
	"log"
	"net/http"
	// "net/url"
	// "path"
	// "strings"

	"github.com/wanliu/go-oauth2-server/models"
)

func (s *Service) index(w http.ResponseWriter, r *http.Request) {
	var (
		currentUser *models.OauthUser
		username    string
	)

	log.Printf("--------------> index ")

	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err == nil {
		session, err := sessionService.GetUserSession()
		if err == nil {
			currentUser, _ = s.GetOauthService().FindUserByUsername(session.Username)
			username = session.Username
		}
	}

	// Render the template
	renderTemplate(w, "index.html", map[string]interface{}{
		"currentUser": currentUser,
		"username":    username,
	})
}
