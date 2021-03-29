package errors

import "net/http"

var (
	ErrServer = PaymentError{
		Code:   http.StatusInternalServerError,
		Title:  "Server error",
		Detail: "",
	}
	ErrTransaction = PaymentError{
		Code:   http.StatusInternalServerError,
		Title:  "Server error",
		Detail: "Transaction Error",
	}
	ErrInvalidBalance = PaymentError{
		Code:   http.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "Balance cannot be zero",
	}
	ErrAccountNotFound = PaymentError{
		Code:   http.StatusNotFound,
		Title:  "Not Found",
		Detail: "Account not fount",
	}
	ErrBadBalance = PaymentError{
		Code:   http.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "Insufficient funds",
	}
)
