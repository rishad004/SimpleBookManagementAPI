package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rishad004/SimpleBookManagementAPI/internal/models"
)

func (h *ApiHandler) CreateBook(c *gin.Context) {
	var book models.Books

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.svc.CreateBook(book.Title, book.Author, book.Category); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(201, gin.H{"message": "Book created successfully"})
}

func (h *ApiHandler) GetBooks(c *gin.Context) {
	title := c.Query("title")

	books, err := h.svc.GetBooks(title)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get books"})
		return
	}

	c.JSON(200, books)
}

func (h *ApiHandler) UpdateBook(c *gin.Context) {
	var book models.Books
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.svc.UpdateBook(id, book.Title, book.Author, book.Category); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(200, gin.H{"message": "Book updated successfully"})
}

func (h *ApiHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := h.svc.DeleteBook(id); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(200, gin.H{"message": "Book deleted successfully"})
}
