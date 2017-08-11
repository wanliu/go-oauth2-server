package database_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanliu/go-oauth2-server/config"
	"github.com/wanliu/go-oauth2-server/database"
)

func TestNewDatabaseTypeNotSupported(t *testing.T) {
	cnf := &config.Config{
		Database: config.DatabaseConfig{
			Type: "bogus",
		},
	}
	_, err := database.NewDatabase(cnf)

	if assert.NotNil(t, err) {
		assert.Equal(t, errors.New("Database type bogus not suppported"), err)
	}
}
