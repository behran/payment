package payment

import (
	"context"

	"payment/internal/domains/payment/contracts"
	"payment/internal/dto"
	"payment/internal/models"
)

//Service ...
type Service struct {
	r contracts.IAccountRepository
}

//New ...
func New(r contracts.IAccountRepository) *Service {
	return &Service{r: r}
}

//FindAccaunt ...
func (s Service) FindAccaunt(ctx context.Context, id int) error {
	return s.r.IsExist(ctx, id)
}

//CreateAccount ...
func (s Service) CreateAccount(ctx context.Context, account dto.Account) (models.AccountAttributes, error) {
	return s.r.Create(ctx, account)
}

//UpdateAccountAmount ...
func (s Service) UpdateAccountAmount(ctx context.Context, id int, payload dto.Payload) (models.AccountAttributes, error) {
	return s.r.UpdateAmount(ctx, id, payload)
}

//RollbackTransaction ...
func (s Service) RollbackTransaction(ctx context.Context, rollback dto.RollBack) error {
	if err := s.r.IsExist(ctx, rollback.AccountID); err != nil {
		return err
	}
	return s.r.RollbackTransaction(ctx, rollback)
}
