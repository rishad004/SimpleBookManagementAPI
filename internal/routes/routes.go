package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishad004/SimpleBookManagementAPI/internal/controllers"
	"github.com/rishad004/SimpleBookManagementAPI/internal/middleware"
)

func Routes(r *gin.Engine, h *controllers.ApiHandler) {
	// User Routes
	r.POST("/register", h.UserRegister)
	r.POST("/login", h.UserLogin)

	// Book Routes
	r.POST("/books", middleware.MiddleWare, h.CreateBook)
	r.GET("/books", h.GetBooks)
	r.PUT("/books/:id", middleware.MiddleWare, h.UpdateBook)
	r.DELETE("/books/:id", middleware.MiddleWare, h.DeleteBook)

	// Category Routes
	r.POST("/categories", middleware.MiddleWare, h.CreateCategory)
	r.GET("/categories", h.GetCategories)
}
