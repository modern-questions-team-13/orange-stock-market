package my

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
)

type Auth struct {
	repos *repository.Repositories
}

func NewAuth(repos *repository.Repositories) *Auth {
	return &Auth{repos: repos}
}

func (a *Auth) GetUserId(ctx context.Context, token string) (int, bool) {
	id, err := a.repos.GetUserId(ctx, token)

	if err != nil {
		return 0, false
	}

	return id, true
}
