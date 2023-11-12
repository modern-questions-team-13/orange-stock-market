package pgx

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
)

type Company struct {
	pg *database.Postgres
}

func (c *Company) GetAll(ctx context.Context) ([]model.Company, error) {
	sql, args, err := c.pg.Sq.Select("*").From("companies").ToSql()

	rows, err := c.pg.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	companies := make([]model.Company, 0)

	for rows.Next() {
		var company model.Company

		err = rows.Scan(&company.Id, &company.Name)

		if err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func NewCompany(pg *database.Postgres) *Company {
	return &Company{pg: pg}
}
