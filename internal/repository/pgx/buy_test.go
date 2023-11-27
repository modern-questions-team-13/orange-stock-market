package pgx

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/pgx/connector"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuyWithGolden(t *testing.T) {
	pg := &database.Postgres{
		Pool: connector.NewGoldenConnector(t, "buy/get_buys"),
		Sq:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
	b := Buy{pg: pg}
	_, err := b.GetBuys(context.Background(), 1, 1, 1)
	require.Error(t, err)
	assert.Equal(t, pgx.ErrNoRows, err)
}
