package routes

import (
	"waysbeans_be/handlers"
	"waysbeans_be/pkg/middleware"
	"waysbeans_be/pkg/mysql"
	"waysbeans_be/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	ProfileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(ProfileRepository)

	e.GET("/profiles", h.FindProfiles, middleware.Auth)
	e.GET("/profile/:id", h.GetProfile, middleware.Auth)
	e.POST("/profile", h.CreateProfile, middleware.Auth, middleware.UploadFile)
	e.PATCH("/profile/:id", h.UpdateProfile, middleware.Auth, middleware.UploadFile)
	e.DELETE("/profile/:id", h.DeleteProfile, middleware.Auth)
}
