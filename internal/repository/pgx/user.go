package pgx

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
)

type User struct {
	pg *database.Postgres
}

func NewUser(pg *database.Postgres) *User {
	return &User{pg: pg}
}

func (u *User) Create(ctx context.Context, login string, wealth int) (model.User, error) {
	sql, args, err := u.pg.Sq.Insert("users").
		Columns("login", "wealth").
		Values(login, wealth).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return model.User{}, err
	}

	var id int

	err = u.pg.Pool.QueryRow(ctx, sql, args).Scan(&id)

	if err != nil {
		return model.User{}, err
	}

	return model.User{Id: id, Login: login, Wealth: wealth}, nil
}
