package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/tcnksm/go-input"
	"github.com/wanliu/go-oauth2-server/services"
)

// LoadData loads fixtures
func Setup(configBackend string) error {
	cnf, db, err := initConfigDB(true, false, configBackend)
	if err != nil {
		return err
	}
	defer db.Close()

	// start the services
	if err := services.Init(cnf, db); err != nil {
		return err
	}
	defer services.Close()

	ui := &input.UI{}

	query := "What is Admin email?"
	name, err := ui.Ask(query, &input.Options{
		// Read the default val from env var
		Default:  "admin@admin",
		Required: true,
		Loop:     true,
	})
	if err != nil {
		log.Fatal(err)
	}

	query = "What is Admin password?"
	pass, err := ui.Ask(query, &input.Options{
		Required:    true,
		Mask:        true,
		MaskDefault: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	services.OauthService.CreateUser("superuser", name, pass)
	log.Printf("Create superuser %s is success.\n", name)

	// create default client for normal signin
	query = "What is login domain?"
	domain, err := ui.Ask(query, &input.Options{
		Default:  "http://localhost:8080/web/",
		Required: true,
		Loop:     true,
		ValidateFunc: func(s string) error {
			_, err = url.Parse(s)
			if err != nil {
				return fmt.Errorf("input must a valid url or domain")
			}
			return nil
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	redirectUri, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
	}

	services.OauthService.CreateClient("normal-client", "normal-client-secret", redirectUri.String())
	services.OauthService.CreateClient("admin-client", "admin-client-secret", redirectUri.String())

	return nil
}
