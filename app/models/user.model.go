package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayPict string
	Name string
	Email string `gorm:"unique"`
	Phonenumber string `gorm:"unique"`
	IdentityImage string
	Password string
	IsVerified bool
	GoogleId string
}
