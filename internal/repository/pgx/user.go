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
	//TODO implement me
	panic("implement me")
}
