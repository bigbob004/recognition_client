package users

import (
	db "FaceRecognitionClient/internal/database"
	"FaceRecognitionClient/internal/domain"
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

func (d Dao) Create(ctx context.Context, p db.PGX, user domain.User) error {
	qb := db.PSQL.
		Insert(db.UsersTableName).
		Columns("name",
			"username",
			"password_hash",
		)

	qb = qb.Values(
		user.Name,
		user.Username,
		user.Password,
	)
	query, args, err := qb.ToSql()
	if err != nil {
		logrus.Error("internal/database/users Create/ToSql ", "err ", err)
		return err
	}
	_, err = p.Exec(ctx, query, args...)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == "23505" {
			return domain.UserNameAlreadyExist
		}
		logrus.Error("internal/database/users Create/Exec ", "err", err, "query ", query, "args ", args)
		return err
	}
	return nil
}
