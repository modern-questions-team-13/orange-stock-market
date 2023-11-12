package pgx

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
)

const length = 20

type Auth struct {
	pg *database.Postgres
}

func NewAuth(pg *database.Postgres) *Auth {
	return &Auth{pg: pg}
}

func (a *Auth) SetToken(ctx context.Context, id int) (token string, err error) {
	token, err = generateSecureToken(length)

	if err != nil {
		return "", err
	}

	sql, args, err := a.pg.Sq.
		Insert("secrets").
		Columns("user_id", "token").
		Values(id, token).
		ToSql()

	if err != nil {
		return "", err
	}

	_, err = a.pg.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return "", err
	}

	return token, err
}

func (a *Auth) GetUserId(ctx context.Context, token string) (int, error) {
	sql, args, err := a.pg.Sq.
		Select("user_id").
		From("secrets").
		Where("token=?", token).
		ToSql()

	if err != nil {
		return 0, err
	}

	var id int

	err = a.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return 0, fmt.Errorf("error get user_id with token=%q: %w", token, repoerrs.ErrNotExists)
			}
		}

		return 0, err
	}

	return id, nil
}

func generateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
