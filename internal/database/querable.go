package database

import (
	"FaceRecognitionClient/pkg/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PGX struct {
	Querable
}

type Querable interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
}

type DataBaseConfig struct {
	DbName   string
	Port     string
	Password string
	Username string
	Host     string
}

func NewPGX(ctx context.Context, cfg DataBaseConfig, maxAttempts int) (PGX, error) {
	var pool *pgxpool.Pool
	var err error
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	err = utils.DoWithRetries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		logrus.Error("internal/database NewPGX", "err", err)
		return PGX{}, err
	}
	if err = pool.Ping(ctx); err != nil {
		logrus.Error("internal/database NewPGX.Ping", "err", err)
		return PGX{}, err
	}

	return PGX{Querable: pool}, nil
}
