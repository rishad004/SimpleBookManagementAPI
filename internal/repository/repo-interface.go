package repository

import (
	"github.com/rishad004/SimpleBookManagementAPI/internal/models"
	"gorm.io/gorm"
)

type Repo interface {
	UserRegister(username string, password string) error
	UserLogin(username string, password string) (string, error)
	CreateBook(bookName string, author string, category string) error
	CheckCategory(category string) error
	GetBooks(title, category string) ([]models.Books, error)
	UpdateBook(id int, bookName string, author string, category string) error
	DeleteBook(id int) error
	CreateCategory(category string) error
	GetCategories() ([]models.Categories, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{
		db: db,
	}
}
