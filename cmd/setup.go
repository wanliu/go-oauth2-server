package cmd

import (
	"log"

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
		Default:  "admin@wanliu.biz",
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

	return nil
}
