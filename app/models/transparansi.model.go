package models

import "gorm.io/gorm"

type Transparansi struct {
	gorm.Model
	AdminId string
	Admin	Admin
	CrowdfundId string
	Crowdfund Crowdfund
	Title string
	ShortDesc string
	Content	string `gorm:"type:text"`
	Thumbnail string
	IsPublished bool
}