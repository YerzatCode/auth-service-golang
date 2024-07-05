package controller

import (
	"github.com/YerzatCode/auth-service/internal/database"
	"github.com/YerzatCode/auth-service/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: database.DB,
	}

	api := r.Group("/api")
	{
		api.POST("/signin", h.Login)
		api.POST("/signup", h.Signup)

		validate := api.Group("/validate", middleware.RequireAuth)
		{
			validate.GET("/me", h.Validate)
		}
	}
}
