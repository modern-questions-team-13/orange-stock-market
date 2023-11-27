package database

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/pgx/connector"
)

type Postgres struct {
	Pool connector.Connector
	Sq   squirrel.StatementBuilderType
}

func NewPostgres(connString string) (*Postgres, error) {
	pool, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		return nil, fmt.Errorf("error connectString format: %w", err)
	}

	err = pool.Ping(context.Background())

	if err != nil {
		return nil, fmt.Errorf("error ping to postgres: %w", err)
	}

	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &Postgres{Pool: pool, Sq: sq}, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
