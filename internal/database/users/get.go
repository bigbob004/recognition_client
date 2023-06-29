package users

import (
	"FaceRecognitionClient/internal/domain"
	"context"
	"github.com/sirupsen/logrus"

	db "FaceRecognitionClient/internal/database"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (d Dao) Get(ctx context.Context, p db.PGX, userName string, passwordHash string) (*domain.User, error) {
	qb := db.PSQL.
		Select("id",
			"username",
			"password_hash",
		).
		From(db.UsersTableName).
		Where(sq.And{
			sq.Eq{"username": userName},
			sq.Eq{"password_hash": passwordHash},
		})

	query, args, err := qb.ToSql()
	if err != nil {
		logrus.Error("internal/database/users Get/ToSql", "err", err)
		return nil, err
	}

	var resUsers []*User
	if err = pgxscan.Select(ctx, p, &resUsers, query, args...); err != nil {
		logrus.Error("internal/database/users Get/Select ", "err", err)
		return nil, err
	}
	if len(resUsers) == 0 {
		return nil, nil
	}
	return mapToDomain(resUsers[0]), nil
}

func mapToDomain(user *User) *domain.User {
	return &domain.User{UserID: user.Id, Username: user.Username, Name: user.Name, Password: user.PasswordHash}
}
