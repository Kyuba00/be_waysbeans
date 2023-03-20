package routes

import (
	"waysbeans_be/handlers"
	"waysbeans_be/pkg/middleware"
	"waysbeans_be/pkg/mysql"
	"waysbeans_be/repositories"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	CartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(CartRepository)

	e.GET("/carts", h.FindCarts, middleware.Auth)
	e.GET("/cart-id", h.FindCartsByTransaction, middleware.Auth)
	e.GET("/cart/:id", h.GetCart, middleware.Auth)
	e.POST("/cart", h.CreateCart, middleware.Auth)
	e.PATCH("/cart/:id", h.UpdateCart, middleware.Auth)
	e.DELETE("/cart/:id", h.DeleteCart, middleware.Auth)
}
