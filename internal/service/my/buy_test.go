package my

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	mock_repository "github.com/modern-questions-team-13/orange-stock-market/internal/repository/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMockBuy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("purchase with save bid", func(t *testing.T) {
		ctx := context.Background()
		user := mock_repository.NewMockUser(ctrl)
		user.EXPECT().Withdraw(gomock.Any(), 2, 2).Return(nil)

		sale := mock_repository.NewMockSale(ctrl)
		sale.EXPECT().GetSales(gomock.Any(), 2, 2, uint64(tryBuyLimit)).Return(nil, nil)
		buy := mock_repository.NewMockBuy(ctrl)
		buy.EXPECT().Create(gomock.Any(), 2, 2, 2)

		buyServ := Buy{repos: &repository.Repositories{
			User: user,
			Sale: sale,
			Buy:  buy,
		}}

		err := buyServ.Create(ctx, 2, 2, 2)
		require.NoError(t, err)
	})
}
