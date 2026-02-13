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
		wantUserId *int64
	}{
		{
			name:       "Valid token with numeric user ID",
			userId:     "12345",
			setupErr:   false,
			wantErr:    false,
			wantUserId: func() *int64 { v := int64(12345); return &v }(),
		},
		{
			name:       "Valid token with large numeric user ID",
			userId:     "9223372036854775807",
			setupErr:   false,
			wantErr:    false,
			wantUserId: func() *int64 { v := int64(9223372036854775807); return &v }(),
		},
		{
			name:       "Invalid token signature",
			userId:     "12345",
			setupErr:   false,
			wantErr:    true,
			wantUserId: nil,
		},
		{
			name:       "Malformed token",
			userId:     "",
			setupErr:   false,
			wantErr:    true,
			wantUserId: nil,
		},
		{
			name:       "User ID not numeric",
			userId:     "user@example.com",
			setupErr:   false,
			wantErr:    true,
			wantUserId: nil,
		},
		{
			name:       "Empty token string",
			userId:     "",
			setupErr:   false,
			wantErr:    true,
			wantUserId: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tokenString string

			switch tt.name {
			case "Invalid token signature":
				tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJleHAiOjk5OTk5OTk5OTl9.invalid"
			case "Malformed token":
				tokenString = "invalid.token.string"
			case "Empty token string":
				tokenString = ""
			default:
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
