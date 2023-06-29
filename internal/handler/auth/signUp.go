package auth

import (
	"FaceRecognitionClient/internal/domain"
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Возможно, нужен динамический salt
const (
	salt = "sglsjkgesnglnskgslgn"
)

func (h Handler) SignUp(ctx context.Context, request domain.SignUpRequest) error {
	passwordHash := generatePasswordHash(request.Password)
	user := domain.User{Username: request.Username, Name: request.Name, Password: passwordHash}
	if err := h.authDao.Create(ctx, h.pgx, user); err != nil {
		logrus.Error("internal/handler SignIn ", "err ", err)
		return err
	}
	return nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
