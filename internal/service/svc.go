package service

import "github.com/rishad004/SimpleBookManagementAPI/internal/models"

func (s *svc) UserRegister(username string, password string) error {
	if err := s.repo.UserRegister(username, password); err != nil {
		return err
	}

	return nil
}

func (s *svc) UserLogin(username string, password string) (string, error) {
	token, err := s.repo.UserLogin(username, password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *svc) CreateBook(bookName string, author string, category string) error {
	if err := s.repo.CheckCategory(category); err != nil {
		return err
	}

	if err := s.repo.CreateBook(bookName, author, category); err != nil {
		return err
	}

	return nil
}

func (s *svc) GetBooks(title, category string) ([]models.Books, error) {
	books, err := s.repo.GetBooks(title, category)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *svc) UpdateBook(id int, bookName string, author string, category string) error {
	if err := s.repo.CheckCategory(category); err != nil {
		return err
	}

	if err := s.repo.UpdateBook(id, bookName, author, category); err != nil {
		return err
	}

	return nil
}

func (s *svc) DeleteBook(id int) error {
	if err := s.repo.DeleteBook(id); err != nil {
		return err
	}

	return nil
}

func (s *svc) CreateCategory(category string) error {
	if err := s.repo.CreateCategory(category); err != nil {
		return err
	}

	return nil
}

func (s *svc) GetCategories() ([]models.Categories, error) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
