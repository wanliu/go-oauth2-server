package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanliu/go-oauth2-server/util"
)

func TestValidateEmail(t *testing.T) {
	assert.False(t, util.ValidateEmail("test@user"))
	assert.True(t, util.ValidateEmail("test@user.com"))
}
