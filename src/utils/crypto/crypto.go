package crypto

import (
	"strconv"
	"time"

	uerror "amartha/src/utils/error"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("amartha")

func CreateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (*int64, error) {
	if tokenString == "" {
		return nil, uerror.ErrUnauthorized
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		num, err := strconv.ParseInt(claims["user_id"].(string), 10, 64)
		if err != nil {
			return nil, err
		}

		return &num, nil
	}

	return nil, err
}
