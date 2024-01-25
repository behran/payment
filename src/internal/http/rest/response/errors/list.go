package errors

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

var (
	ErrServer = ResponseError{
		Code:   fasthttp.StatusInternalServerError,
		Title:  fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
		Detail: "",
	}
	ErrNotFound = ResponseError{
		Code:   fasthttp.StatusNotFound,
		Title:  fasthttp.StatusMessage(fasthttp.StatusNotFound),
		Detail: "Page isn't found",
	}
	ErrInvalidBody = ResponseError{
		Code:   fasthttp.StatusBadRequest,
		Title:  fasthttp.StatusMessage(fasthttp.StatusBadRequest),
		Detail: "Invalid body",
	}
	ErrState = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  fasthttp.StatusMessage(fasthttp.StatusBadRequest),
		Detail: "The `state` field must be `win`|`lost`",
	}
	ErrAmount = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  fasthttp.StatusMessage(fasthttp.StatusBadRequest),
		Detail: "The `amount` field must be decimal",
	}
	ErrSourceType = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  fasthttp.StatusMessage(fasthttp.StatusBadRequest),
		Detail: "The header `Source-Type' must be `game`|`server`|`payment`",
	}
	ErrInvalidAccountID = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  fasthttp.StatusMessage(fasthttp.StatusBadRequest),
		Detail: "Invalid param `id`",
	}
)
