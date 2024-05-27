package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
)

func TestHandle(t *testing.T) {
	t.Run("should retur json with pong and user id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/ping", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{
			Claims: &auth.JwtCustomClaims{
				ID: "123",
			},
		})

		h := NewPingHandler()
		h.Handle(c)

		var res map[string]string
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "pong - user id: 123", res["message"])
	})
}
