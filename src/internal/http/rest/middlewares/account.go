package middlewares

import (
	"encoding/json"
	"strconv"

	"payment/internal/dto"
	"payment/internal/http/rest/response"
	"payment/internal/http/rest/response/errors"
	logger "payment/pkg/log"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

//AccountIDMiddleware ...
func AccountIDMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		paramID := ctx.UserValue(AccountIDKey).(string)

		id, err := strconv.Atoi(paramID)
		if err != nil {
			logger.Logger.Error("invalid param `id`", zap.Error(err))
			response.Error(errors.ErrInvalidAccountID, ctx)
			return
		}
		ctx.SetUserValue(AccountIDKey, id)
		// next
		handle(ctx)
	}
}

//AccountMiddleware ...
func AccountMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var account dto.Account

		if err := json.Unmarshal(ctx.Request.Body(), &account); err != nil {
			logger.Logger.Error("fail unmarshal request body `account`", zap.Error(err))
			response.Error(errors.ErrInvalidBody, ctx)
			return
		}
		if err := validator.New().Struct(account); err != nil {
			logger.Logger.Error("validate failed body `account`", zap.Error(err))
			response.Error(errors.ErrInvalidBody, ctx)
			return
		}
		ctx.SetUserValue(AccountKey, account)
		// next
		handle(ctx)
	}
}
