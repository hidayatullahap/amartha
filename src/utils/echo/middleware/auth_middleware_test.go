package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"amartha/src/utils/constants"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	m := NewAuthMiddleware(nil)
	e := echo.New()

	tests := []struct {
		name     string
		role     interface{}
		allowed  []string
		wantCode int
		wantBody string
	}{
		{
			name:     "missing role",
			role:     nil,
			allowed:  []string{"admin"},
			wantCode: http.StatusUnauthorized,
			wantBody: "\"invalid role\"\n",
		},
		{
			name:     "insufficient permissions",
			role:     "user",
			allowed:  []string{"admin"},
			wantCode: http.StatusForbidden,
			wantBody: "\"insufficient permissions\"\n",
		},
		{
			name:     "allowed role",
			role:     "admin",
			allowed:  []string{"admin"},
			wantCode: http.StatusOK,
			wantBody: "\"ok\"\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			pctx := ctx

			if tt.role != nil {
				pctx.Set(constants.UserRole, tt.role)
			}

			next := func(c *echo.Context) error {
				return c.JSON(http.StatusOK, "ok")
			}

			handler := m.Authorize(tt.allowed)(next)
			err := handler(pctx)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantCode, rec.Code)
			assert.Equal(t, tt.wantBody, rec.Body.String())
		})
	}
}
