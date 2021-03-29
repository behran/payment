package rest

import (
	"payment/internal/http/rest/handlers"
	"payment/internal/http/rest/middlewares"

	"github.com/buaazp/fasthttprouter"
)

//RegisterRoutes Router List ...
func RegisterRoutes(router *fasthttprouter.Router) {
	// Create account
	router.POST("/payment/account",
		middlewares.ApplyMiddleware(
			handlers.CreateAccount,
			middlewares.CreateAccountPayload...,
		),
	)
	// update account balance
	router.POST("/payment/account/:id",
		middlewares.ApplyMiddleware(
			handlers.UpdateAccount,
			middlewares.UpdateAccountPayload...,
		),
	)
	// DOCUMENTATION ...
	router.ServeFiles("/documentation/*filepath", "./documentation")
	// 404 page ...
	router.NotFound = handlers.PageNotFound
}
