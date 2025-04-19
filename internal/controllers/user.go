package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rishad004/SimpleBookManagementAPI/internal/models"
)

func (h *ApiHandler) UserRegister(c *gin.Context) {
	var user models.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.svc.UserRegister(user.Username, user.Password); err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})

}

func (h *ApiHandler) UserLogin(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	token, err := h.svc.UserLogin(user.Username, user.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	c.JSON(200, gin.H{"message": "User logged in successfully",
		"token": token,
	})
}
