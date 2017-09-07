package web

import (
	"fmt"
	"net/http"

	"github.com/wanliu/go-oauth2-server/form"
	"github.com/wanliu/util/rand"
)

func (s *Service) listClient(w http.ResponseWriter, r *http.Request) {
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

	clients, err := s.GetOauthService().ListClientByUserID(currentUser.ID, 0, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "clients.html", map[string]interface{}{
		"error":       errMsg,
		"currentUser": currentUser,
		"clients":     clients,
		"username":    session.Username,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) newClientForm(w http.ResponseWriter, r *http.Request) {
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

	var f createClientForm

	errMsg, _ := sessionService.GetFlashMessage()
	renderTemplate(w, "new_client.html", map[string]interface{}{
		"error":       errMsg,
		"currentUser": currentUser,
		"form":        &f,
		"username":    session.Username,
		"queryString": getQueryString(r.URL.Query()),
	})
}

func (s *Service) createClient(w http.ResponseWriter, r *http.Request) {
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
	r.ParseForm()

	var f createClientForm
	err = form.ParseForm(&f, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errMsg, _ := sessionService.GetFlashMessage()
	if f.Valid() {
		var (
			clientID = rand.String(32)
			secret   = rand.String(32)
		)

		client, err := s.GetOauthService().CreateClientByUserID(currentUser.ID, f.Name, clientID, secret, f.RedirectURI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		redirectWithFragment(fmt.Sprintf("/web/clients/%s", client.ID), r.URL.Query(), w, r)
	} else {
		renderTemplate(w, "new_client.html", map[string]interface{}{
			"error":       errMsg,
			"currentUser": currentUser,
			"form":        &f,
			"username":    session.Username,
			"queryString": getQueryString(r.URL.Query()),
		})
	}

}

func (s *Service) clientDetail(w http.ResponseWriter, r *http.Request) {
	// sessionService, err := getSessionService(r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// session, err := sessionService.GetUserSession()
	// currentUser, err := s.GetOauthService().FindUserByUsername(session.Username)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

}

func (s *Service) deleteClient(w http.ResponseWriter, r *http.Request) {
}
