package client_test

import (
	"context"
	"github.com/jairogloz/pgxpoolmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"text-to-api/internal/domain"
	"text-to-api/internal/repositories/postgres/client"
	"text-to-api/mocks"
	"time"
)

func TestRepository_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mocks.NewMockLogger(ctrl)
	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	repo, err := client.NewClientRepository(logger, mockPool)
	if err != nil {
		t.Fatalf("failed to create client repository: %v", err)
	}

	t.Run("error querying client", func(t *testing.T) {
		row := pgxpoolmock.NewRow([]string{"id", "name"}, "client_id", "client_name")
		mockPool.EXPECT().QueryRow(gomock.Any() /*context*/, client.QueryGetClientByID, "client_id").Return(row)
		logger.EXPECT().Error(gomock.Any() /*context*/, "error scanning client by id", "error", gomock.Any() /*error*/)
		c, err := repo.GetByID(context.Background(), "client_id")
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "failed to scan client by id")
		}
		assert.Nil(t, c)
	})

	t.Run("success", func(t *testing.T) {
		now := time.Now().UTC()
		row := pgxpoolmock.NewRow([]string{"id", "email", "phone", "created_at", "last_sign_in_at", "subscription_id",
			"customer_id"}, "client_id", "client_email", domain.Ptr("client_phone"), now, now,
			domain.Ptr("subscription_id"), domain.Ptr("customer_id"))
		mockPool.EXPECT().QueryRow(gomock.Any() /*context*/, client.QueryGetClientByID, "client_id").Return(row)
		c, err := repo.GetByID(context.Background(), "client_id")
		assert.NoError(t, err)
		if assert.NotNil(t, c) {
			assert.Equal(t, "client_id", c.ID)
			if assert.NotNil(t, c.Phone) {
				assert.Equal(t, "client_phone", *c.Phone)
			}
			if assert.NotNil(t, c.SubscriptionID) {
				assert.Equal(t, "subscription_id", *c.SubscriptionID)
			}
			if assert.NotNil(t, c.CustomerID) {
				assert.Equal(t, "customer_id", *c.CustomerID)
			}
		}
	})
}
