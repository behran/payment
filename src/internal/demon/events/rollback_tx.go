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
	return facade.Service().Payment().RollbackTransaction(context.Background(), dto.RollBack{
		AccountID:  defaultUserID,
		SourceType: rollbackByServer,
	})
}
