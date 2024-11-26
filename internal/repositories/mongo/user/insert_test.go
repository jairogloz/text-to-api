package user_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.uber.org/mock/gomock"
	"testing"
	"text-to-api/internal/domain"
	"text-to-api/internal/repositories/mongo/user"
	"text-to-api/mocks"
)

func TestRepository_Insert(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	ctrl := gomock.NewController(mt)
	defer ctrl.Finish()

	logger := mocks.NewMockLogger(ctrl)

	mt.Run("insertOne succeeds: live", func(mtSubtest *mtest.T) {
		r, err := user.NewUserRepository(mtSubtest.DB.Client(), logger, "test-db")
		if err != nil {
			mtSubtest.Fatalf("error creating new User repository: %s", err.Error())
		}
		mtSubtest.AddMockResponses(mtest.CreateSuccessResponse())
		err = r.Insert(context.Background(), domain.RequestEnvironmentLive, &domain.User{})
		assert.NoError(mtSubtest, err)
	})

	mt.Run("insertOne succeeds: sandbox", func(mtSubtest *mtest.T) {
		r, err := user.NewUserRepository(mtSubtest.DB.Client(), logger, "test-db")
		if err != nil {
			mtSubtest.Fatalf("error creating new User repository: %s", err.Error())
		}
		mtSubtest.AddMockResponses(mtest.CreateSuccessResponse())
		err = r.Insert(context.Background(), domain.RequestEnvironmentSandbox, &domain.User{})
		assert.NoError(mtSubtest, err)
	})

	mt.Run("error inserting: live", func(mtSubtest *mtest.T) {
		r, err := user.NewUserRepository(mtSubtest.DB.Client(), logger, "test-db")
		if err != nil {
			mtSubtest.Fatalf("error creating new User repository: %s", err.Error())
		}
		mtSubtest.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Code: 11000}))
		logger.EXPECT().Error(gomock.Any() /*context*/, "error inserting user", "error", gomock.Any() /*error*/)
		err = r.Insert(context.Background(), domain.RequestEnvironmentLive, &domain.User{})
		assert.Error(mtSubtest, err)
	})

	mt.Run("error inserting: sandbox", func(mtSubtest *mtest.T) {
		r, err := user.NewUserRepository(mtSubtest.DB.Client(), logger, "test-db")
		if err != nil {
			mtSubtest.Fatalf("error creating new User repository: %s", err.Error())
		}
		mtSubtest.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Code: 11000}))
		logger.EXPECT().Error(gomock.Any() /*context*/, "error inserting user", "error", gomock.Any() /*error*/)
		err = r.Insert(context.Background(), domain.RequestEnvironmentSandbox, &domain.User{})
		assert.Error(mtSubtest, err)
	})

}
