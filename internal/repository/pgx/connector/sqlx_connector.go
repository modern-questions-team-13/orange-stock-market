package connector

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLXConnector struct {
	conn *pgxpool.Pool
}

func (s *SQLXConnector) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return s.conn.Exec(ctx, sql, arguments...)
}

func (s *SQLXConnector) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return s.conn.Query(ctx, sql, args...)
}

func (s *SQLXConnector) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return s.conn.QueryRow(ctx, sql, args...)
}

func NewSQLXConnector(conn *pgxpool.Pool) *SQLXConnector {
	return &SQLXConnector{conn: conn}
}
