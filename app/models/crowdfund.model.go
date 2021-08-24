package models

import "gorm.io/gorm"

type Crowdfund struct {
	gorm.Model
	AdminId string
	Admin	Admin
	Title string
	Need int
	Earn int
	Description string `gorm:"type:text"`
	IsActive bool
}