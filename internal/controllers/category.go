package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rishad004/SimpleBookManagementAPI/internal/models"
)

func (h *ApiHandler) CreateCategory(c *gin.Context) {
	var category models.Categories

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.svc.CreateCategory(category.Name); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(201, gin.H{"message": "Category created successfully"})
}

func (h *ApiHandler) GetCategories(c *gin.Context) {
	categories, err := h.svc.GetCategories()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get categories"})
		return
	}

	c.JSON(200, categories)
}
