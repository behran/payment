package middlewares

import (
	"encoding/json"
	"slices"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"payment/internal/dto"
	"payment/internal/http/rest/response"
	"payment/internal/http/rest/response/errors"
	logger "payment/pkg/log"
)

const (
	win  = "win"
	lost = "lost"
)

// PayloadMiddleware ...
func PayloadMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var payload dto.Payload

		if err := json.Unmarshal(ctx.Request.Body(), &payload); err != nil {
			logger.Logger.Error("fail unmarshal request body `payload`", zap.Error(err))
			response.Error(ctx, errors.ErrInvalidBody)
			return
		}
		if err := validator.New().Struct(payload); err != nil {
			logger.Logger.Error("validate failed body `payload`", zap.Error(err))
			response.Error(ctx, errors.ErrInvalidBody)
			return
		}

		if !slices.Contains([]string{
			win, lost,
		}, payload.State) {
			response.Error(ctx, errors.ErrState)
			return
		}
		amount, err := strconv.ParseFloat(payload.Amount, 64)
		if err != nil {
			response.Error(ctx, errors.ErrAmount)
			return
		}
		payload.SetAmountDecimal(amount)

		ctx.SetUserValue(PayloadKey, payload)
		// next
		handle(ctx)
	}
}
