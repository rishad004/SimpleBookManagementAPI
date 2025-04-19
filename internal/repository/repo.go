package repository

import (
	"github.com/rishad004/SimpleBookManagementAPI/internal/models"
	"github.com/rishad004/SimpleBookManagementAPI/internal/utils"
)

func (r *repo) UserRegister(username string, password string) error {

	pass, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	if err := r.db.Create(&models.Users{
		Username: username,
		Password: pass,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) UserLogin(username string, password string) (string, error) {
	var user models.Users
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *repo) CreateBook(bookName string, author string, category string) error {
	if err := r.db.Create(&models.Books{
		Title:    bookName,
		Author:   author,
		Category: category,
	}).Error; err != nil {
		return err
	}

	return nil

}

func (r *repo) CheckCategory(category string) error {
	var cat models.Categories
	if err := r.db.Where("name = ?", category).First(&cat).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) GetBooks(title, category string) ([]models.Books, error) {
	var books []models.Books

	if title != "" && category != "" {
		if err := r.db.Where("title = ? AND category = ?", title, category).Find(&books).Error; err != nil {
			return nil, err
		}
		return books, nil
	}

	if title != "" {
		if err := r.db.Where("title = ?", title).Find(&books).Error; err != nil {
			return nil, err
		}
		return books, nil
	}

	if category != "" {
		if err := r.db.Where("category = ?", category).Find(&books).Error; err != nil {
			return nil, err
		}
		return books, nil
	}

	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repo) UpdateBook(id int, bookName string, author string, category string) error {
	var book models.Books
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}

	if bookName != "" {
		book.Title = bookName
	}
	if author != "" {
		book.Author = author
	}
	if category != "" {
		book.Category = category
	}

	if err := r.db.Save(&book).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteBook(id int) error {
	var book models.Books
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateCategory(category string) error {
	if err := r.db.Create(&models.Categories{
		Name: category,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) GetCategories() ([]models.Categories, error) {
	var categories []models.Categories

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
