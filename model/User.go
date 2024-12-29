package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
}
