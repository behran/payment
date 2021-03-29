package handlers

import (
	"payment/internal/http/rest/response"
	"payment/internal/http/rest/response/errors"

	"github.com/valyala/fasthttp"
)

//PageNotFound ...
func PageNotFound(ctx *fasthttp.RequestCtx) { response.Error(errors.ErrNotFound, ctx) }
