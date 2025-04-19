package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type Books struct {
	gorm.Model
	Title    string `json:"title" gorm:"unique"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

type Categories struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
