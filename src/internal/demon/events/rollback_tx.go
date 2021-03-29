package events

import (
	"context"

	"payment/internal/dto"
	"payment/internal/facade"
)

//RollbackTx ...
type RollbackTx struct{}

//NewRollbackTx ...
func NewRollbackTx() *RollbackTx {
	return &RollbackTx{}
}

const (
	defaultUserID    = 1
	rollbackByServer = "server"
)

//Run ...
func (r RollbackTx) Run() error {
	ctx := context.Background()
	rb := dto.RollBack{
		AccountID:  defaultUserID,
		SourceType: rollbackByServer,
	}
	return facade.Service().Payment().RollbackTransaction(ctx, rb)
}
