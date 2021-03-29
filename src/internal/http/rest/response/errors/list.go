package errors

import "net/http"

var (
	ErrNotFound = ResponseError{
		Code:   http.StatusNotFound,
		Title:  "Not Found",
		Detail: "Page not found",
	}
	ErrInvalidBody = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "Invalid body",
	}
	ErrState = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "The `state` field must be `win`|`lost`",
	}
	ErrAmount = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "The `amount` field must be decimal",
	}
	ErrSourceType = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "The header `Source-Type' must be `game`|`server`|`payment`",
	}
	ErrInvalidAccountID = ResponseError{
		Code:   http.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "Invalid param `id`",
	}
)
