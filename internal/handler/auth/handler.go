package auth

import (
	db "FaceRecognitionClient/internal/database"
	"FaceRecognitionClient/internal/domain"
	"context"
)

type authDao interface {
	Create(ctx context.Context, p db.PGX, user domain.User) error
	Get(ctx context.Context, p db.PGX, userName string, passwordHash string) (*domain.User, error)
}

type Handler struct {
	authDao authDao
	pgx     db.PGX
}

func NewHadler(dao authDao, pgx db.PGX) Handler {
	return Handler{authDao: dao, pgx: pgx}
}
