package response_test

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanliu/go-oauth2-server/util/response"
)

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	obj := map[string]interface{}{
		"foo": "bar",
		"qux": 1,
	}
	response.WriteJSON(w, obj, 201)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	expected, _ := json.Marshal(obj)
	assert.Equal(t, string(expected), strings.TrimSpace(w.Body.String()))
}

func TestNoContent(t *testing.T) {
	w := httptest.NewRecorder()
	response.NoContent(w)

	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", strings.TrimSpace(w.Body.String()))
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	response.Error(w, "something went wrong", 500)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	expected := "{\"error\":\"something went wrong\"}"
	assert.Equal(t, expected, strings.TrimSpace(w.Body.String()))
}
