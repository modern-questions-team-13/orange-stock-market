package pgx

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
)

type User struct {
	pg *database.Postgres
}

func NewUser(pg *database.Postgres) *User {
	return &User{pg: pg}
}

func (u *User) Create(ctx context.Context, login string, wealth int) (id int, err error) {
	sql, args, err := u.pg.Sq.Insert("users").
		Columns("login", "wealth").
		Values(login, wealth).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return 0, err
	}

	err = u.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return 0, fmt.Errorf("error create user with login=%q: %w", login, repoerrs.ErrAlreadyExists)
			}
		}

		return 0, err
	}

	return id, nil
}
