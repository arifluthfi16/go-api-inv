package models

import "gorm.io/gorm"

type Investor struct {
	gorm.Model
	UserId string
	User User
	CrowdfundId string
	Crowdfund Crowdfund
	Fund int
	IsCompleted bool `gorm:"default:false"`
}
