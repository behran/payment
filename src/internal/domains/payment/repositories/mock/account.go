package mock

import (
	"context"

	"payment/internal/dto"
	"payment/internal/models"

	"github.com/stretchr/testify/mock"
)

//AccountRepositoryMock ...
type AccountRepositoryMock struct {
	mock.Mock
}

//IsExist ...
func (a AccountRepositoryMock) IsExist(ctx context.Context, id int) error {
	args := a.Called(ctx, id)
	return args.Error(0)
}

//Create ...
func (a AccountRepositoryMock) Create(ctx context.Context, account dto.Account) (models.AccountAttributes, error) {
	args := a.Called(ctx, account)
	return args.Get(0).(models.AccountAttributes), args.Error(1)
}

//UpdateAmount ...
func (a AccountRepositoryMock) UpdateAmount(ctx context.Context, accountID int, payload dto.Payload) (models.AccountAttributes, error) {
	args := a.Called(ctx, accountID, payload)
	return args.Get(0).(models.AccountAttributes), args.Error(1)
}

//RollbackTransaction ...
func (a AccountRepositoryMock) RollbackTransaction(ctx context.Context, rollback dto.RollBack) error {
	args := a.Called(ctx, rollback)
	return args.Error(0)
}
