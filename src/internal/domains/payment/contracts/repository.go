package contracts

import (
	"context"

	"payment/internal/dto"
	"payment/internal/models"
)

//IPaymentRepository ...
type IAccountRepository interface {
	FindByID(ctx context.Context, id int) error
	Create(ctx context.Context, account dto.Account) (models.AccountAttributes, error)
	UpdateAmount(ctx context.Context, accountID int, payload dto.Payload) (models.AccountAttributes, error)
	RollbackTransaction(ctx context.Context, rollback dto.RollBack) error
}
