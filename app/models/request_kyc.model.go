package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type RequestKYC struct {
	gorm.Model
	UserId string
	User User
	AdminId	sql.NullString
	Admin Admin
	Filepath string
	RejectReason sql.NullString
	IsAccepted bool
}