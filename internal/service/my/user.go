package my

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
)

type User struct {
	repos *repository.Repositories
}

func NewUser(repos *repository.Repositories) *User {
	return &User{repos: repos}
}

func (u *User) Create(ctx context.Context, login string, wealth int) (token string, err error) {
	id, err := u.repos.User.Create(ctx, login, wealth)

	if err != nil {
		return "", err
	}

	return u.repos.Secret.SetToken(ctx, id)
}
