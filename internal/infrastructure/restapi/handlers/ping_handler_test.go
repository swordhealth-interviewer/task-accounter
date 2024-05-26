package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	t.Run("should retur json with pong", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v1/ping", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := NewPingHandler()
		h.Handle(c)

		var res map[string]string
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "pong", res["message"])
	})
}
