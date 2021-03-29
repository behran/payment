package contracts

import (
	"context"

	"payment/internal/dto"
	"payment/internal/models"
)

//IPaymentService ...
type IPaymentService interface {
	FindAccaunt(ctx, id int) error
	CreateAccount(ctx context.Context, account dto.Account) (models.AccountAttributes, error)
	UpdateAccountAmount(ctx context.Context, accountID int, payload dto.Payload) (models.AccountAttributes, error)
	RollbackTransaction(ctx context.Context, rollback dto.RollBack) error
}
