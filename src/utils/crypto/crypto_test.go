package crypto

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	tests := []struct {
		name    string
		userId  string
		wantErr bool
	}{
		{
			name:    "Valid user ID",
			userId:  "user123",
			wantErr: false,
		},
		{
			name:    "Empty user ID",
			userId:  "",
			wantErr: false,
		},
		{
			name:    "User ID with special characters",
			userId:  "user@example.com",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := CreateToken(tt.userId)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, token)

			// Verify token is valid and contains correct user_id
			parsed, errParse := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			assert.NoError(t, errParse)
			assert.True(t, parsed.Valid)

			claims := parsed.Claims.(jwt.MapClaims)
			assert.Equal(t, tt.userId, claims["user_id"])
			assert.NotNil(t, claims["exp"])
		})
	}
}
func TestVerifyToken(t *testing.T) {
	tests := []struct {
		name       string
		userId     string
		setupErr   bool
		wantErr    bool
		wantUserId string
	}{
		{
			name:       "Valid token",
			userId:     "user123",
			setupErr:   false,
			wantErr:    false,
			wantUserId: "user123",
		},
		{
			name:       "Valid token with special characters",
			userId:     "user@example.com",
			setupErr:   false,
			wantErr:    false,
			wantUserId: "user@example.com",
		},
		{
			name:       "Invalid token signature",
			userId:     "user123",
			setupErr:   false,
			wantErr:    true,
			wantUserId: "",
		},
		{
			name:       "Malformed token",
			userId:     "",
			setupErr:   false,
			wantErr:    true,
			wantUserId: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tokenString string

			if tt.name == "Invalid token signature" {
				tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcjEyMyIsImV4cCI6OTk5OTk5OTk5OX0.invalid"
			} else if tt.name == "Malformed token" {
				tokenString = "invalid.token.string"
			} else {
				token, err := CreateToken(tt.userId)
				assert.NoError(t, err)
				tokenString = token
			}

			userId, err := VerifyToken(tokenString)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantUserId, userId)
		})
	}
}
