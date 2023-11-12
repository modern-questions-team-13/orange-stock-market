package my

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
)

type User struct {
	repos *repository.Repositories
}

func (u *User) Get(ctx context.Context, id int) (model.Info, error) {
	user, err := u.repos.User.Get(ctx, id)

	if err != nil {
		return model.Info{}, err
	}

	assets, err := u.repos.Portfolio.Get(ctx, id)

	if err != nil {
		return model.Info{}, err
	}

	return model.Info{Account: user, Assets: assets}, nil
}

func (u *User) Create(ctx context.Context, login string, wealth int) (token string, err error) {
	id, err := u.repos.User.Create(ctx, login, wealth)

	if err != nil {
		return "", err
	}

	return u.repos.Secret.SetToken(ctx, id)
}

func NewUser(repos *repository.Repositories) *User {
	return &User{repos: repos}
}
