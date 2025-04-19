package service

import "github.com/rishad004/SimpleBookManagementAPI/internal/models"

type Repo interface {
	UserRegister(username string, password string) error
	UserLogin(username string, password string) (string, error)
	CreateBook(bookName string, author string, category string) error
	GetBooks(title string) ([]models.Books, error)
	CheckCategory(category string) error
	UpdateBook(id int, bookName string, author string, category string) error
	DeleteBook(id int) error
	CreateCategory(category string) error
	GetCategories(category string) ([]models.Categories, error)
}

type Svc interface {
	UserRegister(username string, password string) error
	UserLogin(username string, password string) (string, error)
	CreateBook(bookName string, author string, category string) error
	GetBooks(title string) ([]models.Books, error)
	UpdateBook(id int, bookName string, author string, category string) error
	DeleteBook(id int) error
	CreateCategory(category string) error
	GetCategories(category string) ([]models.Categories, error)
}

type svc struct {
	repo Repo
}

func NewSvc(repo Repo) Svc {
	return &svc{
		repo: repo,
	}
}
