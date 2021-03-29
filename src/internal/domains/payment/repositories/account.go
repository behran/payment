package repositories

import (
	"context"
	"database/sql"
	"errors"

	"payment/internal/config"
	"payment/internal/database"
	erraccount "payment/internal/domains/payment/errors"
	"payment/internal/dto"
	"payment/internal/models"
	logger "payment/pkg/log"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

//AccountRepository ...
type AccountRepository struct {
	pool *sqlx.DB
}

const (
	startBalance       = 0 // Default balance value when create account ...
	transactionPending = "pending"
)

//NewAccount ...
func NewAccount(manager *database.ConnectManager) (*AccountRepository, error) {
	var err error

	repository := new(AccountRepository)
	repository.pool, err = manager.ConnectPostgreSQL(repository)
	if err != nil {
		return nil, err
	}
	return repository, nil
}

//FindByID ...
func (r AccountRepository) FindByID(ctx context.Context, id int) error {
	query := "SELECT * FROM accounts WHERE account_id = $1"

	return r.pool.QueryRowxContext(ctx, query, id).Err()
}

//Create ...
func (r AccountRepository) Create(ctx context.Context, account dto.Account) (models.AccountAttributes, error) {
	var model models.AccountAttributes

	query := "INSERT INTO accounts (account_name, balance) VALUES ($1, $2) RETURNING account_id,account_name,balance"

	if err := r.pool.QueryRowxContext(ctx, query, account.Name, startBalance).StructScan(&model); err != nil {
		logger.Logger.Error("fail create account", zap.Error(err))
		return models.AccountAttributes{}, err
	}
	return model, nil
}

//UpdateAmount ...
func (r AccountRepository) UpdateAmount(ctx context.Context, accountID int, payload dto.Payload) (models.AccountAttributes, error) {
	var (
		balance float64
		account models.AccountAttributes
	)

	tx, err := r.pool.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return models.AccountAttributes{}, err
	}

	if err := tx.QueryRowxContext(ctx,
		"SELECT balance FROM accounts WHERE account_id = $1",
		accountID,
	).Scan(&balance); err != nil {
		logger.Logger.Error("fail select amount from accounts", zap.Error(err))
		if err := tx.Rollback(); err != nil {
			logger.Logger.Error("fail rollback tx", zap.Error(err))
			return models.AccountAttributes{}, erraccount.ErrServer
		}
		if errors.Is(err, sql.ErrNoRows) {
			return models.AccountAttributes{}, erraccount.ErrAccountNotFound
		}
		return models.AccountAttributes{}, erraccount.ErrServer
	}
	// check amount ...
	if balance+payload.AmountDecimal() < 0 {
		if err := tx.Rollback(); err != nil {
			logger.Logger.Error("fail rollback tx", zap.Error(err))
			return models.AccountAttributes{}, erraccount.ErrServer
		}
		return models.AccountAttributes{}, erraccount.ErrInvalidBalance
	}

	queryInsert := "INSERT INTO transactions (unique_id, state, amount, status, source_type,account_id) " +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	// TODO: чекнуть типы, мог ошибиться  ...
	if _, err := tx.ExecContext(ctx, queryInsert,
		payload.TransactionID,
		payload.State,
		payload.AmountDecimal(),
		transactionPending,
		payload.SourceType(),
		accountID,
	); err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Logger.Error("fail rollback tx", zap.Error(err))
			return models.AccountAttributes{}, erraccount.ErrServer
		}
		logger.Logger.Error("fail insert transaction", zap.Error(err))
		return models.AccountAttributes{}, erraccount.ErrTransaction
	}

	queryUpdate := "UPDATE accounts SET balance = balance + ($1) WHERE account_id = $2"
	if _, err := tx.ExecContext(ctx, queryUpdate,
		payload.AmountDecimal(),
		accountID); err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Logger.Error("fail rollback tx", zap.Error(err))
			return models.AccountAttributes{}, erraccount.ErrServer
		}
		logger.Logger.Error("fail update balance account", zap.Error(err))
		return models.AccountAttributes{}, erraccount.ErrTransaction
	}

	if err := tx.Commit(); err != nil {
		logger.Logger.Error("fail commit tx", zap.Error(err))
		if err := tx.Rollback(); err != nil {
			logger.Logger.Error("fail rollback tx", zap.Error(err))
			return models.AccountAttributes{}, erraccount.ErrServer
		}
		return models.AccountAttributes{}, erraccount.ErrServer
	}

	query := "SELECT * FROM accounts WHERE account_id = $1"
	if err := r.pool.QueryRowxContext(ctx, query, accountID).StructScan(&account); err != nil {
		logger.Logger.Error("fail select account", zap.Error(err))
		return models.AccountAttributes{}, erraccount.ErrTransaction
	}

	return account, nil
}

//RollbackTransaction ...
func (r AccountRepository) RollbackTransaction(ctx context.Context, rollback dto.RollBack) error {
	tx, err := r.pool.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		logger.Logger.Error("fail start tx", zap.Error(err))
		return erraccount.ErrServer
	}
	query := "select * from update_account_balance($1, $2)"

	_, err = tx.ExecContext(ctx, query, rollback.AccountID, rollback.SourceType)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Logger.Error("fail rollback tx", zap.Error(err))
			return erraccount.ErrServer
		}
		return erraccount.ErrServer
	}
	if err := tx.Commit(); err != nil {
		logger.Logger.Error("fail commit tx", zap.Error(err))
		if err := tx.Rollback(); err != nil {
			logger.Logger.Error("fail rollback tx", zap.Error(err))
			return erraccount.ErrServer
		}
		return erraccount.ErrServer
	}

	return erraccount.ErrServer
}

//ConnectName ...
func (AccountRepository) ConnectName() int { return config.PaymentConnect }
