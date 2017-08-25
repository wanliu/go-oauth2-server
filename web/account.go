package web

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func (s *Service) index(w http.ResponseWriter, r *http.Request) {
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := sessionService.GetUserSession()
	currentUser, err := s.GetOauthService().FindUserByUsername(session.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var f = UpdateUserForm{User: currentUser}

	// Render the template
	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "index.html", map[string]interface{}{
		"error":       errMsg,
		"currentUser": currentUser,
		"form":        &f,
		"username":    session.Username,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) updateUser(w http.ResponseWriter, r *http.Request) {
	// Get the session service from the request context
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := sessionService.GetUserSession()
	currentUser, err := s.GetOauthService().FindUserByUsername(session.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.ParseMultipartForm(32 << 20)

	var f = UpdateUserForm{User: currentUser}
	err = parseForm(&f, r.PostForm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f.Valid()

	filename, err := uploadFile("avatar_url", r)

	if err == nil {

		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// f.AddError("AvatarURL", err.Error())
		avatarUrl, err := s.mapAssetToSuite(filename)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			f.AddError("AvatarURL", err.Error())
		}

		if len(avatarUrl) > 0 {
			s.GetOauthService().UpdateUser(currentUser, map[string]interface{}{"AvatarURL": avatarUrl})
		}
	}

	password := r.Form.Get("password")

	if len(password) > 0 {
		if err := s.GetOauthService().SetPassword(currentUser, password); err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			f.AddError("Password", err.Error())
			// return
		}
	}

	params := f.Diff()
	if len(params) > 0 {
		s.GetOauthService().UpdateUser(currentUser, params)
	}

	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "index.html", map[string]interface{}{
		"error":       errMsg,
		"form":        &f,
		"currentUser": currentUser,
		"username":    session.Username,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) mapAssetToSuite(filename string) (uri string, err error) {
	config := s.GetConfig()
	if config == nil {
		return "", fmt.Errorf("invalid config")
	}

	log.Printf("map: %#v", config.AssetsMappings)

	for _, asset := range config.AssetsMappings {
		if strings.Index(filename, asset.Dir) == 0 {
			rest := filename[len(asset.Dir):]
			u, err := url.Parse(asset.Host)
			if err != nil {
				continue
			}

			u.Path = path.Join(u.Path, rest)
			return u.String(), nil
		}
	}

	return "", fmt.Errorf("don't have mapping asset to host path config")
}
