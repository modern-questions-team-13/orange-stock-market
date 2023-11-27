package connector

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"gotest.tools/v3/golden"
	"strings"
	"testing"
)

var goldenReplacer = strings.NewReplacer(
	";", ";\n",
	",", ",\n",
	"SET", "\nSET",
	"WHERE", "\nWHERE",
	"UPDATE", "\nUPDATE",
	"COMMIT;", "\nCOMMIT;",
	"[[[", "\n[[[",
)

type GoldenConnector struct {
	t        *testing.T
	filename string
}

func NewGoldenConnector(t *testing.T, filename string) *GoldenConnector {
	return &GoldenConnector{
		t:        t,
		filename: filename,
	}
}

func (r *GoldenConnector) SetFilename(filename string) *GoldenConnector {
	r.filename = filename
	return r
}

func (r *GoldenConnector) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {

	sql = goldenReplacer.Replace(sql)
	golden.Assert(r.t, sql, r.filename)
	return nil, pgx.ErrNoRows
}

func (r *GoldenConnector) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {

	sql = goldenReplacer.Replace(sql)
	golden.Assert(r.t, sql, r.filename)
	return pgconn.CommandTag{}, nil
}

func (r *GoldenConnector) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {

	sql = goldenReplacer.Replace(sql)
	golden.Assert(r.t, sql, r.filename)
	return nil
}

func (r *GoldenConnector) Close() {
	return
}
