package payment_test

import (
	"context"
	"testing"

	"payment/internal/config"
	"payment/internal/database"
	"payment/internal/domains/payment"
	"payment/internal/domains/payment/repositories"
	"payment/internal/dto"

	"gotest.tools/assert"
)

func TestService_UpdateAccountAmount(t *testing.T) {
	ctx := context.Background()
	conf := config.New()
	connector := database.NewConnectManager(conf)
	if err := database.InitConnections(connector); err != nil {
		assert.Error(t, err, "fail connect postgre")
		t.FailNow()
	}
	repository, err := repositories.NewAccount(connector)
	if err != nil {
		assert.Error(t, err, "fail create repository")
		t.FailNow()
	}

	services := payment.New(repository)
	account, err := services.CreateAccount(context.Background(), dto.Account{
		Name: "Mock",
	})
	if err != nil {
		assert.Error(t, err, "fail create accaunt")
		t.FailNow()
	}
	payload := dto.Payload{
		State: "lost",
		TransactionID: "qwerewqt",
	}
	payload.SetAmountDecimal(200)
	payload.SetSourceType("game")
	_, err = services.UpdateAccountAmount(ctx, account.ID, payload)
	assert.Error(t, err, "fail update amount")
}
