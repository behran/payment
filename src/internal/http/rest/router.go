package rest

import (
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"payment/internal/http/rest/handlers"
	"payment/internal/http/rest/middlewares"

	"github.com/buaazp/fasthttprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterRoutes Router List ...
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
	// 404 PAGE ...
	router.NotFound = handlers.PageNotFound
	// HANDLER PANIC ...
	router.PanicHandler = handlers.PagePanic
	// METRICS ...
	router.GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
}
