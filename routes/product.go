package routes

import (
	"waysbeans_be/handlers"
	"waysbeans_be/pkg/middleware"
	"waysbeans_be/pkg/mysql"
	"waysbeans_be/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	ProductRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(ProductRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProduct)
	e.POST("/product", h.CreateProduct, middleware.UploadFile)
	e.PATCH("/product/:id", h.UpdateProduct, middleware.UploadFile)
	e.DELETE("/product/:id", h.DeleteProduct)
}
