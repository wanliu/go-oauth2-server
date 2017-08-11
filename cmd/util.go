package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/wanliu/go-oauth2-server/config"
	"github.com/wanliu/go-oauth2-server/database"
)

// initConfigDB loads the configuration and connects to the database
func initConfigDB(mustLoadOnce, keepReloading bool, configBackend string) (*config.Config, *gorm.DB, error) {
	// Config
	cnf := config.NewConfig(mustLoadOnce, keepReloading, configBackend)

	// Database
	db, err := database.NewDatabase(cnf)
	if err != nil {
		return nil, nil, err
	}

	return cnf, db, nil
}
