package middlewares

import (
	"encoding/json"
	"strconv"

	"payment/internal/dto"
	"payment/internal/http/rest/response"
	"payment/internal/http/rest/response/errors"
	"payment/internal/tools"
	logger "payment/pkg/log"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

const (
	win  = "win"
	lost = "lost"
)

//PayloadMiddleware ...
func PayloadMiddleware(handle fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var payload dto.Payload

		if err := json.Unmarshal(ctx.Request.Body(), &payload); err != nil {
			logger.Logger.Error("fail unmarshal request body `payload`", zap.Error(err))
			response.Error(errors.ErrInvalidBody, ctx)
			return
		}
		if err := validator.New().Struct(payload); err != nil {
			logger.Logger.Error("validate failed body `payload`", zap.Error(err))
			response.Error(errors.ErrInvalidBody, ctx)
			return
		}

		if !tools.IsExistSlice(payload.State, []string{
			win, lost,
		}) {
			response.Error(errors.ErrState, ctx)
			return
		}
		amount, err := strconv.ParseFloat(payload.Amount, 64)
		if err != nil {
			response.Error(errors.ErrAmount, ctx)
			return
		}

		payload.SetAmountDecimal(amount)

		ctx.SetUserValue(PayloadKey, payload)
		// next
		handle(ctx)
	}
}
