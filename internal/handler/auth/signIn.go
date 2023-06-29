package auth

import (
	"FaceRecognitionClient/internal/domain"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	signingKey = "gdrhaharjta35252rdfwl"
	TokenTTL   = 5 * time.Minute
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (h Handler) SignIn(ctx context.Context, request domain.SignInRequest) (string, error) {
	passwordHash := generatePasswordHash(request.Password)
	user, err := h.authDao.Get(ctx, h.pgx, request.Username, passwordHash)
	if err != nil {
		logrus.Error("internal/SignUp ", "err", err)
		return "", err
	}
	if user == nil {
		return "", domain.UserNameOrPasswordInIncorrectErr
	}
	token, err := generateToken(user.UserID)
	if err != nil {
		logrus.Error("internal/handler/auth SignUp ", "err ", err)
		return "", err
	}
	return token, nil
}
