package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name 		string
	Username 	string `gorm:"unique"`
	Password 	string
	Email		string `gorm:"unique"`
	HakAkses 	string
}