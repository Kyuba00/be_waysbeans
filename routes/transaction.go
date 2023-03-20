package routes

import (
	"waysbeans_be/handlers"
	"waysbeans_be/pkg/middleware"
	"waysbeans_be/pkg/mysql"
	"waysbeans_be/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	e.GET("/transactions", h.FindTransactions, middleware.Auth)
	e.GET("/transaction/:id", h.GetTransaction, middleware.Auth)
	e.POST("/transaction", h.CreateTransaction, middleware.Auth)
	e.PATCH("/transaction", h.UpdateTransaction, middleware.Auth)
	e.DELETE("/transaction/:id", h.DeleteTransaction, middleware.Auth)
	e.POST("/notification", h.Notification)
}
