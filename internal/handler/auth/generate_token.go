package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})
	return token.SignedString([]byte(signingKey))
}
