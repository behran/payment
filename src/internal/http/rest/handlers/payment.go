package handlers

import (
	"payment/internal/dto"
	"payment/internal/facade"
	"payment/internal/http/rest/middlewares"
	"payment/internal/http/rest/response"

	"github.com/valyala/fasthttp"
)

//CreateAccount ...
func CreateAccount(ctx *fasthttp.RequestCtx) {
	account := ctx.UserValue(middlewares.AccountKey).(dto.Account)

	result, err := facade.Service().Payment().CreateAccount(ctx, account)
	if err != nil {
		response.Error(err, ctx)
		return
	}
	response.OK(result, ctx)
}

//UpdateAccount ...
func UpdateAccount(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue(middlewares.AccountIDKey).(int)
	payload := ctx.UserValue(middlewares.PayloadKey).(dto.Payload)

	result, err := facade.Service().Payment().UpdateAccountAmount(ctx, id, payload)
	if err != nil {
		response.Error(err, ctx)
		return
	}
	response.OK(result, ctx)
}
